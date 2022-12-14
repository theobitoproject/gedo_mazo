package application

import (
	"context"
	"fmt"

	"github.com/theobitoproject/gedo_mazo/internal/domain"
	"github.com/theobitoproject/gedo_mazo/tools"
)

// Application encapsulates all available actions
// that the service offers
type Application struct {
	fileStorage FileStorage
}

// NewApplication creates a new instance of Application
func NewApplication(
	fileStorage FileStorage,
) (*Application, error) {
	if tools.IsNil(fileStorage) {
		return nil, fmt.Errorf("file storage cannot be nil")
	}

	return &Application{
		fileStorage,
	}, nil
}

// GenerateDocumentFromTemplate creates a new document
// based on a template document
// inside the specified output folder
// and the it merges data into the new document
func (app *Application) GenerateDocumentFromTemplate(
	ctx context.Context,
	templateDoc *domain.Document,
	outputFolder *domain.Folder,
	data *domain.MergingData,
) error {
	clonedDocument, err := app.fileStorage.CloneDocument(
		ctx,
		templateDoc,
		outputFolder,
	)
	if err != nil {
		return err
	}

	return app.fileStorage.MergeDataIntoDocument(ctx, clonedDocument, data)
}
