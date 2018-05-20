# \DefaultApi

All URIs are relative to *https://localhost/santaba/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AckAlertById**](DefaultApi.md#AckAlertById) | **Post** /alert/alerts/{id}/ack | ack alert by id
[**AckCollectorDownAlertById**](DefaultApi.md#AckCollectorDownAlertById) | **Post** /setting/collectors/{id}/ackdown | ack collector down alert
[**Add**](DefaultApi.md#Add) | **Post** /setting/admins/apitokens | 
[**AddAdmin**](DefaultApi.md#AddAdmin) | **Post** /setting/admins | add admin
[**AddAlertNoteById**](DefaultApi.md#AddAlertNoteById) | **Post** /alert/alerts/{id}/note | add alert note
[**AddAlertRule**](DefaultApi.md#AddAlertRule) | **Post** /setting/alert/rules | add alert rule
[**AddApiTokenByAdminId**](DefaultApi.md#AddApiTokenByAdminId) | **Post** /setting/admins/{adminId}/apitokens | add apiToken by admin
[**AddCollector**](DefaultApi.md#AddCollector) | **Post** /setting/collectors | add collector
[**AddCollectorGroup**](DefaultApi.md#AddCollectorGroup) | **Post** /setting/collectors/groups | add collector group
[**AddDashboard**](DefaultApi.md#AddDashboard) | **Post** /dashboard/dashboards | add dashboard
[**AddDashboardGroup**](DefaultApi.md#AddDashboardGroup) | **Post** /dashboard/groups | add dashboard group
[**AddDevice**](DefaultApi.md#AddDevice) | **Post** /device/devices | add a new device
[**AddDeviceDatasourceInstance**](DefaultApi.md#AddDeviceDatasourceInstance) | **Post** /device/devices/{deviceId}/devicedatasources/{hdsId}/instances | add device instance 
[**AddDeviceDatasourceInstanceGroup**](DefaultApi.md#AddDeviceDatasourceInstanceGroup) | **Post** /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups | add device datasource instance group 
[**AddDeviceGroup**](DefaultApi.md#AddDeviceGroup) | **Post** /device/groups | add device group
[**AddDeviceGroupProperty**](DefaultApi.md#AddDeviceGroupProperty) | **Post** /device/groups/{gid}/properties | add device group property
[**AddDeviceProperty**](DefaultApi.md#AddDeviceProperty) | **Post** /device/devices/{deviceId}/properties | add device property
[**AddEscalationChain**](DefaultApi.md#AddEscalationChain) | **Post** /setting/alert/chains | add escalation chain
[**AddOpsNote**](DefaultApi.md#AddOpsNote) | **Post** /setting/opsnotes | add opsnote
[**AddReport**](DefaultApi.md#AddReport) | **Post** /report/reports | add report
[**AddRole**](DefaultApi.md#AddRole) | **Post** /setting/roles | add role
[**AddSDT**](DefaultApi.md#AddSDT) | **Post** /sdt/sdts | add SDT
[**AddService**](DefaultApi.md#AddService) | **Post** /service/services | add service
[**AddServiceGroup**](DefaultApi.md#AddServiceGroup) | **Post** /service/groups | add service group
[**AddWidget**](DefaultApi.md#AddWidget) | **Post** /dashboard/widgets | add widget
[**AsyncClone**](DefaultApi.md#AsyncClone) | **Post** /dashboard/groups/{id}/asyncclone | 
[**Clone**](DefaultApi.md#Clone) | **Post** /dashboard/groups/{id}/clone | 
[**CreateAutoReports**](DefaultApi.md#CreateAutoReports) | **Post** /report/reports/auto | 
[**DeleteAdminById**](DefaultApi.md#DeleteAdminById) | **Delete** /setting/admins/{id} | delete admin
[**DeleteAlertRuleById**](DefaultApi.md#DeleteAlertRuleById) | **Delete** /setting/alert/rules/{id} | delete alert rule
[**DeleteApiTokenById**](DefaultApi.md#DeleteApiTokenById) | **Delete** /setting/admins/{adminId}/apitokens/{apitokenId} | delete apiToken 
[**DeleteCollectorById**](DefaultApi.md#DeleteCollectorById) | **Delete** /setting/collectors/{id} | delete collector
[**DeleteCollectorGroupById**](DefaultApi.md#DeleteCollectorGroupById) | **Delete** /setting/collectors/groups/{id} | delete collector group
[**DeleteDashboardById**](DefaultApi.md#DeleteDashboardById) | **Delete** /dashboard/dashboards/{id} | delete dashboard
[**DeleteDashboardGroupById**](DefaultApi.md#DeleteDashboardGroupById) | **Delete** /dashboard/groups/{id} | delete dashboard group
[**DeleteDevice**](DefaultApi.md#DeleteDevice) | **Delete** /device/devices/{id} | delete a device
[**DeleteDeviceGroupById**](DefaultApi.md#DeleteDeviceGroupById) | **Delete** /device/groups/{id} | delete device group
[**DeleteDeviceGroupPropertyByName**](DefaultApi.md#DeleteDeviceGroupPropertyByName) | **Delete** /device/groups/{gid}/properties/{name} | delete device group property
[**DeleteDevicePropertyByName**](DefaultApi.md#DeleteDevicePropertyByName) | **Delete** /device/devices/{deviceId}/properties/{name} | delete device  property
[**DeleteEscalationChainById**](DefaultApi.md#DeleteEscalationChainById) | **Delete** /setting/alert/chains/{id} | delete escalation chain by id
[**DeleteOpsNoteById**](DefaultApi.md#DeleteOpsNoteById) | **Delete** /setting/opsnotes/{id} | delete opsnote by id
[**DeleteRoleById**](DefaultApi.md#DeleteRoleById) | **Delete** /setting/roles/{id} | delete role
[**DeleteSDTById**](DefaultApi.md#DeleteSDTById) | **Delete** /sdt/sdts/{id} | delete SDT by id
[**DeleteServiceById**](DefaultApi.md#DeleteServiceById) | **Delete** /service/services/{id} | delete service by id
[**DeleteServiceGroupById**](DefaultApi.md#DeleteServiceGroupById) | **Delete** /service/groups/{id} | delete service group
[**DeleteToken**](DefaultApi.md#DeleteToken) | **Delete** /dashboard/groups/token/{tokenname} | 
[**DeleteWidgetById**](DefaultApi.md#DeleteWidgetById) | **Delete** /dashboard/widgets/{id} | delete widget by id
[**EditTimezoneInUsing**](DefaultApi.md#EditTimezoneInUsing) | **Put** /setting/admins/timezone/{id}/timezoneInUsing | 
[**ExecuteReport**](DefaultApi.md#ExecuteReport) | **Post** /report/reports/{id}/executions | 
[**FetchReport**](DefaultApi.md#FetchReport) | **Get** /report/reports/{id}/tasks/{taskId} | 
[**FinishWizard**](DefaultApi.md#FinishWizard) | **Post** /setting/admins/{id}/finishwizard | 
[**GetAdminById**](DefaultApi.md#GetAdminById) | **Get** /setting/admins/{id} | get admin
[**GetAdminList**](DefaultApi.md#GetAdminList) | **Get** /setting/admins | get admin list
[**GetAlertById**](DefaultApi.md#GetAlertById) | **Get** /alert/alerts/{id} | get alert
[**GetAlertList**](DefaultApi.md#GetAlertList) | **Get** /alert/alerts | get alert list
[**GetAlertListByDeviceGroupId**](DefaultApi.md#GetAlertListByDeviceGroupId) | **Get** /device/groups/{id}/alerts | get device group alerts
[**GetAlertListByDeviceId**](DefaultApi.md#GetAlertListByDeviceId) | **Get** /device/devices/{id}/alerts | get alerts
[**GetAlertRuleById**](DefaultApi.md#GetAlertRuleById) | **Get** /setting/alert/rules/{id} | get alert rule by id
[**GetAlertRuleList**](DefaultApi.md#GetAlertRuleList) | **Get** /setting/alert/rules | get alert rule list
[**GetAllSDTListByDeviceId**](DefaultApi.md#GetAllSDTListByDeviceId) | **Get** /device/devices/{id}/sdts | get all device related SDTs by id
[**GetAllSDTListByServiceGroupId**](DefaultApi.md#GetAllSDTListByServiceGroupId) | **Get** /service/groups/{id}/sdts | get all SDT list by service group id
[**GetApiTokenListByAdminId**](DefaultApi.md#GetApiTokenListByAdminId) | **Get** /setting/admins/{adminId}/apitokens | get apiToken by admin
[**GetApplications**](DefaultApi.md#GetApplications) | **Get** /device/devices/{id}/topTalkers | 
[**GetBandwidthGraphData**](DefaultApi.md#GetBandwidthGraphData) | **Get** /device/devices/{id}/topTalkersGraph | 
[**GetCollectorById**](DefaultApi.md#GetCollectorById) | **Get** /setting/collectors/{id} | get collector
[**GetCollectorGroupById**](DefaultApi.md#GetCollectorGroupById) | **Get** /setting/collectors/groups/{id} | get collector group
[**GetCollectorGroupList**](DefaultApi.md#GetCollectorGroupList) | **Get** /setting/collectors/groups | get collector group list
[**GetCollectorList**](DefaultApi.md#GetCollectorList) | **Get** /setting/collectors | get collector list
[**GetDashboardById**](DefaultApi.md#GetDashboardById) | **Get** /dashboard/dashboards/{id} | get dashboard
[**GetDashboardGroupById**](DefaultApi.md#GetDashboardGroupById) | **Get** /dashboard/groups/{id} | get dashboard group
[**GetDashboardGroupList**](DefaultApi.md#GetDashboardGroupList) | **Get** /dashboard/groups | get dashboard group list
[**GetDashboardGroupListForAutocomplete**](DefaultApi.md#GetDashboardGroupListForAutocomplete) | **Get** /dashboard/groups/autocomplete | get dashboard group list using for autocomplete
[**GetDashboardList**](DefaultApi.md#GetDashboardList) | **Get** /dashboard/dashboards | get dashboard list
[**GetDatasourceById**](DefaultApi.md#GetDatasourceById) | **Get** /setting/datasources/{id} | get datasource by id
[**GetDatasourceList**](DefaultApi.md#GetDatasourceList) | **Get** /setting/datasources | get datasource list
[**GetDeviceById**](DefaultApi.md#GetDeviceById) | **Get** /device/devices/{id} | get device by id
[**GetDeviceConfigSourceConfigById**](DefaultApi.md#GetDeviceConfigSourceConfigById) | **Get** /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config/{id} | get device configSource instance config by id 
[**GetDeviceConfigSourceConfigList**](DefaultApi.md#GetDeviceConfigSourceConfigList) | **Get** /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config | get device configSource instance config list
[**GetDeviceDatasourceById**](DefaultApi.md#GetDeviceDatasourceById) | **Get** /device/devices/{deviceId}/devicedatasources/{id} | get device datasource 
[**GetDeviceDatasourceDataById**](DefaultApi.md#GetDeviceDatasourceDataById) | **Get** /device/devices/{deviceId}/devicedatasources/{id}/data | get device datasource data 
[**GetDeviceDatasourceInstanceAlertSettingById**](DefaultApi.md#GetDeviceDatasourceInstanceAlertSettingById) | **Get** /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id} | get device instance alert setting
[**GetDeviceDatasourceInstanceById**](DefaultApi.md#GetDeviceDatasourceInstanceById) | **Get** /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id} | get device instance 
[**GetDeviceDatasourceInstanceData**](DefaultApi.md#GetDeviceDatasourceInstanceData) | **Get** /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/data | get device instance data
[**GetDeviceDatasourceInstanceGraphData**](DefaultApi.md#GetDeviceDatasourceInstanceGraphData) | **Get** /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/graphs/{graphId}/data | get device instance graph data 
[**GetDeviceDatasourceInstanceGroupById**](DefaultApi.md#GetDeviceDatasourceInstanceGroupById) | **Get** /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{id} | get device datasource instance group 
[**GetDeviceDatasourceInstanceGroupOverviewGraphData**](DefaultApi.md#GetDeviceDatasourceInstanceGroupOverviewGraphData) | **Get** /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{dsigId}/graphs/{ographId}/data | get device instance group overview graph data 
[**GetDeviceDatasourceList**](DefaultApi.md#GetDeviceDatasourceList) | **Get** /device/devices/{deviceId}/devicedatasources | get device datasource list 
[**GetDeviceGroupById**](DefaultApi.md#GetDeviceGroupById) | **Get** /device/groups/{id} | get device group
[**GetDeviceGroupDatasourceAlertSetting**](DefaultApi.md#GetDeviceGroupDatasourceAlertSetting) | **Get** /device/groups/{deviceGroupId}/datasources/{dsId}/alertsettings | get device group datasource alert setting 
[**GetDeviceGroupDatasourceById**](DefaultApi.md#GetDeviceGroupDatasourceById) | **Get** /device/groups/{deviceGroupId}/datasources/{id} | get device group datasource
[**GetDeviceGroupDatasourceList**](DefaultApi.md#GetDeviceGroupDatasourceList) | **Get** /device/groups/{deviceGroupId}/datasources | get device group datasource list
[**GetDeviceGroupList**](DefaultApi.md#GetDeviceGroupList) | **Get** /device/groups | get device group list
[**GetDeviceGroupProperties**](DefaultApi.md#GetDeviceGroupProperties) | **Get** /device/groups/{gid}/properties | get device group properties
[**GetDeviceGroupPropertyByName**](DefaultApi.md#GetDeviceGroupPropertyByName) | **Get** /device/groups/{gid}/properties/{name} | get device group property by name
[**GetDeviceInstanceGraphDataOnlyByInstanceId**](DefaultApi.md#GetDeviceInstanceGraphDataOnlyByInstanceId) | **Get** /device/devicedatasourceinstances/{instanceId}/graphs/{graphId}/data | get device instance graphData
[**GetDeviceList**](DefaultApi.md#GetDeviceList) | **Get** /device/devices | get device list
[**GetDevicePropertiesList**](DefaultApi.md#GetDevicePropertiesList) | **Get** /device/devices/{deviceId}/properties | get device properties
[**GetDevicePropertyByName**](DefaultApi.md#GetDevicePropertyByName) | **Get** /device/devices/{deviceId}/properties/{name} | get device property by name
[**GetEndpoints**](DefaultApi.md#GetEndpoints) | **Get** /device/devices/{id}/endpoints | 
[**GetEscalationChainById**](DefaultApi.md#GetEscalationChainById) | **Get** /setting/alert/chains/{id} | get escalation chain by id
[**GetEscalationChainList**](DefaultApi.md#GetEscalationChainList) | **Get** /setting/alert/chains | get escalation chain list
[**GetFlows**](DefaultApi.md#GetFlows) | **Get** /device/devices/{id}/flows | 
[**GetImmediateDeviceListByDeviceGroupId**](DefaultApi.md#GetImmediateDeviceListByDeviceGroupId) | **Get** /device/groups/{id}/devices | get immediate devices under group
[**GetImmediateServiceListByServiceGroupId**](DefaultApi.md#GetImmediateServiceListByServiceGroupId) | **Get** /service/groups/{id}/services | get immediate service list by service group id
[**GetListForDataSource**](DefaultApi.md#GetListForDataSource) | **Get** /setting/datasources/{id}/updatereasons | 
[**GetOpsNoteById**](DefaultApi.md#GetOpsNoteById) | **Get** /setting/opsnotes/{id} | get opsnote by id
[**GetOpsNoteList**](DefaultApi.md#GetOpsNoteList) | **Get** /setting/opsnotes | get opsnote list
[**GetPorts**](DefaultApi.md#GetPorts) | **Get** /device/devices/{id}/ports | 
[**GetQoSTable**](DefaultApi.md#GetQoSTable) | **Get** /device/devices/{id}/qos | 
[**GetReportById**](DefaultApi.md#GetReportById) | **Get** /report/reports/{id} | get report by id
[**GetReportList**](DefaultApi.md#GetReportList) | **Get** /report/reports | get report list
[**GetRoleById**](DefaultApi.md#GetRoleById) | **Get** /setting/roles/{id} | get role by id
[**GetRoleList**](DefaultApi.md#GetRoleList) | **Get** /setting/roles | get role list
[**GetSDTById**](DefaultApi.md#GetSDTById) | **Get** /sdt/sdts/{id} | get SDT by id
[**GetSDTList**](DefaultApi.md#GetSDTList) | **Get** /sdt/sdts | get SDT list
[**GetServiceAlertListByServiceId**](DefaultApi.md#GetServiceAlertListByServiceId) | **Get** /service/services/{id}/alerts | get service alert list by id
[**GetServiceById**](DefaultApi.md#GetServiceById) | **Get** /service/services/{id} | get service by id
[**GetServiceCheckpointDataById**](DefaultApi.md#GetServiceCheckpointDataById) | **Get** /service/services/{srvId}/checkpoints/{checkId}/data | get service checkpoint data by id
[**GetServiceDataByGraphName**](DefaultApi.md#GetServiceDataByGraphName) | **Get** /service/services/{id}/graphs/{graphName}/data | get service data by graph name
[**GetServiceGraphData**](DefaultApi.md#GetServiceGraphData) | **Get** /service/services/{serviceId}/checkpoints/{checkpointId}/graphs/{graphName}/data | get service graph data
[**GetServiceGroupById**](DefaultApi.md#GetServiceGroupById) | **Get** /service/groups/{id} | get service group
[**GetServiceGroupList**](DefaultApi.md#GetServiceGroupList) | **Get** /service/groups | get service group list
[**GetServiceList**](DefaultApi.md#GetServiceList) | **Get** /service/services | get service list
[**GetServicePropertyListByServiceId**](DefaultApi.md#GetServicePropertyListByServiceId) | **Get** /service/services/{id}/properties | get service property list by id
[**GetServiceSDTListByServiceId**](DefaultApi.md#GetServiceSDTListByServiceId) | **Get** /service/services/{id}/sdts | get service SDT list by id
[**GetSiteMonitorCheckPointList**](DefaultApi.md#GetSiteMonitorCheckPointList) | **Get** /service/smcheckpoints | get site monitor checkpoint list
[**GetUnmonitoredDeviceList**](DefaultApi.md#GetUnmonitoredDeviceList) | **Get** /device/unmonitoreddevices | get unmonitored device list
[**GetValuesOfToken**](DefaultApi.md#GetValuesOfToken) | **Get** /dashboard/groups/{id}/token/{tokenname}/values | 
[**GetWidgetById**](DefaultApi.md#GetWidgetById) | **Get** /dashboard/widgets/{id} | get widget by id
[**GetWidgetDataById**](DefaultApi.md#GetWidgetDataById) | **Get** /dashboard/widgets/{id}/data | get widget data
[**GetWidgetList**](DefaultApi.md#GetWidgetList) | **Get** /dashboard/widgets | get widget list
[**GetWidgetListByDashboardId**](DefaultApi.md#GetWidgetListByDashboardId) | **Get** /dashboard/dashboards/{id}/widgets | get widget list by DashboardId
[**ImportDatasourceWithXML**](DefaultApi.md#ImportDatasourceWithXML) | **Post** /setting/datasources/importxml | import datasource with xml
[**InstallCollector**](DefaultApi.md#InstallCollector) | **Get** /setting/collectors/{collectorId}/installers/{osAndArch} | install collector
[**PatchDeviceById**](DefaultApi.md#PatchDeviceById) | **Patch** /device/devices/{id} | patch a device
[**PatchDeviceGroupById**](DefaultApi.md#PatchDeviceGroupById) | **Patch** /device/groups/{id} | patch device group
[**PatchServiceById**](DefaultApi.md#PatchServiceById) | **Patch** /service/services/{id} | patch service by id
[**PatchServiceGroupById**](DefaultApi.md#PatchServiceGroupById) | **Patch** /service/groups/{id} | patch service group
[**QueryCloneStatus**](DefaultApi.md#QueryCloneStatus) | **Get** /dashboard/groups/asyncclone/{jobId} | 
[**ScheduleAutoDiscoveryByDeviceId**](DefaultApi.md#ScheduleAutoDiscoveryByDeviceId) | **Post** /device/devices/{id}/scheduleAutoDiscovery | schedule auto discovery for a host
[**UpdateAdminById**](DefaultApi.md#UpdateAdminById) | **Put** /setting/admins/{id} | update admin
[**UpdateAlertRuleById**](DefaultApi.md#UpdateAlertRuleById) | **Put** /setting/alert/rules/{id} | update alert rule
[**UpdateApiTokenByAdminId**](DefaultApi.md#UpdateApiTokenByAdminId) | **Put** /setting/admins/{adminId}/apitokens/{apitokenId} | update apiToken by admin
[**UpdateCollectorById**](DefaultApi.md#UpdateCollectorById) | **Put** /setting/collectors/{id} | update collector
[**UpdateCollectorGroupById**](DefaultApi.md#UpdateCollectorGroupById) | **Put** /setting/collectors/groups/{id} | update collector group
[**UpdateDashboardById**](DefaultApi.md#UpdateDashboardById) | **Put** /dashboard/dashboards/{id} | update dashboard
[**UpdateDashboardGroupById**](DefaultApi.md#UpdateDashboardGroupById) | **Put** /dashboard/groups/{id} | update dashboard group
[**UpdateDevice**](DefaultApi.md#UpdateDevice) | **Put** /device/devices/{id} | update a device
[**UpdateDeviceDatasourceInstanceAlertSettingById**](DefaultApi.md#UpdateDeviceDatasourceInstanceAlertSettingById) | **Put** /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id} | update device instance alert setting
[**UpdateDeviceDatasourceInstanceById**](DefaultApi.md#UpdateDeviceDatasourceInstanceById) | **Put** /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id} | update device instance 
[**UpdateDeviceDatasourceInstanceGroupById**](DefaultApi.md#UpdateDeviceDatasourceInstanceGroupById) | **Put** /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{id} | update device datasource instance group 
[**UpdateDeviceGroupById**](DefaultApi.md#UpdateDeviceGroupById) | **Put** /device/groups/{id} | update device group
[**UpdateDeviceGroupDatasourceAlertSetting**](DefaultApi.md#UpdateDeviceGroupDatasourceAlertSetting) | **Put** /device/groups/{deviceGroupId}/datasources/{dsId}/alertsettings | update device group datasource alert setting 
[**UpdateDeviceGroupPropertyByName**](DefaultApi.md#UpdateDeviceGroupPropertyByName) | **Put** /device/groups/{gid}/properties/{name} | update device group property
[**UpdateDevicePropertyByName**](DefaultApi.md#UpdateDevicePropertyByName) | **Put** /device/devices/{deviceId}/properties/{name} | update device  property
[**UpdateEscalationChainById**](DefaultApi.md#UpdateEscalationChainById) | **Put** /setting/alert/chains/{id} | update escalation chain by id
[**UpdateOpsNoteById**](DefaultApi.md#UpdateOpsNoteById) | **Put** /setting/opsnotes/{id} | update opsnote by id
[**UpdateReportById**](DefaultApi.md#UpdateReportById) | **Put** /report/reports/{id} | update report by id
[**UpdateRoleById**](DefaultApi.md#UpdateRoleById) | **Put** /setting/roles/{id} | update role
[**UpdateSDTById**](DefaultApi.md#UpdateSDTById) | **Put** /sdt/sdts/{id} | update SDT by id
[**UpdateServiceById**](DefaultApi.md#UpdateServiceById) | **Put** /service/services/{id} | update service by id
[**UpdateServiceGroupById**](DefaultApi.md#UpdateServiceGroupById) | **Put** /service/groups/{id} | update service group
[**UpdateWidgetById**](DefaultApi.md#UpdateWidgetById) | **Put** /dashboard/widgets/{id} | update widget by id


# **AckAlertById**
> NullObjectResponse AckAlertById(ctx, body, id)
ack alert by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AlertAck**](AlertAck.md)|  | 
  **id** | **string**|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AckCollectorDownAlertById**
> NullObjectResponse AckCollectorDownAlertById(ctx, id, body)
ack collector down alert



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**AckCollectorDown**](AckCollectorDown.md)|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Add**
> Response Add(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeviceDatasourceInstanceConfigPaginationResponse**](DeviceDatasourceInstanceConfigPaginationResponse.md)|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddAdmin**
> AdminResponse AddAdmin(ctx, body)
add admin



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Admin**](Admin.md)|  | 

### Return type

[**AdminResponse**](AdminResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddAlertNoteById**
> NullObjectResponse AddAlertNoteById(ctx, body, id)
add alert note



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AlertAck**](AlertAck.md)|  | 
  **id** | **string**|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddAlertRule**
> AlertRuleResponse AddAlertRule(ctx, body)
add alert rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AlertRule**](AlertRule.md)|  | 

### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddApiTokenByAdminId**
> ApiTokenResponse AddApiTokenByAdminId(ctx, adminId, body)
add apiToken by admin



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **adminId** | **int32**|  | 
  **body** | [**ApiToken**](ApiToken.md)|  | 

### Return type

[**ApiTokenResponse**](ApiTokenResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCollector**
> CollectorResponse AddCollector(ctx, body)
add collector



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Collector**](Collector.md)|  | 

### Return type

[**CollectorResponse**](CollectorResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCollectorGroup**
> CollectorGroupResponse AddCollectorGroup(ctx, body)
add collector group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CollectorGroup**](CollectorGroup.md)|  | 

### Return type

[**CollectorGroupResponse**](CollectorGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDashboard**
> DashboardResponse AddDashboard(ctx, body)
add dashboard



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Dashboard**](Dashboard.md)|  | 

### Return type

[**DashboardResponse**](DashboardResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDashboardGroup**
> DashboardGroupResponse AddDashboardGroup(ctx, body)
add dashboard group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DashboardGroup**](DashboardGroup.md)|  | 

### Return type

[**DashboardGroupResponse**](DashboardGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDevice**
> DeviceResponse AddDevice(ctx, body, optional)
add a new device



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Device**](Device.md)|  | 
 **optional** | ***AddDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **addFromWizard** | **optional.Bool**|  | [default to false]

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDeviceDatasourceInstance**
> DeviceDatasourceInstanceResponse AddDeviceDatasourceInstance(ctx, deviceId, hdsId, body)
add device instance 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **hdsId** | **int32**|  | 
  **body** | [**DeviceDataSourceInstance**](DeviceDataSourceInstance.md)|  | 

### Return type

[**DeviceDatasourceInstanceResponse**](DeviceDatasourceInstanceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDeviceDatasourceInstanceGroup**
> DeviceDatasourceInstanceGroupResponse AddDeviceDatasourceInstanceGroup(ctx, deviceId, deviceDsId, body)
add device datasource instance group 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **deviceDsId** | **int32**|  | 
  **body** | [**DeviceDataSourceInstanceGroup**](DeviceDataSourceInstanceGroup.md)|  | 

### Return type

[**DeviceDatasourceInstanceGroupResponse**](DeviceDatasourceInstanceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDeviceGroup**
> DeviceGroupResponse AddDeviceGroup(ctx, body)
add device group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeviceGroup**](DeviceGroup.md)|  | 

### Return type

[**DeviceGroupResponse**](DeviceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDeviceGroupProperty**
> PropertyResponse AddDeviceGroupProperty(ctx, gid, body)
add device group property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **int32**|  | 
  **body** | [**EntityProperty**](EntityProperty.md)|  | 

### Return type

[**PropertyResponse**](PropertyResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDeviceProperty**
> PropertyResponse AddDeviceProperty(ctx, deviceId, body)
add device property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **body** | [**EntityProperty**](EntityProperty.md)|  | 

### Return type

[**PropertyResponse**](PropertyResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddEscalationChain**
> EscalationChainResponse AddEscalationChain(ctx, body)
add escalation chain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EscalatingChain**](EscalatingChain.md)|  | 

### Return type

[**EscalationChainResponse**](EscalationChainResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddOpsNote**
> OpsNoteResponse AddOpsNote(ctx, body)
add opsnote



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RestOpsNoteV1**](RestOpsNoteV1.md)|  | 

### Return type

[**OpsNoteResponse**](OpsNoteResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddReport**
> ReportResponse AddReport(ctx, body)
add report



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReportBase**](ReportBase.md)|  | 

### Return type

[**ReportResponse**](ReportResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRole**
> RoleResponse AddRole(ctx, body)
add role



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Role**](Role.md)|  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSDT**
> SdtResponse AddSDT(ctx, body)
add SDT



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdtBase**](SdtBase.md)|  | 

### Return type

[**SdtResponse**](SDTResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddService**
> ServiceResponse AddService(ctx, body)
add service



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceBaseVersion1**](ServiceBaseVersion1.md)|  | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceGroup**
> ServiceGroupResponse AddServiceGroup(ctx, body)
add service group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceGroup**](ServiceGroup.md)|  | 

### Return type

[**ServiceGroupResponse**](ServiceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddWidget**
> WidgetResponse AddWidget(ctx, body)
add widget



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WidgetBaseVersion1**](WidgetBaseVersion1.md)|  | 

### Return type

[**WidgetResponse**](WidgetResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AsyncClone**
> Response AsyncClone(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***AsyncCloneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsyncCloneOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DashboardGroup**](DashboardGroup.md)|  | 
 **recursive** | **optional.Bool**|  | [default to true]

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Clone**
> Response Clone(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***CloneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloneOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DashboardGroup**](DashboardGroup.md)|  | 
 **recursive** | **optional.Bool**|  | [default to true]

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAutoReports**
> Response CreateAutoReports(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAutoReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateAutoReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RestCloudAutoVisualInfo**](RestCloudAutoVisualInfo.md)|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAdminById**
> NullObjectResponse DeleteAdminById(ctx, id)
delete admin



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlertRuleById**
> AlertRuleResponse DeleteAlertRuleById(ctx, id)
delete alert rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiTokenById**
> ApiTokenResponse DeleteApiTokenById(ctx, adminId, apitokenId)
delete apiToken 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **adminId** | **int32**|  | 
  **apitokenId** | **int32**|  | 

### Return type

[**ApiTokenResponse**](ApiTokenResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectorById**
> CollectorResponse DeleteCollectorById(ctx, id)
delete collector



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**CollectorResponse**](CollectorResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectorGroupById**
> CollectorGroupResponse DeleteCollectorGroupById(ctx, id)
delete collector group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**CollectorGroupResponse**](CollectorGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDashboardById**
> NullObjectResponse DeleteDashboardById(ctx, id)
delete dashboard



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDashboardGroupById**
> NullObjectResponse DeleteDashboardGroupById(ctx, id, optional)
delete dashboard group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***DeleteDashboardGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteDashboardGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowNonEmptyGroup** | **optional.Bool**|  | [default to false]

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDevice**
> NullObjectResponse DeleteDevice(ctx, id, optional)
delete a device



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***DeleteDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **deleteHard** | **optional.Bool**|  | [default to true]

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeviceGroupById**
> NullObjectResponse DeleteDeviceGroupById(ctx, id, optional)
delete device group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***DeleteDeviceGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteDeviceGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteChildren** | **optional.Bool**|  | [default to false]
 **deleteHard** | **optional.Bool**|  | [default to true]

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeviceGroupPropertyByName**
> NullObjectResponse DeleteDeviceGroupPropertyByName(ctx, gid, name)
delete device group property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **int32**|  | 
  **name** | **string**|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDevicePropertyByName**
> NullObjectResponse DeleteDevicePropertyByName(ctx, deviceId, name)
delete device  property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **name** | **string**|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEscalationChainById**
> NullObjectResponse DeleteEscalationChainById(ctx, id)
delete escalation chain by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOpsNoteById**
> StringResponse DeleteOpsNoteById(ctx, id)
delete opsnote by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**StringResponse**](StringResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleById**
> RoleResponse DeleteRoleById(ctx, id)
delete role



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSDTById**
> SdtResponse DeleteSDTById(ctx, id)
delete SDT by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SdtResponse**](SDTResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceById**
> NullObjectResponse DeleteServiceById(ctx, id)
delete service by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceGroupById**
> NullObjectResponse DeleteServiceGroupById(ctx, id, optional)
delete service group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***DeleteServiceGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteServiceGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteChildren** | **optional.Int32**|  | [default to 0]

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteToken**
> Response DeleteToken(ctx, tokenname, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenname** | **string**|  | 
  **body** | [**[]WidgetTokenValue**](WidgetTokenValue.md)|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWidgetById**
> NullObjectResponse DeleteWidgetById(ctx, id)
delete widget by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**NullObjectResponse**](NullObjectResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTimezoneInUsing**
> Response EditTimezoneInUsing(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***EditTimezoneInUsingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditTimezoneInUsingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timezoneInUsing** | **optional.String**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteReport**
> Response ExecuteReport(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ExecuteReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RestGenerateReportFunc**](RestGenerateReportFunc.md)|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchReport**
> Response FetchReport(ctx, id, taskId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **taskId** | **string**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FinishWizard**
> Response FinishWizard(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdminById**
> AdminResponse GetAdminById(ctx, id, optional)
get admin



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetAdminByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAdminByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**AdminResponse**](AdminResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdminList**
> AdminPaginationResponse GetAdminList(ctx, optional)
get admin list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAdminListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAdminListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**AdminPaginationResponse**](AdminPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertById**
> AlertResponse GetAlertById(ctx, id, optional)
get alert



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GetAlertByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAlertByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **needMessage** | **optional.Bool**|  | [default to false]
 **customColumns** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**AlertResponse**](AlertResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertList**
> AlertPaginationResponse GetAlertList(ctx, optional)
get alert list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAlertListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAlertListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **needMessage** | **optional.Bool**|  | [default to false]
 **customColumns** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **searchId** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**AlertPaginationResponse**](AlertPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertListByDeviceGroupId**
> AlertPaginationResponse GetAlertListByDeviceGroupId(ctx, id, optional)
get device group alerts



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetAlertListByDeviceGroupIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAlertListByDeviceGroupIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **needMessage** | **optional.Bool**|  | [default to false]
 **customColumns** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **searchId** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**AlertPaginationResponse**](AlertPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertListByDeviceId**
> AlertPaginationResponse GetAlertListByDeviceId(ctx, id, optional)
get alerts



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetAlertListByDeviceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAlertListByDeviceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **needMessage** | **optional.Bool**|  | [default to false]
 **customColumns** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **searchId** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**AlertPaginationResponse**](AlertPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertRuleById**
> AlertRuleResponse GetAlertRuleById(ctx, id, optional)
get alert rule by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetAlertRuleByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAlertRuleByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertRuleList**
> AlertRulePaginationResponse GetAlertRuleList(ctx, optional)
get alert rule list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAlertRuleListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAlertRuleListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**AlertRulePaginationResponse**](AlertRulePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSDTListByDeviceId**
> SdtPaginationResponse GetAllSDTListByDeviceId(ctx, id, optional)
get all device related SDTs by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetAllSDTListByDeviceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllSDTListByDeviceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**SdtPaginationResponse**](SDTPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSDTListByServiceGroupId**
> ServicePaginationResponse GetAllSDTListByServiceGroupId(ctx, id, optional)
get all SDT list by service group id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetAllSDTListByServiceGroupIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllSDTListByServiceGroupIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**ServicePaginationResponse**](ServicePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiTokenListByAdminId**
> ApiTokenPaginationResponse GetApiTokenListByAdminId(ctx, adminId, optional)
get apiToken by admin



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **adminId** | **int32**|  | 
 **optional** | ***GetApiTokenListByAdminIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApiTokenListByAdminIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**ApiTokenPaginationResponse**](ApiTokenPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplications**
> Response GetApplications(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApplicationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBandwidthGraphData**
> GetBandwidthGraphData(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetBandwidthGraphDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBandwidthGraphDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **netflowFilter** | **optional.String**|  | 
 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **format** | **optional.String**|  | 
 **keyword** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollectorById**
> CollectorResponse GetCollectorById(ctx, id, optional)
get collector



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetCollectorByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCollectorByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**CollectorResponse**](CollectorResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollectorGroupById**
> CollectorGroupResponse GetCollectorGroupById(ctx, id, optional)
get collector group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetCollectorGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCollectorGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**CollectorGroupResponse**](CollectorGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollectorGroupList**
> CollectorGroupPaginationResponse GetCollectorGroupList(ctx, optional)
get collector group list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCollectorGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCollectorGroupListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**CollectorGroupPaginationResponse**](CollectorGroupPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollectorList**
> CollectorPaginationResponse GetCollectorList(ctx, optional)
get collector list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCollectorListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCollectorListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**CollectorPaginationResponse**](CollectorPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardById**
> DashboardResponse GetDashboardById(ctx, id, optional)
get dashboard



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetDashboardByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDashboardByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**DashboardResponse**](DashboardResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardGroupById**
> DashboardGroupResponse GetDashboardGroupById(ctx, id, optional)
get dashboard group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetDashboardGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDashboardGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **template** | **optional.Bool**|  | [default to false]
 **fields** | **optional.String**|  | 

### Return type

[**DashboardGroupResponse**](DashboardGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardGroupList**
> DashboardGroupPaginationResponse GetDashboardGroupList(ctx, optional)
get dashboard group list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDashboardGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDashboardGroupListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**DashboardGroupPaginationResponse**](DashboardGroupPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardGroupListForAutocomplete**
> DashboardGroupPaginationResponse GetDashboardGroupListForAutocomplete(ctx, )
get dashboard group list using for autocomplete



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DashboardGroupPaginationResponse**](DashboardGroupPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardList**
> DashboardPaginationResponse GetDashboardList(ctx, optional)
get dashboard list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDashboardListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDashboardListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**DashboardPaginationResponse**](DashboardPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatasourceById**
> DatasourceResponse GetDatasourceById(ctx, id, optional)
get datasource by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetDatasourceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDatasourceByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | [default to json]
 **fields** | **optional.String**|  | 

### Return type

[**DatasourceResponse**](DatasourceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatasourceList**
> DatasourcePaginationResponse GetDatasourceList(ctx, optional)
get datasource list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDatasourceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDatasourceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**|  | [default to json]
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**DatasourcePaginationResponse**](DatasourcePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceById**
> DeviceResponse GetDeviceById(ctx, id, optional)
get device by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetDeviceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceConfigSourceConfigById**
> DeviceDatasourceInstanceConfigResponse GetDeviceConfigSourceConfigById(ctx, deviceId, hdsId, instanceId, id, optional)
get device configSource instance config by id 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **hdsId** | **int32**|  | 
  **instanceId** | **int32**|  | 
  **id** | **string**|  | 
 **optional** | ***GetDeviceConfigSourceConfigByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceConfigSourceConfigByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **format** | **optional.String**|  | [default to json]
 **startEpoch** | **optional.Int64**|  | [default to 0]
 **fields** | **optional.String**|  | 

### Return type

[**DeviceDatasourceInstanceConfigResponse**](DeviceDatasourceInstanceConfigResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceConfigSourceConfigList**
> DeviceDatasourceInstanceConfigPaginationResponse GetDeviceConfigSourceConfigList(ctx, deviceId, hdsId, instanceId, optional)
get device configSource instance config list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **hdsId** | **int32**|  | 
  **instanceId** | **int32**|  | 
 **optional** | ***GetDeviceConfigSourceConfigListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceConfigSourceConfigListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**DeviceDatasourceInstanceConfigPaginationResponse**](DeviceDatasourceInstanceConfigPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceDatasourceById**
> DeviceDatasourceResponse GetDeviceDatasourceById(ctx, deviceId, id, optional)
get device datasource 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***GetDeviceDatasourceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceDatasourceByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**DeviceDatasourceResponse**](DeviceDatasourceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceDatasourceDataById**
> DeviceDatasourceDataResponse GetDeviceDatasourceDataById(ctx, deviceId, id, optional)
get device datasource data 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***GetDeviceDatasourceDataByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceDatasourceDataByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **period** | **optional.Float64**|  | [default to 1]
 **start** | **optional.Int64**|  | [default to 0]
 **end** | **optional.Int64**|  | [default to 0]
 **datapoints** | **optional.String**|  | 
 **format** | **optional.String**|  | [default to json]

### Return type

[**DeviceDatasourceDataResponse**](DeviceDatasourceDataResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceDatasourceInstanceAlertSettingById**
> DeviceDatasourceInstanceAlertSettingResponse GetDeviceDatasourceInstanceAlertSettingById(ctx, deviceId, hdsId, instanceId, id, optional)
get device instance alert setting



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **hdsId** | **int32**|  | 
  **instanceId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***GetDeviceDatasourceInstanceAlertSettingByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceDatasourceInstanceAlertSettingByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fields** | **optional.String**|  | 

### Return type

[**DeviceDatasourceInstanceAlertSettingResponse**](DeviceDatasourceInstanceAlertSettingResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceDatasourceInstanceById**
> DeviceDatasourceInstanceResponse GetDeviceDatasourceInstanceById(ctx, deviceId, hdsId, id, optional)
get device instance 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **hdsId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***GetDeviceDatasourceInstanceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceDatasourceInstanceByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **optional.String**|  | 

### Return type

[**DeviceDatasourceInstanceResponse**](DeviceDatasourceInstanceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceDatasourceInstanceData**
> DeviceDatasourceInstanceDataResponse GetDeviceDatasourceInstanceData(ctx, deviceId, hdsId, id, optional)
get device instance data



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **hdsId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***GetDeviceDatasourceInstanceDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceDatasourceInstanceDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **period** | **optional.Float64**|  | [default to 1]
 **start** | **optional.Int64**|  | [default to 0]
 **end** | **optional.Int64**|  | [default to 0]
 **datapoints** | **optional.String**|  | 
 **format** | **optional.String**|  | [default to json]

### Return type

[**DeviceDatasourceInstanceDataResponse**](DeviceDatasourceInstanceDataResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceDatasourceInstanceGraphData**
> GraphPlotResponse GetDeviceDatasourceInstanceGraphData(ctx, deviceId, hdsId, id, graphId, optional)
get device instance graph data 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **hdsId** | **int32**|  | 
  **id** | **int32**|  | 
  **graphId** | **int32**|  | 
 **optional** | ***GetDeviceDatasourceInstanceGraphDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceDatasourceInstanceGraphDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **format** | **optional.String**|  | 

### Return type

[**GraphPlotResponse**](GraphPlotResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceDatasourceInstanceGroupById**
> DeviceDatasourceInstanceGroupResponse GetDeviceDatasourceInstanceGroupById(ctx, deviceId, deviceDsId, id, optional)
get device datasource instance group 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **deviceDsId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***GetDeviceDatasourceInstanceGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceDatasourceInstanceGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **optional.String**|  | 

### Return type

[**DeviceDatasourceInstanceGroupResponse**](DeviceDatasourceInstanceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceDatasourceInstanceGroupOverviewGraphData**
> GraphPlotResponse GetDeviceDatasourceInstanceGroupOverviewGraphData(ctx, deviceId, deviceDsId, dsigId, ographId, optional)
get device instance group overview graph data 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **deviceDsId** | **int32**|  | 
  **dsigId** | **int32**|  | 
  **ographId** | **int32**|  | 
 **optional** | ***GetDeviceDatasourceInstanceGroupOverviewGraphDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceDatasourceInstanceGroupOverviewGraphDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **format** | **optional.String**|  | 

### Return type

[**GraphPlotResponse**](GraphPlotResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceDatasourceList**
> DeviceDatasourcePaginationResponse GetDeviceDatasourceList(ctx, deviceId, optional)
get device datasource list 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***GetDeviceDatasourceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceDatasourceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**DeviceDatasourcePaginationResponse**](DeviceDatasourcePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceGroupById**
> DeviceGroupResponse GetDeviceGroupById(ctx, id, optional)
get device group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetDeviceGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**DeviceGroupResponse**](DeviceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceGroupDatasourceAlertSetting**
> DeviceGroupDatasourceAlertConfigResponse GetDeviceGroupDatasourceAlertSetting(ctx, deviceGroupId, dsId, optional)
get device group datasource alert setting 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceGroupId** | **int32**|  | 
  **dsId** | **int32**|  | 
 **optional** | ***GetDeviceGroupDatasourceAlertSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceGroupDatasourceAlertSettingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**DeviceGroupDatasourceAlertConfigResponse**](DeviceGroupDatasourceAlertConfigResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceGroupDatasourceById**
> DeviceGroupDatasourceResponse GetDeviceGroupDatasourceById(ctx, deviceGroupId, id, optional)
get device group datasource



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceGroupId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***GetDeviceGroupDatasourceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceGroupDatasourceByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**DeviceGroupDatasourceResponse**](DeviceGroupDatasourceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceGroupDatasourceList**
> DeviceGroupDatasourceResponse GetDeviceGroupDatasourceList(ctx, deviceGroupId, optional)
get device group datasource list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceGroupId** | **int32**|  | 
 **optional** | ***GetDeviceGroupDatasourceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceGroupDatasourceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDisabledDataSourceWithoutInstance** | **optional.Bool**|  | [default to false]
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**DeviceGroupDatasourceResponse**](DeviceGroupDatasourceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceGroupList**
> DeviceGroupPaginationResponse GetDeviceGroupList(ctx, optional)
get device group list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDeviceGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceGroupListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**DeviceGroupPaginationResponse**](DeviceGroupPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceGroupProperties**
> PropertyPaginationResponse GetDeviceGroupProperties(ctx, gid, optional)
get device group properties



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **int32**|  | 
 **optional** | ***GetDeviceGroupPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceGroupPropertiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**PropertyPaginationResponse**](PropertyPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceGroupPropertyByName**
> PropertyResponse GetDeviceGroupPropertyByName(ctx, gid, name, optional)
get device group property by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **int32**|  | 
  **name** | **string**|  | 
 **optional** | ***GetDeviceGroupPropertyByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceGroupPropertyByNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**PropertyResponse**](PropertyResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceInstanceGraphDataOnlyByInstanceId**
> GraphPlotResponse GetDeviceInstanceGraphDataOnlyByInstanceId(ctx, instanceId, graphId, optional)
get device instance graphData



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **int32**|  | 
  **graphId** | **int32**|  | 
 **optional** | ***GetDeviceInstanceGraphDataOnlyByInstanceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceInstanceGraphDataOnlyByInstanceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **format** | **optional.String**|  | 

### Return type

[**GraphPlotResponse**](GraphPlotResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceList**
> DevicePaginationResponse GetDeviceList(ctx, optional)
get device list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDeviceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeviceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**DevicePaginationResponse**](DevicePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicePropertiesList**
> PropertyPaginationResponse GetDevicePropertiesList(ctx, deviceId, optional)
get device properties



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***GetDevicePropertiesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDevicePropertiesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**PropertyPaginationResponse**](PropertyPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicePropertyByName**
> PropertyResponse GetDevicePropertyByName(ctx, deviceId, name, optional)
get device property by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **name** | **string**|  | 
 **optional** | ***GetDevicePropertyByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDevicePropertyByNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**PropertyResponse**](PropertyResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEndpoints**
> Response GetEndpoints(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEndpointsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **port** | **optional.String**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEscalationChainById**
> EscalationChainResponse GetEscalationChainById(ctx, id, optional)
get escalation chain by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetEscalationChainByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEscalationChainByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**EscalationChainResponse**](EscalationChainResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEscalationChainList**
> EscalationChainPaginationResponse GetEscalationChainList(ctx, optional)
get escalation chain list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEscalationChainListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEscalationChainListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**EscalationChainPaginationResponse**](EscalationChainPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlows**
> Response GetFlows(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetFlowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFlowsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImmediateDeviceListByDeviceGroupId**
> DevicePaginationResponse GetImmediateDeviceListByDeviceGroupId(ctx, id, optional)
get immediate devices under group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetImmediateDeviceListByDeviceGroupIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetImmediateDeviceListByDeviceGroupIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**DevicePaginationResponse**](DevicePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImmediateServiceListByServiceGroupId**
> ServicePaginationResponse GetImmediateServiceListByServiceGroupId(ctx, id, optional)
get immediate service list by service group id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetImmediateServiceListByServiceGroupIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetImmediateServiceListByServiceGroupIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**ServicePaginationResponse**](ServicePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListForDataSource**
> Response GetListForDataSource(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpsNoteById**
> OpsNoteResponse GetOpsNoteById(ctx, id, optional)
get opsnote by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GetOpsNoteByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpsNoteByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**OpsNoteResponse**](OpsNoteResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpsNoteList**
> OpsNotePaginationResponse GetOpsNoteList(ctx, optional)
get opsnote list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOpsNoteListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpsNoteListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**OpsNotePaginationResponse**](OpsNotePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPorts**
> Response GetPorts(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **ip** | **optional.String**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQoSTable**
> Response GetQoSTable(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetQoSTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetQoSTableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportById**
> ReportResponse GetReportById(ctx, id, optional)
get report by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetReportByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReportByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**ReportResponse**](ReportResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportList**
> ReportPaginationResponse GetReportList(ctx, optional)
get report list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetReportListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReportListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**ReportPaginationResponse**](ReportPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleById**
> RoleResponse GetRoleById(ctx, id, optional)
get role by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetRoleByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRoleByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleList**
> RolePaginationResponse GetRoleList(ctx, optional)
get role list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRoleListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRoleListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**RolePaginationResponse**](RolePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSDTById**
> SdtResponse GetSDTById(ctx, id, optional)
get SDT by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GetSDTByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSDTByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**SdtResponse**](SDTResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSDTList**
> SdtPaginationResponse GetSDTList(ctx, optional)
get SDT list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSDTListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSDTListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**SdtPaginationResponse**](SDTPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceAlertListByServiceId**
> AlertPaginationResponse GetServiceAlertListByServiceId(ctx, id, optional)
get service alert list by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetServiceAlertListByServiceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServiceAlertListByServiceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **needMessage** | **optional.Bool**|  | [default to false]
 **customColumns** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**AlertPaginationResponse**](AlertPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceById**
> ServiceResponse GetServiceById(ctx, id, optional)
get service by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetServiceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServiceByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceCheckpointDataById**
> ServiceCheckpointDataResponse GetServiceCheckpointDataById(ctx, srvId, checkId, optional)
get service checkpoint data by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **srvId** | **int32**|  | 
  **checkId** | **int32**|  | 
 **optional** | ***GetServiceCheckpointDataByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServiceCheckpointDataByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **period** | **optional.Float64**|  | [default to 1]
 **start** | **optional.Int64**|  | [default to 0]
 **end** | **optional.Int64**|  | [default to 0]
 **datapoints** | **optional.String**|  | 
 **format** | **optional.String**|  | [default to json]

### Return type

[**ServiceCheckpointDataResponse**](ServiceCheckpointDataResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceDataByGraphName**
> GraphPlotResponse GetServiceDataByGraphName(ctx, id, graphName, optional)
get service data by graph name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **graphName** | **string**|  | 
 **optional** | ***GetServiceDataByGraphNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServiceDataByGraphNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **optional.Int64**|  | [default to 0]
 **end** | **optional.Int64**|  | [default to 0]
 **format** | **optional.String**|  | 

### Return type

[**GraphPlotResponse**](GraphPlotResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceGraphData**
> GraphPlotResponse GetServiceGraphData(ctx, serviceId, checkpointId, graphName, optional)
get service graph data



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **int32**|  | 
  **checkpointId** | **int32**|  | 
  **graphName** | **string**|  | 
 **optional** | ***GetServiceGraphDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServiceGraphDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **format** | **optional.String**|  | 

### Return type

[**GraphPlotResponse**](GraphPlotResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceGroupById**
> ServiceGroupResponse GetServiceGroupById(ctx, id, optional)
get service group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetServiceGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServiceGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**ServiceGroupResponse**](ServiceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceGroupList**
> ServiceGroupPaginationResponse GetServiceGroupList(ctx, optional)
get service group list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetServiceGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServiceGroupListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**ServiceGroupPaginationResponse**](ServiceGroupPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceList**
> ServicePaginationResponse GetServiceList(ctx, optional)
get service list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServiceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectorIds** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**ServicePaginationResponse**](ServicePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServicePropertyListByServiceId**
> PropertyPaginationResponse GetServicePropertyListByServiceId(ctx, id, optional)
get service property list by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetServicePropertyListByServiceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServicePropertyListByServiceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**PropertyPaginationResponse**](PropertyPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceSDTListByServiceId**
> SdtPaginationResponse GetServiceSDTListByServiceId(ctx, id, optional)
get service SDT list by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetServiceSDTListByServiceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServiceSDTListByServiceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**SdtPaginationResponse**](SDTPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteMonitorCheckPointList**
> SiteMonitorCheckPointPaginationResponse GetSiteMonitorCheckPointList(ctx, optional)
get site monitor checkpoint list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSiteMonitorCheckPointListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSiteMonitorCheckPointListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**SiteMonitorCheckPointPaginationResponse**](SiteMonitorCheckPointPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnmonitoredDeviceList**
> UnmonitoredDevicePaginationResponse GetUnmonitoredDeviceList(ctx, optional)
get unmonitored device list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUnmonitoredDeviceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUnmonitoredDeviceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**UnmonitoredDevicePaginationResponse**](UnmonitoredDevicePaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetValuesOfToken**
> Response GetValuesOfToken(ctx, id, tokenname)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **tokenname** | **string**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWidgetById**
> WidgetResponse GetWidgetById(ctx, id, optional)
get widget by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetWidgetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWidgetByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**WidgetResponse**](WidgetResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWidgetDataById**
> WidgetDataResponse GetWidgetDataById(ctx, id, optional)
get widget data



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetWidgetDataByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWidgetDataByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **format** | **optional.String**|  | 

### Return type

[**WidgetDataResponse**](WidgetDataResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWidgetList**
> WidgetPaginationResponse GetWidgetList(ctx, optional)
get widget list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWidgetListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWidgetListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**WidgetPaginationResponse**](WidgetPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWidgetListByDashboardId**
> WidgetPaginationResponse GetWidgetListByDashboardId(ctx, id, optional)
get widget list by DashboardId



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetWidgetListByDashboardIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWidgetListByDashboardIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 50]
 **offset** | **optional.Int32**|  | [default to 0]
 **filter** | **optional.String**|  | 

### Return type

[**WidgetPaginationResponse**](WidgetPaginationResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportDatasourceWithXML**
> StringResponse ImportDatasourceWithXML(ctx, optional)
import datasource with xml



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportDatasourceWithXMLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportDatasourceWithXMLOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File**|  | 

### Return type

[**StringResponse**](StringResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallCollector**
> *os.File InstallCollector(ctx, collectorId, osAndArch, optional)
install collector



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectorId** | **int32**|  | 
  **osAndArch** | **string**|  | 
 **optional** | ***InstallCollectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallCollectorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectorVersion** | **optional.Int32**|  | 
 **token** | **optional.String**|  | 
 **monitorOthers** | **optional.Bool**|  | [default to true]
 **collectorSize** | **optional.String**|  | [default to medium]
 **useEA** | **optional.Bool**|  | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDeviceById**
> DeviceResponse PatchDeviceById(ctx, body, id, optional)
patch a device



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Device**](Device.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***PatchDeviceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchDeviceByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **opType** | **optional.String**|  | [default to refresh]
 **patchFields** | **optional.String**|  | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDeviceGroupById**
> DeviceGroupResponse PatchDeviceGroupById(ctx, id, body, optional)
patch device group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**DeviceGroup**](DeviceGroup.md)|  | 
 **optional** | ***PatchDeviceGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchDeviceGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opType** | **optional.String**|  | [default to refresh]
 **patchFields** | **optional.String**|  | 

### Return type

[**DeviceGroupResponse**](DeviceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceById**
> ServiceResponse PatchServiceById(ctx, id, body, optional)
patch service by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | **string**| This is a JSON string | 
 **optional** | ***PatchServiceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchServiceByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opType** | **optional.String**|  | [default to refresh]

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceGroupById**
> ServiceGroupResponse PatchServiceGroupById(ctx, id, body, optional)
patch service group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**ServiceGroup**](ServiceGroup.md)|  | 
 **optional** | ***PatchServiceGroupByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchServiceGroupByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opType** | **optional.String**|  | [default to refresh]
 **patchFields** | **optional.String**|  | 

### Return type

[**ServiceGroupResponse**](ServiceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryCloneStatus**
> Response QueryCloneStatus(ctx, jobId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**|  | 

### Return type

[**Response**](Response.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduleAutoDiscoveryByDeviceId**
> StringResponse ScheduleAutoDiscoveryByDeviceId(ctx, id, optional)
schedule auto discovery for a host



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ScheduleAutoDiscoveryByDeviceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleAutoDiscoveryByDeviceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 

### Return type

[**StringResponse**](StringResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAdminById**
> AdminResponse UpdateAdminById(ctx, id, body, optional)
update admin



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**Admin**](Admin.md)|  | 
 **optional** | ***UpdateAdminByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateAdminByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **changePassword** | **optional.Bool**|  | [default to false]

### Return type

[**AdminResponse**](AdminResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAlertRuleById**
> AlertRuleResponse UpdateAlertRuleById(ctx, body, id)
update alert rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AlertRule**](AlertRule.md)|  | 
  **id** | **int32**|  | 

### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApiTokenByAdminId**
> ApiTokenResponse UpdateApiTokenByAdminId(ctx, adminId, apitokenId, body)
update apiToken by admin



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **adminId** | **int32**|  | 
  **apitokenId** | **int32**|  | 
  **body** | [**DeviceDatasourceInstanceConfigPaginationResponse**](DeviceDatasourceInstanceConfigPaginationResponse.md)|  | 

### Return type

[**ApiTokenResponse**](ApiTokenResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCollectorById**
> CollectorResponse UpdateCollectorById(ctx, id, body)
update collector



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**Collector**](Collector.md)|  | 

### Return type

[**CollectorResponse**](CollectorResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCollectorGroupById**
> CollectorGroupResponse UpdateCollectorGroupById(ctx, id, body)
update collector group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**CollectorGroup**](CollectorGroup.md)|  | 

### Return type

[**CollectorGroupResponse**](CollectorGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDashboardById**
> DashboardResponse UpdateDashboardById(ctx, id, body, optional)
update dashboard



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**Dashboard**](Dashboard.md)|  | 
 **optional** | ***UpdateDashboardByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDashboardByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **overwriteGroupFields** | **optional.Bool**|  | [default to false]

### Return type

[**DashboardResponse**](DashboardResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDashboardGroupById**
> DashboardGroupResponse UpdateDashboardGroupById(ctx, id, body)
update dashboard group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**DashboardGroup**](DashboardGroup.md)|  | 

### Return type

[**DashboardGroupResponse**](DashboardGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDevice**
> DeviceResponse UpdateDevice(ctx, body, id, optional)
update a device



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Device**](Device.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***UpdateDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **optional.Int64**|  | 
 **end** | **optional.Int64**|  | 
 **netflowFilter** | **optional.String**|  | 
 **opType** | **optional.String**|  | [default to refresh]

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeviceDatasourceInstanceAlertSettingById**
> DeviceDatasourceInstanceAlertSettingResponse UpdateDeviceDatasourceInstanceAlertSettingById(ctx, deviceId, hdsId, instanceId, id, body)
update device instance alert setting



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **hdsId** | **int32**|  | 
  **instanceId** | **int32**|  | 
  **id** | **int32**|  | 
  **body** | [**DeviceDataSourceInstanceAlertSetting**](DeviceDataSourceInstanceAlertSetting.md)|  | 

### Return type

[**DeviceDatasourceInstanceAlertSettingResponse**](DeviceDatasourceInstanceAlertSettingResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeviceDatasourceInstanceById**
> DeviceDatasourceInstanceResponse UpdateDeviceDatasourceInstanceById(ctx, deviceId, hdsId, id, body, optional)
update device instance 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **hdsId** | **int32**|  | 
  **id** | **int32**|  | 
  **body** | [**DeviceDataSourceInstance**](DeviceDataSourceInstance.md)|  | 
 **optional** | ***UpdateDeviceDatasourceInstanceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDeviceDatasourceInstanceByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opType** | **optional.String**|  | [default to refresh]

### Return type

[**DeviceDatasourceInstanceResponse**](DeviceDatasourceInstanceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeviceDatasourceInstanceGroupById**
> DeviceDatasourceInstanceGroupResponse UpdateDeviceDatasourceInstanceGroupById(ctx, deviceId, deviceDsId, id, body)
update device datasource instance group 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **deviceDsId** | **int32**|  | 
  **id** | **int32**|  | 
  **body** | [**DeviceDataSourceInstanceGroup**](DeviceDataSourceInstanceGroup.md)|  | 

### Return type

[**DeviceDatasourceInstanceGroupResponse**](DeviceDatasourceInstanceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeviceGroupById**
> DeviceGroupResponse UpdateDeviceGroupById(ctx, id, body)
update device group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**DeviceGroup**](DeviceGroup.md)|  | 

### Return type

[**DeviceGroupResponse**](DeviceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeviceGroupDatasourceAlertSetting**
> DeviceGroupDatasourceAlertConfigResponse UpdateDeviceGroupDatasourceAlertSetting(ctx, deviceGroupId, dsId, body)
update device group datasource alert setting 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceGroupId** | **int32**|  | 
  **dsId** | **int32**|  | 
  **body** | [**DeviceGroupDataSourceAlertConfig**](DeviceGroupDataSourceAlertConfig.md)|  | 

### Return type

[**DeviceGroupDatasourceAlertConfigResponse**](DeviceGroupDatasourceAlertConfigResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeviceGroupPropertyByName**
> PropertyResponse UpdateDeviceGroupPropertyByName(ctx, gid, name, body)
update device group property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gid** | **int32**|  | 
  **name** | **string**|  | 
  **body** | [**EntityProperty**](EntityProperty.md)|  | 

### Return type

[**PropertyResponse**](PropertyResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDevicePropertyByName**
> PropertyResponse UpdateDevicePropertyByName(ctx, deviceId, name, body)
update device  property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **name** | **string**|  | 
  **body** | [**EntityProperty**](EntityProperty.md)|  | 

### Return type

[**PropertyResponse**](PropertyResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEscalationChainById**
> EscalationChainResponse UpdateEscalationChainById(ctx, id, optional)
update escalation chain by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***UpdateEscalationChainByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateEscalationChainByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EscalatingChain**](EscalatingChain.md)|  | 

### Return type

[**EscalationChainResponse**](EscalationChainResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOpsNoteById**
> OpsNoteResponse UpdateOpsNoteById(ctx, id, body)
update opsnote by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**RestOpsNoteV1**](RestOpsNoteV1.md)|  | 

### Return type

[**OpsNoteResponse**](OpsNoteResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReportById**
> ReportResponse UpdateReportById(ctx, id, optional)
update report by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***UpdateReportByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateReportByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ReportBase**](ReportBase.md)|  | 

### Return type

[**ReportResponse**](ReportResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleById**
> RoleResponse UpdateRoleById(ctx, id, body)
update role



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**Role**](Role.md)|  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSDTById**
> SdtResponse UpdateSDTById(ctx, id, body)
update SDT by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**SdtBase**](SdtBase.md)|  | 

### Return type

[**SdtResponse**](SDTResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceById**
> ServiceResponse UpdateServiceById(ctx, body, id)
update service by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceBaseVersion1**](ServiceBaseVersion1.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceGroupById**
> ServiceGroupResponse UpdateServiceGroupById(ctx, id, body)
update service group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**ServiceGroup**](ServiceGroup.md)|  | 

### Return type

[**ServiceGroupResponse**](ServiceGroupResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWidgetById**
> WidgetResponse UpdateWidgetById(ctx, id, body)
update widget by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **body** | [**WidgetBaseVersion1**](WidgetBaseVersion1.md)|  | 

### Return type

[**WidgetResponse**](WidgetResponse.md)

### Authorization

[Basic](../README.md#Basic), [LMv1](../README.md#LMv1)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

