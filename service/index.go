package service

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"

	"fmt"
	"path/filepath"
)

// IndexService specifies an API to interact with indexes.
type IndexService interface {
	Create(indexName string, id string, data interface{}) (bleve.Index, error)
	Get(indexName string, id string) (*document.Document, error)
}

// DefaultIndexService is a default implementation of IndexService using bleve.
type DefaultIndexService struct{}

// Create creates an index.
func (*DefaultIndexService) Create(indexName string, id string, data interface{}) (bleve.Index, error) {
	//TODO: Don't open and close the index for every request
	index, err := openOrCreateIndex(indexName)
	if err != nil {
		return nil, err
	}
	defer index.Close()

	if err := index.Index(id, data); err != nil {
		return nil, err
	}

	return index, nil
}

// Get returns a document with id `id`
func (*DefaultIndexService) Get(indexName string, id string) (*document.Document, error) {
	//TODO: Don't open and close the index for every request
	index, err := openOrCreateIndex(indexName)
	if err != nil {
		return nil, err
	}
	defer index.Close()

	document, err := index.Document(id)
	if err != nil {
		return nil, err
	}
	//TODO: return the content of the document instead of the bleve.Document object
	return document, nil

}

func openOrCreateIndex(indexName string) (bleve.Index, error) {
	path := filepath.Join(".db", indexName)
	index, err := bleve.Open(path)
	if err == bleve.ErrorIndexPathDoesNotExist {
		fmt.Println("Creating new index " + indexName)
		mapping := bleve.NewIndexMapping()
		index, err := bleve.New(path, mapping)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return index, nil
	} else if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return index, nil
}
