// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"unicode/utf8"
)

type BrokerInfo struct {
	NodeId               int32
	PartitionsCount      int32
	ClusterSize          int32
	ReplicationFactor    int32
	Addresses            []BrokerInfoAddresses
	PartitionRoles       []BrokerInfoPartitionRoles
	PartitionLeaderTerms []BrokerInfoPartitionLeaderTerms
	PartitionHealth      []BrokerInfoPartitionHealth
	Version              []uint8
}
type BrokerInfoAddresses struct {
	ApiName []uint8
	Address []uint8
}
type BrokerInfoPartitionRoles struct {
	PartitionId int32
	Role        PartitionRoleEnum
}
type BrokerInfoPartitionLeaderTerms struct {
	PartitionId int32
	Term        int64
}
type BrokerInfoPartitionHealth struct {
	PartitionId  int32
	HealthStatus PartitionHealthStatusEnum
}

func (b *BrokerInfo) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := b.RangeCheck(b.SbeSchemaVersion(), b.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, b.NodeId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, b.PartitionsCount); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, b.ClusterSize); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, b.ReplicationFactor); err != nil {
		return err
	}
	var AddressesBlockLength uint16 = 0
	if err := _m.WriteUint16(_w, AddressesBlockLength); err != nil {
		return err
	}
	var AddressesNumInGroup uint8 = uint8(len(b.Addresses))
	if err := _m.WriteUint8(_w, AddressesNumInGroup); err != nil {
		return err
	}
	for _, prop := range b.Addresses {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var PartitionRolesBlockLength uint16 = 5
	if err := _m.WriteUint16(_w, PartitionRolesBlockLength); err != nil {
		return err
	}
	var PartitionRolesNumInGroup uint8 = uint8(len(b.PartitionRoles))
	if err := _m.WriteUint8(_w, PartitionRolesNumInGroup); err != nil {
		return err
	}
	for _, prop := range b.PartitionRoles {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var PartitionLeaderTermsBlockLength uint16 = 12
	if err := _m.WriteUint16(_w, PartitionLeaderTermsBlockLength); err != nil {
		return err
	}
	var PartitionLeaderTermsNumInGroup uint8 = uint8(len(b.PartitionLeaderTerms))
	if err := _m.WriteUint8(_w, PartitionLeaderTermsNumInGroup); err != nil {
		return err
	}
	for _, prop := range b.PartitionLeaderTerms {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var PartitionHealthBlockLength uint16 = 5
	if err := _m.WriteUint16(_w, PartitionHealthBlockLength); err != nil {
		return err
	}
	var PartitionHealthNumInGroup uint8 = uint8(len(b.PartitionHealth))
	if err := _m.WriteUint8(_w, PartitionHealthNumInGroup); err != nil {
		return err
	}
	for _, prop := range b.PartitionHealth {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, uint32(len(b.Version))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, b.Version); err != nil {
		return err
	}
	return nil
}

func (b *BrokerInfo) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !b.NodeIdInActingVersion(actingVersion) {
		b.NodeId = b.NodeIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &b.NodeId); err != nil {
			return err
		}
	}
	if !b.PartitionsCountInActingVersion(actingVersion) {
		b.PartitionsCount = b.PartitionsCountNullValue()
	} else {
		if err := _m.ReadInt32(_r, &b.PartitionsCount); err != nil {
			return err
		}
	}
	if !b.ClusterSizeInActingVersion(actingVersion) {
		b.ClusterSize = b.ClusterSizeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &b.ClusterSize); err != nil {
			return err
		}
	}
	if !b.ReplicationFactorInActingVersion(actingVersion) {
		b.ReplicationFactor = b.ReplicationFactorNullValue()
	} else {
		if err := _m.ReadInt32(_r, &b.ReplicationFactor); err != nil {
			return err
		}
	}
	if actingVersion > b.SbeSchemaVersion() && blockLength > b.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-b.SbeBlockLength()))
	}

	if b.AddressesInActingVersion(actingVersion) {
		var AddressesBlockLength uint16
		if err := _m.ReadUint16(_r, &AddressesBlockLength); err != nil {
			return err
		}
		var AddressesNumInGroup uint8
		if err := _m.ReadUint8(_r, &AddressesNumInGroup); err != nil {
			return err
		}
		if cap(b.Addresses) < int(AddressesNumInGroup) {
			b.Addresses = make([]BrokerInfoAddresses, AddressesNumInGroup)
		}
		b.Addresses = b.Addresses[:AddressesNumInGroup]
		for i := range b.Addresses {
			if err := b.Addresses[i].Decode(_m, _r, actingVersion, uint(AddressesBlockLength)); err != nil {
				return err
			}
		}
	}

	if b.PartitionRolesInActingVersion(actingVersion) {
		var PartitionRolesBlockLength uint16
		if err := _m.ReadUint16(_r, &PartitionRolesBlockLength); err != nil {
			return err
		}
		var PartitionRolesNumInGroup uint8
		if err := _m.ReadUint8(_r, &PartitionRolesNumInGroup); err != nil {
			return err
		}
		if cap(b.PartitionRoles) < int(PartitionRolesNumInGroup) {
			b.PartitionRoles = make([]BrokerInfoPartitionRoles, PartitionRolesNumInGroup)
		}
		b.PartitionRoles = b.PartitionRoles[:PartitionRolesNumInGroup]
		for i := range b.PartitionRoles {
			if err := b.PartitionRoles[i].Decode(_m, _r, actingVersion, uint(PartitionRolesBlockLength)); err != nil {
				return err
			}
		}
	}

	if b.PartitionLeaderTermsInActingVersion(actingVersion) {
		var PartitionLeaderTermsBlockLength uint16
		if err := _m.ReadUint16(_r, &PartitionLeaderTermsBlockLength); err != nil {
			return err
		}
		var PartitionLeaderTermsNumInGroup uint8
		if err := _m.ReadUint8(_r, &PartitionLeaderTermsNumInGroup); err != nil {
			return err
		}
		if cap(b.PartitionLeaderTerms) < int(PartitionLeaderTermsNumInGroup) {
			b.PartitionLeaderTerms = make([]BrokerInfoPartitionLeaderTerms, PartitionLeaderTermsNumInGroup)
		}
		b.PartitionLeaderTerms = b.PartitionLeaderTerms[:PartitionLeaderTermsNumInGroup]
		for i := range b.PartitionLeaderTerms {
			if err := b.PartitionLeaderTerms[i].Decode(_m, _r, actingVersion, uint(PartitionLeaderTermsBlockLength)); err != nil {
				return err
			}
		}
	}

	if b.PartitionHealthInActingVersion(actingVersion) {
		var PartitionHealthBlockLength uint16
		if err := _m.ReadUint16(_r, &PartitionHealthBlockLength); err != nil {
			return err
		}
		var PartitionHealthNumInGroup uint8
		if err := _m.ReadUint8(_r, &PartitionHealthNumInGroup); err != nil {
			return err
		}
		if cap(b.PartitionHealth) < int(PartitionHealthNumInGroup) {
			b.PartitionHealth = make([]BrokerInfoPartitionHealth, PartitionHealthNumInGroup)
		}
		b.PartitionHealth = b.PartitionHealth[:PartitionHealthNumInGroup]
		for i := range b.PartitionHealth {
			if err := b.PartitionHealth[i].Decode(_m, _r, actingVersion, uint(PartitionHealthBlockLength)); err != nil {
				return err
			}
		}
	}

	if b.VersionInActingVersion(actingVersion) {
		var VersionLength uint32
		if err := _m.ReadUint32(_r, &VersionLength); err != nil {
			return err
		}
		if cap(b.Version) < int(VersionLength) {
			b.Version = make([]uint8, VersionLength)
		}
		b.Version = b.Version[:VersionLength]
		if err := _m.ReadBytes(_r, b.Version); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := b.RangeCheck(actingVersion, b.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (b *BrokerInfo) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if b.NodeIdInActingVersion(actingVersion) {
		if b.NodeId < b.NodeIdMinValue() || b.NodeId > b.NodeIdMaxValue() {
			return fmt.Errorf("Range check failed on b.NodeId (%v < %v > %v)", b.NodeIdMinValue(), b.NodeId, b.NodeIdMaxValue())
		}
	}
	if b.PartitionsCountInActingVersion(actingVersion) {
		if b.PartitionsCount < b.PartitionsCountMinValue() || b.PartitionsCount > b.PartitionsCountMaxValue() {
			return fmt.Errorf("Range check failed on b.PartitionsCount (%v < %v > %v)", b.PartitionsCountMinValue(), b.PartitionsCount, b.PartitionsCountMaxValue())
		}
	}
	if b.ClusterSizeInActingVersion(actingVersion) {
		if b.ClusterSize < b.ClusterSizeMinValue() || b.ClusterSize > b.ClusterSizeMaxValue() {
			return fmt.Errorf("Range check failed on b.ClusterSize (%v < %v > %v)", b.ClusterSizeMinValue(), b.ClusterSize, b.ClusterSizeMaxValue())
		}
	}
	if b.ReplicationFactorInActingVersion(actingVersion) {
		if b.ReplicationFactor < b.ReplicationFactorMinValue() || b.ReplicationFactor > b.ReplicationFactorMaxValue() {
			return fmt.Errorf("Range check failed on b.ReplicationFactor (%v < %v > %v)", b.ReplicationFactorMinValue(), b.ReplicationFactor, b.ReplicationFactorMaxValue())
		}
	}
	for _, prop := range b.Addresses {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range b.PartitionRoles {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range b.PartitionLeaderTerms {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range b.PartitionHealth {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(b.Version[:]) {
		return errors.New("b.Version failed UTF-8 validation")
	}
	return nil
}

func BrokerInfoInit(b *BrokerInfo) {
	return
}

func (b *BrokerInfoAddresses) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint32(_w, uint32(len(b.ApiName))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, b.ApiName); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(b.Address))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, b.Address); err != nil {
		return err
	}
	return nil
}

func (b *BrokerInfoAddresses) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if actingVersion > b.SbeSchemaVersion() && blockLength > b.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-b.SbeBlockLength()))
	}

	if b.ApiNameInActingVersion(actingVersion) {
		var ApiNameLength uint32
		if err := _m.ReadUint32(_r, &ApiNameLength); err != nil {
			return err
		}
		if cap(b.ApiName) < int(ApiNameLength) {
			b.ApiName = make([]uint8, ApiNameLength)
		}
		b.ApiName = b.ApiName[:ApiNameLength]
		if err := _m.ReadBytes(_r, b.ApiName); err != nil {
			return err
		}
	}

	if b.AddressInActingVersion(actingVersion) {
		var AddressLength uint32
		if err := _m.ReadUint32(_r, &AddressLength); err != nil {
			return err
		}
		if cap(b.Address) < int(AddressLength) {
			b.Address = make([]uint8, AddressLength)
		}
		b.Address = b.Address[:AddressLength]
		if err := _m.ReadBytes(_r, b.Address); err != nil {
			return err
		}
	}
	return nil
}

