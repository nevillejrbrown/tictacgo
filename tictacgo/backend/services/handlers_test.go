package services

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInvalidServiceReturns404(t *testing.T) {

	Convey("Given an HTTP POST request to /invalid", t, func() {
		req := httptest.NewRequest("POST", "/invalid", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the server", func() {
			NewRouter().ServeHTTP(resp, req)
			Convey("Then I expect a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}
