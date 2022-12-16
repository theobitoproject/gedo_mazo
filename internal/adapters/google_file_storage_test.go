package adapters_test

import (
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/theobitoproject/myoboku/google/messages"
	"github.com/theobitoproject/myoboku/google/mocks"

	"github.com/theobitoproject/gedo_mazo/internal/adapters"
	"github.com/theobitoproject/gedo_mazo/internal/domain"
)

var _ = Describe("GoogleFileStorage", func() {
	var (
		googleFileStorage *adapters.GoogleFileStorage
		err               error

		ctx context.Context

		mockCtrl            *gomock.Controller
		mockDriveClient     *mocks.MockDriveClient
		mockDocumentsClient *mocks.MockDocumentsClient

		baseDocument *domain.Document
		outputFolder *domain.Folder
		docToCreate  *domain.Document
		data         *domain.MergingData

		clonedDocument *domain.Document
	)

	BeforeEach(func() {
		ctx = context.Background()

		mockCtrl = gomock.NewController(GinkgoT())
		mockDriveClient = mocks.NewMockDriveClient(mockCtrl)
		mockDocumentsClient = mocks.NewMockDocumentsClient(mockCtrl)

		baseDocument, err = domain.NewDocument("some doc id", "some base name")
		Expect(err).To(BeNil())

		outputFolder, err = domain.NewFolder("some folder if")
		Expect(err).To(BeNil())

		docToCreate, err = domain.NewUnsavedDocument("doc to create")
		Expect(err).To(BeNil())

		data, err = domain.NewMergingData()
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Creating a google file storage", func() {
		JustBeforeEach(func() {
			googleFileStorage, err = adapters.NewGoogleFileStorage(
				mockDriveClient,
				mockDocumentsClient,
			)
		})

		Context("when drive client is nil", func() {
			BeforeEach(func() {
				mockDriveClient = nil
			})

			It("should return an error", func() {
				Expect(googleFileStorage).To(BeNil())
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when drive client is nil", func() {
			BeforeEach(func() {
				mockDocumentsClient = nil
			})

			It("should return an error", func() {
				Expect(googleFileStorage).To(BeNil())
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when all parameters are defined", func() {
			It("should return an error", func() {
				Expect(googleFileStorage).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Clone document", func() {
		var (
			req *messages.CopyFileRequest
			res *messages.CopyFileResponse
		)

		JustBeforeEach(func() {
			clonedDocument, err = googleFileStorage.CloneDocument(
				ctx,
				baseDocument,
				outputFolder,
				docToCreate,
			)
		})

		BeforeEach(func() {
			googleFileStorage, err = adapters.NewGoogleFileStorage(
				mockDriveClient,
				mockDocumentsClient,
			)
			Expect(err).To(BeNil())

			req = &messages.CopyFileRequest{
				BaseFileId:    baseDocument.Id(),
				OuputFolderId: outputFolder.Id(),
				FileTitle:     docToCreate.Name(),
			}

			res = &messages.CopyFileResponse{
				FileId:    "some cloned id",
				FileTitle: docToCreate.Name(),
			}
		})

		Context("when copying file fails", func() {
			BeforeEach(func() {
				mockDriveClient.
					EXPECT().
					CopyFile(ctx, req).
					Return(nil, fmt.Errorf("error copying file")).
					Times(1)
			})

			It("should throw an error", func() {
				Expect(clonedDocument).To(BeNil())
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when copying file succeeds", func() {
			BeforeEach(func() {
				mockDriveClient.
					EXPECT().
					CopyFile(ctx, req).
					Return(res, nil).
					Times(1)
			})

			It("should return the cloned document", func() {
				Expect(clonedDocument).ToNot(BeNil())
				Expect(err).To(BeNil())

				Expect(clonedDocument.Id()).To(Equal(res.FileId))
				Expect(clonedDocument.Name()).To(Equal(res.FileTitle))
			})
		})
	})

	Describe("Merge data into document", func() {
		var (
			req *messages.MergeDataRequest
		)

		JustBeforeEach(func() {
			err = googleFileStorage.MergeDataIntoDocument(
				ctx,
				clonedDocument,
				data,
			)
		})

		BeforeEach(func() {
			googleFileStorage, err = adapters.NewGoogleFileStorage(
				mockDriveClient,
				mockDocumentsClient,
			)
			Expect(err).To(BeNil())

			clonedDocument, err = domain.NewDocument(
				"some cloned id",
				"some cloned name",
			)
			Expect(err).To(BeNil())

			data.SetKey("some key", "some value")

			req = &messages.MergeDataRequest{
				DocumentId: clonedDocument.Id(),
				Data:       *data,
			}
		})

		Context("when merging data fails", func() {
			BeforeEach(func() {
				mockDocumentsClient.
					EXPECT().
					MergeData(ctx, req).
					Return(nil, fmt.Errorf("error merging data")).
					Times(1)
			})

			It("should throw an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when merging data succeeds", func() {
			BeforeEach(func() {
				mockDocumentsClient.
					EXPECT().
					MergeData(ctx, req).
					Return(nil, nil).
					Times(1)
			})

			It("should NOT throw an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
