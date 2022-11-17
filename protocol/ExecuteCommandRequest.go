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

type ExecuteCommandRequest struct {
	PartitionId uint16
	Key         uint64
	ValueType   ValueTypeEnum
	Intent      uint8
	Value       []uint8
}

func (e *ExecuteCommandRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteUint8(_w, e.Intent); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(e.Value))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Value); err != nil {
		return err
	}
	return nil
}

func (e *ExecuteCommandRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !e.IntentInActingVersion(actingVersion) {
		e.Intent = e.IntentNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.Intent); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}

	if e.ValueInActingVersion(actingVersion) {
		var ValueLength uint32
		if err := _m.ReadUint32(_r, &ValueLength); err != nil {
			return err
		}
		if cap(e.Value) < int(ValueLength) {
			e.Value = make([]uint8, ValueLength)
		}
		e.Value = e.Value[:ValueLength]
		if err := _m.ReadBytes(_r, e.Value); err != nil {
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

func (e *ExecuteCommandRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.IntentInActingVersion(actingVersion) {
		if e.Intent < e.IntentMinValue() || e.Intent > e.IntentMaxValue() {
			return fmt.Errorf("Range check failed on e.Intent (%v < %v > %v)", e.IntentMinValue(), e.Intent, e.IntentMaxValue())
		}
	}
	if !utf8.Valid(e.Value[:]) {
		return errors.New("e.Value failed UTF-8 validation")
	}
	return nil
}

func ExecuteCommandRequestInit(e *ExecuteCommandRequest) {
	return
}

func (*ExecuteCommandRequest) SbeBlockLength() (blockLength uint16) {
	return 12
}

func (*ExecuteCommandRequest) SbeTemplateId() (templateId uint16) {
	return 20
}

func (*ExecuteCommandRequest) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (*ExecuteCommandRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*ExecuteCommandRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*ExecuteCommandRequest) PartitionIdId() uint16 {
	return 1
}

func (*ExecuteCommandRequest) PartitionIdSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandRequest) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartitionIdSinceVersion()
}

func (*ExecuteCommandRequest) PartitionIdDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandRequest) PartitionIdMetaAttribute(meta int) string {
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

func (*ExecuteCommandRequest) PartitionIdMinValue() uint16 {
	return 0
}

func (*ExecuteCommandRequest) PartitionIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecuteCommandRequest) PartitionIdNullValue() uint16 {
	return math.MaxUint16
}

func (*ExecuteCommandRequest) KeyId() uint16 {
	return 4
}

func (*ExecuteCommandRequest) KeySinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandRequest) KeyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.KeySinceVersion()
}

func (*ExecuteCommandRequest) KeyDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandRequest) KeyMetaAttribute(meta int) string {
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

func (*ExecuteCommandRequest) KeyMinValue() uint64 {
	return 0
}

func (*ExecuteCommandRequest) KeyMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecuteCommandRequest) KeyNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecuteCommandRequest) ValueTypeId() uint16 {
	return 5
}

func (*ExecuteCommandRequest) ValueTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandRequest) ValueTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ValueTypeSinceVersion()
}

func (*ExecuteCommandRequest) ValueTypeDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandRequest) ValueTypeMetaAttribute(meta int) string {
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

func (*ExecuteCommandRequest) IntentId() uint16 {
	return 6
}

func (*ExecuteCommandRequest) IntentSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandRequest) IntentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.IntentSinceVersion()
}

func (*ExecuteCommandRequest) IntentDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandRequest) IntentMetaAttribute(meta int) string {
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

func (*ExecuteCommandRequest) IntentMinValue() uint8 {
	return 0
}

func (*ExecuteCommandRequest) IntentMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecuteCommandRequest) IntentNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecuteCommandRequest) ValueMetaAttribute(meta int) string {
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

func (*ExecuteCommandRequest) ValueSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandRequest) ValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ValueSinceVersion()
}

func (*ExecuteCommandRequest) ValueDeprecated() uint16 {
	return 0
}

func (ExecuteCommandRequest) ValueCharacterEncoding() string {
	return "UTF-8"
}

func (ExecuteCommandRequest) ValueHeaderLength() uint64 {
	return 4
}
