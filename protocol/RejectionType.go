// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"fmt"
	"io"
	"reflect"
)

type RejectionTypeEnum uint8
type RejectionTypeValues struct {
	INVALID_ARGUMENT           RejectionTypeEnum
	NOT_FOUND                  RejectionTypeEnum
	ALREADY_EXISTS             RejectionTypeEnum
	INVALID_STATE              RejectionTypeEnum
	PROCESSING_ERROR           RejectionTypeEnum
	EXCEEDED_BATCH_RECORD_SIZE RejectionTypeEnum
	NullValue                  RejectionTypeEnum
}

var RejectionType = RejectionTypeValues{0, 1, 2, 3, 4, 5, 255}

func (r RejectionTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(r)); err != nil {
		return err
	}
	return nil
}

func (r *RejectionTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(r)); err != nil {
		return err
	}
	return nil
}

func (r RejectionTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(RejectionType)
	for idx := 0; idx < value.NumField(); idx++ {
		if r == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on RejectionType, unknown enumeration value %d", r)
}

func (*RejectionTypeEnum) EncodedLength() int64 {
	return 1
}

func (*RejectionTypeEnum) INVALID_ARGUMENTSinceVersion() uint16 {
	return 0
}

func (r *RejectionTypeEnum) INVALID_ARGUMENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.INVALID_ARGUMENTSinceVersion()
}

func (*RejectionTypeEnum) INVALID_ARGUMENTDeprecated() uint16 {
	return 0
}

func (*RejectionTypeEnum) NOT_FOUNDSinceVersion() uint16 {
	return 0
}

func (r *RejectionTypeEnum) NOT_FOUNDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.NOT_FOUNDSinceVersion()
}

func (*RejectionTypeEnum) NOT_FOUNDDeprecated() uint16 {
	return 0
}

func (*RejectionTypeEnum) ALREADY_EXISTSSinceVersion() uint16 {
	return 0
}

func (r *RejectionTypeEnum) ALREADY_EXISTSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ALREADY_EXISTSSinceVersion()
}

func (*RejectionTypeEnum) ALREADY_EXISTSDeprecated() uint16 {
	return 0
}

func (*RejectionTypeEnum) INVALID_STATESinceVersion() uint16 {
	return 0
}

func (r *RejectionTypeEnum) INVALID_STATEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.INVALID_STATESinceVersion()
}

func (*RejectionTypeEnum) INVALID_STATEDeprecated() uint16 {
	return 0
}

func (*RejectionTypeEnum) PROCESSING_ERRORSinceVersion() uint16 {
	return 0
}

func (r *RejectionTypeEnum) PROCESSING_ERRORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.PROCESSING_ERRORSinceVersion()
}

func (*RejectionTypeEnum) PROCESSING_ERRORDeprecated() uint16 {
	return 0
}

func (*RejectionTypeEnum) EXCEEDED_BATCH_RECORD_SIZESinceVersion() uint16 {
	return 0
}

func (r *RejectionTypeEnum) EXCEEDED_BATCH_RECORD_SIZEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.EXCEEDED_BATCH_RECORD_SIZESinceVersion()
}

func (*RejectionTypeEnum) EXCEEDED_BATCH_RECORD_SIZEDeprecated() uint16 {
	return 0
}
