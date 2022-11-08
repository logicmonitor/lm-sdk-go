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

// CollectorBase collector base
//
// swagger:model CollectorBase
type CollectorBase struct {

	// The comment associated with the Collector acknowledgement (if it is in alert and acknowledged)
	// Read Only: true
	AckComment string `json:"ackComment,omitempty"`

	// Whether or not the Collector is currently acknowledged
	// Read Only: true
	Acked *bool `json:"acked,omitempty"`

	// The user that acknowledged the Collector (if it is in alert)
	// Read Only: true
	AckedBy string `json:"ackedBy,omitempty"`

	// The epoch at which the Collector was acknowledged (if it is in alert)
	// Read Only: true
	AckedOn int64 `json:"ackedOn,omitempty"`

	// The time that the Collector was acknowledged (if it is in alert), in local time format
	// Read Only: true
	AckedOnLocal string `json:"ackedOnLocal,omitempty"`

	// This is key value pairs of collector config properties
	// Read Only: true
	AgentConfFields map[string]string `json:"agentConfFields,omitempty"`

	// The collector architecture (Windows | Linux platform followed by 32 | 64 bit)
	// Read Only: true
	Arch string `json:"arch,omitempty"`

	// The details of the Collector's automatic upgrade schedule, if one exists
	AutomaticUpgradeInfo *AutomaticUpgradeInfo `json:"automaticUpgradeInfo,omitempty"`

	// The Id of the backup Collector assigned to the Collector
	// Example: 75
	BackupAgentID int32 `json:"backupAgentId,omitempty"`

	// The Collector version
	// Read Only: true
	Build string `json:"build,omitempty"`

	// Whether the collector can be downgraded to a lower version
	// Read Only: true
	CanDowngrade *bool `json:"canDowngrade,omitempty"`

	// The reason why the collector can be downgraded
	// Read Only: true
	CanDowngradeReason string `json:"canDowngradeReason,omitempty"`

	// Whether or not an alert clear notifcation has been sent for this Collector
	// Read Only: true
	ClearSent *bool `json:"clearSent,omitempty"`

	// The Collector's configuration file
	// Read Only: true
	CollectorConf string `json:"collectorConf,omitempty"`

	// The device id of the collector device
	// Read Only: true
	CollectorDeviceID int32 `json:"collectorDeviceId,omitempty"`

	// The Id of the group the Collector is in
	// Example: 10
	CollectorGroupID int32 `json:"collectorGroupId,omitempty"`

	// The name of the group the Collector is in
	// Read Only: true
	CollectorGroupName string `json:"collectorGroupName,omitempty"`

	// The size of the collector
	// Read Only: true
	CollectorSize string `json:"collectorSize,omitempty"`

	// The version of the agent.conf configuration file
	// Read Only: true
	ConfVersion string `json:"confVersion,omitempty"`

	// The time that the Collector was created, in epoch format
	// Read Only: true
	CreatedOn int64 `json:"createdOn,omitempty"`

	// The time that the Collector was created, in local time format
	// Read Only: true
	CreatedOnLocal string `json:"createdOnLocal,omitempty"`

	// The custom properties defined for the Collector
	CustomProperties []*NameAndValue `json:"customProperties"`

	// The Collector's description
	// Example: Linux Collector
	Description string `json:"description,omitempty"`

	// Whether the collector is in EA version
	// Read Only: true
	Ea *bool `json:"ea,omitempty"`

	// Whether or not automatic failback is enabled for the Collector, the default value is true
	// Example: true
	EnableFailBack bool `json:"enableFailBack,omitempty"`

	// Whether or not the device the Collector is installed on is enabled for fail over
	// Example: true
	EnableFailOverOnCollectorDevice bool `json:"enableFailOverOnCollectorDevice,omitempty"`

	// The Id of the escalation chain associated with this Collector
	// Example: 80
	EscalatingChainID int32 `json:"escalatingChainId,omitempty"`

	// Whether the collector has failover devices
	// Read Only: true
	HasFailOverDevice *bool `json:"hasFailOverDevice,omitempty"`

	// The hostname of the device the Collector is installed on
	// Read Only: true
	Hostname string `json:"hostname,omitempty"`

	// The id of the Collector
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The SDT status of the collector
	// Read Only: true
	InSDT *bool `json:"inSDT,omitempty"`

	// Whether or not the Collector is currently down
	// Read Only: true
	IsDown *bool `json:"isDown,omitempty"`

	// The time, in epoch format, that a notification was last sent for the Collector
	// Read Only: true
	LastSentNotificationOn int64 `json:"lastSentNotificationOn,omitempty"`

	// The time that a notification was last sent for this Collector, in local time format
	// Read Only: true
	LastSentNotificationOnLocal string `json:"lastSentNotificationOnLocal,omitempty"`

	// Whether to create a collector device when instance collector, the default value is true
	// Example: true
	NeedAutoCreateCollectorDevice bool `json:"needAutoCreateCollectorDevice,omitempty"`

	// The Netscan version associated with the Collector
	// Read Only: true
	NetscanVersion string `json:"netscanVersion,omitempty"`

	// The Id of the next recipient to which alert notifications will be sent
	// Read Only: true
	NextRecipient int32 `json:"nextRecipient,omitempty"`

	// The details of the Collector's next upgrade, if one has been scheduled
	// Read Only: true
	NextUpgradeInfo *NextUpgradeInfo `json:"nextUpgradeInfo,omitempty"`

	// The number of devices monitored by the Collector
	// Read Only: true
	NumberOfHosts int32 `json:"numberOfHosts,omitempty"`

	// number of instances
	NumberOfInstances int32 `json:"numberOfInstances,omitempty"`

	// The details of the Collector's automatic downgrade schedule, if one exists
	OnetimeDowngradeInfo *OnetimeUpgradeInfo `json:"onetimeDowngradeInfo,omitempty"`

	// The details of the Collector's one time upgrade, if one has been scheduled
	OnetimeUpgradeInfo *OnetimeUpgradeInfo `json:"onetimeUpgradeInfo,omitempty"`

	// The OS of the Collector device (e.g. Linux, Windows)
	// Read Only: true
	Platform string `json:"platform,omitempty"`

	// Collector configurations
	// Read Only: true
	PredefinedConfig interface{} `json:"predefinedConfig,omitempty"`

	// The previous version of the collector, used for downgrading
	// Read Only: true
	PreviousVersion string `json:"previousVersion,omitempty"`

	// The interval, in minutes, after which alert notifications for the Collector will be resent
	// Example: 15
	ResendIval int32 `json:"resendIval,omitempty"`

	// The Proxy's configuration
	// Read Only: true
	SbproxyConf string `json:"sbproxyConf,omitempty"`

	// The device group id when create a new collector device
	// Example: 0
	SpecifiedCollectorDeviceGroupID int32 `json:"specifiedCollectorDeviceGroupId,omitempty"`

	// The registration status of the Collector. Acceptable values are: 0= unregistered, 1= registered, 2= stopped, 3= suspended
	// Read Only: true
	Status int32 `json:"status,omitempty"`

	// Whether alert clear notifications are suppressed for the Collector
	// Example: true
	SuppressAlertClear bool `json:"suppressAlertClear,omitempty"`

	// Whether the collector can monitor Synthetic devices (Selenium grid property must be defined)
	// Read Only: true
	SyntheticsEnabled *bool `json:"syntheticsEnabled,omitempty"`

	// The time the Collector has been up, in seconds
	// Read Only: true
	UpTime int64 `json:"upTime,omitempty"`

	// The time that the Collector was last updated, in epoch format
	// Read Only: true
	UpdatedOn int64 `json:"updatedOn,omitempty"`

	// The time that the Collector was last updated, in local time format
	// Read Only: true
	UpdatedOnLocal string `json:"updatedOnLocal,omitempty"`

	// The last time the Collector was updated, in epoch format
	// Read Only: true
	UserChangeOn int64 `json:"userChangeOn,omitempty"`

	// The last time the Collector was updated, in the account time zone
	// Read Only: true
	UserChangeOnLocal string `json:"userChangeOnLocal,omitempty"`

	// The permission level of the user that made the API request to get Collector information
	// Read Only: true
	UserPermission string `json:"userPermission,omitempty"`

	// The number of devices monitored by the Collector and visible to the user that made the query
	// Read Only: true
	UserVisibleHostsNum int32 `json:"userVisibleHostsNum,omitempty"`

	// The Watchdog's configuration
	// Read Only: true
	WatchdogConf string `json:"watchdogConf,omitempty"`

	// The time that the Watchdog Services was last updated, in epoch format
	// Read Only: true
	WatchdogUpdatedOn int64 `json:"watchdogUpdatedOn,omitempty"`

	// The time that the Collector Watchdog was last updated, in local time format
	// Read Only: true
	WatchdogUpdatedOnLocal string `json:"watchdogUpdatedOnLocal,omitempty"`

	// The java service wrapper's configuration
	// Read Only: true
	WrapperConf string `json:"wrapperConf,omitempty"`
}

