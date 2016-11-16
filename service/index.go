package service

import "github.com/blevesearch/bleve"

// IndexService specifies an API to interact with indexes.
type IndexService interface {
	Create(indexName string, id string, data interface{}) (bleve.Index, error)
}

// DefaultIndexService is a default implementation of IndexService using bleve.
type DefaultIndexService struct{}

// Create creates an index.
func (*DefaultIndexService) Create(indexName string, id string, data interface{}) (bleve.Index, error) {
	mapping := bleve.NewIndexMapping()

	index, err := bleve.New(".db/"+indexName, mapping)
	if err != nil {
		return nil, err
	}

	if err := index.Index(id, data); err != nil {
		return nil, err
	}

	return index, nil
}
