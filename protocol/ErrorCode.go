// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"fmt"
	"io"
	"reflect"
)

type ErrorCodeEnum uint8
type ErrorCodeValues struct {
	INTERNAL_ERROR               ErrorCodeEnum
	PARTITION_LEADER_MISMATCH    ErrorCodeEnum
	UNSUPPORTED_MESSAGE          ErrorCodeEnum
	INVALID_CLIENT_VERSION       ErrorCodeEnum
	MALFORMED_REQUEST            ErrorCodeEnum
	INVALID_MESSAGE_TEMPLATE     ErrorCodeEnum
	INVALID_DEPLOYMENT_PARTITION ErrorCodeEnum
	PROCESS_NOT_FOUND            ErrorCodeEnum
	RESOURCE_EXHAUSTED           ErrorCodeEnum
	NullValue                    ErrorCodeEnum
}

var ErrorCode = ErrorCodeValues{0, 1, 2, 3, 4, 5, 6, 7, 8, 255}

func (e ErrorCodeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *ErrorCodeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e ErrorCodeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ErrorCode)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ErrorCode, unknown enumeration value %d", e)
}

func (*ErrorCodeEnum) EncodedLength() int64 {
	return 1
}

func (*ErrorCodeEnum) INTERNAL_ERRORSinceVersion() uint16 {
	return 0
}

func (e *ErrorCodeEnum) INTERNAL_ERRORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.INTERNAL_ERRORSinceVersion()
}

func (*ErrorCodeEnum) INTERNAL_ERRORDeprecated() uint16 {
	return 0
}

func (*ErrorCodeEnum) PARTITION_LEADER_MISMATCHSinceVersion() uint16 {
	return 0
}

func (e *ErrorCodeEnum) PARTITION_LEADER_MISMATCHInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PARTITION_LEADER_MISMATCHSinceVersion()
}

func (*ErrorCodeEnum) PARTITION_LEADER_MISMATCHDeprecated() uint16 {
	return 0
}

func (*ErrorCodeEnum) UNSUPPORTED_MESSAGESinceVersion() uint16 {
	return 0
}

func (e *ErrorCodeEnum) UNSUPPORTED_MESSAGEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UNSUPPORTED_MESSAGESinceVersion()
}

func (*ErrorCodeEnum) UNSUPPORTED_MESSAGEDeprecated() uint16 {
	return 0
}

func (*ErrorCodeEnum) INVALID_CLIENT_VERSIONSinceVersion() uint16 {
	return 0
}

func (e *ErrorCodeEnum) INVALID_CLIENT_VERSIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.INVALID_CLIENT_VERSIONSinceVersion()
}

func (*ErrorCodeEnum) INVALID_CLIENT_VERSIONDeprecated() uint16 {
	return 0
}

func (*ErrorCodeEnum) MALFORMED_REQUESTSinceVersion() uint16 {
	return 0
}

func (e *ErrorCodeEnum) MALFORMED_REQUESTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MALFORMED_REQUESTSinceVersion()
}

func (*ErrorCodeEnum) MALFORMED_REQUESTDeprecated() uint16 {
	return 0
}

func (*ErrorCodeEnum) INVALID_MESSAGE_TEMPLATESinceVersion() uint16 {
	return 0
}

func (e *ErrorCodeEnum) INVALID_MESSAGE_TEMPLATEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.INVALID_MESSAGE_TEMPLATESinceVersion()
}

func (*ErrorCodeEnum) INVALID_MESSAGE_TEMPLATEDeprecated() uint16 {
	return 0
}

func (*ErrorCodeEnum) INVALID_DEPLOYMENT_PARTITIONSinceVersion() uint16 {
	return 0
}

func (e *ErrorCodeEnum) INVALID_DEPLOYMENT_PARTITIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.INVALID_DEPLOYMENT_PARTITIONSinceVersion()
}

func (*ErrorCodeEnum) INVALID_DEPLOYMENT_PARTITIONDeprecated() uint16 {
	return 0
}

func (*ErrorCodeEnum) PROCESS_NOT_FOUNDSinceVersion() uint16 {
	return 0
}

func (e *ErrorCodeEnum) PROCESS_NOT_FOUNDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PROCESS_NOT_FOUNDSinceVersion()
}

func (*ErrorCodeEnum) PROCESS_NOT_FOUNDDeprecated() uint16 {
	return 0
}

func (*ErrorCodeEnum) RESOURCE_EXHAUSTEDSinceVersion() uint16 {
	return 0
}

func (e *ErrorCodeEnum) RESOURCE_EXHAUSTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RESOURCE_EXHAUSTEDSinceVersion()
}

func (*ErrorCodeEnum) RESOURCE_EXHAUSTEDDeprecated() uint16 {
	return 0
}
