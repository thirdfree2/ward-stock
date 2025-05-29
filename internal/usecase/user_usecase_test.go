package usecase_test

import (
	"errors"
	"testing"
	"ward-stock-backend/internal/domain"
	"ward-stock-backend/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

type MockRunningNumberRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) List() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRunningNumberRepository) NextNumber(prefix string, format string, length int) (string, error) {
	args := m.Called(prefix, format, length)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) GetPermissions(userID uint) ([]string, error) {
	args := m.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}

func TestCreateUser_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRunningRepo := new(MockRunningNumberRepository)

	uc := usecase.NewUserUsecase(mockUserRepo, mockRunningRepo)

	user := &domain.User{Name: "John"}

	mockRunningRepo.
		On("NextNumber", "USER", "USR", 5).
		Return("USR00001", nil)

	mockUserRepo.
		On("Create", mock.MatchedBy(func(u *domain.User) bool {
			return u.Code == "USR00001" && u.Name == "John"
		})).
		Return(nil)

	err := uc.CreateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, "USR00001", user.Code)

	mockRunningRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestCreate_Failed_NextNumber(t *testing.T) {
	mockUserRepo := new(MockUserRepository)             // เรียกใช้ UserRepo
	mockRunningRepo := new(MockRunningNumberRepository) // เรียนกใช้ RunningRepo

	uc := usecase.NewUserUsecase(mockUserRepo, mockRunningRepo) // สร้าง Usecase จาก struct

	user := &domain.User{Name: "John"} // สร้าง model

	mockRunningRepo.
		On("NextNumber", "USER", "USR", 5).
		Return("", assert.AnError) // สร้าง RunningNumber จาก Repo

	err := uc.CreateUser(user) // สร้าง User

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)

	mockUserRepo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestCreateUser_WhenNextNumberFails_ShouldNotCallCreate(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRunningRepo := new(MockRunningNumberRepository)

	uc := usecase.NewUserUsecase(mockUserRepo, mockRunningRepo)
	user := &domain.User{Name: "John"}

	// Mock การสร้างรหัสล้มเหลว
	mockRunningRepo.
		On("NextNumber", "USER", "USR", 5).
		Return("", errors.New("failed to generate code"))

	err := uc.CreateUser(user)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to generate code")

	// ตรวจว่าฟังก์ชัน Create ไม่ควรถูกเรียก
	mockUserRepo.AssertNotCalled(t, "Create", mock.Anything)
}
