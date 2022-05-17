package server

import (
	"GRPC/models"
	"GRPC/repository"
	"GRPC/studenttpb"
	"GRPC/testpb"
	"context"
	"io"
	"time"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, test *testpb.Test) (*testpb.SetTestResponse, error) {
	err := s.repo.SetTest(ctx, &models.Test{
		Id:   test.GetId(),
		Name: test.GetName(),
	})
	if err != nil {
		return &testpb.SetTestResponse{}, err
	}
	return &testpb.SetTestResponse{
		Id:   test.GetId(),
		Name: test.GetName(),
	}, err
}

func (s TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}
		question := &models.Question{
			Id:       msg.Id,
			Answer:   msg.Answer,
			Question: msg.Question,
			TestId:   msg.TestId,
		}
		err = s.repo.SetQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s TestServer) EnrollmentStudents(stream testpb.TestService_EnrollmentStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}
		enrollment := &models.Enrollments{
			StudentId: msg.StudentId,
			TestId:    msg.TestId,
		}
		err = s.repo.SetEnrollment(context.Background(), enrollment)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s TestServer) GetStudentsPerTest(req *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	test, err := s.repo.GetStudentsPerTest(context.Background(), req.GetId())
	if err != nil {
		return err
	}
	for _, student := range test {
		student := &studenttpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}
		err := stream.Send(student)
		time.Sleep(2 * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
