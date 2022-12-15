package adapters

import (
	"context"
	"fmt"

	"github.com/theobitoproject/gedo_mazo/internal/domain"
	"github.com/theobitoproject/gedo_mazo/tools"
	"github.com/theobitoproject/myoboku/google"
	"github.com/theobitoproject/myoboku/google/messages"
)

// GoogleFileStorage defines a handler to access and manage files
// within Google Drive, Google Documents, Google Sheets and others
type GoogleFileStorage struct {
	driveClient google.DriveClient
	docsClient  google.DocumentsClient
}

// NewGoogleFileStorage creates an instance of GoogleFileStorage
func NewGoogleFileStorage(
	driveClient google.DriveClient,
	docsClient google.DocumentsClient,
) (*GoogleFileStorage, error) {
	if tools.IsNil(driveClient) {
		return nil, fmt.Errorf("drive client cannot be nil")
	}
	if tools.IsNil(docsClient) {
		return nil, fmt.Errorf("docs client cannot be nil")
	}

	return &GoogleFileStorage{
		driveClient,
		docsClient,
	}, nil
}

// CloneDocument clones a document into a specific folder
func (fs *GoogleFileStorage) CloneDocument(
	ctx context.Context,
	baseDocument *domain.Document,
	outputFolder *domain.Folder,
	docToCreate *domain.Document,
) (*domain.Document, error) {
	req := getCopyFileRequest(baseDocument, outputFolder, docToCreate)

	res, err := fs.driveClient.CopyFile(ctx, req)
	if err != nil {
		return nil, err
	}

	return domain.NewDocument(res.FileId, res.FileTitle)
}

// MergeDataIntoDocument merges the passed data into the document
func (fs *GoogleFileStorage) MergeDataIntoDocument(
	ctx context.Context,
	document *domain.Document,
	data *domain.MergingData,
) error {
	req := getMergeDataRequest(document, data)

	_, err := fs.docsClient.MergeData(ctx, req)

	return err
}

func getCopyFileRequest(
	baseDocument *domain.Document,
	outputFolder *domain.Folder,
	docToCreate *domain.Document,
) *messages.CopyFileRequest {
	return &messages.CopyFileRequest{
		BaseFileId:    baseDocument.Id(),
		OuputFolderId: outputFolder.Id(),
		FileTitle:     docToCreate.Name(),
	}
}

func getMergeDataRequest(
	document *domain.Document,
	data *domain.MergingData,
) *messages.MergeDataRequest {
	return &messages.MergeDataRequest{
		DocumentId: document.Id(),
		Data:       *data,
	}
}
