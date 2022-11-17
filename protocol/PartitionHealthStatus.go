// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"fmt"
	"io"
	"reflect"
)

type PartitionHealthStatusEnum uint8
type PartitionHealthStatusValues struct {
	UNHEALTHY PartitionHealthStatusEnum
	HEALTHY   PartitionHealthStatusEnum
	DEAD      PartitionHealthStatusEnum
	NullValue PartitionHealthStatusEnum
}

var PartitionHealthStatus = PartitionHealthStatusValues{0, 1, 2, 255}

func (p PartitionHealthStatusEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(p)); err != nil {
		return err
	}
	return nil
}

func (p *PartitionHealthStatusEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(p)); err != nil {
		return err
	}
	return nil
}

func (p PartitionHealthStatusEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(PartitionHealthStatus)
	for idx := 0; idx < value.NumField(); idx++ {
		if p == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on PartitionHealthStatus, unknown enumeration value %d", p)
}

func (*PartitionHealthStatusEnum) EncodedLength() int64 {
	return 1
}

func (*PartitionHealthStatusEnum) UNHEALTHYSinceVersion() uint16 {
	return 0
}

func (p *PartitionHealthStatusEnum) UNHEALTHYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.UNHEALTHYSinceVersion()
}

func (*PartitionHealthStatusEnum) UNHEALTHYDeprecated() uint16 {
	return 0
}

func (*PartitionHealthStatusEnum) HEALTHYSinceVersion() uint16 {
	return 0
}

func (p *PartitionHealthStatusEnum) HEALTHYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.HEALTHYSinceVersion()
}

func (*PartitionHealthStatusEnum) HEALTHYDeprecated() uint16 {
	return 0
}

func (*PartitionHealthStatusEnum) DEADSinceVersion() uint16 {
	return 0
}

func (p *PartitionHealthStatusEnum) DEADInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.DEADSinceVersion()
}

func (*PartitionHealthStatusEnum) DEADDeprecated() uint16 {
	return 0
}
