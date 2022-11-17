// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"fmt"
	"io"
	"reflect"
)

type RecordTypeEnum uint8
type RecordTypeValues struct {
	EVENT             RecordTypeEnum
	COMMAND           RecordTypeEnum
	COMMAND_REJECTION RecordTypeEnum
	NullValue         RecordTypeEnum
}

var RecordType = RecordTypeValues{0, 1, 2, 255}

func (r RecordTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(r)); err != nil {
		return err
	}
	return nil
}

func (r *RecordTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(r)); err != nil {
		return err
	}
	return nil
}

func (r RecordTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(RecordType)
	for idx := 0; idx < value.NumField(); idx++ {
		if r == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on RecordType, unknown enumeration value %d", r)
}

func (*RecordTypeEnum) EncodedLength() int64 {
	return 1
}

func (*RecordTypeEnum) EVENTSinceVersion() uint16 {
	return 0
}

func (r *RecordTypeEnum) EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.EVENTSinceVersion()
}

func (*RecordTypeEnum) EVENTDeprecated() uint16 {
	return 0
}

func (*RecordTypeEnum) COMMANDSinceVersion() uint16 {
	return 0
}

func (r *RecordTypeEnum) COMMANDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.COMMANDSinceVersion()
}

func (*RecordTypeEnum) COMMANDDeprecated() uint16 {
	return 0
}

func (*RecordTypeEnum) COMMAND_REJECTIONSinceVersion() uint16 {
	return 0
}

func (r *RecordTypeEnum) COMMAND_REJECTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.COMMAND_REJECTIONSinceVersion()
}

func (*RecordTypeEnum) COMMAND_REJECTIONDeprecated() uint16 {
	return 0
}