func (b *BrokerInfoAddresses) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(b.ApiName[:]) {
		return errors.New("b.ApiName failed UTF-8 validation")
	}
	if !utf8.Valid(b.Address[:]) {
		return errors.New("b.Address failed UTF-8 validation")
	}
	return nil
}

func BrokerInfoAddressesInit(b *BrokerInfoAddresses) {
	return
}

func (b *BrokerInfoPartitionRoles) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, b.PartitionId); err != nil {
		return err
	}
	if err := b.Role.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (b *BrokerInfoPartitionRoles) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !b.PartitionIdInActingVersion(actingVersion) {
		b.PartitionId = b.PartitionIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &b.PartitionId); err != nil {
			return err
		}
	}
	if b.RoleInActingVersion(actingVersion) {
		if err := b.Role.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > b.SbeSchemaVersion() && blockLength > b.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-b.SbeBlockLength()))
	}
	return nil
}

func (b *BrokerInfoPartitionRoles) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if b.PartitionIdInActingVersion(actingVersion) {
		if b.PartitionId < b.PartitionIdMinValue() || b.PartitionId > b.PartitionIdMaxValue() {
			return fmt.Errorf("Range check failed on b.PartitionId (%v < %v > %v)", b.PartitionIdMinValue(), b.PartitionId, b.PartitionIdMaxValue())
		}
	}
	if err := b.Role.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func BrokerInfoPartitionRolesInit(b *BrokerInfoPartitionRoles) {
	return
}

