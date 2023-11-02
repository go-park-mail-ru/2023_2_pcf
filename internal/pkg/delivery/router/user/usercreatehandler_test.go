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

func TestUserCreateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User:    mockUserUseCase,
		Session: mockSession,
	}

	fakeUser := &entities.User{
		Id:          1,
		Login:       "testuser@mail.ru",
		Password:    "test312",
		FName:       "test",
		LName:       "test",
		CompanyName: "Yandex",
		Avatar:      "test.jpg",
	}

	fakeSession := &entities.Session{
		Token:  "test",
		UserId: 1,
	}

	userJSON, _ := json.Marshal(fakeUser)

	req, err := http.NewRequest("POST", "/user", strings.NewReader(string(userJSON)))
	if err != nil {
		t.Fatal(err)
	}

	mockUserUseCase.EXPECT().
		UserCreate(gomock.Any()).
		Return(fakeUser, nil)

	mockSession.EXPECT().
		Auth(gomock.Eq(fakeUser)).
		Return(fakeSession, nil)

	rr := httptest.NewRecorder()

	userRouter.UserCreateHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	responseJSON := rr.Body.String()
	assert.JSONEq(t, string(userJSON), responseJSON)
}
