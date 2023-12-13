package main

import (
	"final-project/data"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var pageTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHTML       string
}{
	{name: "home", url: "/", expectedStatusCode: http.StatusOK, handler: testApp.HomePage},
	{name: "login page", url: "/login", expectedStatusCode: http.StatusSeeOther, handler: testApp.LoginPage, expectedHTML: `<h1 class="mt-5">Login</h1>`},
	{name: "logout", url: "/logout", expectedStatusCode: http.StatusOK, handler: testApp.LoginPage, sessionData: map[string]any{
		"userID": 1,
		"user":   data.User{},
	}},
}

func TestConfig_Pages(t *testing.T) {
	pathToTemplates = "./templates"
	for _, e := range pageTests {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", e.url, nil)
		ctx := getCtx(req)
		req = req.WithContext(ctx)

		if len(e.sessionData) > 0 {
			for key, value := range e.sessionData {
				testApp.Session.Put(ctx, key, value)
			}
		}
		e.handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("%s failed: expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
		if len(e.expectedHTML) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, e.expectedHTML) {
				t.Errorf("%s failed: expected to find %s, but did not", e.name, e.expectedHTML)
			}
		}
	}

}

func TestConfig_Home(t *testing.T) {
	pathToTemplates = "./templates"

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	ctx := getCtx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.HomePage)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("failed; expected 200 but got", rr.Code)
	}

}
