package domain

import "fmt"

// Document defines an actual document that
// contains data, can be managed and modified and
// is stored in some location or storage
type Document struct {
	id string
}

// NewDocument creates a new instance of Document
func NewDocument(id string) (*Document, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	return &Document{id}, nil
}

// Id returns the id of the document
func (d *Document) Id() string {
	return d.id
}
