// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EnhancedScriptNetscan enhanced script netscan
//
// swagger:model EnhancedScriptNetscan
type EnhancedScriptNetscan struct {
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

	// The ID of the default group to add discovered devices to
	// Example: 1
	DefaultGroup int32 `json:"defaultGroup,omitempty"`

	// The full path of the default group to add discovered devices to
	// Example: lm/prod/servers
	// Read Only: true
	DefaultGroupFullPath string `json:"defaultGroupFullPath,omitempty"`

	// The filter to be applied to filter out the reported devices
	Filters []*DeviceFilter `json:"filters"`

	// For embedded script scans, the groovy script contents
	// Example: testing
	GroovyScript string `json:"groovyScript,omitempty"`

	// For embedded script scans, the groovy script parameters
	// Example: prod=true
	GroovyScriptParams string `json:"groovyScriptParams,omitempty"`

	// linux script
	LinuxScript string `json:"linuxScript,omitempty"`

	// linux script params
	LinuxScriptParams string `json:"linuxScriptParams,omitempty"`

	// The parameters for an external script
	// Example: prod=true
	ScriptParams string `json:"scriptParams,omitempty"`

	// The script path for an external script
	// Example: C://scripts
	ScriptPath string `json:"scriptPath,omitempty"`

	// For script scans, the type of script. Options are embeded and external
	// Example: embeded
	// Required: true
	ScriptType *string `json:"scriptType"`

	// windows script
	WindowsScript string `json:"windowsScript,omitempty"`

	// windows script params
	WindowsScriptParams string `json:"windowsScriptParams,omitempty"`
}

// Collector gets the collector of this subtype
func (m *EnhancedScriptNetscan) Collector() int32 {
	return m.collectorField
}

// SetCollector sets the collector of this subtype
func (m *EnhancedScriptNetscan) SetCollector(val int32) {
	m.collectorField = val
}

// CollectorDescription gets the collector description of this subtype
func (m *EnhancedScriptNetscan) CollectorDescription() string {
	return m.collectorDescriptionField
}

// SetCollectorDescription sets the collector description of this subtype
func (m *EnhancedScriptNetscan) SetCollectorDescription(val string) {
	m.collectorDescriptionField = val
}

// CollectorGroup gets the collector group of this subtype
func (m *EnhancedScriptNetscan) CollectorGroup() int32 {
	return m.collectorGroupField
}

// SetCollectorGroup sets the collector group of this subtype
func (m *EnhancedScriptNetscan) SetCollectorGroup(val int32) {
	m.collectorGroupField = val
}

// CollectorGroupName gets the collector group name of this subtype
func (m *EnhancedScriptNetscan) CollectorGroupName() string {
	return m.collectorGroupNameField
}

// SetCollectorGroupName sets the collector group name of this subtype
func (m *EnhancedScriptNetscan) SetCollectorGroupName(val string) {
	m.collectorGroupNameField = val
}

// Creator gets the creator of this subtype
func (m *EnhancedScriptNetscan) Creator() string {
	return m.creatorField
}

// SetCreator sets the creator of this subtype
func (m *EnhancedScriptNetscan) SetCreator(val string) {
	m.creatorField = val
}

// Description gets the description of this subtype
func (m *EnhancedScriptNetscan) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *EnhancedScriptNetscan) SetDescription(val string) {
	m.descriptionField = val
}

// Duplicate gets the duplicate of this subtype
func (m *EnhancedScriptNetscan) Duplicate() *ExcludeDuplicateIps {
	return m.duplicateField
}

// SetDuplicate sets the duplicate of this subtype
func (m *EnhancedScriptNetscan) SetDuplicate(val *ExcludeDuplicateIps) {
	m.duplicateField = val
}

// Group gets the group of this subtype
func (m *EnhancedScriptNetscan) Group() string {
	return m.groupField
}

// SetGroup sets the group of this subtype
func (m *EnhancedScriptNetscan) SetGroup(val string) {
	m.groupField = val
}

// ID gets the id of this subtype
func (m *EnhancedScriptNetscan) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *EnhancedScriptNetscan) SetID(val int32) {
	m.idField = val
}

// IgnoreSystemIPsDuplicates gets the ignore system i ps duplicates of this subtype
func (m *EnhancedScriptNetscan) IgnoreSystemIPsDuplicates() *bool {
	return m.ignoreSystemIPsDuplicatesField
}

