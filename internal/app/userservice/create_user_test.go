package userservice

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/GalahadKingsman/messenger_users/internal/models"
	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type userRepoCreateInterface interface {
	CreateUser(user models.User) (int, error)
}

type testServiceCreate struct {
	userRepo userRepoCreateInterface
}

func (s *testServiceCreate) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	if req.GetFirstName() == "" || req.GetEmail() == "" {
		return nil, errors.New("имя и email обязательны для заполнения")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Login:     req.GetLogin(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
		Password:  string(hashedPassword),
	}

	id, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании пользователя: %v", err)
	}

	return &pb.CreateResponse{
		Success: fmt.Sprintf("Пользователь успешно создан с ID: %d", id),
	}, nil
}

type mockUserRepoCreate struct {
	mock.Mock
}

func (m *mockUserRepoCreate) CreateUser(user models.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mockUserRepoCreate)

	mockRepo.On("CreateUser", mock.AnythingOfType("models.User")).Return(101, nil)

	s := &testServiceCreate{userRepo: mockRepo}

	req := &pb.CreateRequest{
		Login:     "testuser",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Phone:     "1234567890",
		Password:  "secret",
	}

	resp, err := s.CreateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Contains(t, resp.Success, "Пользователь успешно создан с ID: 101")
}

func TestCreateUser_ValidationError(t *testing.T) {
	mockRepo := new(mockUserRepoCreate)

	s := &testServiceCreate{userRepo: mockRepo}

	req := &pb.CreateRequest{
		Login:     "testuser",
		FirstName: "",
		Email:     "", // <== ошибка валидации
		Password:  "secret",
	}

	resp, err := s.CreateUser(context.Background(), req)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "имя и email обязательны для заполнения")
}

func TestCreateUser_RepoError(t *testing.T) {
	mockRepo := new(mockUserRepoCreate)

	mockRepo.On("CreateUser", mock.AnythingOfType("models.User")).Return(0, errors.New("db error"))

	s := &testServiceCreate{userRepo: mockRepo}

	req := &pb.CreateRequest{
		Login:     "user2",
		FirstName: "Alice",
		LastName:  "Smith",
		Email:     "alice@example.com",
		Phone:     "9876543210",
		Password:  "mypassword",
	}

	resp, err := s.CreateUser(context.Background(), req)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "ошибка при создании пользователя: db error")
}
