package http

//func TestHandler_HealthCheckHandler(t *testing.T) {
//	ctrl := gomock.NewController(t)
//
//	defer ctrl.Finish()
//
//	errNotFound := errors.New("not found")
//	resp := &status.Status{Db: true, Cache: true}
//	tests := []struct {
//		name       string
//		err        error
//		statusCode int
//		resp       *status.Status
//	}{
//		{
//			name:       "Success",
//			err:        nil,
//			statusCode: http.StatusOK,
//			resp:       resp,
//		},
//		{
//			name:       "Failure",
//			err:        errNotFound,
//			statusCode: http.StatusInternalServerError,
//			resp:       nil,
//		},
//	}
//
//	l := logger.NewMock()
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			m := mock.NewMockService(ctrl)
//
//			m.
//				EXPECT().
//				CheckStatus(gomock.Any()).
//				Return(test.resp, test.err).
//				Times(1)
//
//			h := Handler{
//				log:  l,
//				serv: m,
//			}
//
//			w := httptest.NewRecorder()
//			r, err := http.NewRequest(http.MethodGet, "/health", nil)
//
//			if err != nil {
//				assert.NoError(t, err)
//				t.Fail()
//			}
//			mux := chi.NewRouter()
//			mux.Get("/health", h.HealthCheckHandler)
//
//			mux.ServeHTTP(w, r)
//			assert.Equal(t, test.statusCode, w.Code)
//		})
//	}
//}
//
//func TestHandler_SayHelloHandler(t *testing.T) {
//	resp := &status.Message{Body: "Hello From the Server!"}
//	tests := []struct {
//		name       string
//		statusCode int
//		resp       *status.Message
//	}{
//		{
//			name:       "Success",
//			statusCode: http.StatusOK,
//			resp:       resp,
//		},
//	}
//
//	l := logger.NewMock()
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			h := Handler{
//				log: l,
//			}
//
//			w := httptest.NewRecorder()
//			r, err := http.NewRequest(http.MethodGet, "/hello", nil)
//
//			if err != nil {
//				assert.NoError(t, err)
//				t.Fail()
//			}
//			mux := chi.NewRouter()
//			mux.Get("/hello", h.SayHelloHandler)
//
//			mux.ServeHTTP(w, r)
//			assert.Equal(t, test.statusCode, w.Code)
//
//			msg := status.Message{}
//			err = json.NewDecoder(w.Body).Decode(&msg)
//			if err != nil {
//				assert.NoError(t, err)
//				t.Fail()
//			}
//
//			assert.Equal(t, test.resp.Body, msg.Body)
//		})
//	}
//}
