package middleware

//
//func TestAuthMiddlewareWithCsrf(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockSession := mock_entities2.NewMockSessionUseCaseInterface(ctrl)
//	mockCsrf := mock_entities.NewMockCsrfUseCaseInterface(ctrl)
//
//	authMiddleware := Auth(mockSession, mockCsrf)
//
//	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		userID := r.Context().Value("userId")
//		if userID != nil {
//			w.WriteHeader(http.StatusOK)
//		} else {
//			w.WriteHeader(http.StatusUnauthorized)
//		}
//	})
//
//	server := httptest.NewServer(authMiddleware(testHandler))
//	defer server.Close()
//
//	userId := 1
//	sessionTokenValue := "testToken"
//	csrfTokenValue := "testCsrfToken"
//	newCsrfTokenValue := "newTestCsrfToken"
//
//	// Mocking session token validation
//	//mockSession.EXPECT().GetUserID(sessionTokenValue).Return(userId, nil).AnyTimes()
//
//	// Mocking CSRF token validation
//	csrfEntity := &entities.Csrf{Token: csrfTokenValue}
//	mockCsrf.EXPECT().GetByUserId(userId).Return(csrfEntity, nil).AnyTimes()
//	mockCsrf.EXPECT().CsrfRemove(csrfEntity).Return(nil).AnyTimes()
//	mockCsrf.EXPECT().CsrfCreate(userId).Return(&entities.Csrf{Token: newCsrfTokenValue}, nil).AnyTimes()
//
//	// Creating request with valid session and CSRF tokens
//	req, err := http.NewRequest("GET", server.URL, nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	req.AddCookie(&http.Cookie{Name: "session_token", Value: sessionTokenValue})
//	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: csrfTokenValue})
//
//	// Sending request
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer resp.Body.Close()
//
//	// Verifying the status code
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//
//	// Verifying the CSRF cookie
//	csrfCookieFound := false
//	for _, cookie := range resp.Cookies() {
//		if cookie.Name == "csrf_token" && cookie.Value == newCsrfTokenValue {
//			csrfCookieFound = true
//			break
//		}
//	}
//	assert.True(t, csrfCookieFound)
//}