func (b *BrokerInfoPartitionLeaderTerms) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, b.PartitionId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, b.Term); err != nil {
		return err
	}
	return nil
}

func (b *BrokerInfoPartitionLeaderTerms) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !b.PartitionIdInActingVersion(actingVersion) {
		b.PartitionId = b.PartitionIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &b.PartitionId); err != nil {
			return err
		}
	}
	if !b.TermInActingVersion(actingVersion) {
		b.Term = b.TermNullValue()
	} else {
		if err := _m.ReadInt64(_r, &b.Term); err != nil {
			return err
		}
	}
	if actingVersion > b.SbeSchemaVersion() && blockLength > b.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-b.SbeBlockLength()))
	}
	return nil
}

func (b *BrokerInfoPartitionLeaderTerms) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if b.PartitionIdInActingVersion(actingVersion) {
		if b.PartitionId < b.PartitionIdMinValue() || b.PartitionId > b.PartitionIdMaxValue() {
			return fmt.Errorf("Range check failed on b.PartitionId (%v < %v > %v)", b.PartitionIdMinValue(), b.PartitionId, b.PartitionIdMaxValue())
		}
	}
	if b.TermInActingVersion(actingVersion) {
		if b.Term < b.TermMinValue() || b.Term > b.TermMaxValue() {
			return fmt.Errorf("Range check failed on b.Term (%v < %v > %v)", b.TermMinValue(), b.Term, b.TermMaxValue())
		}
	}
	return nil
}

