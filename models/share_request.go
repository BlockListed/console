// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

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

// ShareRequest share request
//
// swagger:model shareRequest
type ShareRequest struct {

	// access key
	// Required: true
	AccessKey *string `json:"access_key"`

	// expires
	Expires string `json:"expires,omitempty"`

	// prefix
	// Required: true
	Prefix *string `json:"prefix"`

	// secret key
	// Required: true
	SecretKey *string `json:"secret_key"`

	// version id
	// Required: true
	VersionID *string `json:"version_id"`
}

// Validate validates this share request
func (m *ShareRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShareRequest) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("access_key", "body", m.AccessKey); err != nil {
		return err
	}

	return nil
}

func (m *ShareRequest) validatePrefix(formats strfmt.Registry) error {

	if err := validate.Required("prefix", "body", m.Prefix); err != nil {
		return err
	}

	return nil
}

func (m *ShareRequest) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secret_key", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

func (m *ShareRequest) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("version_id", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this share request based on context it is used
func (m *ShareRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShareRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShareRequest) UnmarshalBinary(b []byte) error {
	var res ShareRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
