// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type ErrorResponse struct {
	ErrorCode errorCodeEnum
	ErrorData []uint8
}

func (e *ErrorResponse) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := e.ErrorCode.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(e.ErrorData))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.ErrorData); err != nil {
		return err
	}
	return nil
}

func (e *ErrorResponse) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if e.ErrorCodeInActingVersion(actingVersion) {
		if err := e.ErrorCode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}

	if e.ErrorDataInActingVersion(actingVersion) {
		var ErrorDataLength uint32
		if err := _m.ReadUint32(_r, &ErrorDataLength); err != nil {
			return err
		}
		if cap(e.ErrorData) < int(ErrorDataLength) {
			e.ErrorData = make([]uint8, ErrorDataLength)
		}
		e.ErrorData = e.ErrorData[:ErrorDataLength]
		if err := _m.ReadBytes(_r, e.ErrorData); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *ErrorResponse) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := e.ErrorCode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(e.ErrorData[:]) {
		return errors.New("e.ErrorData failed UTF-8 validation")
	}
	return nil
}

func ErrorResponseInit(e *ErrorResponse) {
	return
}

func (*ErrorResponse) SbeBlockLength() (blockLength uint16) {
	return 1
}

func (*ErrorResponse) SbeTemplateId() (templateId uint16) {
	return 10
}

func (*ErrorResponse) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (*ErrorResponse) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*ErrorResponse) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*ErrorResponse) ErrorCodeId() uint16 {
	return 1
}

func (*ErrorResponse) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (e *ErrorResponse) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ErrorCodeSinceVersion()
}

func (*ErrorResponse) ErrorCodeDeprecated() uint16 {
	return 0
}

func (*ErrorResponse) ErrorCodeMetaAttribute(meta int) string {
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

func (*ErrorResponse) ErrorDataMetaAttribute(meta int) string {
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

func (*ErrorResponse) ErrorDataSinceVersion() uint16 {
	return 0
}

func (e *ErrorResponse) ErrorDataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ErrorDataSinceVersion()
}

func (*ErrorResponse) ErrorDataDeprecated() uint16 {
	return 0
}

func (ErrorResponse) ErrorDataCharacterEncoding() string {
	return "UTF-8"
}

func (ErrorResponse) ErrorDataHeaderLength() uint64 {
	return 4
}
