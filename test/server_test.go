package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Clever/inter-service-api-testing/codegen-poc/generated/client"
	"github.com/Clever/inter-service-api-testing/codegen-poc/generated/models"

	"golang.org/x/net/context"
)

func TestBasicEndToEnd(t *testing.T) {
	s := setupServer()

	c := client.New(s.URL)

	bookID := int64(123)
	bookName := "Test"

	createOutput, err := c.CreateBook(
		context.Background(), &models.CreateBookInput{NewBook: models.Book{ID: bookID, Name: bookName}})
	assert.NoError(t, err)
	createdBook, ok := createOutput.(models.CreateBook200Output)
	require.True(t, ok)
	assert.Equal(t, bookID, createdBook.ID)
	assert.Equal(t, bookName, createdBook.Name)

	booksOutput, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	assert.NoError(t, err)
	getBooks, ok := booksOutput.(models.GetBooks200Output)
	require.True(t, ok)
	require.Equal(t, 1, len(getBooks))
	assert.Equal(t, bookID, getBooks[0].ID)
	assert.Equal(t, bookName, getBooks[0].Name)

	singleBookOutput, err := c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: 123})
	assert.NoError(t, err)
	singleBook, ok := singleBookOutput.(models.GetBookByID200Output)
	require.True(t, ok)
	assert.Equal(t, bookID, singleBook.ID)
	assert.Equal(t, bookName, singleBook.Name)
}

func TestUserDefinedErrorResponse(t *testing.T) {
	// The 404 generated by the code
	s := setupServer()

	c := client.New(s.URL)

	_, err := c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: 123})
	assert.Error(t, err)
	_, ok := err.(models.GetBookByID404Output)
	assert.True(t, ok)
}

func TestValidationErrorResponse(t *testing.T) {
	// TODO: We don't have any test cases for this yet...
}

func TestClientSideError(t *testing.T) {
	c := client.New("badServer")

	_, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	assert.Error(t, err)
	_, ok := err.(models.DefaultInternalError)
	assert.True(t, ok)
}
