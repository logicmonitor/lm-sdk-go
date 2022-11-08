// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NMapNetscan n map netscan
//
// swagger:model NMapNetscan
type NMapNetscan struct {
	collectorField int32

	collectorDescriptionField string

	collectorGroupField int32

	collectorGroupNameField string

	creatorField string

	descriptionField string

	duplicateField *ExcludeDuplicateIps

	groupField string

	idField int32

	ignoreSystemIPsDuplicatesField *bool

	nameField *string

	nextStartField string

	nextStartEpochField int64

	nsgIdField int32

	scheduleField *RestSchedule

	versionField int32

	// The credentials to be used for the scan
	// Example: 2
	Credentials *RestNMapNetscanPolicyCredential `json:"credentials,omitempty"`

	// Information related to including / excluding discovered devices in / from monitoring
	Ddr *NMapDDR `json:"ddr,omitempty"`

	// The subnet to exclude from scanning from nmap scans
	// Example: 10.35.41.1-10.35.41.254
	Exclude string `json:"exclude,omitempty"`

	// Include Network & Broadcast Address for CIDR based netscan
	// Example: false
	// Required: true
	IncludeNetworkAndBroadcast *bool `json:"includeNetworkAndBroadcast"`

	// The ports that should be used in the Netscan
	Ports *RestNetscanPorts `json:"ports,omitempty"`

	// The subnet to scan for nmap scans
	// Example: 10.35.41.1-10.35.41.254
	// Required: true
	Subnet *string `json:"subnet"`
}

// Collector gets the collector of this subtype
func (m *NMapNetscan) Collector() int32 {
	return m.collectorField
}

// SetCollector sets the collector of this subtype
func (m *NMapNetscan) SetCollector(val int32) {
	m.collectorField = val
}

// CollectorDescription gets the collector description of this subtype
func (m *NMapNetscan) CollectorDescription() string {
	return m.collectorDescriptionField
}

// SetCollectorDescription sets the collector description of this subtype
func (m *NMapNetscan) SetCollectorDescription(val string) {
	m.collectorDescriptionField = val
}

// CollectorGroup gets the collector group of this subtype
func (m *NMapNetscan) CollectorGroup() int32 {
	return m.collectorGroupField
}

// SetCollectorGroup sets the collector group of this subtype
func (m *NMapNetscan) SetCollectorGroup(val int32) {
	m.collectorGroupField = val
}

// CollectorGroupName gets the collector group name of this subtype
func (m *NMapNetscan) CollectorGroupName() string {
	return m.collectorGroupNameField
}

// SetCollectorGroupName sets the collector group name of this subtype
func (m *NMapNetscan) SetCollectorGroupName(val string) {
	m.collectorGroupNameField = val
}

// Creator gets the creator of this subtype
func (m *NMapNetscan) Creator() string {
	return m.creatorField
}

// SetCreator sets the creator of this subtype
func (m *NMapNetscan) SetCreator(val string) {
	m.creatorField = val
}

// Description gets the description of this subtype
func (m *NMapNetscan) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *NMapNetscan) SetDescription(val string) {
	m.descriptionField = val
}

// Duplicate gets the duplicate of this subtype
func (m *NMapNetscan) Duplicate() *ExcludeDuplicateIps {
	return m.duplicateField
}

// SetDuplicate sets the duplicate of this subtype
func (m *NMapNetscan) SetDuplicate(val *ExcludeDuplicateIps) {
	m.duplicateField = val
}

// Group gets the group of this subtype
func (m *NMapNetscan) Group() string {
	return m.groupField
}

// SetGroup sets the group of this subtype
func (m *NMapNetscan) SetGroup(val string) {
	m.groupField = val
}

// ID gets the id of this subtype
func (m *NMapNetscan) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *NMapNetscan) SetID(val int32) {
	m.idField = val
}

// IgnoreSystemIPsDuplicates gets the ignore system i ps duplicates of this subtype
func (m *NMapNetscan) IgnoreSystemIPsDuplicates() *bool {
	return m.ignoreSystemIPsDuplicatesField
}

// SetIgnoreSystemIPsDuplicates sets the ignore system i ps duplicates of this subtype
func (m *NMapNetscan) SetIgnoreSystemIPsDuplicates(val *bool) {
	m.ignoreSystemIPsDuplicatesField = val
}

// Method gets the method of this subtype
func (m *NMapNetscan) Method() string {
	return "NMapNetscan"
}

// SetMethod sets the method of this subtype
func (m *NMapNetscan) SetMethod(val string) {
}

// Name gets the name of this subtype
func (m *NMapNetscan) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *NMapNetscan) SetName(val *string) {
	m.nameField = val
}

// NextStart gets the next start of this subtype
func (m *NMapNetscan) NextStart() string {
	return m.nextStartField
}

// SetNextStart sets the next start of this subtype
func (m *NMapNetscan) SetNextStart(val string) {
	m.nextStartField = val
}

