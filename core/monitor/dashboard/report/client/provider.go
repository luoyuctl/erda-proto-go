// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: report.proto

package client

import (
	fmt "fmt"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/dashboard/report/pb"
	grpc1 "google.golang.org/grpc"
	reflect "reflect"
	strings "strings"
)

var dependencies = []string{
	"grpc-client@erda.core.monitor.dashboard.report",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                      = reflect.TypeOf((*Client)(nil)).Elem()
	dashboardReportServiceClientType = reflect.TypeOf((*pb.DashboardReportServiceClient)(nil)).Elem()
	dashboardReportServiceServerType = reflect.TypeOf((*pb.DashboardReportServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.monitor.dashboard.report-client":
		return p.client
	case "erda.core.monitor.dashboard.report.DashboardReportService":
		return &dashboardReportServiceWrapper{client: p.client.DashboardReportService(), opts: opts}
	case "erda.core.monitor.dashboard.report.DashboardReportService.client":
		return p.client.DashboardReportService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case dashboardReportServiceClientType:
		return p.client.DashboardReportService()
	case dashboardReportServiceServerType:
		return &dashboardReportServiceWrapper{client: p.client.DashboardReportService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.monitor.dashboard.report-client", &servicehub.Spec{
		Services: []string{
			"erda.core.monitor.dashboard.report.DashboardReportService",
			"erda.core.monitor.dashboard.report-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			dashboardReportServiceClientType,
			// server types
			dashboardReportServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
