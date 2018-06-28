/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta2/generated.proto
// DO NOT EDIT!

/*
	Package v1beta2 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta2/generated.proto

	It has these top-level messages:
		MetricIdent
		MetricListOptions
		MetricValue
		MetricValueList
*/
package v1beta2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *MetricIdent) Reset()                    { *m = MetricIdent{} }
func (*MetricIdent) ProtoMessage()               {}
func (*MetricIdent) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *MetricListOptions) Reset()                    { *m = MetricListOptions{} }
func (*MetricListOptions) ProtoMessage()               {}
func (*MetricListOptions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *MetricValue) Reset()                    { *m = MetricValue{} }
func (*MetricValue) ProtoMessage()               {}
func (*MetricValue) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *MetricValueList) Reset()                    { *m = MetricValueList{} }
func (*MetricValueList) ProtoMessage()               {}
func (*MetricValueList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*MetricIdent)(nil), "k8s.io.metrics.pkg.apis.custom_metrics.v1beta2.MetricIdent")
	proto.RegisterType((*MetricListOptions)(nil), "k8s.io.metrics.pkg.apis.custom_metrics.v1beta2.MetricListOptions")
	proto.RegisterType((*MetricValue)(nil), "k8s.io.metrics.pkg.apis.custom_metrics.v1beta2.MetricValue")
	proto.RegisterType((*MetricValueList)(nil), "k8s.io.metrics.pkg.apis.custom_metrics.v1beta2.MetricValueList")
}
func (m *MetricIdent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricIdent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Selector)))
	i += copy(dAtA[i:], m.Selector)
	return i, nil
}

func (m *MetricListOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricListOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.LabelSelector)))
	i += copy(dAtA[i:], m.LabelSelector)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.MetricLabelSelector)))
	i += copy(dAtA[i:], m.MetricLabelSelector)
	return i, nil
}

