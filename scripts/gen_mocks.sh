#!/bin/bash

# app mocks
mockgen -package=mocks -source=internal/application/file_storage.go -destination=internal/application/mocks/mock_file_storage.go
