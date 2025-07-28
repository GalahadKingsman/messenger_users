package userservice

import (
	"context"
	"errors"
	"testing"

	"github.com/GalahadKingsman/messenger_users/internal/models"
	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ===== Интерфейс и тестовый сервис =====

type userRepoGetInterface interface {
	GetUsers(ctx context.Context, filter *models.GetUserFilter) ([]models.User, error)
}

type testServiceGet struct {
	userRepo userRepoGetInterface
}

func (s *testServiceGet) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	filter := &models.GetUserFilter{
		Id:        req.Id,
		Login:     req.Login,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}

	users, err := s.userRepo.GetUsers(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbUsers := make([]*pb.GetUserResponse_User, 0, len(users))
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.GetUserResponse_User{
			Id:        user.ID,
			Login:     user.Login,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		})
	}
	return &pb.GetUserResponse{
		Users: pbUsers,
	}, nil
}

// ===== Мок =====

type mockUserRepoGet struct {
	mock.Mock
}

func (m *mockUserRepoGet) GetUsers(ctx context.Context, filter *models.GetUserFilter) ([]models.User, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).([]models.User), args.Error(1)
}

// ===== Тесты =====

func TestGetUser_Success(t *testing.T) {
	mockRepo := new(mockUserRepoGet)

	login := "john_doe"

	mockRepo.On("GetUsers", mock.Anything, mock.MatchedBy(func(f *models.GetUserFilter) bool {
		return f != nil && f.Login != nil && *f.Login == login
	})).Return([]models.User{
		{
			ID:        1,
			Login:     "john_doe",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			Phone:     "1234567890",
		},
	}, nil)

	service := &testServiceGet{userRepo: mockRepo}

	req := &pb.GetUserRequest{
		Login: &login,
	}

	resp, err := service.GetUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Len(t, resp.Users, 1)
	assert.Equal(t, "john_doe", resp.Users[0].Login)
	assert.Equal(t, "John", resp.Users[0].FirstName)
}

func TestGetUser_RepoError(t *testing.T) {
	mockRepo := new(mockUserRepoGet)

	email := "notfound@example.com"

	mockRepo.On("GetUsers", mock.Anything, mock.MatchedBy(func(f *models.GetUserFilter) bool {
		return f != nil && f.Email != nil && *f.Email == email
	})).Return([]models.User(nil), errors.New("db failure"))

	service := &testServiceGet{userRepo: mockRepo}

	req := &pb.GetUserRequest{
		Email: &email,
	}

	resp, err := service.GetUser(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)

	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Internal, st.Code())
	assert.Equal(t, "db failure", st.Message())
}
