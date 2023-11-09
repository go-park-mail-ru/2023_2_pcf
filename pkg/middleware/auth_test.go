package middleware

import (
	"AdHub/internal/pkg/entities/mock_entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	authMiddleware := Auth(mockSession)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userId")
		if userID != nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	server := httptest.NewServer(authMiddleware(testHandler))
	defer server.Close()

	mockSession.EXPECT().GetUserId("testToken").Return(1, nil)
	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	cookie := &http.Cookie{Name: "session_token", Value: "testToken"}
	req.AddCookie(cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	reqWithoutToken, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	respWithoutToken, err := http.DefaultClient.Do(reqWithoutToken)
	if err != nil {
		t.Fatal(err)
	}
	defer respWithoutToken.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, respWithoutToken.StatusCode)
}
