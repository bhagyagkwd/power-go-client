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

// PowerVSInstance PowerVS Instance for a specific IBM Cloud Account
//
// swagger:model PowerVSInstance
type PowerVSInstance struct {

	// capabilities
	Capabilities InstanceCapabilities `json:"capabilities,omitempty"`

	// The timestamp in which the PowerVS service instance was created
	// Example: 2022-04-04T16:20:15.581636275Z
	CreationDate string `json:"creationDate,omitempty"`

	// The PowerVS Service Instance CRN
	// Example: crn:v1:bluemix:public:power-iaas:dal12:a/2bc3df23c0d14ebe921397bd8aa2573a:3a5798f1-4d2b-4e0a-9311-9b0fd6b94698::
	// Required: true
	Crn *string `json:"crn"`

	// The PowerVS IBM Cloud URL path for UI (Tentative, still need verification that this is possible)
	// Example: https://cloud.ibm.com/services/power-iaas/crn%3Av1%3Abluemix%3Apublic%3Apower-iaas%3Adal12%3Aa%2F2bc3df23c0d14ebe921397bd8aa2573a%3A5de8348d-bc6a-466e-854f-661d1e86b230%3A%3A
	Href string `json:"href,omitempty"`

	// Location of the PowerVS Instance
	// Example: dal12
	// Required: true
	Location *string `json:"location"`

	// The PowerVS URL path to access specific service instance information
	// Example: https://us-south.power-iaas.cloud.ibm.com
	// Required: true
	LocationURL *string `json:"locationUrl"`

	// The name of the service instance (This field will be empty for old accounts as PowerVS did not previously saved the names)
	// Example: Test Name
	Name string `json:"name,omitempty"`

	// IBM Resource Group ID associated with the PowerVS Service Instance (This field will be empty for old accounts as PowerVS did not previously saved the Resource Group ID)
	// Example: 2bf1887bf5c947b1966de2bd88220489
	ResourceGroupID string `json:"resourceGroupId,omitempty"`

	// The status of the service instance (PowerVS behavior, if Service Instance exists then then status is active)
	// Example: Active
	Status string `json:"status,omitempty"`
}

// Validate validates this power v s instance
func (m *PowerVSInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerVSInstance) validateCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	if m.Capabilities != nil {
		if err := m.Capabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *PowerVSInstance) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *PowerVSInstance) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	return nil
}

func (m *PowerVSInstance) validateLocationURL(formats strfmt.Registry) error {

	if err := validate.Required("locationUrl", "body", m.LocationURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this power v s instance based on the context it is used
func (m *PowerVSInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerVSInstance) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Capabilities.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("capabilities")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("capabilities")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerVSInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerVSInstance) UnmarshalBinary(b []byte) error {
	var res PowerVSInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
