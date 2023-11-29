package router

import (
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFileHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFile := mock_entities.NewMockFileUseCaseInterface(ctrl)

	userRouter := UserRouter{
		File: mockFile,
	}

	fakeFilename := "example.png"
	fakeFileData := []byte("fake image data")

	mockFile.EXPECT().Get(fakeFilename).Return(fakeFileData, nil)

	server := httptest.NewServer(http.HandlerFunc(userRouter.FileHandler))
	defer server.Close()

	response, err := http.Get(server.URL + "/file?file=" + fakeFilename)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", response.StatusCode)
	}

	expectedContentType := "image/png"
	actualContentType := response.Header.Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, but got %s", expectedContentType, actualContentType)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(responseBytes, fakeFileData) {
		t.Error("Response data does not match expected data")
	}
}
