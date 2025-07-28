package userservice

import (
	"context"
	"errors"
	"testing"

	"github.com/GalahadKingsman/messenger_users/internal/jwt"
	"github.com/GalahadKingsman/messenger_users/internal/models"
	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type userRepoInterface interface {
	GetUserByLogin(login string) (*models.User, error)
}

type testService struct {
	userRepo userRepoInterface
}

func (s *testService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userRepo.GetUserByLogin(req.Login)
	if err != nil || user.Password == "" {
		return &pb.LoginResponse{Message: "Некорректный логин или пароль"}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return &pb.LoginResponse{Message: "Некорректный логин или пароль"}, nil
	}

	token, err := jwt.GenerateToken(string(rune(user.ID)))
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Message: "Успешный вход",
		UserId:  int32(user.ID),
		Token:   token,
	}, nil
}

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) GetUserByLogin(login string) (*models.User, error) {
	args := m.Called(login)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*models.User), args.Error(1)
}

func TestLogin_Success(t *testing.T) {
	password := "qwerty"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockRepo := new(mockUserRepo)
	mockRepo.On("GetUserByLogin", "testuser").Return(&models.User{
		ID:       42,
		Login:    "testuser",
		Password: string(hashedPassword),
	}, nil)

	s := &testService{userRepo: mockRepo}

	req := &pb.LoginRequest{
		Login:    "testuser",
		Password: password,
	}

	resp, err := s.Login(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Успешный вход", resp.Message)
	assert.Equal(t, int32(42), resp.UserId)
	assert.NotEmpty(t, resp.Token)
}

func TestLogin_WrongPassword(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correct"), bcrypt.DefaultCost)

	mockRepo := new(mockUserRepo)
	mockRepo.On("GetUserByLogin", "testuser").Return(&models.User{
		ID:       43,
		Login:    "testuser",
		Password: string(hashedPassword),
	}, nil)

	s := &testService{userRepo: mockRepo}

	req := &pb.LoginRequest{
		Login:    "testuser",
		Password: "wrong",
	}

	resp, err := s.Login(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Некорректный логин или пароль", resp.Message)
	assert.Empty(t, resp.Token)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(mockUserRepo)
	mockRepo.On("GetUserByLogin", "nouser").Return(nil, errors.New("not found"))

	s := &testService{userRepo: mockRepo}

	req := &pb.LoginRequest{
		Login:    "nouser",
		Password: "pass",
	}

	resp, err := s.Login(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Некорректный логин или пароль", resp.Message)
	assert.Empty(t, resp.Token)
}
