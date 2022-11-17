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

type ExecuteCommandResponse struct {
	PartitionId     uint16
	Key             uint64
	RecordType      RecordTypeEnum
	ValueType       ValueTypeEnum
	Intent          uint8
	RejectionType   RejectionTypeEnum
	Value           []uint8
	RejectionReason []uint8
}

func (e *ExecuteCommandResponse) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := e.RecordType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ValueType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, e.Intent); err != nil {
		return err
	}
	if err := e.RejectionType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(e.Value))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Value); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(e.RejectionReason))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.RejectionReason); err != nil {
		return err
	}
	return nil
}

func (e *ExecuteCommandResponse) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if e.RecordTypeInActingVersion(actingVersion) {
		if err := e.RecordType.Decode(_m, _r, actingVersion); err != nil {
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
	if e.RejectionTypeInActingVersion(actingVersion) {
		if err := e.RejectionType.Decode(_m, _r, actingVersion); err != nil {
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

	if e.RejectionReasonInActingVersion(actingVersion) {
		var RejectionReasonLength uint32
		if err := _m.ReadUint32(_r, &RejectionReasonLength); err != nil {
			return err
		}
		if cap(e.RejectionReason) < int(RejectionReasonLength) {
			e.RejectionReason = make([]uint8, RejectionReasonLength)
		}
		e.RejectionReason = e.RejectionReason[:RejectionReasonLength]
		if err := _m.ReadBytes(_r, e.RejectionReason); err != nil {
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

func (e *ExecuteCommandResponse) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if err := e.RecordType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ValueType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.IntentInActingVersion(actingVersion) {
		if e.Intent < e.IntentMinValue() || e.Intent > e.IntentMaxValue() {
			return fmt.Errorf("Range check failed on e.Intent (%v < %v > %v)", e.IntentMinValue(), e.Intent, e.IntentMaxValue())
		}
	}
	if err := e.RejectionType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(e.Value[:]) {
		return errors.New("e.Value failed UTF-8 validation")
	}
	if !utf8.Valid(e.RejectionReason[:]) {
		return errors.New("e.RejectionReason failed UTF-8 validation")
	}
	return nil
}

func ExecuteCommandResponseInit(e *ExecuteCommandResponse) {
	return
}

func (*ExecuteCommandResponse) SbeBlockLength() (blockLength uint16) {
	return 14
}

func (*ExecuteCommandResponse) SbeTemplateId() (templateId uint16) {
	return 21
}

func (*ExecuteCommandResponse) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (*ExecuteCommandResponse) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*ExecuteCommandResponse) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*ExecuteCommandResponse) PartitionIdId() uint16 {
	return 1
}

func (*ExecuteCommandResponse) PartitionIdSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandResponse) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartitionIdSinceVersion()
}

func (*ExecuteCommandResponse) PartitionIdDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandResponse) PartitionIdMetaAttribute(meta int) string {
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

func (*ExecuteCommandResponse) PartitionIdMinValue() uint16 {
	return 0
}

func (*ExecuteCommandResponse) PartitionIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecuteCommandResponse) PartitionIdNullValue() uint16 {
	return math.MaxUint16
}

func (*ExecuteCommandResponse) KeyId() uint16 {
	return 2
}

func (*ExecuteCommandResponse) KeySinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandResponse) KeyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.KeySinceVersion()
}

func (*ExecuteCommandResponse) KeyDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandResponse) KeyMetaAttribute(meta int) string {
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

func (*ExecuteCommandResponse) KeyMinValue() uint64 {
	return 0
}

func (*ExecuteCommandResponse) KeyMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecuteCommandResponse) KeyNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecuteCommandResponse) RecordTypeId() uint16 {
	return 3
}

func (*ExecuteCommandResponse) RecordTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandResponse) RecordTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RecordTypeSinceVersion()
}

func (*ExecuteCommandResponse) RecordTypeDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandResponse) RecordTypeMetaAttribute(meta int) string {
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

func (*ExecuteCommandResponse) ValueTypeId() uint16 {
	return 4
}

func (*ExecuteCommandResponse) ValueTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandResponse) ValueTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ValueTypeSinceVersion()
}

func (*ExecuteCommandResponse) ValueTypeDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandResponse) ValueTypeMetaAttribute(meta int) string {
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

func (*ExecuteCommandResponse) IntentId() uint16 {
	return 5
}

func (*ExecuteCommandResponse) IntentSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandResponse) IntentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.IntentSinceVersion()
}

func (*ExecuteCommandResponse) IntentDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandResponse) IntentMetaAttribute(meta int) string {
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

func (*ExecuteCommandResponse) IntentMinValue() uint8 {
	return 0
}

func (*ExecuteCommandResponse) IntentMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecuteCommandResponse) IntentNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecuteCommandResponse) RejectionTypeId() uint16 {
	return 6
}

func (*ExecuteCommandResponse) RejectionTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandResponse) RejectionTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RejectionTypeSinceVersion()
}

func (*ExecuteCommandResponse) RejectionTypeDeprecated() uint16 {
	return 0
}

func (*ExecuteCommandResponse) RejectionTypeMetaAttribute(meta int) string {
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

func (*ExecuteCommandResponse) ValueMetaAttribute(meta int) string {
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

func (*ExecuteCommandResponse) ValueSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandResponse) ValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ValueSinceVersion()
}

func (*ExecuteCommandResponse) ValueDeprecated() uint16 {
	return 0
}

func (ExecuteCommandResponse) ValueCharacterEncoding() string {
	return "UTF-8"
}

func (ExecuteCommandResponse) ValueHeaderLength() uint64 {
	return 4
}

func (*ExecuteCommandResponse) RejectionReasonMetaAttribute(meta int) string {
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

func (*ExecuteCommandResponse) RejectionReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecuteCommandResponse) RejectionReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RejectionReasonSinceVersion()
}

func (*ExecuteCommandResponse) RejectionReasonDeprecated() uint16 {
	return 0
}

func (ExecuteCommandResponse) RejectionReasonCharacterEncoding() string {
	return "UTF-8"
}

func (ExecuteCommandResponse) RejectionReasonHeaderLength() uint64 {
	return 4
}
