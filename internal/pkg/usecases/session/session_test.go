package session

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSessionUseCase_SessionCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockSessionRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeSession := &entities.Session{
		Token:  "fakeToken",
		UserId: 1,
	}

	mockRepo.EXPECT().Create(gomock.Eq(fakeSession)).Return(fakeSession, nil)

	createdSession, err := useCase.SessionCreate(fakeSession)
	assert.NoError(t, err)
	assert.Equal(t, fakeSession, createdSession)
}

func TestSessionUseCase_SessionRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockSessionRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeSession := &entities.Session{
		Token:  "fakeToken",
		UserId: 1,
	}

	mockRepo.EXPECT().Read(gomock.Eq(fakeSession)).Return(fakeSession, nil)

	createdSession, err := useCase.SessionRead(fakeSession)
	assert.NoError(t, err)
	assert.Equal(t, fakeSession, createdSession)
}

func TestSessionUseCase_SessionRemove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockSessionRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeSession := &entities.Session{
		Token:  "fakeToken",
		UserId: 1,
	}

	mockRepo.EXPECT().Remove(gomock.Eq(fakeSession)).Return(nil)

	err := useCase.SessionRemove(fakeSession)
	assert.NoError(t, err)
}

func TestSessionUseCase_SessionContains(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockSessionRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeSession := &entities.Session{
		Token:  "fakeToken",
		UserId: 1,
	}

	mockRepo.EXPECT().Contains(gomock.Eq(fakeSession)).Return(true, nil)

	_, err := useCase.SessionContains(fakeSession)
	assert.NoError(t, err)
}

func TestSessionUseCase_Auth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockSessionRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeUser := &entities.User{
		Id:       1,
		Login:    "user@example.com",
		Password: "1234556",
		LName:    "joe",
		FName:    "joe",
	}

	fakeSession := &entities.Session{
		Token:  "fakeToken",
		UserId: 1,
	}

	mockRepo.EXPECT().Contains(gomock.Any()).Return(false, nil).AnyTimes()

	mockRepo.EXPECT().Create(gomock.Any()).Return(fakeSession, nil)

	createdSession, err := useCase.Auth(fakeUser)
	assert.NoError(t, err)
	assert.Equal(t, fakeSession, createdSession)
}

func TestSessionUseCase_GetUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockSessionRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeToken := "fakeToken"

	fakeSession := &entities.Session{
		Token:  fakeToken,
		UserId: 1,
	}

	mockRepo.EXPECT().Read(gomock.Eq(&entities.Session{Token: fakeToken, UserId: 0})).Return(fakeSession, nil)

	userId, err := useCase.GetUserId(fakeToken)
	assert.NoError(t, err)
	assert.Equal(t, fakeSession.UserId, userId)
}
