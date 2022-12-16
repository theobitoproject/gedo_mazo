package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/theobitoproject/myoboku/google"
	"github.com/theobitoproject/myoboku/google/auth"

	"github.com/theobitoproject/gedo_mazo/internal/adapters"
	"github.com/theobitoproject/gedo_mazo/internal/application"
	"github.com/theobitoproject/gedo_mazo/internal/ports"
	gedomazopb "github.com/theobitoproject/gedo_mazo/internal/ports/gedo_mazo"
)

func main() {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return
	}

	authHandler, err := auth.NewFileAuthHandler("token.json", b)
	if err != nil {
		log.Fatalf("Unable to create auth handler: %v", err)
		return
	}

	driveClient, err := google.NewDefaultDriveClient(authHandler)
	if err != nil {
		log.Fatalf("Unable to create drive client: %v", err)
		return
	}

	docsClient, err := google.NewDefaultDocumentsClient(authHandler)
	if err != nil {
		log.Fatalf("Unable to create docs client: %v", err)
		return
	}

	fs, err := adapters.NewGoogleFileStorage(driveClient, docsClient)
	if err != nil {
		panic(err)
	}

	app, err := application.NewApplication(fs)
	if err != nil {
		panic(err)
	}

	server := ports.NewGrpcServer(app)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	gedomazopb.RegisterGedoMazoServer(s, &server)
	reflection.Register(s)

	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
