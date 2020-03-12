package services

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/nevillejrbrown/tictacgo/backend/domain"

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

func TestValidGameCreation(t *testing.T) {
	Convey("Given an HTTP POST request to /game", t, func() {
		aGame := domain.Game{-1, "TestPlayer1"}
		gameMarshalled, _ := json.Marshal(aGame)
		req := httptest.NewRequest("POST", "/game", bytes.NewReader(gameMarshalled))
		resp := httptest.NewRecorder()

		Convey("when the request is handled by the server", func() {
			NewRouter().ServeHTTP(resp, req)
			Convey("Then I expect a 201", func() {
				So(resp.Code, ShouldEqual, 201)
			})
		})
	})
}
