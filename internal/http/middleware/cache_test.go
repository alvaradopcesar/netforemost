package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/golang/mock/gomock"

	mock "gitlab.com/prettytechnical/oryx-backend-core/pkg/cache/mock"
)

func TestWithCache(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	msg := "hello world"
	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(msg))
	})

	url := "/hello"
	exp := 50 * time.Second
	keys := []string{"/hello"}
	tests := []struct {
		name      string
		code      int
		withCache bool
		method    string
		getReq    string
		getResp   []byte
		getError  error
		getTimes  int
		setErr    error
		setTimes  int
		keysReq   string
		keysResp  []string
		keysErr   error
		keysTimes int
		delReq    string
		delErr    error
		delTimes  int
	}{
		{
			name:      "WithoutCache",
			code:      http.StatusOK,
			withCache: false,
			method:    http.MethodGet,
			getReq:    url,
			getResp:   nil,
			getError:  nil,
			getTimes:  0,
			setErr:    nil,
			setTimes:  0,
			keysReq:   "",
			keysResp:  nil,
			keysErr:   nil,
			keysTimes: 0,
			delReq:    "",
			delErr:    nil,
			delTimes:  0,
		},
		{
			name:      "WithCache-Found",
			code:      http.StatusOK,
			withCache: true,
			method:    http.MethodGet,
			getReq:    url,
			getResp:   []byte(msg),
			getError:  nil,
			getTimes:  1,
			setErr:    nil,
			setTimes:  0,
			keysReq:   "",
			keysResp:  nil,
			keysErr:   nil,
			keysTimes: 0,
			delReq:    "",
			delErr:    nil,
			delTimes:  0,
		},
		{
			name:      "WithCache-NotFound",
			code:      http.StatusOK,
			withCache: true,
			method:    http.MethodGet,
			getReq:    url,
			getResp:   nil,
			getError:  nil,
			getTimes:  1,
			setErr:    nil,
			setTimes:  1,
			keysReq:   "",
			keysResp:  nil,
			keysErr:   nil,
			keysTimes: 0,
			delReq:    "",
			delErr:    nil,
			delTimes:  0,
		},
		{
			name:      "WithCache-Write",
			code:      http.StatusOK,
			withCache: true,
			method:    http.MethodPost,
			getReq:    "",
			getResp:   nil,
			getError:  nil,
			getTimes:  0,
			setErr:    nil,
			setTimes:  0,
			keysReq:   "*hello*",
			keysResp:  keys,
			keysErr:   nil,
			keysTimes: 1,
			delReq:    keys[0],
			delErr:    nil,
			delTimes:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, url, nil)
			if err != nil {
				t.Log(err)
				t.Fail()
			}
			if !test.withCache {
				req.Header.Set("Cache-Control", "no-cache")
			}

			resp := httptest.NewRecorder()

			mux := http.NewServeMux()

			m := mock.NewMockCache(ctrl)
			m.
				EXPECT().
				Get(test.getReq).
				Return(test.getResp, test.getError).
				Times(test.getTimes)
			m.
				EXPECT().
				Set(url, []byte(msg), exp).
				Return(test.setErr).
				Times(test.setTimes)
			m.
				EXPECT().
				Keys("*hello*").
				Return(test.keysResp, test.keysErr).
				Times(test.keysTimes)
			m.
				EXPECT().
				Del(test.delReq).
				Return(test.delErr).
				Times(test.delTimes)

			mux.Handle(url, WithCache(exp, m)(f))
			mux.ServeHTTP(resp, req)

			assert.Equal(t, test.code, resp.Code)
		})
	}
}
