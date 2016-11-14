package service

import "github.com/blevesearch/bleve"

// IndexService specifies an API to interact with indexes.
type IndexService interface {
	Create(path string) (bleve.Index, error)
}

// DefaultIndexService is a default implementation of IndexService using bleve.
type DefaultIndexService struct{}

// Create creates an index.
func (*DefaultIndexService) Create(path string) (bleve.Index, error) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(path, mapping)
	if err != nil {
		return nil, err
	}
	return index, nil
}
