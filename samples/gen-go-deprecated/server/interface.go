package server

import (
	"context"

	"github.com/Clever/wag/samples/gen-go-deprecated/models"
)

//go:generate $GOPATH/bin/mockgen -source=$GOFILE -destination=mock_controller.go -package=server

// Controller defines the interface for the swagger-test service.
type Controller interface {

	// GetBook makes a GET request to /books/{id}
	//
	// 200: nil
	// 400: *models.BadRequest
	// 404: *models.NotFound
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	GetBook(ctx context.Context, i *models.GetBookInput) error
}
