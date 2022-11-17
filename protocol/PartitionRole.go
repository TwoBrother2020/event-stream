// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"fmt"
	"io"
	"reflect"
)

type PartitionRoleEnum uint8
type PartitionRoleValues struct {
	LEADER    PartitionRoleEnum
	FOLLOWER  PartitionRoleEnum
	INACTIVE  PartitionRoleEnum
	NullValue PartitionRoleEnum
}

var PartitionRole = PartitionRoleValues{0, 1, 2, 255}

func (p PartitionRoleEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(p)); err != nil {
		return err
	}
	return nil
}

func (p *PartitionRoleEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(p)); err != nil {
		return err
	}
	return nil
}

func (p PartitionRoleEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(PartitionRole)
	for idx := 0; idx < value.NumField(); idx++ {
		if p == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on PartitionRole, unknown enumeration value %d", p)
}

func (*PartitionRoleEnum) EncodedLength() int64 {
	return 1
}

func (*PartitionRoleEnum) LEADERSinceVersion() uint16 {
	return 0
}

func (p *PartitionRoleEnum) LEADERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.LEADERSinceVersion()
}

func (*PartitionRoleEnum) LEADERDeprecated() uint16 {
	return 0
}

func (*PartitionRoleEnum) FOLLOWERSinceVersion() uint16 {
	return 0
}

func (p *PartitionRoleEnum) FOLLOWERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.FOLLOWERSinceVersion()
}

func (*PartitionRoleEnum) FOLLOWERDeprecated() uint16 {
	return 0
}

func (*PartitionRoleEnum) INACTIVESinceVersion() uint16 {
	return 0
}

func (p *PartitionRoleEnum) INACTIVEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.INACTIVESinceVersion()
}

func (*PartitionRoleEnum) INACTIVEDeprecated() uint16 {
	return 0
}
