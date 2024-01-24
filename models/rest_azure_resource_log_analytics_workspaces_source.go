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

// RestAzureResourceLogAnalyticsWorkspacesSource rest azure resource log analytics workspaces source
//
// swagger:model RestAzureResourceLogAnalyticsWorkspacesSource
type RestAzureResourceLogAnalyticsWorkspacesSource struct {
	alertBodyTemplateField string

	alertEffectiveIvalField *int32

	alertLevelField string

	alertSubjectTemplateField string

	appliesToField string

	auditVersionField int64

	checksumField string

	clearAfterAckField bool

	descriptionField string

	filtersField []*RestEventSourceFilter

	groupField string

	idField int32

	installationMetadataField *IntegrationMetadata

	lineageIdField string

	nameField *string

	suppressDuplicatesESField bool

	tagsField string

	technologyField string

	versionField int64

	// Column Instance Name
	ColumnInstanceName string `json:"columnInstanceName,omitempty"`

	// Azure Log Analytics Workspaces Query
	Query string `json:"query,omitempty"`

	// The polling interval for the EventSource
	Schedule int32 `json:"schedule,omitempty"`
}

// AlertBodyTemplate gets the alert body template of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) AlertBodyTemplate() string {
	return m.alertBodyTemplateField
}

// SetAlertBodyTemplate sets the alert body template of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetAlertBodyTemplate(val string) {
	m.alertBodyTemplateField = val
}

// AlertEffectiveIval gets the alert effective ival of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) AlertEffectiveIval() *int32 {
	return m.alertEffectiveIvalField
}

// SetAlertEffectiveIval sets the alert effective ival of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetAlertEffectiveIval(val *int32) {
	m.alertEffectiveIvalField = val
}

// AlertLevel gets the alert level of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) AlertLevel() string {
	return m.alertLevelField
}

// SetAlertLevel sets the alert level of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetAlertLevel(val string) {
	m.alertLevelField = val
}

// AlertSubjectTemplate gets the alert subject template of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) AlertSubjectTemplate() string {
	return m.alertSubjectTemplateField
}

// SetAlertSubjectTemplate sets the alert subject template of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetAlertSubjectTemplate(val string) {
	m.alertSubjectTemplateField = val
}

// AppliesTo gets the applies to of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) AppliesTo() string {
	return m.appliesToField
}

// SetAppliesTo sets the applies to of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetAppliesTo(val string) {
	m.appliesToField = val
}

// AuditVersion gets the audit version of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) AuditVersion() int64 {
	return m.auditVersionField
}

// SetAuditVersion sets the audit version of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetAuditVersion(val int64) {
	m.auditVersionField = val
}

// Checksum gets the checksum of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Checksum() string {
	return m.checksumField
}

// SetChecksum sets the checksum of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetChecksum(val string) {
	m.checksumField = val
}

// ClearAfterAck gets the clear after ack of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) ClearAfterAck() bool {
	return m.clearAfterAckField
}

// SetClearAfterAck sets the clear after ack of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetClearAfterAck(val bool) {
	m.clearAfterAckField = val
}

// Collector gets the collector of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Collector() string {
	return "RestAzureResourceLogAnalyticsWorkspacesSource"
}

// SetCollector sets the collector of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetCollector(val string) {
}

// Description gets the description of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetDescription(val string) {
	m.descriptionField = val
}

// Filters gets the filters of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Filters() []*RestEventSourceFilter {
	return m.filtersField
}

// SetFilters sets the filters of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetFilters(val []*RestEventSourceFilter) {
	m.filtersField = val
}

// Group gets the group of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Group() string {
	return m.groupField
}

// SetGroup sets the group of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetGroup(val string) {
	m.groupField = val
}

// ID gets the id of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetID(val int32) {
	m.idField = val
}

// InstallationMetadata gets the installation metadata of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) InstallationMetadata() *IntegrationMetadata {
	return m.installationMetadataField
}

