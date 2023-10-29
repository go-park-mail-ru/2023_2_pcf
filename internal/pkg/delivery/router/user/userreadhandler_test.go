package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserReadHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User: mockUserUseCase,
	}

	req, err := http.NewRequest("GET", "/user?login=testuser", nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedUser := &entities.User{
		Id:       1,
		Login:    "testuser",
		Password: "test",
		FName:    "test",
		LName:    "test",
	}

	mockUserUseCase.EXPECT().UserRead("testuser").Return(expectedUser, nil)
	rr := httptest.NewRecorder()

	userRouter.UserReadHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseUser entities.User
	if err := json.NewDecoder(rr.Body).Decode(&responseUser); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedUser, &responseUser)
}
