package server

import (
	"encoding/json"
	"errors"
	"github.com/Clever/wag/generated/models"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

var _ = strconv.ParseInt
var _ = strfmt.Default
var _ = swag.ConvertInt32

var controller Controller

var formats = strfmt.Default
var _ = formats

func ConvertBase64(input string) (strfmt.Base64, error) {
	temp, err := formats.Parse("byte", input)
	if err != nil {
		return strfmt.Base64{}, err
	}
	return *temp.(*strfmt.Base64), nil
}

func ConvertDateTime(input string) (strfmt.DateTime, error) {
	temp, err := formats.Parse("date-time", input)
	if err != nil {
		return strfmt.DateTime{}, err
	}
	return *temp.(*strfmt.DateTime), nil
}

func ConvertDate(input string) (strfmt.Date, error) {
	temp, err := formats.Parse("date", input)
	if err != nil {
		return strfmt.Date{}, err
	}
	return *temp.(*strfmt.Date), nil
}

func jsonMarshalNoError(i interface{}) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		// This should never happen
		return ""
	}
	return string(bytes)
}
func GetBooksHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	input, err := NewGetBooksInput(r)
	if err != nil {
		http.Error(w, jsonMarshalNoError(models.DefaultBadRequest{Msg: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()
	if err != nil {
		http.Error(w, jsonMarshalNoError(models.DefaultBadRequest{Msg: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := controller.GetBooks(ctx, input)
	if err != nil {
		if respErr, ok := err.(models.GetBooksError); ok {
			http.Error(w, respErr.Error(), respErr.GetBooksStatusCode())
			return
		} else {
			http.Error(w, jsonMarshalNoError(models.DefaultInternalError{Msg: err.Error()}), http.StatusInternalServerError)
			return
		}
	}

	respBytes, err := json.Marshal(resp.GetBooksData())
	if err != nil {
		http.Error(w, jsonMarshalNoError(models.DefaultInternalError{Msg: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}
func NewGetBooksInput(r *http.Request) (*models.GetBooksInput, error) {
	var input models.GetBooksInput

	var err error
	_ = err

	authorStr := r.URL.Query().Get("author")
	if len(authorStr) != 0 {
		var authorTmp string
		authorTmp, err = authorStr, error(nil)
		input.Author = &authorTmp

	}
	if err != nil && len(authorStr) != 0 {
		return nil, err
	}
	availableStr := r.URL.Query().Get("available")
	if len(availableStr) != 0 {
		var availableTmp bool
		availableTmp, err = strconv.ParseBool(availableStr)
		input.Available = &availableTmp

	}
	if err != nil && len(availableStr) != 0 {
		return nil, err
	}
	stateStr := r.URL.Query().Get("state")
	if len(stateStr) != 0 {
		var stateTmp string
		stateTmp, err = stateStr, error(nil)
		input.State = &stateTmp

	}
	if err != nil && len(stateStr) != 0 {
		return nil, err
	}
	publishedStr := r.URL.Query().Get("published")
	if len(publishedStr) != 0 {
		var publishedTmp strfmt.Date
		publishedTmp, err = ConvertDate(publishedStr)
		input.Published = &publishedTmp

	}
	if err != nil && len(publishedStr) != 0 {
		return nil, err
	}
	completedStr := r.URL.Query().Get("completed")
	if len(completedStr) != 0 {
		var completedTmp strfmt.DateTime
		completedTmp, err = ConvertDateTime(completedStr)
		input.Completed = &completedTmp

	}
	if err != nil && len(completedStr) != 0 {
		return nil, err
	}
	maxPagesStr := r.URL.Query().Get("maxPages")
	if len(maxPagesStr) != 0 {
		var maxPagesTmp float64
		maxPagesTmp, err = swag.ConvertFloat64(maxPagesStr)
		input.MaxPages = &maxPagesTmp

	}
	if err != nil && len(maxPagesStr) != 0 {
		return nil, err
	}
	minPagesStr := r.URL.Query().Get("minPages")
	if len(minPagesStr) != 0 {
		var minPagesTmp int32
		minPagesTmp, err = swag.ConvertInt32(minPagesStr)
		input.MinPages = &minPagesTmp

	}
	if err != nil && len(minPagesStr) != 0 {
		return nil, err
	}
	pagesToTimeStr := r.URL.Query().Get("pagesToTime")
	if len(pagesToTimeStr) != 0 {
		var pagesToTimeTmp float32
		pagesToTimeTmp, err = swag.ConvertFloat32(pagesToTimeStr)
		input.PagesToTime = &pagesToTimeTmp

	}
	if err != nil && len(pagesToTimeStr) != 0 {
		return nil, err
	}

	return &input, nil
}

func GetBookByIDHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	input, err := NewGetBookByIDInput(r)
	if err != nil {
		http.Error(w, jsonMarshalNoError(models.DefaultBadRequest{Msg: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()
	if err != nil {
		http.Error(w, jsonMarshalNoError(models.DefaultBadRequest{Msg: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := controller.GetBookByID(ctx, input)
	if err != nil {
		if respErr, ok := err.(models.GetBookByIDError); ok {
			http.Error(w, respErr.Error(), respErr.GetBookByIDStatusCode())
			return
		} else {
			http.Error(w, jsonMarshalNoError(models.DefaultInternalError{Msg: err.Error()}), http.StatusInternalServerError)
			return
		}
	}

	respBytes, err := json.Marshal(resp.GetBookByIDData())
	if err != nil {
		http.Error(w, jsonMarshalNoError(models.DefaultInternalError{Msg: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}
func NewGetBookByIDInput(r *http.Request) (*models.GetBookByIDInput, error) {
	var input models.GetBookByIDInput

	var err error
	_ = err

	bookIDStr := mux.Vars(r)["bookID"]
	if len(bookIDStr) == 0 {
		return nil, errors.New("Parameter must be specified")
	}
	if len(bookIDStr) != 0 {
		var bookIDTmp int64
		bookIDTmp, err = swag.ConvertInt64(bookIDStr)
		input.BookID = bookIDTmp

	}
	if err != nil {
		return nil, err
	}
	authorizationStr := r.Header.Get("authorization")
	if len(authorizationStr) != 0 {
		var authorizationTmp string
		authorizationTmp, err = authorizationStr, error(nil)
		input.Authorization = &authorizationTmp

	}
	if err != nil && len(authorizationStr) != 0 {
		return nil, err
	}
	randomBytesStr := r.URL.Query().Get("randomBytes")
	if len(randomBytesStr) != 0 {
		var randomBytesTmp strfmt.Base64
		randomBytesTmp, err = ConvertBase64(randomBytesStr)
		input.RandomBytes = &randomBytesTmp

	}
	if err != nil && len(randomBytesStr) != 0 {
		return nil, err
	}

	return &input, nil
}

func CreateBookHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	input, err := NewCreateBookInput(r)
	if err != nil {
		http.Error(w, jsonMarshalNoError(models.DefaultBadRequest{Msg: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()
	if err != nil {
		http.Error(w, jsonMarshalNoError(models.DefaultBadRequest{Msg: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := controller.CreateBook(ctx, input)
	if err != nil {
		if respErr, ok := err.(models.CreateBookError); ok {
			http.Error(w, respErr.Error(), respErr.CreateBookStatusCode())
			return
		} else {
			http.Error(w, jsonMarshalNoError(models.DefaultInternalError{Msg: err.Error()}), http.StatusInternalServerError)
			return
		}
	}

	respBytes, err := json.Marshal(resp.CreateBookData())
	if err != nil {
		http.Error(w, jsonMarshalNoError(models.DefaultInternalError{Msg: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}
func NewCreateBookInput(r *http.Request) (*models.CreateBookInput, error) {
	var input models.CreateBookInput

	var err error
	_ = err

	input.NewBook = &models.Book{}
	err = json.NewDecoder(r.Body).Decode(input.NewBook)
	if err != nil {
		return nil, err
	}

	return &input, nil
}
