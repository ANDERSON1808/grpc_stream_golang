package server

import (
	"GRPC/models"
	"GRPC/repository"
	"GRPC/studenttpb"
	"context"
)

type Server struct {
	repo repository.Repository
	studenttpb.UnimplementedStudentServiceServer
}

func NewServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s Server) GetStudent(ctx context.Context, req *studenttpb.GetStudentRequest) (*studenttpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &studenttpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s Server) SetStudent(ctx context.Context, student *studenttpb.Student) (*studenttpb.SetStudentResponse, error) {
	err := s.repo.SetStudent(ctx, &models.Student{
		Id:   student.GetId(),
		Name: student.GetName(),
		Age:  student.GetAge(),
	})
	if err != nil {
		return &studenttpb.SetStudentResponse{}, err
	}
	return &studenttpb.SetStudentResponse{
		Id: student.Id,
	}, nil
}
