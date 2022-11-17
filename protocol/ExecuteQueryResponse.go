// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type ExecuteQueryResponse struct {
	BpmnProcessId []uint8
}

func (e *ExecuteQueryResponse) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, uint32(len(e.BpmnProcessId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.BpmnProcessId); err != nil {
		return err
	}
	return nil
}

func (e *ExecuteQueryResponse) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}

	if e.BpmnProcessIdInActingVersion(actingVersion) {
		var BpmnProcessIdLength uint32
		if err := _m.ReadUint32(_r, &BpmnProcessIdLength); err != nil {
			return err
		}
		if cap(e.BpmnProcessId) < int(BpmnProcessIdLength) {
			e.BpmnProcessId = make([]uint8, BpmnProcessIdLength)
		}
		e.BpmnProcessId = e.BpmnProcessId[:BpmnProcessIdLength]
		if err := _m.ReadBytes(_r, e.BpmnProcessId); err != nil {
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

func (e *ExecuteQueryResponse) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(e.BpmnProcessId[:]) {
		return errors.New("e.BpmnProcessId failed UTF-8 validation")
	}
	return nil
}

func ExecuteQueryResponseInit(e *ExecuteQueryResponse) {
	return
}

func (*ExecuteQueryResponse) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*ExecuteQueryResponse) SbeTemplateId() (templateId uint16) {
	return 31
}

func (*ExecuteQueryResponse) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (*ExecuteQueryResponse) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*ExecuteQueryResponse) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*ExecuteQueryResponse) BpmnProcessIdMetaAttribute(meta int) string {
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

func (*ExecuteQueryResponse) BpmnProcessIdSinceVersion() uint16 {
	return 0
}

func (e *ExecuteQueryResponse) BpmnProcessIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BpmnProcessIdSinceVersion()
}

func (*ExecuteQueryResponse) BpmnProcessIdDeprecated() uint16 {
	return 0
}

func (ExecuteQueryResponse) BpmnProcessIdCharacterEncoding() string {
	return "UTF-8"
}

func (ExecuteQueryResponse) BpmnProcessIdHeaderLength() uint64 {
	return 4
}
