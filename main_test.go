package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello World", rec.Body.String())
}

func TestParamShow(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest("GET", "/Ryoukata", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello Ryoukata", rec.Body.String())
}

func TestJsonGet(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest("GET", "/json", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"foo": "bar", "hoge": "fuga"}`, rec.Body.String())
}