func BrokerInfoPartitionLeaderTermsInit(b *BrokerInfoPartitionLeaderTerms) {
	return
}

func (b *BrokerInfoPartitionHealth) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, b.PartitionId); err != nil {
		return err
	}
	if err := b.HealthStatus.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (b *BrokerInfoPartitionHealth) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !b.PartitionIdInActingVersion(actingVersion) {
		b.PartitionId = b.PartitionIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &b.PartitionId); err != nil {
			return err
		}
	}
	if b.HealthStatusInActingVersion(actingVersion) {
		if err := b.HealthStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > b.SbeSchemaVersion() && blockLength > b.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-b.SbeBlockLength()))
	}
	return nil
}

func (b *BrokerInfoPartitionHealth) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if b.PartitionIdInActingVersion(actingVersion) {
		if b.PartitionId < b.PartitionIdMinValue() || b.PartitionId > b.PartitionIdMaxValue() {
			return fmt.Errorf("Range check failed on b.PartitionId (%v < %v > %v)", b.PartitionIdMinValue(), b.PartitionId, b.PartitionIdMaxValue())
		}
	}
	if err := b.HealthStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func BrokerInfoPartitionHealthInit(b *BrokerInfoPartitionHealth) {
	return
}

func (*BrokerInfo) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*BrokerInfo) SbeTemplateId() (templateId uint16) {
	return 201
}

func (*BrokerInfo) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (*BrokerInfo) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*BrokerInfo) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*BrokerInfo) NodeIdId() uint16 {
	return 1
}

func (*BrokerInfo) NodeIdSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfo) NodeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.NodeIdSinceVersion()
}

func (*BrokerInfo) NodeIdDeprecated() uint16 {
	return 0
}

func (*BrokerInfo) NodeIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfo) NodeIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*BrokerInfo) NodeIdMaxValue() int32 {
	return math.MaxInt32
}

func (*BrokerInfo) NodeIdNullValue() int32 {
	return math.MinInt32
}

func (*BrokerInfo) PartitionsCountId() uint16 {
	return 2
}

func (*BrokerInfo) PartitionsCountSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfo) PartitionsCountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.PartitionsCountSinceVersion()
}

func (*BrokerInfo) PartitionsCountDeprecated() uint16 {
	return 0
}

func (*BrokerInfo) PartitionsCountMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfo) PartitionsCountMinValue() int32 {
	return math.MinInt32 + 1
}

