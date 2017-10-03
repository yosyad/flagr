// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EvalResult eval result
// swagger:model evalResult

type EvalResult struct {

	// eval context
	// Required: true
	EvalContext *EvalContext `json:"evalContext"`

	// eval debug log
	EvalDebugLog *EvalDebugLog `json:"evalDebugLog,omitempty"`

	// flag ID
	// Required: true
	FlagID *int64 `json:"flagID"`

	// segment ID
	// Required: true
	SegmentID *int64 `json:"segmentID"`

	// timestamp
	// Required: true
	// Min Length: 1
	Timestamp *string `json:"timestamp"`

	// variant ID
	// Required: true
	VariantID *int64 `json:"variantID"`
}

/* polymorph evalResult evalContext false */

/* polymorph evalResult evalDebugLog false */

/* polymorph evalResult flagID false */

/* polymorph evalResult segmentID false */

/* polymorph evalResult timestamp false */

/* polymorph evalResult variantID false */

// Validate validates this eval result
func (m *EvalResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvalContext(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEvalDebugLog(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFlagID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSegmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVariantID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvalResult) validateEvalContext(formats strfmt.Registry) error {

	if err := validate.Required("evalContext", "body", m.EvalContext); err != nil {
		return err
	}

	if m.EvalContext != nil {

		if err := m.EvalContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evalContext")
			}
			return err
		}
	}

	return nil
}

func (m *EvalResult) validateEvalDebugLog(formats strfmt.Registry) error {

	if swag.IsZero(m.EvalDebugLog) { // not required
		return nil
	}

	if m.EvalDebugLog != nil {

		if err := m.EvalDebugLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evalDebugLog")
			}
			return err
		}
	}

	return nil
}

func (m *EvalResult) validateFlagID(formats strfmt.Registry) error {

	if err := validate.Required("flagID", "body", m.FlagID); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validateSegmentID(formats strfmt.Registry) error {

	if err := validate.Required("segmentID", "body", m.SegmentID); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.MinLength("timestamp", "body", string(*m.Timestamp), 1); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validateVariantID(formats strfmt.Registry) error {

	if err := validate.Required("variantID", "body", m.VariantID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvalResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvalResult) UnmarshalBinary(b []byte) error {
	var res EvalResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
