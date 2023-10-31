package router

import (
	"AdHub/internal/pkg/entities/mock_entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserDeleteHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User: mockUserUseCase,
	}

	req, err := http.NewRequest("DELETE", "/user?login=testuser@mail.ru", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockUserUseCase.EXPECT().UserDelete("testuser@mail.ru").Return(nil)

	rr := httptest.NewRecorder()

	userRouter.UserDeleteHandler(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}
