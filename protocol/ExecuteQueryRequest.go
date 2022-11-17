// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecuteQueryRequest struct {
	PartitionId uint16
	Key         uint64
	ValueType   ValueTypeEnum
}

func (e *ExecuteQueryRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint16(_w, e.PartitionId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.Key); err != nil {
		return err
	}
	if err := e.ValueType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *ExecuteQueryRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.PartitionIdInActingVersion(actingVersion) {
		e.PartitionId = e.PartitionIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.PartitionId); err != nil {
			return err
		}
	}
	if !e.KeyInActingVersion(actingVersion) {
		e.Key = e.KeyNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.Key); err != nil {
			return err
		}
	}
	if e.ValueTypeInActingVersion(actingVersion) {
		if err := e.ValueType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExecuteQueryRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.PartitionIdInActingVersion(actingVersion) {
		if e.PartitionId < e.PartitionIdMinValue() || e.PartitionId > e.PartitionIdMaxValue() {
			return fmt.Errorf("Range check failed on e.PartitionId (%v < %v > %v)", e.PartitionIdMinValue(), e.PartitionId, e.PartitionIdMaxValue())
		}
	}
	if e.KeyInActingVersion(actingVersion) {
		if e.Key < e.KeyMinValue() || e.Key > e.KeyMaxValue() {
			return fmt.Errorf("Range check failed on e.Key (%v < %v > %v)", e.KeyMinValue(), e.Key, e.KeyMaxValue())
		}
	}
	if err := e.ValueType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func ExecuteQueryRequestInit(e *ExecuteQueryRequest) {
	return
}

func (*ExecuteQueryRequest) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*ExecuteQueryRequest) SbeTemplateId() (templateId uint16) {
	return 30
}

func (*ExecuteQueryRequest) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (*ExecuteQueryRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*ExecuteQueryRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*ExecuteQueryRequest) PartitionIdId() uint16 {
	return 1
}

func (*ExecuteQueryRequest) PartitionIdSinceVersion() uint16 {
	return 0
}

func (e *ExecuteQueryRequest) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartitionIdSinceVersion()
}

func (*ExecuteQueryRequest) PartitionIdDeprecated() uint16 {
	return 0
}

func (*ExecuteQueryRequest) PartitionIdMetaAttribute(meta int) string {
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

func (*ExecuteQueryRequest) PartitionIdMinValue() uint16 {
	return 0
}

func (*ExecuteQueryRequest) PartitionIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecuteQueryRequest) PartitionIdNullValue() uint16 {
	return math.MaxUint16
}

func (*ExecuteQueryRequest) KeyId() uint16 {
	return 2
}

func (*ExecuteQueryRequest) KeySinceVersion() uint16 {
	return 0
}

func (e *ExecuteQueryRequest) KeyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.KeySinceVersion()
}

func (*ExecuteQueryRequest) KeyDeprecated() uint16 {
	return 0
}

func (*ExecuteQueryRequest) KeyMetaAttribute(meta int) string {
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

func (*ExecuteQueryRequest) KeyMinValue() uint64 {
	return 0
}

func (*ExecuteQueryRequest) KeyMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecuteQueryRequest) KeyNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecuteQueryRequest) ValueTypeId() uint16 {
	return 3
}

func (*ExecuteQueryRequest) ValueTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecuteQueryRequest) ValueTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ValueTypeSinceVersion()
}

func (*ExecuteQueryRequest) ValueTypeDeprecated() uint16 {
	return 0
}

func (*ExecuteQueryRequest) ValueTypeMetaAttribute(meta int) string {
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
