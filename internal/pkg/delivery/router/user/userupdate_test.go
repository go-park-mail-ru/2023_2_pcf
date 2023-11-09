package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserUpdateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)
	mockFileUseCase := mock_entities.NewMockFileUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User: mockUserUseCase,
		File: mockFileUseCase,
	}

	// Создаем тестовый файл (аватар)
	avatarData := []byte("test_avatar_data")
	avatarFilename := "test_avatar.jpg"

	// Создаем мультипарт форму с файлом
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("avatar", avatarFilename)
	if err != nil {
		t.Fatal(err)
	}
	part.Write(avatarData)
	writer.Close()

	req, err := http.NewRequest("POST", "/user/update", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Создаем тестового пользователя
	fakeUser := &entities.User{
		Id:          1,
		Login:       "testuser",
		Password:    "test",
		FName:       "test",
		LName:       "test",
		CompanyName: "Yandex",
		Avatar:      "test.jpg",
	}

	// Устанавливаем ожидания для моков
	mockFileUseCase.EXPECT().Save(gomock.Any(), avatarFilename).Return("avatar_path", nil)
	mockUserUseCase.EXPECT().UserUpdate(gomock.Any()).Return(nil)
	mockUserUseCase.EXPECT().UserReadById(gomock.Any()).Return(fakeUser, nil)

	ctx := context.WithValue(req.Context(), "userId", 1)
	req = req.WithContext(ctx)
	// Вызываем обработчик
	rr := httptest.NewRecorder()
	userRouter.UserUpdateHandler(rr, req)

	// Проверяем статус HTTP-ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Дополнительные проверки, если необходимо
}