func (*BrokerInfo) PartitionsCountMaxValue() int32 {
	return math.MaxInt32
}

func (*BrokerInfo) PartitionsCountNullValue() int32 {
	return math.MinInt32
}

func (*BrokerInfo) ClusterSizeId() uint16 {
	return 3
}

func (*BrokerInfo) ClusterSizeSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfo) ClusterSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.ClusterSizeSinceVersion()
}

func (*BrokerInfo) ClusterSizeDeprecated() uint16 {
	return 0
}

func (*BrokerInfo) ClusterSizeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfo) ClusterSizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*BrokerInfo) ClusterSizeMaxValue() int32 {
	return math.MaxInt32
}

func (*BrokerInfo) ClusterSizeNullValue() int32 {
	return math.MinInt32
}

func (*BrokerInfo) ReplicationFactorId() uint16 {
	return 4
}

func (*BrokerInfo) ReplicationFactorSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfo) ReplicationFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.ReplicationFactorSinceVersion()
}

func (*BrokerInfo) ReplicationFactorDeprecated() uint16 {
	return 0
}

func (*BrokerInfo) ReplicationFactorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfo) ReplicationFactorMinValue() int32 {
	return math.MinInt32 + 1
}

func (*BrokerInfo) ReplicationFactorMaxValue() int32 {
	return math.MaxInt32
}

func (*BrokerInfo) ReplicationFactorNullValue() int32 {
	return math.MinInt32
}

func (*BrokerInfoAddresses) ApiNameMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfoAddresses) ApiNameSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfoAddresses) ApiNameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.ApiNameSinceVersion()
}

func (*BrokerInfoAddresses) ApiNameDeprecated() uint16 {
	return 0
}

func (BrokerInfoAddresses) ApiNameCharacterEncoding() string {
	return "UTF-8"
}

func (BrokerInfoAddresses) ApiNameHeaderLength() uint64 {
	return 4
}

func (*BrokerInfoAddresses) AddressMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfoAddresses) AddressSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfoAddresses) AddressInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.AddressSinceVersion()
}

func (*BrokerInfoAddresses) AddressDeprecated() uint16 {
	return 0
}

func (BrokerInfoAddresses) AddressCharacterEncoding() string {
	return "UTF-8"
}

func (BrokerInfoAddresses) AddressHeaderLength() uint64 {
	return 4
}

func (*BrokerInfoPartitionRoles) PartitionIdId() uint16 {
	return 9
}

func (*BrokerInfoPartitionRoles) PartitionIdSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfoPartitionRoles) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.PartitionIdSinceVersion()
}

func (*BrokerInfoPartitionRoles) PartitionIdDeprecated() uint16 {
	return 0
}

func (*BrokerInfoPartitionRoles) PartitionIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfoPartitionRoles) PartitionIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*BrokerInfoPartitionRoles) PartitionIdMaxValue() int32 {
	return math.MaxInt32
}

func (*BrokerInfoPartitionRoles) PartitionIdNullValue() int32 {
	return math.MinInt32
}

func (*BrokerInfoPartitionRoles) RoleId() uint16 {
	return 10
}

func (*BrokerInfoPartitionRoles) RoleSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfoPartitionRoles) RoleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.RoleSinceVersion()
}

func (*BrokerInfoPartitionRoles) RoleDeprecated() uint16 {
	return 0
}

func (*BrokerInfoPartitionRoles) RoleMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfoPartitionLeaderTerms) PartitionIdId() uint16 {
	return 12
}

func (*BrokerInfoPartitionLeaderTerms) PartitionIdSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfoPartitionLeaderTerms) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.PartitionIdSinceVersion()
}

func (*BrokerInfoPartitionLeaderTerms) PartitionIdDeprecated() uint16 {
	return 0
}

func (*BrokerInfoPartitionLeaderTerms) PartitionIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfoPartitionLeaderTerms) PartitionIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*BrokerInfoPartitionLeaderTerms) PartitionIdMaxValue() int32 {
	return math.MaxInt32
}

func (*BrokerInfoPartitionLeaderTerms) PartitionIdNullValue() int32 {
	return math.MinInt32
}

