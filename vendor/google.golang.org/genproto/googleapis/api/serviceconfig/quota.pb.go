// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/quota.proto

package serviceconfig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Quota configuration helps to achieve fairness and budgeting in service
// usage.
//
// The quota configuration works this way:
// - The service configuration defines a set of metrics.
// - For API calls, the quota.metric_rules maps methods to metrics with
//   corresponding costs.
// - The quota.limits defines limits on the metrics, which will be used for
//   quota checks at runtime.
//
// An example quota configuration in yaml format:
//
//    quota:
//      limits:
//
//      - name: apiWriteQpsPerProject
//        metric: library.googleapis.com/write_calls
//        unit: "1/min/{project}"  # rate limit for consumer projects
//        values:
//          STANDARD: 10000
//
//      # The metric rules bind all methods to the read_calls metric,
//      # except for the UpdateBook and DeleteBook methods. These two methods
//      # are mapped to the write_calls metric, with the UpdateBook method
//      # consuming at twice rate as the DeleteBook method.
//      metric_rules:
//      - selector: "*"
//        metric_costs:
//          library.googleapis.com/read_calls: 1
//      - selector: google.example.library.v1.LibraryService.UpdateBook
//        metric_costs:
//          library.googleapis.com/write_calls: 2
//      - selector: google.example.library.v1.LibraryService.DeleteBook
//        metric_costs:
//          library.googleapis.com/write_calls: 1
//
//  Corresponding Metric definition:
//
//      metrics:
//      - name: library.googleapis.com/read_calls
//        display_name: Read requests
//        metric_kind: DELTA
//        value_type: INT64
//
//      - name: library.googleapis.com/write_calls
//        display_name: Write requests
//        metric_kind: DELTA
//        value_type: INT64
//
type Quota struct {
	// List of `QuotaLimit` definitions for the service.
	//
	// Used by metric-based quotas only.
	Limits []*QuotaLimit `protobuf:"bytes,3,rep,name=limits,proto3" json:"limits,omitempty"`
	// List of `MetricRule` definitions, each one mapping a selected method to one
	// or more metrics.
	//
	// Used by metric-based quotas only.
	MetricRules          []*MetricRule `protobuf:"bytes,4,rep,name=metric_rules,json=metricRules,proto3" json:"metric_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Quota) Reset()         { *m = Quota{} }
func (m *Quota) String() string { return proto.CompactTextString(m) }
func (*Quota) ProtoMessage()    {}
func (*Quota) Descriptor() ([]byte, []int) {
	return fileDescriptor_6822ef0454b3845a, []int{0}
}

func (m *Quota) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quota.Unmarshal(m, b)
}
func (m *Quota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quota.Marshal(b, m, deterministic)
}
func (m *Quota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quota.Merge(m, src)
}
func (m *Quota) XXX_Size() int {
	return xxx_messageInfo_Quota.Size(m)
}
func (m *Quota) XXX_DiscardUnknown() {
	xxx_messageInfo_Quota.DiscardUnknown(m)
}

var xxx_messageInfo_Quota proto.InternalMessageInfo

func (m *Quota) GetLimits() []*QuotaLimit {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *Quota) GetMetricRules() []*MetricRule {
	if m != nil {
		return m.MetricRules
	}
	return nil
}

// Bind API methods to metrics. Binding a method to a metric causes that
// metric's configured quota, billing, and monitoring behaviors to apply to the
// method call.
//
// Used by metric-based quotas only.
type MetricRule struct {
	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Metrics to update when the selected methods are called, and the associated
	// cost applied to each metric.
	//
	// The key of the map is the metric name, and the values are the amount
	// increased for the metric against which the quota limits are defined.
	// The value must not be negative.
	MetricCosts          map[string]int64 `protobuf:"bytes,2,rep,name=metric_costs,json=metricCosts,proto3" json:"metric_costs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MetricRule) Reset()         { *m = MetricRule{} }
func (m *MetricRule) String() string { return proto.CompactTextString(m) }
func (*MetricRule) ProtoMessage()    {}
func (*MetricRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6822ef0454b3845a, []int{1}
}

func (m *MetricRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricRule.Unmarshal(m, b)
}
func (m *MetricRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricRule.Marshal(b, m, deterministic)
}
func (m *MetricRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricRule.Merge(m, src)
}
func (m *MetricRule) XXX_Size() int {
	return xxx_messageInfo_MetricRule.Size(m)
}
func (m *MetricRule) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricRule.DiscardUnknown(m)
}

