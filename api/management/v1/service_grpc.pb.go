// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: management/v1/service.proto

package managementv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ManagementService_AddAnnotation_FullMethodName         = "/management.v1.ManagementService/AddAnnotation"
	ManagementService_ListAgents_FullMethodName            = "/management.v1.ManagementService/ListAgents"
	ManagementService_ListAgentVersions_FullMethodName     = "/management.v1.ManagementService/ListAgentVersions"
	ManagementService_RegisterNode_FullMethodName          = "/management.v1.ManagementService/RegisterNode"
	ManagementService_UnregisterNode_FullMethodName        = "/management.v1.ManagementService/UnregisterNode"
	ManagementService_ListNodes_FullMethodName             = "/management.v1.ManagementService/ListNodes"
	ManagementService_GetNode_FullMethodName               = "/management.v1.ManagementService/GetNode"
	ManagementService_AddService_FullMethodName            = "/management.v1.ManagementService/AddService"
	ManagementService_ListServices_FullMethodName          = "/management.v1.ManagementService/ListServices"
	ManagementService_DiscoverRDS_FullMethodName           = "/management.v1.ManagementService/DiscoverRDS"
	ManagementService_DiscoverAzureDatabase_FullMethodName = "/management.v1.ManagementService/DiscoverAzureDatabase"
	ManagementService_AddAzureDatabase_FullMethodName      = "/management.v1.ManagementService/AddAzureDatabase"
	ManagementService_RemoveService_FullMethodName         = "/management.v1.ManagementService/RemoveService"
)

// ManagementServiceClient is the client API for ManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ManagementService provides public methods for managing and querying Services.
type ManagementServiceClient interface {
	// AddAnnotation adds an annotation.
	AddAnnotation(ctx context.Context, in *AddAnnotationRequest, opts ...grpc.CallOption) (*AddAnnotationResponse, error)
	// ListAgents returns a list of Agents filtered by service_id or node_id.
	ListAgents(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error)
	// ListAgentVersions returns a list of PMM Agent versions and their update severity.
	ListAgentVersions(ctx context.Context, in *ListAgentVersionsRequest, opts ...grpc.CallOption) (*ListAgentVersionsResponse, error)
	// RegisterNode registers a new Node and a pmm-agent.
	RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error)
	// UnregisterNode unregisters a Node, pmm-agent and removes the service account and its token.
	UnregisterNode(ctx context.Context, in *UnregisterNodeRequest, opts ...grpc.CallOption) (*UnregisterNodeResponse, error)
	// ListNode returns a list of nodes.
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
	// GetNode returns a single Node by ID.
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	// AddService adds a Service and starts several Agents.
	AddService(ctx context.Context, in *AddServiceRequest, opts ...grpc.CallOption) (*AddServiceResponse, error)
	// ListServices returns a list of Services with a rich set of properties.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	// DiscoverRDS discovers RDS instances.
	DiscoverRDS(ctx context.Context, in *DiscoverRDSRequest, opts ...grpc.CallOption) (*DiscoverRDSResponse, error)
	// DiscoverAzureDatabase discovers Azure Database for MySQL, MariaDB and PostgreSQL Server instances.
	DiscoverAzureDatabase(ctx context.Context, in *DiscoverAzureDatabaseRequest, opts ...grpc.CallOption) (*DiscoverAzureDatabaseResponse, error)
	// AddAzureDatabase adds Azure Database instance.
	AddAzureDatabase(ctx context.Context, in *AddAzureDatabaseRequest, opts ...grpc.CallOption) (*AddAzureDatabaseResponse, error)
	// RemoveService removes a Service along with its Agents.
	RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error)
}

type managementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagementServiceClient(cc grpc.ClientConnInterface) ManagementServiceClient {
	return &managementServiceClient{cc}
}