func (m *MetricValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.DescribedObject.Size()))
	n1, err := m.DescribedObject.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Metric.Size()))
	n2, err := m.Metric.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Timestamp.Size()))
	n3, err := m.Timestamp.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if m.WindowSeconds != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.WindowSeconds))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Value.Size()))
	n4, err := m.Value.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *MetricValueList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricValueList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n5, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MetricIdent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Selector)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *MetricListOptions) Size() (n int) {
	var l int
	_ = l
	l = len(m.LabelSelector)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.MetricLabelSelector)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *MetricValue) Size() (n int) {
	var l int
	_ = l
	l = m.DescribedObject.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Metric.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Timestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.WindowSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.WindowSeconds))
	}
	l = m.Value.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *MetricValueList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MetricIdent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MetricIdent{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Selector:` + fmt.Sprintf("%v", this.Selector) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MetricListOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MetricListOptions{`,
		`LabelSelector:` + fmt.Sprintf("%v", this.LabelSelector) + `,`,
		`MetricLabelSelector:` + fmt.Sprintf("%v", this.MetricLabelSelector) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MetricValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MetricValue{`,
		`DescribedObject:` + strings.Replace(strings.Replace(this.DescribedObject.String(), "ObjectReference", "k8s_io_api_core_v1.ObjectReference", 1), `&`, ``, 1) + `,`,
		`Metric:` + strings.Replace(strings.Replace(this.Metric.String(), "MetricIdent", "MetricIdent", 1), `&`, ``, 1) + `,`,
		`Timestamp:` + strings.Replace(strings.Replace(this.Timestamp.String(), "Time", "k8s_io_apimachinery_pkg_apis_meta_v1.Time", 1), `&`, ``, 1) + `,`,
		`WindowSeconds:` + valueToStringGenerated(this.WindowSeconds) + `,`,
		`Value:` + strings.Replace(strings.Replace(this.Value.String(), "Quantity", "k8s_io_apimachinery_pkg_api_resource.Quantity", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MetricValueList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MetricValueList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "MetricValue", "MetricValue", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MetricIdent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricIdent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricIdent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Selector = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MetricListOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricListOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricListOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LabelSelector", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LabelSelector = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricLabelSelector", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricLabelSelector = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MetricValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DescribedObject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DescribedObject.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metric.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowSeconds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WindowSeconds = &v
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MetricValueList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricValueList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricValueList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, MetricValue{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta2/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xee, 0x52, 0x8a, 0x30, 0x88, 0xc8, 0x12, 0x63, 0x83, 0xc9, 0xd2, 0xd4, 0x1b, 0x34, 0x32,
	0x13, 0xd0, 0x18, 0x13, 0xee, 0x36, 0xde, 0x90, 0x50, 0x89, 0x0b, 0x91, 0xc4, 0x9f, 0xe8, 0xec,
	0xee, 0xa1, 0x8c, 0xed, 0xce, 0x6c, 0x66, 0x66, 0x4b, 0xb8, 0xf3, 0x11, 0x7c, 0x03, 0x1f, 0xc3,
	0x57, 0xe0, 0x12, 0xef, 0xb8, 0x22, 0x52, 0x5f, 0xc4, 0xec, 0xec, 0x2c, 0x6d, 0x29, 0xa8, 0x78,
	0xb7, 0x7b, 0xe6, 0xfb, 0xbe, 0xf3, 0x9d, 0x39, 0xdf, 0xa0, 0xbd, 0xce, 0x0b, 0x85, 0x99, 0x20,
	0x9d, 0x2c, 0x04, 0xc9, 0x41, 0x83, 0x22, 0x3d, 0xe0, 0xb1, 0x90, 0xc4, 0x1e, 0x24, 0xa0, 0x25,
	0x8b, 0x14, 0x49, 0x3b, 0x6d, 0x42, 0x53, 0xa6, 0x48, 0x94, 0x29, 0x2d, 0x92, 0x8f, 0x65, 0xbd,
	0xb7, 0x16, 0x82, 0xa6, 0xeb, 0xa4, 0x0d, 0x1c, 0x24, 0xd5, 0x10, 0xe3, 0x54, 0x0a, 0x2d, 0x5c,
	0x5c, 0xf0, 0xb1, 0xc5, 0xe1, 0xb4, 0xd3, 0xc6, 0x39, 0x1f, 0x8f, 0xf2, 0xb1, 0xe5, 0x2f, 0xad,
	0xb6, 0x99, 0x3e, 0xc8, 0x42, 0x1c, 0x89, 0x84, 0xb4, 0x45, 0x5b, 0x10, 0x23, 0x13, 0x66, 0xfb,
	0xe6, 0xcf, 0xfc, 0x98, 0xaf, 0x42, 0x7e, 0xa9, 0x69, 0xed, 0xd1, 0x94, 0x91, 0x48, 0x48, 0x20,
	0xbd, 0xb5, 0xcb, 0x16, 0x96, 0x9e, 0x0d, 0x30, 0x09, 0x8d, 0x0e, 0x18, 0x07, 0x79, 0x54, 0xce,
	0x41, 0x24, 0x28, 0x91, 0xc9, 0x08, 0x6e, 0xc4, 0x52, 0xf9, 0x75, 0xd0, 0xab, 0x7a, 0x91, 0xeb,
	0x58, 0x32, 0xe3, 0x9a, 0x25, 0xe3, 0x6d, 0x9e, 0xff, 0x8d, 0xa0, 0xa2, 0x03, 0x48, 0xe8, 0x18,
	0xef, 0xe9, 0x75, 0xbc, 0x4c, 0xb3, 0x2e, 0x61, 0x5c, 0x2b, 0x2d, 0x2f, 0x93, 0x9a, 0x1f, 0xd0,
	0x6c, 0xcb, 0xdc, 0xf7, 0x66, 0x0c, 0x5c, 0xbb, 0x0d, 0x34, 0xc9, 0x69, 0x02, 0x75, 0xa7, 0xe1,
	0xac, 0xcc, 0xf8, 0xb7, 0x8f, 0xcf, 0x96, 0x2b, 0xfd, 0xb3, 0xe5, 0xc9, 0x57, 0x34, 0x81, 0xc0,
	0x9c, 0xb8, 0x4f, 0xd0, 0xb4, 0x82, 0x2e, 0x44, 0x5a, 0xc8, 0xfa, 0x84, 0x41, 0xdd, 0xb5, 0xa8,
	0xe9, 0x1d, 0x5b, 0x0f, 0x2e, 0x10, 0xcd, 0x6f, 0x0e, 0x5a, 0x28, 0xf4, 0xb7, 0x98, 0xd2, 0xdb,
	0xa9, 0x66, 0x82, 0x2b, 0x77, 0x03, 0xcd, 0x75, 0x69, 0x08, 0xdd, 0x92, 0x60, 0xdb, 0xdd, 0xb3,
	0x42, 0x73, 0x5b, 0xc3, 0x87, 0xc1, 0x28, 0xd6, 0x6d, 0xa1, 0xc5, 0x22, 0x21, 0x23, 0x28, 0xeb,
	0xe5, 0x81, 0x95, 0x58, 0x6c, 0x8d, 0x43, 0x82, 0xab, 0x78, 0xcd, 0xef, 0xd5, 0xf2, 0x06, 0xde,
	0xd0, 0x6e, 0x06, 0xee, 0x3e, 0x9a, 0x8f, 0x41, 0x45, 0x92, 0x85, 0x10, 0x6f, 0x87, 0x9f, 0x21,
	0xd2, 0xc6, 0xdd, 0xec, 0xfa, 0xc3, 0x32, 0xb7, 0x34, 0x65, 0x38, 0x0f, 0x16, 0xee, 0xad, 0xe1,
	0x02, 0x11, 0xc0, 0x3e, 0x48, 0xe0, 0x11, 0xf8, 0xf7, 0x6d, 0xff, 0xf9, 0x97, 0xa3, 0x1a, 0xc1,
	0x65, 0x51, 0x37, 0x42, 0x53, 0x85, 0x1d, 0xe3, 0x7c, 0x76, 0x7d, 0xe3, 0x86, 0xcf, 0x02, 0x0f,
	0xad, 0xcd, 0xbf, 0x63, 0xdb, 0x4e, 0x15, 0xc5, 0xc0, 0x4a, 0xbb, 0xef, 0xd0, 0x4c, 0x9e, 0x18,
	0xa5, 0x69, 0x92, 0xd6, 0xab, 0xa6, 0xcf, 0xe3, 0xa1, 0x31, 0x2e, 0x62, 0x32, 0x68, 0x96, 0xa7,
	0x38, 0x1f, 0x6c, 0x97, 0x25, 0xe0, 0x2f, 0x58, 0xd9, 0x99, 0xdd, 0x52, 0x24, 0x18, 0xe8, 0xb9,
	0x8f, 0xd0, 0xd4, 0x21, 0xe3, 0xb1, 0x38, 0xac, 0x4f, 0x36, 0x9c, 0x95, 0xaa, 0xbf, 0x90, 0xaf,
	0x6e, 0xcf, 0x54, 0x76, 0x20, 0x12, 0x3c, 0x56, 0x81, 0x05, 0xb8, 0x3b, 0xa8, 0xd6, 0xcb, 0x6f,
	0xb7, 0x5e, 0x33, 0x1e, 0xf0, 0x9f, 0x3c, 0xe0, 0xf2, 0xfd, 0xe1, 0xd7, 0x19, 0xe5, 0x9a, 0xe9,
	0x23, 0x7f, 0xce, 0xfa, 0xa8, 0x99, 0x15, 0x05, 0x85, 0x56, 0xf3, 0x87, 0x83, 0xe6, 0x87, 0x36,
	0x97, 0x07, 0xcc, 0x7d, 0x8f, 0xa6, 0xf3, 0x09, 0x62, 0xaa, 0xa9, 0x5d, 0x1b, 0xfe, 0xb7, 0x79,
	0x73, 0x76, 0x0b, 0x34, 0x1d, 0xa4, 0xb9, 0xac, 0x04, 0x17, 0x8a, 0xee, 0x27, 0x54, 0x63, 0x1a,
	0x12, 0x55, 0x9f, 0x68, 0x54, 0xff, 0x7f, 0x65, 0xc6, 0xed, 0x60, 0xa6, 0xcd, 0x5c, 0x31, 0x28,
	0x84, 0xfd, 0xd5, 0xe3, 0x73, 0xaf, 0x72, 0x72, 0xee, 0x55, 0x4e, 0xcf, 0xbd, 0xca, 0x97, 0xbe,
	0xe7, 0x1c, 0xf7, 0x3d, 0xe7, 0xa4, 0xef, 0x39, 0xa7, 0x7d, 0xcf, 0xf9, 0xd9, 0xf7, 0x9c, 0xaf,
	0xbf, 0xbc, 0xca, 0xdb, 0x5b, 0x56, 0xf0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0xd0, 0xb1,
	0xae, 0xa4, 0x05, 0x00, 0x00,
}
