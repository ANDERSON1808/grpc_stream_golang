package database

import (
	"GRPC/models"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type PosgresRepository struct {
	db *sql.DB
}

func NewPosgresRepository(url string) (*PosgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PosgresRepository{db}, nil
}

func (receiver PosgresRepository) SetStudent(ctx context.Context, student *models.Student) (err error) {
	execContext, err := receiver.db.ExecContext(ctx, "INSERT INTO students (id,name,age) VALUES ($1,$2,$3)", student.Id, student.Name, student.Age)
	if err != nil {
		return
	}
	affected, err := execContext.RowsAffected()
	if err != nil || affected < 1 {
		return
	}
	return
}

func (receiver PosgresRepository) GetStudent(ctx context.Context, id string) (student models.Student, err error) {
	rows, err := receiver.db.QueryContext(ctx, "SELECT id ,name ,age FROM students WHERE id = $1", id)
	if err != nil {
		return student, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			return
		}
	}()
	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return student, err
		}
	}
	return
}

func (receiver PosgresRepository) SetTest(ctx context.Context, test *models.Test) (err error) {
	execContext, err := receiver.db.ExecContext(ctx, "INSERT INTO tests (id,name) VALUES ($1,$2)", test.Id, test.Name)
	if err != nil {
		return
	}
	affected, err := execContext.RowsAffected()
	if err != nil || affected < 1 {
		return
	}
	return
}

func (receiver PosgresRepository) GetTest(ctx context.Context, id string) (test models.Test, err error) {
	rows, err := receiver.db.QueryContext(ctx, "SELECT id ,name  FROM tests WHERE id = $1", id)
	if err != nil {
		return test, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			return
		}
	}()
	for rows.Next() {
		err := rows.Scan(&test.Id, &test.Name)
		if err != nil {
			return test, err
		}
	}
	return
}

func (receiver PosgresRepository) SetQuestion(ctx context.Context, question *models.Question) (err error) {
	execContext, err := receiver.db.ExecContext(ctx, "INSERT INTO questions (id,answer,question,test_id) VALUES ($1,$2,$3,$4)", question.Id, question.Answer, question.Question, question.TestId)
	if err != nil {
		return
	}
	affected, err := execContext.RowsAffected()
	if err != nil || affected < 1 {
		return
	}
	return
}

func (receiver PosgresRepository) GetStudentsPerTest(ctx context.Context, testId string) (students []models.Student, err error) {
	rows, err := receiver.db.QueryContext(ctx, "SELECT id,name, age FROM students WHERE id in(SELECT student_id FROM enrollments WHERE test_id = $1)", testId)
	if err != nil {
		return
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			return
		}
	}()
	for rows.Next() {
		var student = models.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (receiver PosgresRepository) SetEnrollment(ctx context.Context, enrollments *models.Enrollments) (err error) {
	execContext, err := receiver.db.ExecContext(ctx, "INSERT INTO enrollments (student_id, test_id) VALUES ($1,$2)", enrollments.StudentId, enrollments.TestId)
	if err != nil {
		return
	}
	affected, err := execContext.RowsAffected()
	if err != nil || affected < 1 {
		return
	}
	return
}
