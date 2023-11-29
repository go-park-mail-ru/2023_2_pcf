package router

import (
	mock_entities2 "AdHub/auth/pkg/entities/mock_entities"
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"AdHub/proto/api"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)
	mockSession := mock_entities2.NewMockSessionUseCaseInterface(ctrl)
	mockCsrf := mock_entities.NewMockCsrfUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User:    mockUserUseCase,
		Session: mockSession,
		Csrf:    mockCsrf,
	}

	fakeUser := &entities.User{
		Login:    "test@example.com",
		Password: "password123",
		FName:    "John",
		LName:    "Doe",
	}

	//fakeSession := &entities2.Session{
	//	Token:  "test",
	//	UserId: 1,
	//}

	// Mocking the UserReadByLogin method to return fakeUser
	mockUserUseCase.EXPECT().
		UserReadByLogin(fakeUser.Login).
		Return(fakeUser, nil)

	// Mocking the Auth method to return fakeSession
	mockSession.EXPECT().
		Auth(context.Background(), gomock.Any(), gomock.Any()).
		Return(&api.AuthResponse{}, nil)

	// Mocking the CsrfCreate
	mockCsrf.EXPECT().CsrfCreate(0).Return(&entities.Csrf{}, nil)

	userJSON, _ := json.Marshal(fakeUser)

	req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(userJSON))
	rr := httptest.NewRecorder()

	userRouter.AuthHandler(rr, req)

	// Check for the expected status code
	assert.Equal(t, http.StatusOK, rr.Code)
}
