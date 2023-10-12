/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package generator

import (
	"bytes"
	"strings"
	"testing"
)

import (
	"github.com/emicklei/proto"

	"github.com/stretchr/testify/assert"
)

func init() {
	protoFile := bytes.NewBufferString(testProto)
	parser := proto.NewParser(protoFile)
	p, _ := parser.Parse()

	g := NewGenerator(Context{
		GoModuleName: "protoc-gen-triple",
		Src:          "../../.",
	})
	data, _ = g.parseProtoToTriple(p)
}

var data TripleGo

//func TestPreamble(t *testing.T) {
//	err := TplPreamble.Execute(os.Stdout, data)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func TestPackage(t *testing.T) {
//	err := TplPackage.Execute(os.Stdout, data)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func TestImport(t *testing.T) {
//	err := TplImport.Execute(os.Stdout, data)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func TestTotalTpl(t *testing.T) {
//	err := TplTotal.Execute(os.Stdout, data)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func TestClientInterfaceTemplate(t *testing.T) {
//	err := TplClientInterface.Execute(os.Stdout, data)
//	if err != nil {
//		t.Fatalf("Failed to execute template: %v", err)
//	}
//}
//
//func TestClientInterfaceImplTpl(t *testing.T) {
//	err := TplClientInterfaceImpl.Execute(os.Stdout, data)
//	if err != nil {
//		t.Fatalf("Failed to execute template: %v", err)
//	}
//}
//
//func TestClientImplTpl(t *testing.T) {
//	err := TplClientImpl.Execute(os.Stdout, data)
//	if err != nil {
//		t.Fatalf("Failed to execute template: %v", err)
//	}
//}
//
//func TestMethodInfoTpl(t *testing.T) {
//	err := TplMethodInfo.Execute(os.Stdout, data)
//	if err != nil {
//		t.Fatalf("Failed to execute template: %v", err)
//	}
//}
//
//func TestHandlerTpl(t *testing.T) {
//	err := TplHandler.Execute(os.Stdout, data)
//	if err != nil {
//		t.Fatalf("Failed to execute template: %v", err)
//	}
//}
//
//func TestServerImplTpl(t *testing.T) {
//	err := TplServerImpl.Execute(os.Stdout, data)
//	if err != nil {
//		t.Fatalf("Failed to execute template: %v", err)
//	}
//}
//
//func TestServiceInfoImplTpl(t *testing.T) {
//	err := TplServerInfo.Execute(os.Stdout, data)
//	if err != nil {
//		t.Fatalf("Failed to execute template: %v", err)
//	}
//}
//
//func TestUnImplTpl(t *testing.T) {
//	err := TplUnImpl.Execute(os.Stdout, data)
//	if err != nil {
//		t.Fatalf("Failed to execute template: %v", err)
//	}
//}

func TestAll(t *testing.T) {
	var builder strings.Builder
	for _, tpl := range Tpls {
		err := tpl.Execute(&builder, data)
		if err != nil {
			t.Fatalf("Failed to execute template: %v", err)
		}
	}
	assert.Equal(t, testTripleGo, builder.String())
}

const testProto = `syntax = "proto3";

package greet;

option go_package = "/proto;proto";

message GreetRequest {
  string name = 1;
}

message GreetResponse {
  string greeting = 1;
}

message GreetStreamRequest {
  string name = 1;
}

message GreetStreamResponse {
  string greeting = 1;
}

message GreetClientStreamRequest {
  string name = 1;
}

message GreetClientStreamResponse {
  string greeting = 1;
}

message GreetServerStreamRequest {
  string name = 1;
}

message GreetServerStreamResponse {
  string greeting = 1;
}

service GreetService {
  rpc Greet(GreetRequest) returns (GreetResponse) {}
  rpc GreetStream(stream GreetStreamRequest) returns (stream GreetStreamResponse) {}
  rpc GreetClientStream(stream GreetClientStreamRequest) returns (GreetClientStreamResponse) {}
  rpc GreetServerStream(GreetServerStreamRequest) returns (stream GreetServerStreamResponse) {}
}`

