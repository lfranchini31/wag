package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// NotFoundError not found error
// swagger:model NotFoundError
type NotFoundError struct {

	// msg
	Msg string `json:"msg,omitempty"`
}

// Validate validates this not found error
func (m *NotFoundError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