// SetInstallationMetadata sets the installation metadata of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetInstallationMetadata(val *IntegrationMetadata) {
	m.installationMetadataField = val
}

// LineageID gets the lineage Id of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) LineageID() string {
	return m.lineageIdField
}

// SetLineageID sets the lineage Id of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetLineageID(val string) {
	m.lineageIdField = val
}

// Name gets the name of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetName(val *string) {
	m.nameField = val
}

// SuppressDuplicatesES gets the suppress duplicates e s of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SuppressDuplicatesES() bool {
	return m.suppressDuplicatesESField
}

// SetSuppressDuplicatesES sets the suppress duplicates e s of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetSuppressDuplicatesES(val bool) {
	m.suppressDuplicatesESField = val
}

// Tags gets the tags of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Tags() string {
	return m.tagsField
}

// SetTags sets the tags of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetTags(val string) {
	m.tagsField = val
}

// Technology gets the technology of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Technology() string {
	return m.technologyField
}

// SetTechnology sets the technology of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetTechnology(val string) {
	m.technologyField = val
}

// Version gets the version of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Version() int64 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) SetVersion(val int64) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Column Instance Name
		ColumnInstanceName string `json:"columnInstanceName,omitempty"`

		// Azure Log Analytics Workspaces Query
		Query string `json:"query,omitempty"`

		// The polling interval for the EventSource
		Schedule int32 `json:"schedule,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		AlertBodyTemplate string `json:"alertBodyTemplate,omitempty"`

		AlertEffectiveIval *int32 `json:"alertEffectiveIval"`

		AlertLevel string `json:"alertLevel,omitempty"`

		AlertSubjectTemplate string `json:"alertSubjectTemplate,omitempty"`

		AppliesTo string `json:"appliesTo,omitempty"`

		AuditVersion int64 `json:"auditVersion,omitempty"`

		Checksum string `json:"checksum,omitempty"`

		ClearAfterAck bool `json:"clearAfterAck,omitempty"`

		Collector string `json:"collector,omitempty"`

		Description string `json:"description,omitempty"`

		Filters []*RestEventSourceFilter `json:"filters,omitempty"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		InstallationMetadata *IntegrationMetadata `json:"installationMetadata,omitempty"`

		LineageID string `json:"lineageId,omitempty"`

		Name *string `json:"name"`

		SuppressDuplicatesES bool `json:"suppressDuplicatesES,omitempty"`

		Tags string `json:"tags,omitempty"`

		Technology string `json:"technology,omitempty"`

		Version int64 `json:"version,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result RestAzureResourceLogAnalyticsWorkspacesSource

	result.alertBodyTemplateField = base.AlertBodyTemplate

	result.alertEffectiveIvalField = base.AlertEffectiveIval

	result.alertLevelField = base.AlertLevel

	result.alertSubjectTemplateField = base.AlertSubjectTemplate

	result.appliesToField = base.AppliesTo

	result.auditVersionField = base.AuditVersion

	result.checksumField = base.Checksum

	result.clearAfterAckField = base.ClearAfterAck

	if base.Collector != result.Collector() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid collector value: %q", base.Collector)
	}
	result.descriptionField = base.Description

	result.filtersField = base.Filters

	result.groupField = base.Group

	result.idField = base.ID

	result.installationMetadataField = base.InstallationMetadata

	result.lineageIdField = base.LineageID

	result.nameField = base.Name

	result.suppressDuplicatesESField = base.SuppressDuplicatesES

	result.tagsField = base.Tags

	result.technologyField = base.Technology

	result.versionField = base.Version

	result.ColumnInstanceName = data.ColumnInstanceName
	result.Query = data.Query
	result.Schedule = data.Schedule

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m RestAzureResourceLogAnalyticsWorkspacesSource) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Column Instance Name
		ColumnInstanceName string `json:"columnInstanceName,omitempty"`

		// Azure Log Analytics Workspaces Query
		Query string `json:"query,omitempty"`

		// The polling interval for the EventSource
		Schedule int32 `json:"schedule,omitempty"`
	}{

		ColumnInstanceName: m.ColumnInstanceName,

		Query: m.Query,

		Schedule: m.Schedule,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AlertBodyTemplate string `json:"alertBodyTemplate,omitempty"`

		AlertEffectiveIval *int32 `json:"alertEffectiveIval"`

		AlertLevel string `json:"alertLevel,omitempty"`

		AlertSubjectTemplate string `json:"alertSubjectTemplate,omitempty"`

		AppliesTo string `json:"appliesTo,omitempty"`

		AuditVersion int64 `json:"auditVersion,omitempty"`

		Checksum string `json:"checksum,omitempty"`

		ClearAfterAck bool `json:"clearAfterAck,omitempty"`

		Collector string `json:"collector,omitempty"`

		Description string `json:"description,omitempty"`

		Filters []*RestEventSourceFilter `json:"filters,omitempty"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		InstallationMetadata *IntegrationMetadata `json:"installationMetadata,omitempty"`

		LineageID string `json:"lineageId,omitempty"`

		Name *string `json:"name"`

		SuppressDuplicatesES bool `json:"suppressDuplicatesES,omitempty"`

		Tags string `json:"tags,omitempty"`

		Technology string `json:"technology,omitempty"`

		Version int64 `json:"version,omitempty"`
	}{

		AlertBodyTemplate: m.AlertBodyTemplate(),

		AlertEffectiveIval: m.AlertEffectiveIval(),

		AlertLevel: m.AlertLevel(),

		AlertSubjectTemplate: m.AlertSubjectTemplate(),

		AppliesTo: m.AppliesTo(),

		AuditVersion: m.AuditVersion(),

		Checksum: m.Checksum(),

		ClearAfterAck: m.ClearAfterAck(),

		Collector: m.Collector(),

		Description: m.Description(),

		Filters: m.Filters(),

		Group: m.Group(),

		ID: m.ID(),

		InstallationMetadata: m.InstallationMetadata(),

		LineageID: m.LineageID(),

		Name: m.Name(),

		SuppressDuplicatesES: m.SuppressDuplicatesES(),

		Tags: m.Tags(),

		Technology: m.Technology(),

		Version: m.Version(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this rest azure resource log analytics workspaces source
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertEffectiveIval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallationMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) validateAlertEffectiveIval(formats strfmt.Registry) error {

	if err := validate.Required("alertEffectiveIval", "body", m.AlertEffectiveIval()); err != nil {
		return err
	}

	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) validateFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.Filters()) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters()); i++ {
		if swag.IsZero(m.filtersField[i]) { // not required
			continue
		}

		if m.filtersField[i] != nil {
			if err := m.filtersField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) validateInstallationMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.InstallationMetadata()) { // not required
		return nil
	}

	if m.InstallationMetadata() != nil {
		if err := m.InstallationMetadata().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installationMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rest azure resource log analytics workspaces source based on the context it is used
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChecksum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallationMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLineageID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) contextValidateAuditVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "auditVersion", "body", int64(m.AuditVersion())); err != nil {
		return err
	}

	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) contextValidateChecksum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "checksum", "body", string(m.Checksum())); err != nil {
		return err
	}

	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters()); i++ {

		if m.filtersField[i] != nil {
			if err := m.filtersField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) contextValidateInstallationMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.InstallationMetadata() != nil {
		if err := m.InstallationMetadata().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installationMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) contextValidateLineageID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lineageId", "body", string(m.LineageID())); err != nil {
		return err
	}

	return nil
}

func (m *RestAzureResourceLogAnalyticsWorkspacesSource) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version", "body", int64(m.Version())); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestAzureResourceLogAnalyticsWorkspacesSource) UnmarshalBinary(b []byte) error {
	var res RestAzureResourceLogAnalyticsWorkspacesSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}