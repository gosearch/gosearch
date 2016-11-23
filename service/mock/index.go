package mock

import "github.com/blevesearch/bleve"
import "github.com/blevesearch/bleve/document"

// MockIndexService is a default mocked IndexService.
type MockIndexService struct {
	CreateFunc    func(indexName string, id string, data interface{}) (bleve.Index, error)
	CreateInvoked bool
	GetFunc       func(indexName string, id string) (*document.Document, error)
	GetInvoked    bool
}

func (s *MockIndexService) Create(indexName string, id string, data interface{}) (bleve.Index, error) {
	s.CreateInvoked = true
	return s.CreateFunc(indexName, id, data)
}

func (s *MockIndexService) Get(indexName string, id string) (*document.Document, error) {
	s.GetInvoked = true
	return s.GetFunc(indexName, id)
}
