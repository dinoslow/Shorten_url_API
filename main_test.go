package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestPostDriven(t *testing.T) {
	var tests = []struct {
		json []byte	
		code int
	} {
		{[]byte(`{"url":"http://www.google.com","expireAt":"2024-04-04T09:20:41Z"}`), 201},
		{[]byte(`{"url":"http://www.google.com","expireAt":"2020-04-04T09:20:41Z"}`), 201},
		{[]byte(``), 400},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("test")
		t.Run(testName, func (t *testing.T)  {
			request, _ := http.NewRequest("POST", "/api/v1/urls", bytes.NewBuffer(tt.json))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
		
			Post(response, request)
		
			fmt.Println(response.Body)
		
			assert.Equal(t, tt.code, response.Code, "Post url failed")
		})
	}
}

func TestGetDriven(t *testing.T) {
	var tests = []struct {
		id, url string
		code int
	} {
		{"1", "http://www.google.com", 303},
		{"2", "", 404},
		{"3", "", 404},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("test_%s", tt.id)
		t.Run(testName, func(t *testing.T) {
			request, _ := http.NewRequest("GET", `/`, nil)
			response := httptest.NewRecorder()

			vars := map[string]string {
				"url_id": tt.id,
			}
		
			request = mux.SetURLVars(request, vars)
		
			Get(response, request)

			assert.Equal(t, tt.code, response.Code, fmt.Sprintf("Redirect to %s", tt.url))

			expectedURL := tt.url
			assert.Equal(t, expectedURL, response.Result().Header.Get("Location"))
		})
	}
}