// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: menu.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	http1 "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// MenuServiceHandler is the server API for MenuService service.
type MenuServiceHandler interface {
	// GET /api/tmc/micro-service/menu/tenantGroup/{tenantGroup}
	GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error)
	// GET /api/tmc/micro-service/setting/tenantGroup/{tenantGroup}
	GetSetting(context.Context, *GetSettingRequest) (*GetSettingResponse, error)
}

// RegisterMenuServiceHandler register MenuServiceHandler to http.Router.
func RegisterMenuServiceHandler(r http.Router, srv MenuServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
	}

	add_GetMenu := func(method, path string, fn func(context.Context, *GetMenuRequest) (*GetMenuResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetMenuRequest))
		}
		var GetMenu_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetMenu_info = transport.NewServiceInfo("erda.msp.menu.MenuService", "GetMenu", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in GetMenuRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "tenantGroup":
							in.TenantGroup = val
						}
					}
				}
				ctx := context.WithValue(r.Context(), http.RequestContextKey, r)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetMenu_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetSetting := func(method, path string, fn func(context.Context, *GetSettingRequest) (*GetSettingResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetSettingRequest))
		}
		var GetSetting_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetSetting_info = transport.NewServiceInfo("erda.msp.menu.MenuService", "GetSetting", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in GetSettingRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "tenantGroup":
							in.TenantGroup = val
						}
					}
				}
				ctx := context.WithValue(r.Context(), http.RequestContextKey, r)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetSetting_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetMenu("GET", "/api/tmc/micro-service/menu/tenantGroup/{tenantGroup}", srv.GetMenu)
	add_GetSetting("GET", "/api/tmc/micro-service/setting/tenantGroup/{tenantGroup}", srv.GetSetting)
}
