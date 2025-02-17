// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SubjectSet subject set
//
// swagger:model subjectSet
type SubjectSet struct {

	// Namespace of the Subject Set
	// Required: true
	Namespace *string `json:"namespace"`

	// Object of the Subject Set
	// Required: true
	Object *string `json:"object"`

	// Relation of the Subject Set
	// Required: true
	Relation *string `json:"relation"`
}

// Validate validates this subject set
func (m *SubjectSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubjectSet) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *SubjectSet) validateObject(formats strfmt.Registry) error {

	if err := validate.Required("object", "body", m.Object); err != nil {
		return err
	}

	return nil
}

func (m *SubjectSet) validateRelation(formats strfmt.Registry) error {

	if err := validate.Required("relation", "body", m.Relation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this subject set based on context it is used
func (m *SubjectSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubjectSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectSet) UnmarshalBinary(b []byte) error {
	var res SubjectSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
