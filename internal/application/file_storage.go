package application

import "github.com/theobitoproject/gedo_mazo/internal/domain"

// FileStorage defines a handler to access and manage files
type FileStorage interface {
	// CloneDocument clones a document into a specific folder
	CloneDocument(
		*domain.Document,
		*domain.Folder,
	) (*domain.Document, error)
	// MergeDataIntoDocument merges the passed data into the document
	MergeDataIntoDocument(
		*domain.Document,
		*domain.MergingData,
	) error
}
