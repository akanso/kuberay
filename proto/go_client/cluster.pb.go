// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: cluster.proto

package go_client

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Optional field.
type Cluster_Environment int32

const (
	Cluster_DEV        Cluster_Environment = 0
	Cluster_TESTING    Cluster_Environment = 1
	Cluster_STAGING    Cluster_Environment = 2
	Cluster_PRODUCTION Cluster_Environment = 3
)

// Enum value maps for Cluster_Environment.
var (
	Cluster_Environment_name = map[int32]string{
		0: "DEV",
		1: "TESTING",
		2: "STAGING",
		3: "PRODUCTION",
	}
	Cluster_Environment_value = map[string]int32{
		"DEV":        0,
		"TESTING":    1,
		"STAGING":    2,
		"PRODUCTION": 3,
	}
)

func (x Cluster_Environment) Enum() *Cluster_Environment {
	p := new(Cluster_Environment)
	*p = x
	return p
}

func (x Cluster_Environment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_Environment) Descriptor() protoreflect.EnumDescriptor {
	return file_cluster_proto_enumTypes[0].Descriptor()
}

func (Cluster_Environment) Type() protoreflect.EnumType {
	return &file_cluster_proto_enumTypes[0]
}

func (x Cluster_Environment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_Environment.Descriptor instead.
func (Cluster_Environment) EnumDescriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{5, 0}
}

type CreateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Cluster to be created.
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClusterRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type GetClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Cluster to be retrieved.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Cluster to be retrieved.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetClusterRequest) Reset() {
	*x = GetClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRequest) ProtoMessage() {}

func (x *GetClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterRequest.ProtoReflect.Descriptor instead.
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *GetClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListClustersRequest) Reset() {
	*x = ListClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersRequest) ProtoMessage() {}

func (x *ListClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersRequest.ProtoReflect.Descriptor instead.
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{2}
}

type ListClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of Clusters returned.
	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *ListClustersResponse) Reset() {
	*x = ListClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersResponse) ProtoMessage() {}

func (x *ListClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersResponse.ProtoReflect.Descriptor instead.
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *ListClustersResponse) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type DeleteClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the cluster to be deleted.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Cluster to be retrieved.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteClusterRequest) Reset() {
	*x = DeleteClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterRequest) ProtoMessage() {}

