// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: management/v1/haproxy.proto

package managementv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v1 "github.com/percona/pmm/api/inventory/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddHAProxyServiceParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node identifier on which an external exporter is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node name on which a service and node is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Create a new Node with those parameters.
	// add_node always should be passed with address.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	AddNode *AddNodeParams `protobuf:"bytes,3,opt,name=add_node,json=addNode,proto3" json:"add_node,omitempty"`
	// Node and Exporter access address (DNS name or IP).
	// address always should be passed with add_node.
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,5,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// HTTP basic auth username for collecting metrics.
	Username string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	// HTTP basic auth password for collecting metrics.
	Password string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	// Scheme to generate URI to exporter metrics endpoints.
	Scheme string `protobuf:"bytes,8,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// Path under which metrics are exposed, used to generate URI.
	MetricsPath string `protobuf:"bytes,9,opt,name=metrics_path,json=metricsPath,proto3" json:"metrics_path,omitempty"`
	// Listen port for scraping metrics.
	ListenPort uint32 `protobuf:"varint,10,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,11,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,12,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,13,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `protobuf:"bytes,14,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Defines metrics flow model for this exporter.
	// Metrics could be pushed to the server with vmagent,
	// pulled by the server, or the server could choose behavior automatically.
	// Node with registered pmm_agent_id must present at pmm-server
	// in case of push metrics_mode.
	MetricsMode MetricsMode `protobuf:"varint,15,opt,name=metrics_mode,json=metricsMode,proto3,enum=management.v1.MetricsMode" json:"metrics_mode,omitempty"`
	// Skip connection check.
	SkipConnectionCheck bool `protobuf:"varint,16,opt,name=skip_connection_check,json=skipConnectionCheck,proto3" json:"skip_connection_check,omitempty"`
}

func (x *AddHAProxyServiceParams) Reset() {
	*x = AddHAProxyServiceParams{}
	mi := &file_management_v1_haproxy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddHAProxyServiceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHAProxyServiceParams) ProtoMessage() {}

func (x *AddHAProxyServiceParams) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_haproxy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHAProxyServiceParams.ProtoReflect.Descriptor instead.
func (*AddHAProxyServiceParams) Descriptor() ([]byte, []int) {
	return file_management_v1_haproxy_proto_rawDescGZIP(), []int{0}
}

func (x *AddHAProxyServiceParams) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetAddNode() *AddNodeParams {
	if x != nil {
		return x.AddNode
	}
	return nil
}

func (x *AddHAProxyServiceParams) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetMetricsPath() string {
	if x != nil {
		return x.MetricsPath
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetListenPort() uint32 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *AddHAProxyServiceParams) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetReplicationSet() string {
	if x != nil {
		return x.ReplicationSet
	}
	return ""
}

func (x *AddHAProxyServiceParams) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *AddHAProxyServiceParams) GetMetricsMode() MetricsMode {
	if x != nil {
		return x.MetricsMode
	}
	return MetricsMode_METRICS_MODE_UNSPECIFIED
}

func (x *AddHAProxyServiceParams) GetSkipConnectionCheck() bool {
	if x != nil {
		return x.SkipConnectionCheck
	}
	return false
}

type HAProxyServiceResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service          *v1.HAProxyService   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ExternalExporter *v1.ExternalExporter `protobuf:"bytes,2,opt,name=external_exporter,json=externalExporter,proto3" json:"external_exporter,omitempty"`
}

func (x *HAProxyServiceResult) Reset() {
	*x = HAProxyServiceResult{}
	mi := &file_management_v1_haproxy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HAProxyServiceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HAProxyServiceResult) ProtoMessage() {}

func (x *HAProxyServiceResult) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_haproxy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HAProxyServiceResult.ProtoReflect.Descriptor instead.
func (*HAProxyServiceResult) Descriptor() ([]byte, []int) {
	return file_management_v1_haproxy_proto_rawDescGZIP(), []int{1}
}

func (x *HAProxyServiceResult) GetService() *v1.HAProxyService {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *HAProxyServiceResult) GetExternalExporter() *v1.ExternalExporter {
	if x != nil {
		return x.ExternalExporter
	}
	return nil
}

var File_management_v1_haproxy_proto protoreflect.FileDescriptor

var file_management_v1_haproxy_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x05, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x48, 0x41, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f,
	0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0b, 0xfa,
	0x42, 0x08, 0x2a, 0x06, 0x10, 0x80, 0x80, 0x04, 0x20, 0x00, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x5d, 0x0a, 0x0d, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x6b, 0x69,
	0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x3f, 0x0a,
	0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9b,
	0x01, 0x0a, 0x14, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4b, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x42, 0xad, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x48, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa,
	0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x19, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_v1_haproxy_proto_rawDescOnce sync.Once
	file_management_v1_haproxy_proto_rawDescData = file_management_v1_haproxy_proto_rawDesc
)

func file_management_v1_haproxy_proto_rawDescGZIP() []byte {
	file_management_v1_haproxy_proto_rawDescOnce.Do(func() {
		file_management_v1_haproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_v1_haproxy_proto_rawDescData)
	})
	return file_management_v1_haproxy_proto_rawDescData
}

var (
	file_management_v1_haproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_management_v1_haproxy_proto_goTypes  = []any{
		(*AddHAProxyServiceParams)(nil), // 0: management.v1.AddHAProxyServiceParams
		(*HAProxyServiceResult)(nil),    // 1: management.v1.HAProxyServiceResult
		nil,                             // 2: management.v1.AddHAProxyServiceParams.CustomLabelsEntry
		(*AddNodeParams)(nil),           // 3: management.v1.AddNodeParams
		(MetricsMode)(0),                // 4: management.v1.MetricsMode
		(*v1.HAProxyService)(nil),       // 5: inventory.v1.HAProxyService
		(*v1.ExternalExporter)(nil),     // 6: inventory.v1.ExternalExporter
	}
)

var file_management_v1_haproxy_proto_depIdxs = []int32{
	3, // 0: management.v1.AddHAProxyServiceParams.add_node:type_name -> management.v1.AddNodeParams
	2, // 1: management.v1.AddHAProxyServiceParams.custom_labels:type_name -> management.v1.AddHAProxyServiceParams.CustomLabelsEntry
	4, // 2: management.v1.AddHAProxyServiceParams.metrics_mode:type_name -> management.v1.MetricsMode
	5, // 3: management.v1.HAProxyServiceResult.service:type_name -> inventory.v1.HAProxyService
	6, // 4: management.v1.HAProxyServiceResult.external_exporter:type_name -> inventory.v1.ExternalExporter
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_management_v1_haproxy_proto_init() }
func file_management_v1_haproxy_proto_init() {
	if File_management_v1_haproxy_proto != nil {
		return
	}
	file_management_v1_metrics_proto_init()
	file_management_v1_node_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_management_v1_haproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_management_v1_haproxy_proto_goTypes,
		DependencyIndexes: file_management_v1_haproxy_proto_depIdxs,
		MessageInfos:      file_management_v1_haproxy_proto_msgTypes,
	}.Build()
	File_management_v1_haproxy_proto = out.File
	file_management_v1_haproxy_proto_rawDesc = nil
	file_management_v1_haproxy_proto_goTypes = nil
	file_management_v1_haproxy_proto_depIdxs = nil
}