var xxx_messageInfo_MetricRule proto.InternalMessageInfo

func (m *MetricRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *MetricRule) GetMetricCosts() map[string]int64 {
	if m != nil {
		return m.MetricCosts
	}
	return nil
}

// `QuotaLimit` defines a specific limit that applies over a specified duration
// for a limit type. There can be at most one limit for a duration and limit
// type combination defined within a `QuotaGroup`.
type QuotaLimit struct {
	// Name of the quota limit. The name is used to refer to the limit when
	// overriding the default limit on per-consumer basis.
	//
	// For group-based quota limits, the name must be unique within the quota
	// group. If a name is not provided, it will be generated from the limit_by
	// and duration fields.
	//
	// For metric-based quota limits, the name must be provided, and it must be
	// unique within the service. The name can only include alphanumeric
	// characters as well as '-'.
	//
	// The maximum length of the limit name is 64 characters.
	//
	// The name of a limit is used as a unique identifier for this limit.
	// Therefore, once a limit has been put into use, its name should be
	// immutable. You can use the display_name field to provide a user-friendly
	// name for the limit. The display name can be evolved over time without
	// affecting the identity of the limit.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. User-visible, extended description for this quota limit.
	// Should be used only when more context is needed to understand this limit
	// than provided by the limit's display name (see: `display_name`).
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Default number of tokens that can be consumed during the specified
	// duration. This is the number of tokens assigned when a client
	// application developer activates the service for his/her project.
	//
	// Specifying a value of 0 will block all requests. This can be used if you
	// are provisioning quota to selected consumers and blocking others.
	// Similarly, a value of -1 will indicate an unlimited quota. No other
	// negative values are allowed.
	//
	// Used by group-based quotas only.
	DefaultLimit int64 `protobuf:"varint,3,opt,name=default_limit,json=defaultLimit,proto3" json:"default_limit,omitempty"`
	// Maximum number of tokens that can be consumed during the specified
	// duration. Client application developers can override the default limit up
	// to this maximum. If specified, this value cannot be set to a value less
	// than the default limit. If not specified, it is set to the default limit.
	//
	// To allow clients to apply overrides with no upper bound, set this to -1,
	// indicating unlimited maximum quota.
	//
	// Used by group-based quotas only.
	MaxLimit int64 `protobuf:"varint,4,opt,name=max_limit,json=maxLimit,proto3" json:"max_limit,omitempty"`
	// Free tier value displayed in the Developers Console for this limit.
	// The free tier is the number of tokens that will be subtracted from the
	// billed amount when billing is enabled.
	// This field can only be set on a limit with duration "1d", in a billable
	// group; it is invalid on any other limit. If this field is not set, it
	// defaults to 0, indicating that there is no free tier for this service.
	//
	// Used by group-based quotas only.
	FreeTier int64 `protobuf:"varint,7,opt,name=free_tier,json=freeTier,proto3" json:"free_tier,omitempty"`
	// Duration of this limit in textual notation. Example: "100s", "24h", "1d".
	// For duration longer than a day, only multiple of days is supported. We
	// support only "100s" and "1d" for now. Additional support will be added in
	// the future. "0" indicates indefinite duration.
	//
	// Used by group-based quotas only.
	Duration string `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	// The name of the metric this quota limit applies to. The quota limits with
	// the same metric will be checked together during runtime. The metric must be
	// defined within the service config.
	//
	// Used by metric-based quotas only.
	Metric string `protobuf:"bytes,8,opt,name=metric,proto3" json:"metric,omitempty"`
	// Specify the unit of the quota limit. It uses the same syntax as
	// [Metric.unit][]. The supported unit kinds are determined by the quota
	// backend system.
	//
	// The [Google Service Control](https://cloud.google.com/service-control)
	// supports the following unit components:
	// * One of the time intevals:
	//   * "/min"  for quota every minute.
	//   * "/d"  for quota every 24 hours, starting 00:00 US Pacific Time.
	//   * Otherwise the quota won't be reset by time, such as storage limit.
	// * One and only one of the granted containers:
	//   * "/{organization}" quota for an organization.
	//   * "/{project}" quota for a project.
	//   * "/{folder}" quota for a folder.
	//   * "/{resource}" quota for a universal resource.
	// * Zero or more quota segmentation dimension. Not all combos are valid.
	//   * "/{region}" quota for every region. Not to be used with time intervals.
	//   * Otherwise the resources granted on the target is not segmented.
	//   * "/{zone}" quota for every zone. Not to be used with time intervals.
	//   * Otherwise the resources granted on the target is not segmented.
	//   * "/{resource}" quota for a resource associated with a project or org.
	//
	// Here are some examples:
	// * "1/min/{project}" for quota per minute per project.
	// * "1/min/{user}" for quota per minute per user.
	// * "1/min/{organization}" for quota per minute per organization.
	//
	// Note: the order of unit components is insignificant.
	// The "1" at the beginning is required to follow the metric unit syntax.
	//
	// Used by metric-based quotas only.
	Unit string `protobuf:"bytes,9,opt,name=unit,proto3" json:"unit,omitempty"`
	// Tiered limit values. Also allows for regional or zone overrides for these
	// values if "/{region}" or "/{zone}" is specified in the unit field.
	//
	// Currently supported tiers from low to high:
	// VERY_LOW, LOW, STANDARD, HIGH, VERY_HIGH
	//
	// To apply different limit values for users according to their tiers, specify
	// the values for the tiers you want to differentiate. For example:
	// {LOW:100, STANDARD:500, HIGH:1000, VERY_HIGH:5000}
	//
	// The limit value for each tier is optional except for the tier STANDARD.
	// The limit value for an unspecified tier falls to the value of its next
	// tier towards tier STANDARD. For the above example, the limit value for tier
	// STANDARD is 500.
	//
	// To apply the same limit value for all users, just specify limit value for
	// tier STANDARD. For example: {STANDARD:500}.
	//
	// To apply a regional overide for a tier, add a map entry with key
	// "<TIER>/<region>", where <region> is a region name. Similarly, for a zone
	// override, add a map entry with key "<TIER>/{zone}".
	// Further, a wildcard can be used at the end of a zone name in order to
	// specify zone level overrides. For example:
	// LOW: 10, STANDARD: 50, HIGH: 100,
	// LOW/us-central1: 20, STANDARD/us-central1: 60, HIGH/us-central1: 200,
	// LOW/us-central1-*: 10, STANDARD/us-central1-*: 20, HIGH/us-central1-*: 80
	//
	// The regional overrides tier set for each region must be the same as
	// the tier set for default limit values. Same rule applies for zone overrides
	// tier as well.
	//
	// Used by metric-based quotas only.
	Values map[string]int64 `protobuf:"bytes,10,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// User-visible display name for this limit.
	// Optional. If not set, the UI will provide a default display name based on
	// the quota configuration. This field can be used to override the default
	// display name generated from the configuration.
	DisplayName          string   `protobuf:"bytes,12,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaLimit) Reset()         { *m = QuotaLimit{} }
func (m *QuotaLimit) String() string { return proto.CompactTextString(m) }
func (*QuotaLimit) ProtoMessage()    {}
func (*QuotaLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_6822ef0454b3845a, []int{2}
}

func (m *QuotaLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaLimit.Unmarshal(m, b)
}
func (m *QuotaLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaLimit.Marshal(b, m, deterministic)
}
func (m *QuotaLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaLimit.Merge(m, src)
}
func (m *QuotaLimit) XXX_Size() int {
	return xxx_messageInfo_QuotaLimit.Size(m)
}
func (m *QuotaLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaLimit.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaLimit proto.InternalMessageInfo

func (m *QuotaLimit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QuotaLimit) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *QuotaLimit) GetDefaultLimit() int64 {
	if m != nil {
		return m.DefaultLimit
	}
	return 0
}

func (m *QuotaLimit) GetMaxLimit() int64 {
	if m != nil {
		return m.MaxLimit
	}
	return 0
}

func (m *QuotaLimit) GetFreeTier() int64 {
	if m != nil {
		return m.FreeTier
	}
	return 0
}

func (m *QuotaLimit) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *QuotaLimit) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *QuotaLimit) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *QuotaLimit) GetValues() map[string]int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *QuotaLimit) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func init() {
	proto.RegisterType((*Quota)(nil), "google.api.Quota")
	proto.RegisterType((*MetricRule)(nil), "google.api.MetricRule")
	proto.RegisterMapType((map[string]int64)(nil), "google.api.MetricRule.MetricCostsEntry")
	proto.RegisterType((*QuotaLimit)(nil), "google.api.QuotaLimit")
	proto.RegisterMapType((map[string]int64)(nil), "google.api.QuotaLimit.ValuesEntry")
}

func init() { proto.RegisterFile("google/api/quota.proto", fileDescriptor_6822ef0454b3845a) }

var fileDescriptor_6822ef0454b3845a = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0x9a, 0xb6, 0xb4, 0xd3, 0x82, 0x56, 0x16, 0xaa, 0xac, 0xc2, 0xa1, 0x94, 0x03, 0x3d,
	0xa5, 0x12, 0x5c, 0xd8, 0x45, 0x42, 0x62, 0xd1, 0x0a, 0x81, 0x00, 0x95, 0x08, 0x71, 0xe0, 0x52,
	0x99, 0x74, 0x1a, 0x59, 0x38, 0x71, 0xb0, 0x9d, 0xd5, 0xf6, 0xcc, 0x9f, 0xf0, 0x0d, 0x7c, 0x20,
	0xf2, 0xd8, 0xdb, 0x16, 0xd8, 0xcb, 0xde, 0x66, 0xe6, 0xbd, 0xe7, 0x17, 0x3f, 0x4f, 0x60, 0x52,
	0x6a, 0x5d, 0x2a, 0x5c, 0x8a, 0x46, 0x2e, 0x7f, 0xb4, 0xda, 0x89, 0xac, 0x31, 0xda, 0x69, 0x06,
	0x61, 0x9e, 0x89, 0x46, 0x4e, 0x1f, 0x1e, 0x71, 0x44, 0x5d, 0x6b, 0x27, 0x9c, 0xd4, 0xb5, 0x0d,
	0xcc, 0xb9, 0x81, 0xde, 0x27, 0x2f, 0x64, 0x19, 0xf4, 0x95, 0xac, 0xa4, 0xb3, 0x3c, 0x9d, 0xa5,
	0x8b, 0xd1, 0xd3, 0x49, 0x76, 0x38, 0x23, 0x23, 0xca, 0x7b, 0x0f, 0xe7, 0x91, 0xc5, 0x4e, 0x61,
	0x5c, 0xa1, 0x33, 0xb2, 0x58, 0x9b, 0x56, 0xa1, 0xe5, 0xdd, 0xff, 0x55, 0x1f, 0x08, 0xcf, 0x5b,
	0x85, 0xf9, 0xa8, 0xda, 0xd7, 0x76, 0xfe, 0x3b, 0x01, 0x38, 0x60, 0x6c, 0x0a, 0x03, 0x8b, 0x0a,
	0x0b, 0xa7, 0x0d, 0x4f, 0x66, 0xc9, 0x62, 0x98, 0xef, 0x7b, 0xf6, 0x6e, 0xef, 0x52, 0x68, 0xeb,
	0x2c, 0xef, 0x90, 0xcb, 0x93, 0x9b, 0x5d, 0x62, 0xf9, 0xda, 0x33, 0x2f, 0x6a, 0x67, 0x76, 0xd7,
	0xb6, 0x34, 0x99, 0xbe, 0x84, 0x93, 0x7f, 0x09, 0xec, 0x04, 0xd2, 0xef, 0xb8, 0x8b, 0xb6, 0xbe,
	0x64, 0xf7, 0xa1, 0x77, 0x29, 0x54, 0x8b, 0xbc, 0x33, 0x4b, 0x16, 0x69, 0x1e, 0x9a, 0xb3, 0xce,
	0xf3, 0x64, 0xfe, 0x33, 0x05, 0x38, 0x04, 0xc1, 0x18, 0x74, 0x6b, 0x51, 0x21, 0xef, 0x93, 0x96,
	0x6a, 0x36, 0x83, 0xd1, 0x06, 0x6d, 0x61, 0x64, 0xe3, 0x33, 0xa6, 0x23, 0x86, 0xf9, 0xf1, 0x88,
	0x3d, 0x86, 0xbb, 0x1b, 0xdc, 0x8a, 0x56, 0xb9, 0x35, 0x05, 0xc9, 0x53, 0xb2, 0x19, 0xc7, 0x61,
	0x38, 0xfa, 0x01, 0x0c, 0x2b, 0x71, 0x15, 0x09, 0x5d, 0x22, 0x0c, 0x2a, 0x71, 0xb5, 0x07, 0xb7,
	0x06, 0x71, 0xed, 0x24, 0x1a, 0x7e, 0x27, 0x80, 0x7e, 0xf0, 0x59, 0xa2, 0xf1, 0x59, 0x6e, 0x5a,
	0x43, 0x2f, 0xcc, 0x7b, 0x21, 0xcb, 0xeb, 0x9e, 0x4d, 0xa0, 0x1f, 0xe2, 0xe0, 0x03, 0x42, 0x62,
	0xe7, 0x2f, 0xd2, 0xd6, 0xd2, 0xf1, 0x61, 0xb8, 0x88, 0xaf, 0xd9, 0x19, 0xf4, 0xe9, 0xe2, 0x96,
	0x03, 0x25, 0x3e, 0xbf, 0x79, 0x1b, 0xb2, 0x2f, 0x44, 0x0a, 0x61, 0x47, 0x05, 0x7b, 0x04, 0xe3,
	0x8d, 0xb4, 0x8d, 0x12, 0xbb, 0x35, 0x05, 0x34, 0x8e, 0x29, 0x84, 0xd9, 0x47, 0x51, 0xe1, 0xf4,
	0x14, 0x46, 0x47, 0xca, 0xdb, 0xbc, 0xc2, 0xb9, 0x82, 0x7b, 0x85, 0xae, 0x8e, 0x3e, 0xe7, 0x3c,
	0x3c, 0xca, 0xca, 0xaf, 0xf3, 0x2a, 0xf9, 0x7a, 0x11, 0x91, 0x52, 0x2b, 0x51, 0x97, 0x99, 0x36,
	0xe5, 0xb2, 0xc4, 0x9a, 0x96, 0x7d, 0x19, 0x20, 0xd1, 0x48, 0x4b, 0x7f, 0x83, 0x45, 0x73, 0x29,
	0x0b, 0x2c, 0x74, 0xbd, 0x95, 0xe5, 0x8b, 0xbf, 0xba, 0x5f, 0x9d, 0xee, 0x9b, 0x57, 0xab, 0xb7,
	0xdf, 0xfa, 0x24, 0x7c, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x90, 0x7e, 0xf5, 0xab, 0x69, 0x03,
	0x00, 0x00,
}