func (x *DeleteClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterRequest.ProtoReflect.Descriptor instead.
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required input field. Unique Cluster name provided by user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required input field. Cluster's namespace provided by user
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Required field. This field indicates the user who owns the cluster.
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// Optional input field. Ray cluster version
	Version     string              `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Environment Cluster_Environment `protobuf:"varint,5,opt,name=environment,proto3,enum=proto.Cluster_Environment" json:"environment,omitempty"`
	// Required field. This field indicates ray cluster configuration
	ClusterSpec *ClusterSpec `protobuf:"bytes,6,opt,name=cluster_spec,json=clusterSpec,proto3" json:"cluster_spec,omitempty"`
	// Output. The time that the Cluster created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Output. The time that the Cluster deleted.
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Cluster) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Cluster) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Cluster) GetEnvironment() Cluster_Environment {
	if x != nil {
		return x.Environment
	}
	return Cluster_DEV
}

func (x *Cluster) GetClusterSpec() *ClusterSpec {
	if x != nil {
		return x.ClusterSpec
	}
	return nil
}

func (x *Cluster) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Cluster) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type ClusterSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The head group configuration
	HeadGroupSpec *HeadGroupSpec `protobuf:"bytes,1,opt,name=head_group_spec,json=headGroupSpec,proto3" json:"head_group_spec,omitempty"`
	// The worker group configurations
	WorkerGroupSepc []*WorkerGroupSpec `protobuf:"bytes,2,rep,name=worker_group_sepc,json=workerGroupSepc,proto3" json:"worker_group_sepc,omitempty"`
}

func (x *ClusterSpec) Reset() {
	*x = ClusterSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterSpec) ProtoMessage() {}

func (x *ClusterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterSpec.ProtoReflect.Descriptor instead.
func (*ClusterSpec) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{6}
}

func (x *ClusterSpec) GetHeadGroupSpec() *HeadGroupSpec {
	if x != nil {
		return x.HeadGroupSpec
	}
	return nil
}

func (x *ClusterSpec) GetWorkerGroupSepc() []*WorkerGroupSpec {
	if x != nil {
		return x.WorkerGroupSepc
	}
	return nil
}

type HeadGroupSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The computeTemplate of head node group
	ComputeTemplate string `protobuf:"bytes,1,opt,name=compute_template,json=computeTemplate,proto3" json:"compute_template,omitempty"`
	// Optional field. This field will be used to retrieve right ray container
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// Optional. The service type (ClusterIP, NodePort, Load balancer) of the head node
	ServiceType string `protobuf:"bytes,3,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	// Optional. The ray start parames of head node group
	RayStartParams map[string]string `protobuf:"bytes,4,rep,name=ray_start_params,json=rayStartParams,proto3" json:"ray_start_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HeadGroupSpec) Reset() {
	*x = HeadGroupSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeadGroupSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeadGroupSpec) ProtoMessage() {}

func (x *HeadGroupSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeadGroupSpec.ProtoReflect.Descriptor instead.
func (*HeadGroupSpec) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{7}
}

func (x *HeadGroupSpec) GetComputeTemplate() string {
	if x != nil {
		return x.ComputeTemplate
	}
	return ""
}

func (x *HeadGroupSpec) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *HeadGroupSpec) GetServiceType() string {
	if x != nil {
		return x.ServiceType
	}
	return ""
}

func (x *HeadGroupSpec) GetRayStartParams() map[string]string {
	if x != nil {
		return x.RayStartParams
	}
	return nil
}

type WorkerGroupSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Group name of the current worker group
	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	// Optional. The computeTemplate of head node group
	ComputeTemplate string `protobuf:"bytes,2,opt,name=compute_template,json=computeTemplate,proto3" json:"compute_template,omitempty"`
	// Optional field. This field will be used to retrieve right ray container
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	// Required. Desired replicas of the worker group
	Replicas int32 `protobuf:"varint,4,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// Optional. Min replicas of the worker group
	MinReplicas int32 `protobuf:"varint,5,opt,name=min_replicas,json=minReplicas,proto3" json:"min_replicas,omitempty"`
	// Optional. Max replicas of the worker group
	MaxReplicas int32 `protobuf:"varint,6,opt,name=max_replicas,json=maxReplicas,proto3" json:"max_replicas,omitempty"`
	// Optional. The ray start parames of worker node group
	RayStartParams map[string]string `protobuf:"bytes,7,rep,name=ray_start_params,json=rayStartParams,proto3" json:"ray_start_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WorkerGroupSpec) Reset() {
	*x = WorkerGroupSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerGroupSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerGroupSpec) ProtoMessage() {}

func (x *WorkerGroupSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerGroupSpec.ProtoReflect.Descriptor instead.
func (*WorkerGroupSpec) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{8}
}

func (x *WorkerGroupSpec) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *WorkerGroupSpec) GetComputeTemplate() string {
	if x != nil {
		return x.ComputeTemplate
	}
	return ""
}

func (x *WorkerGroupSpec) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *WorkerGroupSpec) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *WorkerGroupSpec) GetMinReplicas() int32 {
	if x != nil {
		return x.MinReplicas
	}
	return 0
}

func (x *WorkerGroupSpec) GetMaxReplicas() int32 {
	if x != nil {
		return x.MaxReplicas
	}
	return 0
}

func (x *WorkerGroupSpec) GetRayStartParams() map[string]string {
	if x != nil {
		return x.RayStartParams
	}
	return nil
}

var File_cluster_proto protoreflect.FileDescriptor

var file_cluster_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x40, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x96, 0x03, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x3c, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x35,
	0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x40, 0x0a, 0x0b, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x45,
	0x56, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x22, 0x8f, 0x01,
	0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3c, 0x0a,
	0x0f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0d, 0x68, 0x65,
	0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x42, 0x0a, 0x11, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x70, 0x63,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0f,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x70, 0x63, 0x22,
	0x8a, 0x02, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x72, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x61, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x61, 0x79, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x41, 0x0a, 0x13, 0x52, 0x61, 0x79,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xec, 0x02, 0x0a,
	0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x73, 0x12, 0x54, 0x0a, 0x10, 0x72, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x61, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x61, 0x79, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x41, 0x0a, 0x13, 0x52, 0x61, 0x79, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xab, 0x03, 0x0a, 0x0e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x28, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x22, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x3a, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x67, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x6a, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x1c, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x54, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x92, 0x41, 0x21, 0x2a, 0x01,
	0x01, 0x52, 0x1c, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x11, 0x12, 0x0f,
	0x0a, 0x0d, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cluster_proto_rawDescOnce sync.Once
	file_cluster_proto_rawDescData = file_cluster_proto_rawDesc
)

func file_cluster_proto_rawDescGZIP() []byte {
	file_cluster_proto_rawDescOnce.Do(func() {
		file_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_cluster_proto_rawDescData)
	})
	return file_cluster_proto_rawDescData
}

var file_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_cluster_proto_goTypes = []interface{}{
	(Cluster_Environment)(0),      // 0: proto.Cluster.Environment
	(*CreateClusterRequest)(nil),  // 1: proto.CreateClusterRequest
	(*GetClusterRequest)(nil),     // 2: proto.GetClusterRequest
	(*ListClustersRequest)(nil),   // 3: proto.ListClustersRequest
	(*ListClustersResponse)(nil),  // 4: proto.ListClustersResponse
	(*DeleteClusterRequest)(nil),  // 5: proto.DeleteClusterRequest
	(*Cluster)(nil),               // 6: proto.Cluster
	(*ClusterSpec)(nil),           // 7: proto.ClusterSpec
	(*HeadGroupSpec)(nil),         // 8: proto.HeadGroupSpec
	(*WorkerGroupSpec)(nil),       // 9: proto.WorkerGroupSpec
	nil,                           // 10: proto.HeadGroupSpec.RayStartParamsEntry
	nil,                           // 11: proto.WorkerGroupSpec.RayStartParamsEntry
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 13: google.protobuf.Empty
}
var file_cluster_proto_depIdxs = []int32{
	6,  // 0: proto.CreateClusterRequest.cluster:type_name -> proto.Cluster
	6,  // 1: proto.ListClustersResponse.clusters:type_name -> proto.Cluster
	0,  // 2: proto.Cluster.environment:type_name -> proto.Cluster.Environment
	7,  // 3: proto.Cluster.cluster_spec:type_name -> proto.ClusterSpec
	12, // 4: proto.Cluster.created_at:type_name -> google.protobuf.Timestamp
	12, // 5: proto.Cluster.deleted_at:type_name -> google.protobuf.Timestamp
	8,  // 6: proto.ClusterSpec.head_group_spec:type_name -> proto.HeadGroupSpec
	9,  // 7: proto.ClusterSpec.worker_group_sepc:type_name -> proto.WorkerGroupSpec
	10, // 8: proto.HeadGroupSpec.ray_start_params:type_name -> proto.HeadGroupSpec.RayStartParamsEntry
	11, // 9: proto.WorkerGroupSpec.ray_start_params:type_name -> proto.WorkerGroupSpec.RayStartParamsEntry
	1,  // 10: proto.ClusterService.CreateCluster:input_type -> proto.CreateClusterRequest
	2,  // 11: proto.ClusterService.GetCluster:input_type -> proto.GetClusterRequest
	3,  // 12: proto.ClusterService.ListCluster:input_type -> proto.ListClustersRequest
	5,  // 13: proto.ClusterService.DeleteCluster:input_type -> proto.DeleteClusterRequest
	6,  // 14: proto.ClusterService.CreateCluster:output_type -> proto.Cluster
	6,  // 15: proto.ClusterService.GetCluster:output_type -> proto.Cluster
	4,  // 16: proto.ClusterService.ListCluster:output_type -> proto.ListClustersResponse
	13, // 17: proto.ClusterService.DeleteCluster:output_type -> google.protobuf.Empty
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cluster_proto_init() }
func file_cluster_proto_init() {
	if File_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClustersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClustersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeadGroupSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerGroupSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cluster_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cluster_proto_goTypes,
		DependencyIndexes: file_cluster_proto_depIdxs,
		EnumInfos:         file_cluster_proto_enumTypes,
		MessageInfos:      file_cluster_proto_msgTypes,
	}.Build()
	File_cluster_proto = out.File
	file_cluster_proto_rawDesc = nil
	file_cluster_proto_goTypes = nil
	file_cluster_proto_depIdxs = nil
}
