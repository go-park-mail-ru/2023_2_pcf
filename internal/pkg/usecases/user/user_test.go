package user

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserUseCase_UserCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockUserRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeUser := &entities.User{
		Id:          1,
		Login:       "user@example.com",
		Password:    "1234556",
		LName:       "Doe",
		FName:       "John",
		CompanyName: "Smith",
		Avatar:      "def.jpg",
		BalanceId:   1,
	}

	mockRepo.EXPECT().Create(gomock.Eq(fakeUser)).Return(fakeUser, nil)

	createdUser, err := useCase.UserCreate(fakeUser)
	assert.NoError(t, err)
	assert.Equal(t, fakeUser, createdUser)
}

func TestUserUseCase_UserDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockUserRepoInterface(ctrl)

	useCase := New(mockRepo)

	userMail := "user@example.com"

	mockRepo.EXPECT().Remove(userMail).Return(nil)

	err := useCase.UserDelete(userMail)
	assert.NoError(t, err)
}

func TestUserUseCase_UserRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockUserRepoInterface(ctrl)

	useCase := New(mockRepo)

	login := "user@example.com"

	fakeUser := &entities.User{
		Id:          1,
		Login:       login,
		Password:    "1234556",
		LName:       "Doe",
		FName:       "John",
		CompanyName: "Smith",
		Avatar:      "defaul.jpg",
		BalanceId:   1,
	}

	mockRepo.EXPECT().ReadByLogin(login).Return(fakeUser, nil)

	user, err := useCase.UserReadByLogin(login)
	assert.NoError(t, err)
	assert.Equal(t, fakeUser, user)
}

func TestUserUseCase_UserReadById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockUserRepoInterface(ctrl)

	useCase := New(mockRepo)

	id := 1

	fakeUser := &entities.User{
		Id:          1,
		Login:       "login@login.ru",
		Password:    "1234556",
		LName:       "Doe",
		FName:       "John",
		CompanyName: "Smith",
		Avatar:      "defaul.jpg",
		BalanceId:   1,
	}

	mockRepo.EXPECT().ReadById(id).Return(fakeUser, nil)

	user, err := useCase.UserReadById(id)
	assert.NoError(t, err)
	assert.Equal(t, fakeUser, user)
}
