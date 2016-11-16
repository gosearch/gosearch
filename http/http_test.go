package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"

	"github.com/blevesearch/bleve"
	"github.com/gosearch/gosearch/service/mock"
	"github.com/hooklift/assert"
)

func TestCreateIndex(t *testing.T) {
	mockService := &mock.MockIndexService{
		CreateFunc: func(name string) (bleve.Index, error) {
			return nil, nil
		},
	}
	target := "/index/parent.child"
	handler := createIndex(mockService)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, target, nil)
	handler(w, r)
	assert.Cond(t, mockService.CreateInvoked, "Create should be called.")
	assert.Equals(t, http.StatusCreated, w.Code)
}

func TestCreateIndexMissingName(t *testing.T) {
	mockService := &mock.MockIndexService{}
	// No index name specified.
	target := "/index/"
	handler := createIndex(mockService)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, target, nil)
	handler(w, r)
	assert.Cond(t, !mockService.CreateInvoked, "Create shouldn't be called.")
	assert.Equals(t, http.StatusBadRequest, w.Code)
}

func TestCreateIndexError(t *testing.T) {
	mockService := &mock.MockIndexService{
		CreateFunc: func(name string) (bleve.Index, error) {
			return nil, errors.New("Scary error")
		},
	}
	target := "/index/some.index"
	handler := createIndex(mockService)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, target, nil)
	handler(w, r)
	assert.Cond(t, mockService.CreateInvoked, "Create should be called.")
	assert.Equals(t, http.StatusBadRequest, w.Code)
}
