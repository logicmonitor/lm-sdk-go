// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Admin admin
//
// swagger:model Admin
type Admin struct {

	// Whether or not the user is required to accept the EULA (end user license agreement)
	// Example: true
	AcceptEULA bool `json:"acceptEULA,omitempty"`

	// The time, in epoch format, that the user accepted the EULA (if required to)
	// Read Only: true
	AcceptEULAOn int64 `json:"acceptEULAOn,omitempty"`

	// Any API Tokens associated with the user
	APITokens []*APIToken `json:"apiTokens,omitempty"`

	// Whether it is a API only user
	// Example: false
	Apionly bool `json:"apionly,omitempty"`

	// email | smsemail
	// Example: email
	ContactMethod string `json:"contactMethod,omitempty"`

	// Who created the user. This may be another user, SAML or LogicMonitor
	// Example: Chief Admin
	CreatedBy string `json:"createdBy,omitempty"`

	// The email address associated with the user
	// Example: john.doe@logicmonitor.com
	// Required: true
	Email *string `json:"email"`

	// The first name associated with the user
	// Example: John
	FirstName string `json:"firstName,omitempty"`

	// Whether or not the user should be forced to change their password on the next login
	// Example: true
	ForcePasswordChange bool `json:"forcePasswordChange,omitempty"`

	// The Id of the user
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The last action taken by the user
	// Read Only: true
	LastAction string `json:"lastAction,omitempty"`

	// The time, in epoch format, of the user's last action
	// Read Only: true
	LastActionOn int64 `json:"lastActionOn,omitempty"`

	// The time, in local format, of the user's last action
	// Read Only: true
	LastActionOnLocal string `json:"lastActionOnLocal,omitempty"`

	// The time that the user last logged in, in epoch format
	// Read Only: true
	LastLoginOn int64 `json:"lastLoginOn,omitempty"`

	// The last name associated with the user
	// Example: Doe
	LastName string `json:"lastName,omitempty"`

	// Any notes assocaited with the user
	// Example: John Doe is an Admin on this Portal
	Note string `json:"note,omitempty"`

	// The password associated with the user
	// Example: JohnDoe1
	// Required: true
	Password *string `json:"password"`

	// The phone number associated with the user
	// Example: 8054445555
	Phone string `json:"phone,omitempty"`

	// The roles assigned to the user
	// Required: true
	// Unique: true
	Roles []*Role `json:"roles"`

	// The sms email address associated with the user
	// Example: 8054445555@logicmonitor.com
	SmsEmail string `json:"smsEmail,omitempty"`

	// sms | fullText, where sms = 160 characters and fullText= all characters
	// Example: sms
	SmsEmailFormat string `json:"smsEmailFormat,omitempty"`

	// The user's status. Should be one of active and suspended
	// Example: active
	Status string `json:"status,omitempty"`

	// The timezone of the user
	// Example: America/Los Angeles
	Timezone string `json:"timezone,omitempty"`

	// The email address for user's Training account
	// Read Only: true
	TrainingEmail string `json:"trainingEmail,omitempty"`

	// Whether or not two factor authentication is enabled for the user
	// Example: false
	TwoFAEnabled bool `json:"twoFAEnabled,omitempty"`

	// The username associated with the user
	// Example: John
	// Required: true
	Username *string `json:"username"`

	// The account tabs that will be visible to the user
	// Example: {\n\n\"Hosts\" : true,\n\"Services\" : true,\n\"Reports\" : true,\n\"Dashboards\" : true,\n\"Alerts\" : true,\n\"Settings\" : true\n}
	ViewPermission interface{} `json:"viewPermission,omitempty"`
}

// Validate validates this admin
func (m *Admin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPITokens(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Admin) validateAPITokens(formats strfmt.Registry) error {
	if swag.IsZero(m.APITokens) { // not required
		return nil
	}

	for i := 0; i < len(m.APITokens); i++ {
		if swag.IsZero(m.APITokens[i]) { // not required
			continue
		}

		if m.APITokens[i] != nil {
			if err := m.APITokens[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiTokens" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Admin) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *Admin) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *Admin) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	if err := validate.UniqueItems("roles", "body", m.Roles); err != nil {
		return err
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Admin) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this admin based on the context it is used
func (m *Admin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcceptEULAOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPITokens(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastActionOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastActionOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastLoginOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrainingEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Admin) contextValidateAcceptEULAOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "acceptEULAOn", "body", int64(m.AcceptEULAOn)); err != nil {
		return err
	}

	return nil
}

func (m *Admin) contextValidateAPITokens(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.APITokens); i++ {

		if m.APITokens[i] != nil {
			if err := m.APITokens[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiTokens" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Admin) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Admin) contextValidateLastAction(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastAction", "body", string(m.LastAction)); err != nil {
		return err
	}

	return nil
}

func (m *Admin) contextValidateLastActionOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastActionOn", "body", int64(m.LastActionOn)); err != nil {
		return err
	}

	return nil
}

func (m *Admin) contextValidateLastActionOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastActionOnLocal", "body", string(m.LastActionOnLocal)); err != nil {
		return err
	}

	return nil
}

func (m *Admin) contextValidateLastLoginOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastLoginOn", "body", int64(m.LastLoginOn)); err != nil {
		return err
	}

	return nil
}

func (m *Admin) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {
			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Admin) contextValidateTrainingEmail(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "trainingEmail", "body", string(m.TrainingEmail)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Admin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Admin) UnmarshalBinary(b []byte) error {
	var res Admin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}