// NextStartEpoch gets the next start epoch of this subtype
func (m *NMapNetscan) NextStartEpoch() int64 {
	return m.nextStartEpochField
}

// SetNextStartEpoch sets the next start epoch of this subtype
func (m *NMapNetscan) SetNextStartEpoch(val int64) {
	m.nextStartEpochField = val
}

// NsgID gets the nsg Id of this subtype
func (m *NMapNetscan) NsgID() int32 {
	return m.nsgIdField
}

// SetNsgID sets the nsg Id of this subtype
func (m *NMapNetscan) SetNsgID(val int32) {
	m.nsgIdField = val
}

// Schedule gets the schedule of this subtype
func (m *NMapNetscan) Schedule() *RestSchedule {
	return m.scheduleField
}

// SetSchedule sets the schedule of this subtype
func (m *NMapNetscan) SetSchedule(val *RestSchedule) {
	m.scheduleField = val
}

// Version gets the version of this subtype
func (m *NMapNetscan) Version() int32 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *NMapNetscan) SetVersion(val int32) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *NMapNetscan) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The credentials to be used for the scan
		// Example: 2
		Credentials *RestNMapNetscanPolicyCredential `json:"credentials,omitempty"`

		// Information related to including / excluding discovered devices in / from monitoring
		Ddr *NMapDDR `json:"ddr,omitempty"`

		// The subnet to exclude from scanning from nmap scans
		// Example: 10.35.41.1-10.35.41.254
		Exclude string `json:"exclude,omitempty"`

		// Include Network & Broadcast Address for CIDR based netscan
		// Example: false
		// Required: true
		IncludeNetworkAndBroadcast *bool `json:"includeNetworkAndBroadcast"`

		// The ports that should be used in the Netscan
		Ports *RestNetscanPorts `json:"ports,omitempty"`

		// The subnet to scan for nmap scans
		// Example: 10.35.41.1-10.35.41.254
		// Required: true
		Subnet *string `json:"subnet"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Collector int32 `json:"collector,omitempty"`

		CollectorDescription string `json:"collectorDescription,omitempty"`

		CollectorGroup int32 `json:"collectorGroup,omitempty"`

		CollectorGroupName string `json:"collectorGroupName,omitempty"`

		Creator string `json:"creator,omitempty"`

		Description string `json:"description,omitempty"`

		Duplicate *ExcludeDuplicateIps `json:"duplicate"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		IgnoreSystemIPsDuplicates *bool `json:"ignoreSystemIPsDuplicates,omitempty"`

		Method string `json:"method"`

		Name *string `json:"name"`

		NextStart string `json:"nextStart,omitempty"`

		NextStartEpoch int64 `json:"nextStartEpoch,omitempty"`

		NsgID int32 `json:"nsgId,omitempty"`

		Schedule *RestSchedule `json:"schedule"`

		Version int32 `json:"version,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result NMapNetscan

	result.collectorField = base.Collector

	result.collectorDescriptionField = base.CollectorDescription

	result.collectorGroupField = base.CollectorGroup

	result.collectorGroupNameField = base.CollectorGroupName

	result.creatorField = base.Creator

	result.descriptionField = base.Description

	result.duplicateField = base.Duplicate

	result.groupField = base.Group

	result.idField = base.ID

	result.ignoreSystemIPsDuplicatesField = base.IgnoreSystemIPsDuplicates

	if base.Method != result.Method() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid method value: %q", base.Method)
	}
	result.nameField = base.Name

	result.nextStartField = base.NextStart

	result.nextStartEpochField = base.NextStartEpoch

	result.nsgIdField = base.NsgID

	result.scheduleField = base.Schedule

	result.versionField = base.Version

	result.Credentials = data.Credentials
	result.Ddr = data.Ddr
	result.Exclude = data.Exclude
	result.IncludeNetworkAndBroadcast = data.IncludeNetworkAndBroadcast
	result.Ports = data.Ports
	result.Subnet = data.Subnet

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m NMapNetscan) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The credentials to be used for the scan
		// Example: 2
		Credentials *RestNMapNetscanPolicyCredential `json:"credentials,omitempty"`

		// Information related to including / excluding discovered devices in / from monitoring
		Ddr *NMapDDR `json:"ddr,omitempty"`

		// The subnet to exclude from scanning from nmap scans
		// Example: 10.35.41.1-10.35.41.254
		Exclude string `json:"exclude,omitempty"`

		// Include Network & Broadcast Address for CIDR based netscan
		// Example: false
		// Required: true
		IncludeNetworkAndBroadcast *bool `json:"includeNetworkAndBroadcast"`

		// The ports that should be used in the Netscan
		Ports *RestNetscanPorts `json:"ports,omitempty"`

		// The subnet to scan for nmap scans
		// Example: 10.35.41.1-10.35.41.254
		// Required: true
		Subnet *string `json:"subnet"`
	}{

		Credentials: m.Credentials,

		Ddr: m.Ddr,

		Exclude: m.Exclude,

		IncludeNetworkAndBroadcast: m.IncludeNetworkAndBroadcast,

		Ports: m.Ports,

		Subnet: m.Subnet,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Collector int32 `json:"collector,omitempty"`

		CollectorDescription string `json:"collectorDescription,omitempty"`

		CollectorGroup int32 `json:"collectorGroup,omitempty"`

		CollectorGroupName string `json:"collectorGroupName,omitempty"`

		Creator string `json:"creator,omitempty"`

		Description string `json:"description,omitempty"`

		Duplicate *ExcludeDuplicateIps `json:"duplicate"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		IgnoreSystemIPsDuplicates *bool `json:"ignoreSystemIPsDuplicates,omitempty"`

		Method string `json:"method"`

		Name *string `json:"name"`

		NextStart string `json:"nextStart,omitempty"`

		NextStartEpoch int64 `json:"nextStartEpoch,omitempty"`

		NsgID int32 `json:"nsgId,omitempty"`

		Schedule *RestSchedule `json:"schedule"`

		Version int32 `json:"version,omitempty"`
	}{

		Collector: m.Collector(),

		CollectorDescription: m.CollectorDescription(),

		CollectorGroup: m.CollectorGroup(),

		CollectorGroupName: m.CollectorGroupName(),

		Creator: m.Creator(),

		Description: m.Description(),

		Duplicate: m.Duplicate(),

		Group: m.Group(),

		ID: m.ID(),

		IgnoreSystemIPsDuplicates: m.IgnoreSystemIPsDuplicates(),

		Method: m.Method(),

		Name: m.Name(),

		NextStart: m.NextStart(),

		NextStartEpoch: m.NextStartEpoch(),

		NsgID: m.NsgID(),

		Schedule: m.Schedule(),

		Version: m.Version(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this n map netscan
func (m *NMapNetscan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuplicate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDdr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeNetworkAndBroadcast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NMapNetscan) validateDuplicate(formats strfmt.Registry) error {

	if err := validate.Required("duplicate", "body", m.Duplicate()); err != nil {
		return err
	}

	if m.Duplicate() != nil {
		if err := m.Duplicate().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duplicate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("duplicate")
			}
			return err
		}
	}

	return nil
}

