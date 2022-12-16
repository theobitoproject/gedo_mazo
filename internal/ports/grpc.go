package ports

import (
	"context"

	"github.com/theobitoproject/gedo_mazo/internal/application"
	gedomazopb "github.com/theobitoproject/gedo_mazo/internal/ports/gedo_mazo"
)

// GrpcServer is the server that holds all entrypoints
// to interact with the application through gRPC
type GrpcServer struct {
	app *application.Application
}

// NewGrpcServer creates an instance of GrpcServer
func NewGrpcServer(app *application.Application) GrpcServer {
	return GrpcServer{app}
}

// Takes a template as the base to clone and create a new document
// and then modifies/replaces context inside the cloned document
func (gs *GrpcServer) GenerateDocumentFromTemplate(
	ctx context.Context,
	req *gedomazopb.GenerateDocumentFromTemplateRequest,
) (*gedomazopb.GenerateDocumentFromTemplateResponse, error) {
	return nil, nil
}
