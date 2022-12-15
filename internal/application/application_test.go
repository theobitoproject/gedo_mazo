package application_test

import (
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/theobitoproject/gedo_mazo/internal/application"
	"github.com/theobitoproject/gedo_mazo/internal/application/mocks"
	"github.com/theobitoproject/gedo_mazo/internal/domain"
)

var _ = Describe("Application", func() {
	var (
		app *application.Application
		err error

		ctx context.Context

		mockCtrl        *gomock.Controller
		mockFileStorage *mocks.MockFileStorage

		templateDoc  *domain.Document
		outputFolder *domain.Folder
		docToCreate  *domain.Document
		data         *domain.MergingData

		clonedDocument *domain.Document
	)

	BeforeEach(func() {
		ctx = context.Background()

		mockCtrl = gomock.NewController(GinkgoT())
		mockFileStorage = mocks.NewMockFileStorage(mockCtrl)

		templateDoc, err = domain.NewDocument("some doc id", "some template name")
		Expect(err).To(BeNil())

		outputFolder, err = domain.NewFolder("some folder id")
		Expect(err).To(BeNil())

		docToCreate, err = domain.NewUnsavedDocument("doc to create")
		Expect(err).To(BeNil())

		data, err = domain.NewMergingData()
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Creating an application", func() {
		JustBeforeEach(func() {
			app, err = application.NewApplication(mockFileStorage)
		})

		Context("when file storage is nil", func() {
			BeforeEach(func() {
				mockFileStorage = nil
			})

			It("should return an error", func() {
				Expect(app).To(BeNil())
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when all parameters are defined", func() {
			It("should return an error", func() {
				Expect(app).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Generate document from template", func() {
		JustBeforeEach(func() {
			err = app.GenerateDocumentFromTemplate(
				ctx,
				templateDoc,
				outputFolder,
				docToCreate,
				data,
			)
		})

		BeforeEach(func() {
			app, err = application.NewApplication(mockFileStorage)
			Expect(err).To(BeNil())
		})

		Context("when doc to create is already saved", func() {
			BeforeEach(func() {
				docToCreate, err = domain.NewDocument(
					"id already in place",
					"doc to create",
				)
				Expect(err).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when doc to create is NOT saved yet", func() {
			Context("when cloning document fails", func() {
				BeforeEach(func() {
					mockFileStorage.
						EXPECT().
						CloneDocument(
							ctx,
							templateDoc,
							outputFolder,
							docToCreate,
						).
						Return(nil, fmt.Errorf("error cloning document")).
						Times(1)
				})

				It("should return an error", func() {
					Expect(err).ToNot(BeNil())
				})
			})

			Context("when cloning document succeeds", func() {
				BeforeEach(func() {
					clonedDocument, err = domain.NewDocument(
						"some cloned doc id",
						"some cloned doc name",
					)
					Expect(err).To(BeNil())

					mockFileStorage.
						EXPECT().
						CloneDocument(
							ctx,
							templateDoc,
							outputFolder,
							docToCreate,
						).
						Return(clonedDocument, nil).
						Times(1)
				})

				Context("when merging data into document fails", func() {
					BeforeEach(func() {
						mockFileStorage.
							EXPECT().
							MergeDataIntoDocument(ctx, clonedDocument, data).
							Return(fmt.Errorf("error merging data")).
							Times(1)
					})

					It("should return an error", func() {
						Expect(err).ToNot(BeNil())
					})
				})

				Context("when merging data into document succeeds", func() {
					BeforeEach(func() {
						mockFileStorage.
							EXPECT().
							MergeDataIntoDocument(ctx, clonedDocument, data).
							Return(nil).
							Times(1)
					})

					It("should NOT return an error", func() {
						Expect(err).To(BeNil())
					})
				})
			})
		})
	})
})
