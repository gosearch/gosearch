package mock

import "github.com/blevesearch/bleve"

// MockIndexService is a default mocked IndexService.
type MockIndexService struct {
	CreateFunc    func(path string) (bleve.Index, error)
	CreateInvoked bool
}

func (s *MockIndexService) Create(path string) (bleve.Index, error) {
	s.CreateInvoked = true
	return s.CreateFunc(path)
}
