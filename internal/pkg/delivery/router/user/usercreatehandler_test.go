package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserCreateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User: mockUserUseCase,
	}

	fakeUser := &entities.User{
		Id:       1,
		Login:    "testuser@mail.ru",
		Password: "test12",
		FName:    "test12",
		LName:    "test12",
	}

	userJSON, err := json.Marshal(fakeUser)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/user", bytes.NewReader(userJSON))
	rr := httptest.NewRecorder()

	mockUserUseCase.EXPECT().UserCreate(fakeUser).Return(fakeUser, nil)

	userRouter.UserCreateHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var responseUser entities.User
	if err := json.NewDecoder(rr.Body).Decode(&responseUser); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, fakeUser, &responseUser)
}