// Validate validates this collector base
func (m *CollectorBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutomaticUpgradeInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextUpgradeInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnetimeDowngradeInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnetimeUpgradeInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectorBase) validateAutomaticUpgradeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AutomaticUpgradeInfo) { // not required
		return nil
	}

	if m.AutomaticUpgradeInfo != nil {
		if err := m.AutomaticUpgradeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("automaticUpgradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("automaticUpgradeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CollectorBase) validateCustomProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomProperties); i++ {
		if swag.IsZero(m.CustomProperties[i]) { // not required
			continue
		}

		if m.CustomProperties[i] != nil {
			if err := m.CustomProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CollectorBase) validateNextUpgradeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.NextUpgradeInfo) { // not required
		return nil
	}

	if m.NextUpgradeInfo != nil {
		if err := m.NextUpgradeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nextUpgradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nextUpgradeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CollectorBase) validateOnetimeDowngradeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.OnetimeDowngradeInfo) { // not required
		return nil
	}

	if m.OnetimeDowngradeInfo != nil {
		if err := m.OnetimeDowngradeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onetimeDowngradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onetimeDowngradeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CollectorBase) validateOnetimeUpgradeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.OnetimeUpgradeInfo) { // not required
		return nil
	}

	if m.OnetimeUpgradeInfo != nil {
		if err := m.OnetimeUpgradeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onetimeUpgradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onetimeUpgradeInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this collector base based on the context it is used
func (m *CollectorBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAckComment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAcked(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAckedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAckedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAckedOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgentConfFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateArch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutomaticUpgradeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCanDowngrade(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCanDowngradeReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClearSent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorConf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorDeviceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorGroupName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHasFailOverDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInSDT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsDown(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastSentNotificationOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastSentNotificationOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetscanVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextRecipient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextUpgradeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumberOfHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnetimeDowngradeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnetimeUpgradeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreviousVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSbproxyConf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyntheticsEnabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserChangeOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserChangeOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserVisibleHostsNum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWatchdogConf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWatchdogUpdatedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWatchdogUpdatedOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrapperConf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectorBase) contextValidateAckComment(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ackComment", "body", string(m.AckComment)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateAcked(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "acked", "body", m.Acked); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateAckedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ackedBy", "body", string(m.AckedBy)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateAckedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ackedOn", "body", int64(m.AckedOn)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateAckedOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ackedOnLocal", "body", string(m.AckedOnLocal)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateAgentConfFields(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *CollectorBase) contextValidateArch(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "arch", "body", string(m.Arch)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateAutomaticUpgradeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.AutomaticUpgradeInfo != nil {
		if err := m.AutomaticUpgradeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("automaticUpgradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("automaticUpgradeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CollectorBase) contextValidateBuild(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "build", "body", string(m.Build)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateCanDowngrade(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "canDowngrade", "body", m.CanDowngrade); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateCanDowngradeReason(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "canDowngradeReason", "body", string(m.CanDowngradeReason)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateClearSent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "clearSent", "body", m.ClearSent); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateCollectorConf(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorConf", "body", string(m.CollectorConf)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateCollectorDeviceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorDeviceId", "body", int32(m.CollectorDeviceID)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateCollectorGroupName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorGroupName", "body", string(m.CollectorGroupName)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateCollectorSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorSize", "body", string(m.CollectorSize)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateConfVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "confVersion", "body", string(m.ConfVersion)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateCreatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdOn", "body", int64(m.CreatedOn)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateCreatedOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdOnLocal", "body", string(m.CreatedOnLocal)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateCustomProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomProperties); i++ {

		if m.CustomProperties[i] != nil {
			if err := m.CustomProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CollectorBase) contextValidateEa(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ea", "body", m.Ea); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateHasFailOverDevice(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "hasFailOverDevice", "body", m.HasFailOverDevice); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateHostname(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "hostname", "body", string(m.Hostname)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateInSDT(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "inSDT", "body", m.InSDT); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateIsDown(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "isDown", "body", m.IsDown); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateLastSentNotificationOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastSentNotificationOn", "body", int64(m.LastSentNotificationOn)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateLastSentNotificationOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastSentNotificationOnLocal", "body", string(m.LastSentNotificationOnLocal)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateNetscanVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "netscanVersion", "body", string(m.NetscanVersion)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateNextRecipient(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nextRecipient", "body", int32(m.NextRecipient)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateNextUpgradeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.NextUpgradeInfo != nil {
		if err := m.NextUpgradeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nextUpgradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nextUpgradeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CollectorBase) contextValidateNumberOfHosts(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numberOfHosts", "body", int32(m.NumberOfHosts)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateOnetimeDowngradeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.OnetimeDowngradeInfo != nil {
		if err := m.OnetimeDowngradeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onetimeDowngradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onetimeDowngradeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CollectorBase) contextValidateOnetimeUpgradeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.OnetimeUpgradeInfo != nil {
		if err := m.OnetimeUpgradeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onetimeUpgradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onetimeUpgradeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CollectorBase) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "platform", "body", string(m.Platform)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidatePreviousVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "previousVersion", "body", string(m.PreviousVersion)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateSbproxyConf(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sbproxyConf", "body", string(m.SbproxyConf)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", int32(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateSyntheticsEnabled(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "syntheticsEnabled", "body", m.SyntheticsEnabled); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateUpTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "upTime", "body", int64(m.UpTime)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateUpdatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedOn", "body", int64(m.UpdatedOn)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateUpdatedOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedOnLocal", "body", string(m.UpdatedOnLocal)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateUserChangeOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userChangeOn", "body", int64(m.UserChangeOn)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateUserChangeOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userChangeOnLocal", "body", string(m.UserChangeOnLocal)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateUserVisibleHostsNum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userVisibleHostsNum", "body", int32(m.UserVisibleHostsNum)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateWatchdogConf(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "watchdogConf", "body", string(m.WatchdogConf)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateWatchdogUpdatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "watchdogUpdatedOn", "body", int64(m.WatchdogUpdatedOn)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateWatchdogUpdatedOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "watchdogUpdatedOnLocal", "body", string(m.WatchdogUpdatedOnLocal)); err != nil {
		return err
	}

	return nil
}

func (m *CollectorBase) contextValidateWrapperConf(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wrapperConf", "body", string(m.WrapperConf)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CollectorBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectorBase) UnmarshalBinary(b []byte) error {
	var res CollectorBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