func (c *managementServiceClient) AddAnnotation(ctx context.Context, in *AddAnnotationRequest, opts ...grpc.CallOption) (*AddAnnotationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddAnnotationResponse)
	err := c.cc.Invoke(ctx, ManagementService_AddAnnotation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) ListAgents(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAgentsResponse)
	err := c.cc.Invoke(ctx, ManagementService_ListAgents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) ListAgentVersions(ctx context.Context, in *ListAgentVersionsRequest, opts ...grpc.CallOption) (*ListAgentVersionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAgentVersionsResponse)
	err := c.cc.Invoke(ctx, ManagementService_ListAgentVersions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterNodeResponse)
	err := c.cc.Invoke(ctx, ManagementService_RegisterNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) UnregisterNode(ctx context.Context, in *UnregisterNodeRequest, opts ...grpc.CallOption) (*UnregisterNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnregisterNodeResponse)
	err := c.cc.Invoke(ctx, ManagementService_UnregisterNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNodesResponse)
	err := c.cc.Invoke(ctx, ManagementService_ListNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNodeResponse)
	err := c.cc.Invoke(ctx, ManagementService_GetNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) AddService(ctx context.Context, in *AddServiceRequest, opts ...grpc.CallOption) (*AddServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddServiceResponse)
	err := c.cc.Invoke(ctx, ManagementService_AddService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, ManagementService_ListServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) DiscoverRDS(ctx context.Context, in *DiscoverRDSRequest, opts ...grpc.CallOption) (*DiscoverRDSResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DiscoverRDSResponse)
	err := c.cc.Invoke(ctx, ManagementService_DiscoverRDS_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) DiscoverAzureDatabase(ctx context.Context, in *DiscoverAzureDatabaseRequest, opts ...grpc.CallOption) (*DiscoverAzureDatabaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DiscoverAzureDatabaseResponse)
	err := c.cc.Invoke(ctx, ManagementService_DiscoverAzureDatabase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) AddAzureDatabase(ctx context.Context, in *AddAzureDatabaseRequest, opts ...grpc.CallOption) (*AddAzureDatabaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddAzureDatabaseResponse)
	err := c.cc.Invoke(ctx, ManagementService_AddAzureDatabase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveServiceResponse)
	err := c.cc.Invoke(ctx, ManagementService_RemoveService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServiceServer is the server API for ManagementService service.
// All implementations must embed UnimplementedManagementServiceServer
// for forward compatibility.
//
// ManagementService provides public methods for managing and querying Services.
type ManagementServiceServer interface {
	// AddAnnotation adds an annotation.
	AddAnnotation(context.Context, *AddAnnotationRequest) (*AddAnnotationResponse, error)
	// ListAgents returns a list of Agents filtered by service_id or node_id.
	ListAgents(context.Context, *ListAgentsRequest) (*ListAgentsResponse, error)
	// ListAgentVersions returns a list of PMM Agent versions and their update severity.
	ListAgentVersions(context.Context, *ListAgentVersionsRequest) (*ListAgentVersionsResponse, error)
	// RegisterNode registers a new Node and a pmm-agent.
	RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error)
	// UnregisterNode unregisters a Node, pmm-agent and removes the service account and its token.
	UnregisterNode(context.Context, *UnregisterNodeRequest) (*UnregisterNodeResponse, error)
	// ListNode returns a list of nodes.
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
	// GetNode returns a single Node by ID.
	GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	// AddService adds a Service and starts several Agents.
	AddService(context.Context, *AddServiceRequest) (*AddServiceResponse, error)
	// ListServices returns a list of Services with a rich set of properties.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	// DiscoverRDS discovers RDS instances.
	DiscoverRDS(context.Context, *DiscoverRDSRequest) (*DiscoverRDSResponse, error)
	// DiscoverAzureDatabase discovers Azure Database for MySQL, MariaDB and PostgreSQL Server instances.
	DiscoverAzureDatabase(context.Context, *DiscoverAzureDatabaseRequest) (*DiscoverAzureDatabaseResponse, error)
	// AddAzureDatabase adds Azure Database instance.
	AddAzureDatabase(context.Context, *AddAzureDatabaseRequest) (*AddAzureDatabaseResponse, error)
	// RemoveService removes a Service along with its Agents.
	RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error)
	mustEmbedUnimplementedManagementServiceServer()
}

// UnimplementedManagementServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedManagementServiceServer struct{}

func (UnimplementedManagementServiceServer) AddAnnotation(context.Context, *AddAnnotationRequest) (*AddAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAnnotation not implemented")
}

func (UnimplementedManagementServiceServer) ListAgents(context.Context, *ListAgentsRequest) (*ListAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgents not implemented")
}

func (UnimplementedManagementServiceServer) ListAgentVersions(context.Context, *ListAgentVersionsRequest) (*ListAgentVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgentVersions not implemented")
}