func (*BrokerInfoPartitionLeaderTerms) TermId() uint16 {
	return 13
}

func (*BrokerInfoPartitionLeaderTerms) TermSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfoPartitionLeaderTerms) TermInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.TermSinceVersion()
}

func (*BrokerInfoPartitionLeaderTerms) TermDeprecated() uint16 {
	return 0
}

func (*BrokerInfoPartitionLeaderTerms) TermMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfoPartitionLeaderTerms) TermMinValue() int64 {
	return math.MinInt64 + 1
}

func (*BrokerInfoPartitionLeaderTerms) TermMaxValue() int64 {
	return math.MaxInt64
}

func (*BrokerInfoPartitionLeaderTerms) TermNullValue() int64 {
	return math.MinInt64
}

func (*BrokerInfoPartitionHealth) PartitionIdId() uint16 {
	return 16
}

func (*BrokerInfoPartitionHealth) PartitionIdSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfoPartitionHealth) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.PartitionIdSinceVersion()
}

func (*BrokerInfoPartitionHealth) PartitionIdDeprecated() uint16 {
	return 0
}

func (*BrokerInfoPartitionHealth) PartitionIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfoPartitionHealth) PartitionIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*BrokerInfoPartitionHealth) PartitionIdMaxValue() int32 {
	return math.MaxInt32
}

func (*BrokerInfoPartitionHealth) PartitionIdNullValue() int32 {
	return math.MinInt32
}

func (*BrokerInfoPartitionHealth) HealthStatusId() uint16 {
	return 17
}

func (*BrokerInfoPartitionHealth) HealthStatusSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfoPartitionHealth) HealthStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.HealthStatusSinceVersion()
}

func (*BrokerInfoPartitionHealth) HealthStatusDeprecated() uint16 {
	return 0
}

func (*BrokerInfoPartitionHealth) HealthStatusMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfo) AddressesId() uint16 {
	return 5
}

func (*BrokerInfo) AddressesSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfo) AddressesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.AddressesSinceVersion()
}

func (*BrokerInfo) AddressesDeprecated() uint16 {
	return 0
}

func (*BrokerInfoAddresses) SbeBlockLength() (blockLength uint) {
	return 0
}

func (*BrokerInfoAddresses) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*BrokerInfo) PartitionRolesId() uint16 {
	return 8
}

func (*BrokerInfo) PartitionRolesSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfo) PartitionRolesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.PartitionRolesSinceVersion()
}

func (*BrokerInfo) PartitionRolesDeprecated() uint16 {
	return 0
}

func (*BrokerInfoPartitionRoles) SbeBlockLength() (blockLength uint) {
	return 5
}

func (*BrokerInfoPartitionRoles) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*BrokerInfo) PartitionLeaderTermsId() uint16 {
	return 11
}

func (*BrokerInfo) PartitionLeaderTermsSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfo) PartitionLeaderTermsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.PartitionLeaderTermsSinceVersion()
}

func (*BrokerInfo) PartitionLeaderTermsDeprecated() uint16 {
	return 0
}

func (*BrokerInfoPartitionLeaderTerms) SbeBlockLength() (blockLength uint) {
	return 12
}

func (*BrokerInfoPartitionLeaderTerms) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*BrokerInfo) PartitionHealthId() uint16 {
	return 15
}

func (*BrokerInfo) PartitionHealthSinceVersion() uint16 {
	return 3
}

func (b *BrokerInfo) PartitionHealthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.PartitionHealthSinceVersion()
}

func (*BrokerInfo) PartitionHealthDeprecated() uint16 {
	return 0
}

func (*BrokerInfoPartitionHealth) SbeBlockLength() (blockLength uint) {
	return 5
}

func (*BrokerInfoPartitionHealth) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*BrokerInfo) VersionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*BrokerInfo) VersionSinceVersion() uint16 {
	return 0
}

func (b *BrokerInfo) VersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.VersionSinceVersion()
}

func (*BrokerInfo) VersionDeprecated() uint16 {
	return 0
}

func (BrokerInfo) VersionCharacterEncoding() string {
	return "UTF-8"
}

func (BrokerInfo) VersionHeaderLength() uint64 {
	return 4
}