// SetIgnoreSystemIPsDuplicates sets the ignore system i ps duplicates of this subtype
func (m *EnhancedScriptNetscan) SetIgnoreSystemIPsDuplicates(val *bool) {
	m.ignoreSystemIPsDuplicatesField = val
}

// Method gets the method of this subtype
func (m *EnhancedScriptNetscan) Method() string {
	return "EnhancedScriptNetscan"
}

// SetMethod sets the method of this subtype
func (m *EnhancedScriptNetscan) SetMethod(val string) {
}

// Name gets the name of this subtype
func (m *EnhancedScriptNetscan) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *EnhancedScriptNetscan) SetName(val *string) {
	m.nameField = val
}

// NextStart gets the next start of this subtype
func (m *EnhancedScriptNetscan) NextStart() string {
	return m.nextStartField
}

// SetNextStart sets the next start of this subtype
func (m *EnhancedScriptNetscan) SetNextStart(val string) {
	m.nextStartField = val
}

// NextStartEpoch gets the next start epoch of this subtype
func (m *EnhancedScriptNetscan) NextStartEpoch() int64 {
	return m.nextStartEpochField
}

// SetNextStartEpoch sets the next start epoch of this subtype
func (m *EnhancedScriptNetscan) SetNextStartEpoch(val int64) {
	m.nextStartEpochField = val
}

// NsgID gets the nsg Id of this subtype
func (m *EnhancedScriptNetscan) NsgID() int32 {
	return m.nsgIdField
}

// SetNsgID sets the nsg Id of this subtype
func (m *EnhancedScriptNetscan) SetNsgID(val int32) {
	m.nsgIdField = val
}

// Schedule gets the schedule of this subtype
func (m *EnhancedScriptNetscan) Schedule() *RestSchedule {
	return m.scheduleField
}

// SetSchedule sets the schedule of this subtype
func (m *EnhancedScriptNetscan) SetSchedule(val *RestSchedule) {
	m.scheduleField = val
}

