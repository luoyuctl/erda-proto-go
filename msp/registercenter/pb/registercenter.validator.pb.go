// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registercenter.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListInterfaceRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	if this.TenantGroup == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantGroup", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantGroup))
	}
	return nil
}
func (this *ListInterfaceResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetHTTPServicesRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	if this.ClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.ClusterName))
	}
	if this.TenantGroup == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantGroup", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantGroup))
	}
	return nil
}
func (this *GetHTTPServicesResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *EnableHTTPServiceRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	if this.ClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.ClusterName))
	}
	if this.TenantGroup == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantGroup", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantGroup))
	}
	if nil == this.Service {
		return github_com_mwitkow_go_proto_validators.FieldError("Service", fmt.Errorf("message must exist"))
	}
	if this.Service != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Service); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Service", err)
		}
	}
	return nil
}
func (this *EnableHTTPServiceResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *EnableHTTPService) Validate() error {
	return nil
}
func (this *GetRouteRuleRequest) Validate() error {
	if this.InterfaceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InterfaceName", fmt.Errorf(`value '%v' must not be an empty string`, this.InterfaceName))
	}
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *GetRouteRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateRouteRuleRequest) Validate() error {
	if this.InterfaceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InterfaceName", fmt.Errorf(`value '%v' must not be an empty string`, this.InterfaceName))
	}
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	if nil == this.Rule {
		return github_com_mwitkow_go_proto_validators.FieldError("Rule", fmt.Errorf("message must exist"))
	}
	if this.Rule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Rule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Rule", err)
		}
	}
	return nil
}
func (this *CreateRouteRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DeleteRouteRuleRequest) Validate() error {
	if this.InterfaceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InterfaceName", fmt.Errorf(`value '%v' must not be an empty string`, this.InterfaceName))
	}
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *DeleteRouteRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CetHostRuleRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *CetHostRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateHostRuleRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	if nil == this.Rules {
		return github_com_mwitkow_go_proto_validators.FieldError("Rules", fmt.Errorf("message must exist"))
	}
	if this.Rules != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Rules); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
		}
	}
	return nil
}
func (this *CreateHostRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DeleteHostRuleRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *DeleteHostRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetHostRuntimeRuleRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	if this.Host == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Host", fmt.Errorf(`value '%v' must not be an empty string`, this.Host))
	}
	return nil
}
func (this *GetHostRuntimeRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateHostRuntimeRuleRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	if this.Host == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Host", fmt.Errorf(`value '%v' must not be an empty string`, this.Host))
	}
	if nil == this.Rules {
		return github_com_mwitkow_go_proto_validators.FieldError("Rules", fmt.Errorf("message must exist"))
	}
	if this.Rules != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Rules); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
		}
	}
	return nil
}
func (this *CreateHostRuntimeRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetAllHostRuntimeRulesRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *GetAllHostRuntimeRulesResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetDubboInterfaceTimeRequest) Validate() error {
	if this.InterfaceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InterfaceName", fmt.Errorf(`value '%v' must not be an empty string`, this.InterfaceName))
	}
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *GetDubboInterfaceTimeResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DubboInterfaceTime) Validate() error {
	return nil
}
func (this *GetDubboInterfaceQPSRequest) Validate() error {
	if this.InterfaceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InterfaceName", fmt.Errorf(`value '%v' must not be an empty string`, this.InterfaceName))
	}
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *GetDubboInterfaceQPSResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetDubboInterfaceFailedRequest) Validate() error {
	if this.InterfaceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InterfaceName", fmt.Errorf(`value '%v' must not be an empty string`, this.InterfaceName))
	}
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *GetDubboInterfaceFailedResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetDubboInterfaceAvgTimeRequest) Validate() error {
	if this.InterfaceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InterfaceName", fmt.Errorf(`value '%v' must not be an empty string`, this.InterfaceName))
	}
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *GetDubboInterfaceAvgTimeResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Interface) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *InterfaceOwner) Validate() error {
	return nil
}
func (this *RequestRule) Validate() error {
	return nil
}
func (this *HostRules) Validate() error {
	if this.Rule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Rule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Rule", err)
		}
	}
	return nil
}
func (this *HostRoute) Validate() error {
	return nil
}
func (this *HostRuntimeRules) Validate() error {
	for _, item := range this.Rule {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rule", err)
			}
		}
	}
	return nil
}
func (this *HostRuntimeRule) Validate() error {
	return nil
}
func (this *HostRuntimeInterfaces) Validate() error {
	for _, item := range this.Node {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
			}
		}
	}
	return nil
}
func (this *HTTPServices) Validate() error {
	for _, item := range this.ServiceList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ServiceList", err)
			}
		}
	}
	return nil
}
func (this *HTTPService) Validate() error {
	for _, item := range this.HttpServiceDto {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HttpServiceDto", err)
			}
		}
	}
	return nil
}
func (this *HTTPServiceItem) Validate() error {
	return nil
}
func (this *DubboInterface) Validate() error {
	for _, item := range this.Results {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Results", err)
			}
		}
	}
	return nil
}
func (this *DubboInterfaceResult) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *DubboMointorMap) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *DubboMointor) Validate() error {
	return nil
}
