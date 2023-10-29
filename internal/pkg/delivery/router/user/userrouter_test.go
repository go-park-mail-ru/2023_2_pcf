package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestUserRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)

	userRouter := UserRouter{
		router: mux.NewRouter(),
		User:   mockUserUseCase,
	}

	ConfigureRouter(&userRouter)

	req_read := httptest.NewRequest("GET", "/user", nil)
	req_delete := httptest.NewRequest("DELETE", "/user", nil)

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

	req_create := httptest.NewRequest("POST", "/user", bytes.NewReader(userJSON))
	rr := httptest.NewRecorder()
	rd := httptest.NewRecorder()
	rc := httptest.NewRecorder()

	mockUserUseCase.EXPECT().UserRead(gomock.Any()).Return(nil, nil)
	mockUserUseCase.EXPECT().UserDelete(gomock.Any()).Return(nil)
	mockUserUseCase.EXPECT().UserCreate(gomock.Any()).Return(fakeUser, nil)

	userRouter.router.ServeHTTP(rr, req_read)
	assert.Equal(t, http.StatusOK, rr.Code)

	userRouter.router.ServeHTTP(rr, req_delete)
	assert.Equal(t, http.StatusOK, rd.Code)

	userRouter.router.ServeHTTP(rr, req_create)
	assert.Equal(t, http.StatusOK, rc.Code)

}
