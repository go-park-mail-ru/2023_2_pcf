package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User:    mockUserUseCase,
		Session: mockSession,
	}

	fakeUser := &entities.User{
		Login:    "test@example.com",
		Password: "password123",
		FName:    "John",
		LName:    "Doe",
	}

	mockUserUseCase.EXPECT().UserRead(gomock.Any()).Return(fakeUser, nil)

	fakeSession := &entities.Session{
		Token:  "test",
		UserId: 1,
	}

	mockSession.EXPECT().
		Auth(gomock.Eq(fakeUser)).
		Return(fakeSession, nil)

	userJSON, _ := json.Marshal(fakeUser)

	req, _ := http.NewRequest("POST", "/auth", strings.NewReader(string(userJSON)))
	rr := httptest.NewRecorder()

	userRouter.AuthHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
