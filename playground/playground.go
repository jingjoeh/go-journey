package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

type UserRepository interface {
	FindByID(id int) (User, error)
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUser(id int) (User, error) {
	return s.repo.FindByID(id)
}

// Mock
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(id int) (User, error) {
	args := m.Called(id)

	user := args.Get(0).(User)
	err := args.Error(1)

	return user, err
}

func TestUserService_GetUser(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.
		On("FindByID", 99).
		Return(User{ID: 99, Name: "Joe"}, nil).
		Once()

	service := UserService{
		repo: mockRepo,
	}

	user, err := service.GetUser(99)

	// TODO 1: ถ้ามี error ให้ test หยุดทันที
	require.NoError(t, err)
	// TODO 2: ตรวจว่า user.Name == "Joe"
	assert.Equal(t, "Joe", user.Name)

	mockRepo.AssertExpectations(t)
}

var ErrUserNotFound = errors.New("user not found")
var ErrInactiveUser = errors.New("inactive user")

func TestUserService_GetUser_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.
		On("FindByID", 99).
		Return(User{}, ErrUserNotFound).
		Once()

	service := UserService{
		repo: mockRepo,
	}

	_, err := service.GetUser(99)

	assert.ErrorIs(t, err, ErrUserNotFound)
	require.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func (s *UserService) GetActiveUser(id int) (User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return User{}, err
	}

	if !user.Active {
		return User{}, ErrInactiveUser
	}

	return user, nil
}

func TestUserService_GetActiveUser(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.
		On("FindByID", 99).
		Return(User{
			ID:     99,
			Name:   "Joe",
			Active: true,
		}, nil).
		Once()

	service := UserService{
		repo: mockRepo,
	}

	user, err := service.GetActiveUser(99)

	require.NoError(t, err)
	assert.Equal(t, "Joe", user.Name)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetActiveUser_Inactive(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.
		On("FindByID", 99).
		Return(User{
			ID:     99,
			Name:   "Joe",
			Active: false,
		}, nil).
		Once()

	service := UserService{
		repo: mockRepo,
	}

	// TODO 1: เรียก GetActiveUser
	_, err := service.GetActiveUser(99)

	assert.ErrorIs(t, ErrInactiveUser, err)

	mockRepo.AssertExpectations(t)
}
