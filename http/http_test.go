package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"

	"bytes"
	"github.com/blevesearch/bleve"
	"github.com/gosearch/gosearch/service/mock"
	"github.com/hooklift/assert"
)

func TestCreateIndex(t *testing.T) {
	mockService := &mock.MockIndexService{
		CreateFunc: func(indexName string, id string, data interface{}) (bleve.Index, error) {
			return nil, nil
		},
	}
	target := "/some-index/some-id"
	handler := createIndex(mockService)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, target, bytes.NewBufferString("{\"hello\":\"world\"}"))
	handler(w, r)
	assert.Cond(t, mockService.CreateInvoked, "Create should be called.")
	assert.Equals(t, http.StatusCreated, w.Code)
}

func TestCreateIndexMissingName(t *testing.T) {
	mockService := &mock.MockIndexService{}
	// No index name specified.
	target := "/"
	handler := createIndex(mockService)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, target, nil)
	handler(w, r)
	assert.Cond(t, !mockService.CreateInvoked, "Create shouldn't be called.")
	assert.Equals(t, http.StatusUnprocessableEntity, w.Code)
}

func TestCreateIndexError(t *testing.T) {
	mockService := &mock.MockIndexService{
		CreateFunc: func(indexName string, id string, data interface{}) (bleve.Index, error) {
			return nil, errors.New("Scary error")
		},
	}
	target := "/some-index/some-id"
	handler := createIndex(mockService)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, target, bytes.NewBufferString("{\"hello\":\"world\"}"))
	handler(w, r)
	assert.Cond(t, mockService.CreateInvoked, "Create should be called.")
	assert.Equals(t, http.StatusUnprocessableEntity, w.Code)
}
