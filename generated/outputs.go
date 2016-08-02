package generated

import "github.com/Clever/inter-service-api-testing/codegen-poc/generated/models"


// DefaultInternalError represents a generic 500 response.
type DefaultInternalError struct {
	Msg string `json:"msg"`
}

func (d DefaultInternalError) Error() string {
	return d.Msg
}

type DefaultBadRequest struct {
	Msg string `json:"msg"`
}

func (d DefaultBadRequest) Error() string {
	return d.Msg
}

type GetBookByIDOutput interface {
	GetBookByIDStatus() int
	// Data is the underlying model object. We know it is json serializable
	GetBookByIDData() interface{}
}

type GetBookByIDError interface {
	error // Extend the error interface
	GetBookByIDStatusCode() int
}

type GetBookByID200Output models.Book

func (o GetBookByID200Output) GetBookByIDData() interface{} {
	return o
}

func (o GetBookByID200Output) GetBookByIDStatus() int {
	return 200
}

type GetBookByID404Output models.Error

func (o GetBookByID404Output) GetBookByIDData() interface{} {
	return o
}

func (o GetBookByID404Output) Error() string {
	return "Status Code: " + "404"
}

func (o GetBookByID404Output) GetBookByIDStatusCode() int {
	return 404
}

type CreateBookOutput interface {
	CreateBookStatus() int
	// Data is the underlying model object. We know it is json serializable
	CreateBookData() interface{}
}

type CreateBookError interface {
	error // Extend the error interface
	CreateBookStatusCode() int
}

type CreateBook200Output models.Book

func (o CreateBook200Output) CreateBookData() interface{} {
	return o
}

func (o CreateBook200Output) CreateBookStatus() int {
	return 200
}

type GetBooksOutput interface {
	GetBooksStatus() int
	// Data is the underlying model object. We know it is json serializable
	GetBooksData() interface{}
}

type GetBooksError interface {
	error // Extend the error interface
	GetBooksStatusCode() int
}

type GetBooks200Output []models.Book

func (o GetBooks200Output) GetBooksData() interface{} {
	return o
}

func (o GetBooks200Output) GetBooksStatus() int {
	return 200
}