// Version gets the version of this subtype
func (m *EnhancedScriptNetscan) Version() int32 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *EnhancedScriptNetscan) SetVersion(val int32) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *EnhancedScriptNetscan) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The credentials to be used for the scan
		// Example: 2
		Credentials *RestNMapNetscanPolicyCredential `json:"credentials,omitempty"`

		// The ID of the default group to add discovered devices to
		// Example: 1
		DefaultGroup int32 `json:"defaultGroup,omitempty"`

		// The full path of the default group to add discovered devices to
		// Example: lm/prod/servers
		// Read Only: true
		DefaultGroupFullPath string `json:"defaultGroupFullPath,omitempty"`

		// The filter to be applied to filter out the reported devices
		Filters []*DeviceFilter `json:"filters"`

		// For embedded script scans, the groovy script contents
		// Example: testing
		GroovyScript string `json:"groovyScript,omitempty"`

		// For embedded script scans, the groovy script parameters
		// Example: prod=true
		GroovyScriptParams string `json:"groovyScriptParams,omitempty"`

		// linux script
		LinuxScript string `json:"linuxScript,omitempty"`

		// linux script params
		LinuxScriptParams string `json:"linuxScriptParams,omitempty"`

		// The parameters for an external script
		// Example: prod=true
		ScriptParams string `json:"scriptParams,omitempty"`

		// The script path for an external script
		// Example: C://scripts
		ScriptPath string `json:"scriptPath,omitempty"`

		// For script scans, the type of script. Options are embeded and external
		// Example: embeded
		// Required: true
		ScriptType *string `json:"scriptType"`

		// windows script
		WindowsScript string `json:"windowsScript,omitempty"`

		// windows script params
		WindowsScriptParams string `json:"windowsScriptParams,omitempty"`
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

	var result EnhancedScriptNetscan

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
	result.DefaultGroup = data.DefaultGroup
	result.DefaultGroupFullPath = data.DefaultGroupFullPath
	result.Filters = data.Filters
	result.GroovyScript = data.GroovyScript
	result.GroovyScriptParams = data.GroovyScriptParams
	result.LinuxScript = data.LinuxScript
	result.LinuxScriptParams = data.LinuxScriptParams
	result.ScriptParams = data.ScriptParams
	result.ScriptPath = data.ScriptPath
	result.ScriptType = data.ScriptType
	result.WindowsScript = data.WindowsScript
	result.WindowsScriptParams = data.WindowsScriptParams

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m EnhancedScriptNetscan) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The credentials to be used for the scan
		// Example: 2
		Credentials *RestNMapNetscanPolicyCredential `json:"credentials,omitempty"`

		// The ID of the default group to add discovered devices to
		// Example: 1
		DefaultGroup int32 `json:"defaultGroup,omitempty"`

		// The full path of the default group to add discovered devices to
		// Example: lm/prod/servers
		// Read Only: true
		DefaultGroupFullPath string `json:"defaultGroupFullPath,omitempty"`

		// The filter to be applied to filter out the reported devices
		Filters []*DeviceFilter `json:"filters"`

		// For embedded script scans, the groovy script contents
		// Example: testing
		GroovyScript string `json:"groovyScript,omitempty"`

		// For embedded script scans, the groovy script parameters
		// Example: prod=true
		GroovyScriptParams string `json:"groovyScriptParams,omitempty"`

		// linux script
		LinuxScript string `json:"linuxScript,omitempty"`

		// linux script params
		LinuxScriptParams string `json:"linuxScriptParams,omitempty"`

		// The parameters for an external script
		// Example: prod=true
		ScriptParams string `json:"scriptParams,omitempty"`

		// The script path for an external script
		// Example: C://scripts
		ScriptPath string `json:"scriptPath,omitempty"`

		// For script scans, the type of script. Options are embeded and external
		// Example: embeded
		// Required: true
		ScriptType *string `json:"scriptType"`

		// windows script
		WindowsScript string `json:"windowsScript,omitempty"`

		// windows script params
		WindowsScriptParams string `json:"windowsScriptParams,omitempty"`
	}{

		Credentials: m.Credentials,

		DefaultGroup: m.DefaultGroup,

		DefaultGroupFullPath: m.DefaultGroupFullPath,

		Filters: m.Filters,

		GroovyScript: m.GroovyScript,

		GroovyScriptParams: m.GroovyScriptParams,

		LinuxScript: m.LinuxScript,

		LinuxScriptParams: m.LinuxScriptParams,

		ScriptParams: m.ScriptParams,

		ScriptPath: m.ScriptPath,

		ScriptType: m.ScriptType,

		WindowsScript: m.WindowsScript,

		WindowsScriptParams: m.WindowsScriptParams,
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

// Validate validates this enhanced script netscan
func (m *EnhancedScriptNetscan) Validate(formats strfmt.Registry) error {
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

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScriptType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnhancedScriptNetscan) validateDuplicate(formats strfmt.Registry) error {

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

func (m *EnhancedScriptNetscan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) validateSchedule(formats strfmt.Registry) error {

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

func (m *EnhancedScriptNetscan) validateCredentials(formats strfmt.Registry) error {

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

func (m *EnhancedScriptNetscan) validateFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnhancedScriptNetscan) validateScriptType(formats strfmt.Registry) error {

	if err := validate.Required("scriptType", "body", m.ScriptType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this enhanced script netscan based on the context it is used
func (m *EnhancedScriptNetscan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateDefaultGroupFullPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnhancedScriptNetscan) contextValidateCollectorDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorDescription", "body", string(m.CollectorDescription())); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) contextValidateCollectorGroup(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorGroup", "body", int32(m.CollectorGroup())); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) contextValidateCollectorGroupName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorGroupName", "body", string(m.CollectorGroupName())); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) contextValidateCreator(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "creator", "body", string(m.Creator())); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) contextValidateDuplicate(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EnhancedScriptNetscan) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) contextValidateIgnoreSystemIPsDuplicates(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ignoreSystemIPsDuplicates", "body", m.IgnoreSystemIPsDuplicates()); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) contextValidateNextStart(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nextStart", "body", string(m.NextStart())); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) contextValidateNextStartEpoch(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nextStartEpoch", "body", int64(m.NextStartEpoch())); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EnhancedScriptNetscan) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EnhancedScriptNetscan) contextValidateDefaultGroupFullPath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "defaultGroupFullPath", "body", string(m.DefaultGroupFullPath)); err != nil {
		return err
	}

	return nil
}

func (m *EnhancedScriptNetscan) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {
			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnhancedScriptNetscan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnhancedScriptNetscan) UnmarshalBinary(b []byte) error {
	var res EnhancedScriptNetscan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
