package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"gorm.io/plugin/opentelemetry/provider"
	"net"
	"tiktok/cmd/friend/dal"
	"tiktok/cmd/friend/rpc"
	"tiktok/kitex_gen/friend/friendservice"
	"tiktok/pkg/consts"
	"tiktok/pkg/mw"
)

func Init() {
	dal.Init()
	rpc.Init()
	// klog init
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.FriendServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.FriendServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	svr := friendservice.NewServer(new(FriendServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.FriendServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
