// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alert.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *QueryCustomizeMetricRequest) Validate() error {
	return nil
}
func (this *QueryCustomizeMetricResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CustomizeMetrics) Validate() error {
	for _, item := range this.Metrics {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Metrics", err)
			}
		}
	}
	for _, item := range this.FunctionOperators {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FunctionOperators", err)
			}
		}
	}
	for _, item := range this.FilterOperators {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FilterOperators", err)
			}
		}
	}
	for _, item := range this.Aggregator {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Aggregator", err)
			}
		}
	}
	return nil
}
func (this *MetricMeta) Validate() error {
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	for _, item := range this.Fields {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Fields", err)
			}
		}
	}
	for _, item := range this.Tags {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Tags", err)
			}
		}
	}
	return nil
}
func (this *DisplayKey) Validate() error {
	return nil
}
func (this *FieldMeta) Validate() error {
	if this.Field != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Field); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Field", err)
		}
	}
	return nil
}
func (this *TagMeta) Validate() error {
	if this.Tag != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Tag); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Tag", err)
		}
	}
	return nil
}
func (this *Operator) Validate() error {
	return nil
}
func (this *QueryCustomizeNotifyTargetRequest) Validate() error {
	return nil
}
func (this *QueryCustomizeNotifyTargetResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryCustomizeNotifyTargetData) Validate() error {
	for _, item := range this.Targets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Targets", err)
			}
		}
	}
	return nil
}
func (this *QueryOrgCustomizeNotifyTargetRequest) Validate() error {
	return nil
}
func (this *QueryOrgCustomizeNotifyTargetResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryCustomizeAlertRequest) Validate() error {
	if !(this.PageNo > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNo", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNo))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.PageSize < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be less than '101'`, this.PageSize))
	}
	return nil
}
func (this *QueryCustomizeAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryCustomizeAlertData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *CustomizeAlertOverview) Validate() error {
	return nil
}
func (this *GetCustomizeAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetCustomizeAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CustomizeAlertDetail) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	for _, item := range this.Notifies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notifies", err)
			}
		}
	}
	for _, item := range this.Lang {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Lang", err)
			}
		}
	}
	return nil
}
func (this *CustomizeAlertRule) Validate() error {
	for _, item := range this.Functions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Functions", err)
			}
		}
	}
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CustomizeAlertRuleFunction) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *CustomizeAlertRuleFilter) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *CustomizeAlertNotifyTemplates) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *LanguageCode) Validate() error {
	return nil
}
func (this *GetCustomizeAlertDetailRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetCustomizeAlertDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateCustomizeAlertRequest) Validate() error {
	if this.Alert != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alert); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alert", err)
		}
	}
	return nil
}
func (this *CreateCustomizeAlertResponse) Validate() error {
	return nil
}
func (this *UpdateCustomizeAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateCustomizeAlertResponse) Validate() error {
	return nil
}
func (this *UpdateCustomizeAlertEnableRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateCustomizeAlertEnableResponse) Validate() error {
	return nil
}
func (this *DeleteCustomizeAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *DeleteCustomizeAlertResponse) Validate() error {
	return nil
}
func (this *QueryOrgCustomizeMetricRequest) Validate() error {
	return nil
}
func (this *QueryOrgCustomizeMetricResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryOrgCustomizeAlertsRequest) Validate() error {
	if !(this.PageNo > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNo", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNo))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.PageSize < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be less than '101'`, this.PageSize))
	}
	return nil
}
func (this *QueryOrgCustomizeAlertsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryOrgCustomizeAlertsData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *GetOrgCustomizeAlertDetailRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetOrgCustomizeAlertDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateOrgCustomizeAlertRequest) Validate() error {
	if this.Alert != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alert); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alert", err)
		}
	}
	return nil
}
func (this *CreateOrgCustomizeAlertResponse) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateOrgCustomizeAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	if this.NewAlert != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NewAlert); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NewAlert", err)
		}
	}
	return nil
}
func (this *UpdateOrgCustomizeAlertResponse) Validate() error {
	return nil
}
func (this *UpdateOrgCustomizeAlertEnableRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateOrgCustomizeAlertEnableResponse) Validate() error {
	return nil
}
func (this *DeleteOrgCustomizeAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *DeleteOrgCustomizeAlertResponse) Validate() error {
	return nil
}
func (this *QueryDashboardByAlertRequest) Validate() error {
	if this.Alert != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alert); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alert", err)
		}
	}
	return nil
}
func (this *QueryDashboardByAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *View) Validate() error {
	if this.StaticData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StaticData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StaticData", err)
		}
	}
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	if this.Api != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Api); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Api", err)
		}
	}
	if this.Controls != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Controls); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Controls", err)
		}
	}
	return nil
}
func (this *Config) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	if this.DataSourceConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DataSourceConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DataSourceConfig", err)
		}
	}
	if this.Option != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Option); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Option", err)
		}
	}
	return nil
}
func (this *API) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *QueryOrgDashboardByAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryAlertRuleRequest) Validate() error {
	return nil
}
func (this *QueryAlertRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *AlertTypeRuleResp) Validate() error {
	for _, item := range this.AlertTypeRules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AlertTypeRules", err)
			}
		}
	}
	for _, item := range this.Operators {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Operators", err)
			}
		}
	}
	for _, item := range this.Aggregator {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Aggregator", err)
			}
		}
	}
	for _, item := range this.Silence {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Silence", err)
			}
		}
	}
	return nil
}
func (this *AlertTypeRule) Validate() error {
	if this.AlertType != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AlertType); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AlertType", err)
		}
	}
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	return nil
}
func (this *AlertRule) Validate() error {
	if this.AlertIndex != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AlertIndex); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AlertIndex", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Functions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Functions", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AlertRuleFunction) Validate() error {
	if this.Field != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Field); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Field", err)
		}
	}
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *NotifySilence) Validate() error {
	if this.Unit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Unit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Unit", err)
		}
	}
	return nil
}
func (this *QueryAlertRequest) Validate() error {
	if !(this.PageNo > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNo", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNo))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.PageSize < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be less than '101'`, this.PageSize))
	}
	return nil
}
func (this *QueryAlertsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryAlertsData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *Alert) Validate() error {
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	for _, item := range this.Notifies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notifies", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AlertExpression) Validate() error {
	for _, item := range this.Functions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Functions", err)
			}
		}
	}
	return nil
}
func (this *AlertExpression_AlertExpressionFunction) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *AlertNotify) Validate() error {
	if this.NotifyGroup != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NotifyGroup); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NotifyGroup", err)
		}
	}
	if this.Silence != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Silence); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Silence", err)
		}
	}
	return nil
}
func (this *NotifyGroup) Validate() error {
	for _, item := range this.Targets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Targets", err)
			}
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	return nil
}
func (this *NotifyTarget) Validate() error {
	for _, item := range this.Values {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Values", err)
			}
		}
	}
	return nil
}
func (this *Target) Validate() error {
	return nil
}
func (this *AlertNotifySilence) Validate() error {
	return nil
}
func (this *GetAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetAlertDetailRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetAlertDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateAlertResponse) Validate() error {
	return nil
}
func (this *UpdateAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateAlertResponse) Validate() error {
	return nil
}
func (this *UpdateAlertEnableRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateAlertEnableResponse) Validate() error {
	return nil
}
func (this *DeleteAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *DeleteAlertResponse) Validate() error {
	return nil
}
func (this *QueryOrgAlertRuleRequest) Validate() error {
	return nil
}
func (this *QueryOrgAlertRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryOrgAlertRequest) Validate() error {
	if !(this.PageNo > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNo", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNo))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.PageSize < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be less than '101'`, this.PageSize))
	}
	return nil
}
func (this *QueryOrgAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryOrgAlertData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *GetOrgAlertDetailRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetOrgAlertDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateOrgAlertResponse) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateOrgAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	if this.Alert != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alert); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alert", err)
		}
	}
	return nil
}
func (this *UpdateOrgAlertResponse) Validate() error {
	return nil
}
func (this *UpdateOrgAlertEnableRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateOrgAlertEnableResponse) Validate() error {
	return nil
}
func (this *DeleteOrgAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *DeleteOrgAlertResponse) Validate() error {
	return nil
}
func (this *GetAlertRecordAttrRequest) Validate() error {
	return nil
}
func (this *GetAlertRecordAttrResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *AlertRecordAttr) Validate() error {
	for _, item := range this.AlertState {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AlertState", err)
			}
		}
	}
	for _, item := range this.AlertType {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AlertType", err)
			}
		}
	}
	for _, item := range this.HandleState {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HandleState", err)
			}
		}
	}
	return nil
}
func (this *QueryAlertRecordRequest) Validate() error {
	if !(this.PageNo > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNo", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNo))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.PageSize < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be less than '101'`, this.PageSize))
	}
	return nil
}
func (this *QueryAlertRecordResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListResult) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *AlertRecord) Validate() error {
	return nil
}
func (this *GetAlertRecordRequest) Validate() error {
	return nil
}
func (this *GetAlertRecordResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryAlertHistoryRequest) Validate() error {
	if !(this.Start > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Start", fmt.Errorf(`value '%v' must be greater than '0'`, this.Start))
	}
	if !(this.End > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("End", fmt.Errorf(`value '%v' must be greater than '0'`, this.End))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	return nil
}
func (this *QueryAlertHistoryResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *AlertHistory) Validate() error {
	return nil
}
func (this *CreateAlertIssueRequest) Validate() error {
	if this.PlanStartedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlanStartedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlanStartedAt", err)
		}
	}
	if this.PlanFinishedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlanFinishedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlanFinishedAt", err)
		}
	}
	if this.IssueManHour != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IssueManHour); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IssueManHour", err)
		}
	}
	return nil
}
func (this *IssueManHour) Validate() error {
	return nil
}
func (this *CreateAlertIssueResponse) Validate() error {
	return nil
}
func (this *UpdateAlertIssueRequest) Validate() error {
	if this.PlanStartedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlanStartedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlanStartedAt", err)
		}
	}
	if this.PlanFinishedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlanFinishedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlanFinishedAt", err)
		}
	}
	if this.IssueManHour != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IssueManHour); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IssueManHour", err)
		}
	}
	return nil
}
func (this *UpdateAlertIssueResponse) Validate() error {
	return nil
}
func (this *GetOrgAlertRecordAttrRequest) Validate() error {
	return nil
}
func (this *GetOrgAlertRecordAttrResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryOrgAlertRecordRequest) Validate() error {
	if !(this.PageNo > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNo", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNo))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.PageSize < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be less than '101'`, this.PageSize))
	}
	return nil
}
func (this *QueryOrgHostsAlertRecordRequest) Validate() error {
	for _, item := range this.Clusters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Clusters", err)
			}
		}
	}
	if !(this.PageNo > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNo", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNo))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.PageSize < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be less than '101'`, this.PageSize))
	}
	return nil
}
func (this *QueryOrgHostsAlertRecordRequestClusterReq) Validate() error {
	return nil
}
func (this *QueryOrgAlertRecordResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetOrgAlertRecordRequest) Validate() error {
	return nil
}
func (this *GetOrgAlertRecordResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryOrgAlertHistoryRequest) Validate() error {
	if !(this.Start > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Start", fmt.Errorf(`value '%v' must be greater than '0'`, this.Start))
	}
	if !(this.End > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("End", fmt.Errorf(`value '%v' must be greater than '0'`, this.End))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	return nil
}
func (this *QueryOrgAlertHistoryResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CreateOrgAlertIssueRequest) Validate() error {
	if this.PlanStartedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlanStartedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlanStartedAt", err)
		}
	}
	if this.PlanFinishedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlanFinishedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlanFinishedAt", err)
		}
	}
	if this.IssueManHour != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IssueManHour); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IssueManHour", err)
		}
	}
	return nil
}
func (this *CreateOrgAlertIssueResponse) Validate() error {
	return nil
}
func (this *UpdateOrgAlertIssueRequest) Validate() error {
	if this.PlanStartedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlanStartedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlanStartedAt", err)
		}
	}
	if this.PlanFinishedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlanFinishedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlanFinishedAt", err)
		}
	}
	if this.IssueManHour != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IssueManHour); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IssueManHour", err)
		}
	}
	return nil
}
func (this *UpdateOrgAlertIssueResponse) Validate() error {
	return nil
}
