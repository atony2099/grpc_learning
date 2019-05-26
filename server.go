package main

import (
	"fmt"
	"net"

	pb "example/proto" // 引入编译生成的包

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	logrusLogger *logrus.Logger
	customFunc   grpc_logrus.CodeToLevel
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	var inter grpc.UnaryServerInterceptor
	inter = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("hello")
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	opts = append(opts, grpc.UnaryInterceptor(inter))

	// 实例化grpc Server
	s := grpc.NewServer(opts...)

	//
	pb.RegisterHelloServer(s, HelloService)

	grpclog.Println("Listen on " + Address)
	s.Serve(listen)

	// golang rpc;

}