const testTripleGo = `// Code generated by protoc-gen-triple. DO NOT EDIT.
//
// Source: 
package greettriple

import (
	context "context"
	errors "errors"
	http "net/http"

	client "dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/config"
	proto "protoc-gen-triple/proto"
	triple_protocol "dubbo.apache.org/dubbo-go/v3/protocol/triple/triple_protocol"
	"dubbo.apache.org/dubbo-go/v3/provider"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = triple_protocol.IsAtLeastVersion0_1_0

const (
	// GreetServiceName is the fully-qualified name of the GreetService service.
	GreetServiceName = "greet.GreetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreetServiceGreetProcedure is the fully-qualified name of the GreetService's Greet RPC.
	GreetServiceGreetProcedure = "/greet.GreetService/Greet"
	// GreetServiceGreetStreamProcedure is the fully-qualified name of the GreetService's GreetStream RPC.
	GreetServiceGreetStreamProcedure = "/greet.GreetService/GreetStream"
	// GreetServiceGreetClientStreamProcedure is the fully-qualified name of the GreetService's GreetClientStream RPC.
	GreetServiceGreetClientStreamProcedure = "/greet.GreetService/GreetClientStream"
	// GreetServiceGreetServerStreamProcedure is the fully-qualified name of the GreetService's GreetServerStream RPC.
	GreetServiceGreetServerStreamProcedure = "/greet.GreetService/GreetServerStream"
)

//GreetServiceClient is a client for the greet.GreetService service.
type GreetServiceClient interface {
	
		Greet(ctx context.Context, req *proto.GreetRequest, opt ...client.CallOption) (*proto.GreetResponse, error)
	
		GreetStream(ctx context.Context, opt ...client.CallOption) (GreetService_GreetStreamClient, error)
	
		GreetClientStream(ctx context.Context, opt ...client.CallOption) (GreetService_GreetClientStreamClient, error)
	
		GreetServerStream(ctx context.Context, req *proto.GreetServerStreamRequest, opt ...client.CallOption) (GreetService_GreetServerStreamClient, error)
	
}

// NewGreetServiceClient constructs a client for the greettriple.GreetService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGreetServiceClient(cli *client.Client) (GreetServiceClient, error) {
	if err := cli.Init(&GreetService_ClientInfo); err != nil {
		return nil, err
	}
	return &GreetServiceClientImpl{
		cli: cli,
	}, nil
}

func SetConsumerService(srv common.RPCService) {
	config.SetClientInfoService(&GreetService_ClientInfo, srv)
}

// GreetServiceClientImpl implements GreetServiceClient.
type GreetServiceClientImpl struct {
	cli *client.Client
}

func (c *GreetServiceClientImpl) Greet(ctx context.Context, req *proto.GreetRequest, opts ...client.CallOption) (*proto.GreetResponse, error) {
	triReq := triple_protocol.NewRequest(req)
	resp := new(proto.GreetResponse)
	triResp := triple_protocol.NewResponse(resp)
	if err := c.cli.CallUnary(ctx, triReq, triResp, "greet.GreetService", "Greet", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *GreetServiceClientImpl) GreetStream(ctx context.Context, opts ...client.CallOption)(GreetService_GreetStreamClient,error) {
	stream, err := c.cli.CallBidiStream(ctx, "greet.GreetService", "GreetStream", opts...)
	if err != nil {
		return nil, err
	}
	rawStream := stream.(*triple_protocol.BidiStreamForClient)
	return &greetServiceGreetStreamClient{rawStream}, nil
}

func (c *GreetServiceClientImpl) GreetClientStream(ctx context.Context, opts ...client.CallOption) (GreetService_GreetClientStreamClient, error) {
	stream, err := c.cli.CallClientStream(ctx, "greet.GreetService", "GreetClientStream", opts...)
	if err != nil {
		return nil, err
	}
	rawStream := stream.(*triple_protocol.ClientStreamForClient)
	return &greetServiceGreetClientStreamClient{rawStream}, nil
}

func (c *GreetServiceClientImpl) GreetServerStream(ctx context.Context, req *proto.GreetServerStreamRequest, opts ...client.CallOption) (GreetService_GreetServerStreamClient, error) {
	triReq := triple_protocol.NewRequest(req)
	stream, err := c.cli.CallServerStream(ctx, triReq, "greet.GreetService", "GreetServerStream", opts...)
	if err != nil {
		return nil, err
	}
	rawStream := stream.(*triple_protocol.ServerStreamForClient)
	return &greetServiceGreetServerStreamClient{rawStream}, nil
}


type GreetService_GreetClient interface {
	Recv() bool
	ResponseHeader() http.Header
	ResponseTrailer() http.Header
	Msg() *proto.GreetResponse
	Err() error
	Conn() (triple_protocol.StreamingClientConn, error)
	Close() error
}

type greetServiceGreetClient struct {
	*triple_protocol.ServerStreamForClient
}

func (cli *greetServiceGreetClient) Recv() bool {
	msg := new(proto.GreetServerStreamResponse)
	return cli.ServerStreamForClient.Receive(msg)
}

func (cli *greetServiceGreetClient) Msg() *proto.GreetResponse {
	msg := cli.ServerStreamForClient.Msg()
	if msg == nil {
		return new(proto.GreetResponse)
	}
	return msg.(*proto.GreetResponse)
}

func (cli *greetServiceGreetClient) Conn() (triple_protocol.StreamingClientConn, error) {
	return cli.ServerStreamForClient.Conn()
}
 
type GreetService_GreetStreamClient interface {
	Spec() triple_protocol.Spec
	Peer() triple_protocol.Peer
	Send(*proto.GreetStreamRequest) error
	RequestHeader() http.Header
	CloseRequest() error
	Recv() (*proto.GreetStreamResponse, error)
	ResponseHeader() http.Header
	ResponseTrailer() http.Header
	CloseResponse() error
}

type greetServiceGreetStreamClient struct {
	*triple_protocol.BidiStreamForClient
}

func (cli *greetServiceGreetStreamClient) Send(msg *proto.GreetStreamRequest) error {
	return cli.BidiStreamForClient.Send(msg)
}

func (cli *greetServiceGreetStreamClient) Recv() (*proto.GreetStreamResponse, error) {
	msg := new(proto.GreetStreamResponse)
	if err := cli.BidiStreamForClient.Receive(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

type GreetService_GreetClientStreamClient interface {
	Spec() triple_protocol.Spec
	Peer() triple_protocol.Peer
	Send(*proto.GreetClientStreamRequest) error
	RequestHeader() http.Header
	CloseAndRecv() (*proto.GreetClientStreamResponse, error)
	Conn() (triple_protocol.StreamingClientConn, error)
}

type greetServiceGreetClientStreamClient struct {
	*triple_protocol.ClientStreamForClient
}

func (cli *greetServiceGreetClientStreamClient) Send(msg *proto.GreetClientStreamRequest) error {
	return cli.ClientStreamForClient.Send(msg)
}

func (cli *greetServiceGreetClientStreamClient) CloseAndRecv() (*proto.GreetClientStreamResponse, error) {
	msg := new(proto.GreetClientStreamResponse)
	resp := triple_protocol.NewResponse(msg)
	if err := cli.ClientStreamForClient.CloseAndReceive(resp); err != nil {
		return nil, err
	}
	return msg, nil
}

func (cli *greetServiceGreetClientStreamClient) Conn() (triple_protocol.StreamingClientConn, error) {
	return cli.ClientStreamForClient.Conn()
}

type GreetService_GreetServerStreamClient interface {
	Recv() bool
	ResponseHeader() http.Header
	ResponseTrailer() http.Header
	Msg() *proto.GreetServerStreamResponse
	Err() error
	Conn() (triple_protocol.StreamingClientConn, error)
	Close() error
}

type greetServiceGreetServerStreamClient struct {
	*triple_protocol.ServerStreamForClient
}

func (cli *greetServiceGreetServerStreamClient) Recv() bool {
	msg := new(proto.GreetServerStreamResponse)
	return cli.ServerStreamForClient.Receive(msg)
}

func (cli *greetServiceGreetServerStreamClient) Msg() *proto.GreetServerStreamResponse {
	msg := cli.ServerStreamForClient.Msg()
	if msg == nil {
		return new(proto.GreetServerStreamResponse)
	}
	return msg.(*proto.GreetServerStreamResponse)
}

func (cli *greetServiceGreetServerStreamClient) Conn() (triple_protocol.StreamingClientConn, error) {
	return cli.ServerStreamForClient.Conn()
}


var GreetService_ClientInfo = client.ClientInfo{
	InterfaceName : "greettriple.GreetService",
	MethodNames :   []string{"Greet","GreetStream","GreetClientStream","GreetServerStream"},
	ClientInjectFunc : func(dubboCliRaw interface{}, cli *client.Client) {
		dubboCli := dubboCliRaw.(GreetServiceClientImpl)
		dubboCli.cli = cli
	},
}

// GreetServiceHandler is an implementation of the greet.GreetService service.
type GreetServiceHandler interface {
	
		Greet(context.Context, *proto.GreetRequest) (*proto.GreetResponse, error)
	
		GreetStream(context.Context, GreetService_GreetStreamServer) error
	
		GreetClientStream(context.Context, GreetService_GreetClientStreamServer) (*proto.GreetClientStreamResponse, error)
	
		GreetServerStream(context.Context, *proto.GreetServerStreamRequest,GreetService_GreetServerStreamServer) error
	
}

func ProvideGreetServiceHandler(pro *provider.Provider, hdlr GreetServiceHandler) error {
	return pro.Provide(hdlr, &GreetService_ServiceInfo)
}

type GreetService_GreetServer interface {
	Send(*proto.GreetServerStreamResponse) error
	ResponseHeader() http.Header
	ResponseTrailer() http.Header
	Conn() triple_protocol.StreamingHandlerConn
}

type greetServiceGreetServer struct {
	*triple_protocol.ServerStream
}

func (g *greetServiceGreetServer) Send(msg *proto.GreetResponse) error {
	return g.ServerStream.Send(msg)
}

type GreetService_GreetStreamServer interface {
	Send(*proto.GreetStreamResponse) error
	Recv() (*proto.GreetStreamRequest, error)
	Spec() triple_protocol.Spec
	Peer() triple_protocol.Peer
	RequestHeader() http.Header
	ResponseHeader() http.Header
	ResponseTrailer() http.Header
	Conn() triple_protocol.StreamingHandlerConn
}

type greetServiceGreetStreamServer struct {
	*triple_protocol.BidiStream
}

func (srv *greetServiceGreetStreamServer) Send(msg *proto.GreetStreamResponse) error {
	return srv.BidiStream.Send(msg)
}

func (srv greetServiceGreetStreamServer) Recv() (*proto.GreetStreamRequest, error) {
	msg := new(proto.GreetStreamRequest)
	if err := srv.BidiStream.Receive(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

type GreetService_GreetClientStreamServer interface {
	Spec() triple_protocol.Spec
	Peer() triple_protocol.Peer
	Recv() bool
	RequestHeader() http.Header
	Msg() *proto.GreetClientStreamRequest
	Err() error
	Conn() triple_protocol.StreamingHandlerConn
}

type greetServiceGreetClientStreamServer struct {
	*triple_protocol.ClientStream
}

func (srv *greetServiceGreetClientStreamServer) Recv() bool {
	msg := new(proto.GreetClientStreamRequest)
	return srv.ClientStream.Receive(msg)
}

func (srv *greetServiceGreetClientStreamServer) Msg() *proto.GreetClientStreamRequest {
	msgRaw := srv.ClientStream.Msg()
	if msgRaw == nil {
		return new(proto.GreetClientStreamRequest)
	}
	return msgRaw.(*proto.GreetClientStreamRequest)
}

type GreetService_GreetServerStreamServer interface {
	Send(*proto.GreetServerStreamResponse) error
	ResponseHeader() http.Header
	ResponseTrailer() http.Header
	Conn() triple_protocol.StreamingHandlerConn
}

type greetServiceGreetServerStreamServer struct {
	*triple_protocol.ServerStream
}

func (g *greetServiceGreetServerStreamServer) Send(msg *proto.GreetServerStreamResponse) error {
	return g.ServerStream.Send(msg)
}
var GreetService_ServiceInfo = provider.ServiceInfo{
	InterfaceName: "greet.GreetService",
	ServiceType:   (*GreetServiceHandler)(nil),
	Methods: []provider.MethodInfo{
		{
			Name : "Greet",
			Type : constant.CallUnary,
			ReqInitFunc : func() interface{} {
				return new(proto.GreetRequest)
			},
			MethodFunc : func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*proto.GreetRequest)
				res, err := handler.(GreetServiceHandler).Greet(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
		{
			Name : "GreetStream",
			Type : constant.CallBidiStream,
			StreamInitFunc : func(baseStream interface{}) interface{} {
				return &greetServiceGreetStreamServer{baseStream.(*triple_protocol.BidiStream)}
			},
			MethodFunc : func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				stream := args[0].(GreetService_GreetStreamServer)
				if err := handler.(GreetServiceHandler).GreetStream(ctx, stream); err != nil {
					return nil, err
				}
				return nil, nil
			},
		},
		{
			Name : "GreetClientStream",
			Type : constant.CallClientStream,
			StreamInitFunc: func(baseStream interface{}) interface{} {
				return &greetServiceGreetClientStreamServer{baseStream.(*triple_protocol.ClientStream)}
			},
			MethodFunc : func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				stream := args[0].(GreetService_GreetClientStreamServer)
				res, err := handler.(GreetServiceHandler).GreetClientStream(ctx, stream)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
		{
			Name : "GreetServerStream",
			Type : constant.CallServerStream,
			ReqInitFunc : func() interface{} {
				return new(proto.GreetServerStreamRequest)
			},
			StreamInitFunc : func(baseStream interface{}) interface{} {
				return &greetServiceGreetServerStreamServer{baseStream.(*triple_protocol.ServerStream)}
			},
			MethodFunc : func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*proto.GreetServerStreamRequest)
				stream := args[1].(GreetService_GreetServerStreamServer)
				if err := handler.(GreetServiceHandler).GreetServerStream(ctx, req, stream); err != nil {
					return nil, err
				}
				return nil, nil
			},
		},
	},
}
// UnimplementedGreetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGreetServiceHandler struct{}

func (UnimplementedGreetServiceHandler) Greet(context.Context, *proto.GreetRequest) (*proto.GreetResponse, error) {
	return nil, triple_protocol.NewError(triple_protocol.CodeUnimplemented, errors.New("greettriple.GreetService.Greet is not implemented"))
}

func (UnimplementedGreetServiceHandler) GreetStream(context.Context, *triple_protocol.BidiStream) error {
	return triple_protocol.NewError(triple_protocol.CodeUnimplemented, errors.New("greettriple.GreetService.GreetStream is not implemented"))
}

func (UnimplementedGreetServiceHandler) GreetClientStream(context.Context, *triple_protocol.ClientStream) (*triple_protocol.Response, error) {
	return nil, triple_protocol.NewError(triple_protocol.CodeUnimplemented, errors.New("greettriple.GreetService.GreetClientStream is not implemented"))
}

func (UnimplementedGreetServiceHandler) GreetServerStream(context.Context, *triple_protocol.Request, *triple_protocol.ServerStream) error {
	return triple_protocol.NewError(triple_protocol.CodeUnimplemented, errors.New("greettriple.GreetService.GreetServerStream is not implemented"))
}

`