func (m *NMapNetscan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule()); err != nil {
		return err
	}

	if m.Schedule() != nil {
		if err := m.Schedule().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *NMapNetscan) validateCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *NMapNetscan) validateDdr(formats strfmt.Registry) error {

	if swag.IsZero(m.Ddr) { // not required
		return nil
	}

	if m.Ddr != nil {
		if err := m.Ddr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ddr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ddr")
			}
			return err
		}
	}

	return nil
}

func (m *NMapNetscan) validateIncludeNetworkAndBroadcast(formats strfmt.Registry) error {

	if err := validate.Required("includeNetworkAndBroadcast", "body", m.IncludeNetworkAndBroadcast); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	if m.Ports != nil {
		if err := m.Ports.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ports")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ports")
			}
			return err
		}
	}

	return nil
}

func (m *NMapNetscan) validateSubnet(formats strfmt.Registry) error {

	if err := validate.Required("subnet", "body", m.Subnet); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this n map netscan based on the context it is used
func (m *NMapNetscan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectorDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorGroupName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDuplicate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIgnoreSystemIPsDuplicates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextStart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextStartEpoch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDdr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NMapNetscan) contextValidateCollectorDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorDescription", "body", string(m.CollectorDescription())); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) contextValidateCollectorGroup(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorGroup", "body", int32(m.CollectorGroup())); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) contextValidateCollectorGroupName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorGroupName", "body", string(m.CollectorGroupName())); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) contextValidateCreator(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "creator", "body", string(m.Creator())); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) contextValidateDuplicate(ctx context.Context, formats strfmt.Registry) error {

	if m.Duplicate() != nil {
		if err := m.Duplicate().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duplicate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("duplicate")
			}
			return err
		}
	}

	return nil
}

func (m *NMapNetscan) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) contextValidateIgnoreSystemIPsDuplicates(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ignoreSystemIPsDuplicates", "body", m.IgnoreSystemIPsDuplicates()); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) contextValidateNextStart(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nextStart", "body", string(m.NextStart())); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) contextValidateNextStartEpoch(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nextStartEpoch", "body", int64(m.NextStartEpoch())); err != nil {
		return err
	}

	return nil
}

func (m *NMapNetscan) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule() != nil {
		if err := m.Schedule().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *NMapNetscan) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *NMapNetscan) contextValidateDdr(ctx context.Context, formats strfmt.Registry) error {

	if m.Ddr != nil {
		if err := m.Ddr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ddr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ddr")
			}
			return err
		}
	}

	return nil
}

func (m *NMapNetscan) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	if m.Ports != nil {
		if err := m.Ports.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ports")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ports")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NMapNetscan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NMapNetscan) UnmarshalBinary(b []byte) error {
	var res NMapNetscan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