func (UnimplementedManagementServiceServer) RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}

func (UnimplementedManagementServiceServer) UnregisterNode(context.Context, *UnregisterNodeRequest) (*UnregisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterNode not implemented")
}

func (UnimplementedManagementServiceServer) ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}

func (UnimplementedManagementServiceServer) GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}

func (UnimplementedManagementServiceServer) AddService(context.Context, *AddServiceRequest) (*AddServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddService not implemented")
}

func (UnimplementedManagementServiceServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}

func (UnimplementedManagementServiceServer) DiscoverRDS(context.Context, *DiscoverRDSRequest) (*DiscoverRDSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverRDS not implemented")
}

func (UnimplementedManagementServiceServer) DiscoverAzureDatabase(context.Context, *DiscoverAzureDatabaseRequest) (*DiscoverAzureDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverAzureDatabase not implemented")
}

func (UnimplementedManagementServiceServer) AddAzureDatabase(context.Context, *AddAzureDatabaseRequest) (*AddAzureDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAzureDatabase not implemented")
}

func (UnimplementedManagementServiceServer) RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveService not implemented")
}
func (UnimplementedManagementServiceServer) mustEmbedUnimplementedManagementServiceServer() {}
func (UnimplementedManagementServiceServer) testEmbeddedByValue()                           {}

// UnsafeManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagementServiceServer will
// result in compilation errors.
type UnsafeManagementServiceServer interface {
	mustEmbedUnimplementedManagementServiceServer()
}

func RegisterManagementServiceServer(s grpc.ServiceRegistrar, srv ManagementServiceServer) {
	// If the following call pancis, it indicates UnimplementedManagementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ManagementService_ServiceDesc, srv)
}

func _ManagementService_AddAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).AddAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_AddAnnotation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).AddAnnotation(ctx, req.(*AddAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_ListAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).ListAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_ListAgents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).ListAgents(ctx, req.(*ListAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_ListAgentVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).ListAgentVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_ListAgentVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).ListAgentVersions(ctx, req.(*ListAgentVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_RegisterNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).RegisterNode(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_UnregisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).UnregisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_UnregisterNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).UnregisterNode(ctx, req.(*UnregisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_ListNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_GetNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_AddService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).AddService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_AddService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).AddService(ctx, req.(*AddServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_ListServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_DiscoverRDS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverRDSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).DiscoverRDS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_DiscoverRDS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).DiscoverRDS(ctx, req.(*DiscoverRDSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_DiscoverAzureDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverAzureDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).DiscoverAzureDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_DiscoverAzureDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).DiscoverAzureDatabase(ctx, req.(*DiscoverAzureDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_AddAzureDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAzureDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).AddAzureDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_AddAzureDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).AddAzureDatabase(ctx, req.(*AddAzureDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_RemoveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).RemoveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_RemoveService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).RemoveService(ctx, req.(*RemoveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagementService_ServiceDesc is the grpc.ServiceDesc for ManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.v1.ManagementService",
	HandlerType: (*ManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAnnotation",
			Handler:    _ManagementService_AddAnnotation_Handler,
		},
		{
			MethodName: "ListAgents",
			Handler:    _ManagementService_ListAgents_Handler,
		},
		{
			MethodName: "ListAgentVersions",
			Handler:    _ManagementService_ListAgentVersions_Handler,
		},
		{
			MethodName: "RegisterNode",
			Handler:    _ManagementService_RegisterNode_Handler,
		},
		{
			MethodName: "UnregisterNode",
			Handler:    _ManagementService_UnregisterNode_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _ManagementService_ListNodes_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _ManagementService_GetNode_Handler,
		},
		{
			MethodName: "AddService",
			Handler:    _ManagementService_AddService_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _ManagementService_ListServices_Handler,
		},
		{
			MethodName: "DiscoverRDS",
			Handler:    _ManagementService_DiscoverRDS_Handler,
		},
		{
			MethodName: "DiscoverAzureDatabase",
			Handler:    _ManagementService_DiscoverAzureDatabase_Handler,
		},
		{
			MethodName: "AddAzureDatabase",
			Handler:    _ManagementService_AddAzureDatabase_Handler,
		},
		{
			MethodName: "RemoveService",
			Handler:    _ManagementService_RemoveService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "management/v1/service.proto",
}