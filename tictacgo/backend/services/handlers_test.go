package services

import (
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestInvalidServiceReturns404(t *testing.T) {

	convey.Convey("Given an HTTP POST request to /invalid", t, func() {
		req := httptest.NewRequest("POST", "/invalid", nil)
		resp := httptest.NewRecorder()

		convey.Convey("When the request is handled by the server", func() {
			NewRouter().ServeHTTP(resp, req)
		})
	})
}
