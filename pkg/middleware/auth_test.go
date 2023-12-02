package middleware

import (
	mock_entities2 "AdHub/auth/pkg/entities/mock_entities"
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"AdHub/proto/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessionClient := mock_entities2.NewMockSessionUseCaseInterface(ctrl)
	mockCsrfUseCase := mock_entities.NewMockCsrfUseCaseInterface(ctrl)

	authMiddleware := Auth(mockSessionClient, mockCsrfUseCase)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test handler logic
		w.WriteHeader(http.StatusOK)
	})

	// Set up test server
	server := httptest.NewServer(authMiddleware(testHandler))
	defer server.Close()

	// Define test cases
	testCases := []struct {
		name           string
		path           string
		setupMocks     func()
		expectedStatus int
	}{
		{
			name: "Valid Session and CSRF Token",
			path: "/api/v1/test",
			setupMocks: func() {
				mockSessionClient.EXPECT().GetUserId(gomock.Any(), gomock.Any()).Return(&api.GetResponse{Id: 1}, nil)
				mockCsrfUseCase.EXPECT().GetByUserId(1).Return(&entities.Csrf{Token: "validCsrfToken"}, nil)
				mockCsrfUseCase.EXPECT().CsrfRemove(gomock.Any()).Return(nil)
				mockCsrfUseCase.EXPECT().CsrfCreate(1).Return(&entities.Csrf{Token: "newCsrfToken"}, nil)
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks()

			req, _ := http.NewRequest("GET", server.URL+tc.path, nil)
			req.AddCookie(&http.Cookie{Name: "session_token", Value: "validSessionToken"})
			req.AddCookie(&http.Cookie{Name: "csrf_token", Value: "validCsrfToken"})

			resp, _ := http.DefaultClient.Do(req)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode)
		})
	}
}
