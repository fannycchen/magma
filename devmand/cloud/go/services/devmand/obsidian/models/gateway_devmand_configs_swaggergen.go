// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayDevmandConfigs Configuration for devmand
// swagger:model gateway_devmand_configs
type GatewayDevmandConfigs struct {

	// managed devices
	ManagedDevices []string `json:"managed_devices"`
}

// Validate validates this gateway devmand configs
func (m *GatewayDevmandConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagedDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayDevmandConfigs) validateManagedDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagedDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.ManagedDevices); i++ {

		if err := validate.MinLength("managed_devices"+"."+strconv.Itoa(i), "body", string(m.ManagedDevices[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayDevmandConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayDevmandConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayDevmandConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
