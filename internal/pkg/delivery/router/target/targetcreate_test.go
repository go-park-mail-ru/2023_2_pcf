package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateTargetHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTargetUseCase := mock_entities.NewMockTargetUseCaseInterface(ctrl)

	targetRouter := TargetRouter{
		Target: mockTargetUseCase,
	}

	// Подготовка тестовых данных
	newTarget := entities.Target{
		Name:      "Test Target",
		Owner_id:  1,
		Gender:    "any",
		Min_age:   18,
		Max_age:   35,
		Interests: []string{"music", "sports"},
		Tags:      []string{"tag1", "tag2"},
		Keys:      []string{"key1", "key2"},
		Regions:   []string{"region1", "region2"},
	}

	//todo: саша, эта структурка специально для тебя:
	//reqData := struct {
	//	Name      string `json:"name"`
	//	Gender    string `json:"gender"`
	//	MinAge    string `json:"min_age"`
	//	MaxAge    string `json:"max_age"`
	//	Interests string `json:"interests"`
	//	Tags      string `json:"tags"`
	//	Keys      string `json:"keys"`
	//	Regions   string `json:"regions"`
	//}{
	//	Name:      newTarget.Name,
	//	Gender:    newTarget.Gender,
	//	MinAge:    strconv.Itoa(newTarget.Min_age),
	//	MaxAge:    strconv.Itoa(newTarget.Max_age),
	//	Interests: strings.Join(newTarget.Interests, ", "),
	//	Tags:      strings.Join(newTarget.Tags, ", "),
	//	Keys:      strings.Join(newTarget.Keys, ", "),
	//	Regions:   strings.Join(newTarget.Regions, ", "),
	//}
	//todo: ее нужно подсунуть сюда, но оно не пускает ее:
	mockTargetUseCase.EXPECT().TargetCreate(gomock.Any()).Return(&newTarget, nil)

	requestBody, _ := json.Marshal(newTarget)
	req := httptest.NewRequest("POST", "/target/create", bytes.NewReader(requestBody))
	req = req.WithContext(context.WithValue(req.Context(), "userId", newTarget.Owner_id))
	rec := httptest.NewRecorder()

	targetRouter.CreateTargetHandler(rec, req)

	// Проверка ответа
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusCreated, rec.Code)
	}

	expectedBody, _ := json.Marshal(&newTarget)
	if rec.Body.String() != string(expectedBody) {
		t.Errorf("Response body does not match the expected value.\nExpected: %s\nActual: %s", string(expectedBody), rec.Body.String())
	}
}
