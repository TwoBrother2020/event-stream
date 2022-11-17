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

type RecordMetadata struct {
	RecordType      RecordTypeEnum
	RequestStreamId int32
	RequestId       uint64
	ProtocolVersion uint16
	ValueType       ValueTypeEnum
	Intent          uint8
	RejectionType   RejectionTypeEnum
	BrokerVersion   Version
	RejectionReason []uint8
}

func (r *RecordMetadata) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := r.RangeCheck(r.SbeSchemaVersion(), r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := r.RecordType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, r.RequestStreamId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.RequestId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, r.ProtocolVersion); err != nil {
		return err
	}
	if err := r.ValueType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, r.Intent); err != nil {
		return err
	}
	if err := r.RejectionType.Encode(_m, _w); err != nil {
		return err
	}
	if err := r.BrokerVersion.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(r.RejectionReason))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, r.RejectionReason); err != nil {
		return err
	}
	return nil
}

func (r *RecordMetadata) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if r.RecordTypeInActingVersion(actingVersion) {
		if err := r.RecordType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !r.RequestStreamIdInActingVersion(actingVersion) {
		r.RequestStreamId = r.RequestStreamIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &r.RequestStreamId); err != nil {
			return err
		}
	}
	if !r.RequestIdInActingVersion(actingVersion) {
		r.RequestId = r.RequestIdNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.RequestId); err != nil {
			return err
		}
	}
	if !r.ProtocolVersionInActingVersion(actingVersion) {
		r.ProtocolVersion = r.ProtocolVersionNullValue()
	} else {
		if err := _m.ReadUint16(_r, &r.ProtocolVersion); err != nil {
			return err
		}
	}
	if r.ValueTypeInActingVersion(actingVersion) {
		if err := r.ValueType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !r.IntentInActingVersion(actingVersion) {
		r.Intent = r.IntentNullValue()
	} else {
		if err := _m.ReadUint8(_r, &r.Intent); err != nil {
			return err
		}
	}
	if r.RejectionTypeInActingVersion(actingVersion) {
		if err := r.RejectionType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if r.BrokerVersionInActingVersion(actingVersion) {
		if err := r.BrokerVersion.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > r.SbeSchemaVersion() && blockLength > r.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-r.SbeBlockLength()))
	}

	if r.RejectionReasonInActingVersion(actingVersion) {
		var RejectionReasonLength uint32
		if err := _m.ReadUint32(_r, &RejectionReasonLength); err != nil {
			return err
		}
		if cap(r.RejectionReason) < int(RejectionReasonLength) {
			r.RejectionReason = make([]uint8, RejectionReasonLength)
		}
		r.RejectionReason = r.RejectionReason[:RejectionReasonLength]
		if err := _m.ReadBytes(_r, r.RejectionReason); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := r.RangeCheck(actingVersion, r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (r *RecordMetadata) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := r.RecordType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if r.RequestStreamIdInActingVersion(actingVersion) {
		if r.RequestStreamId < r.RequestStreamIdMinValue() || r.RequestStreamId > r.RequestStreamIdMaxValue() {
			return fmt.Errorf("Range check failed on r.RequestStreamId (%v < %v > %v)", r.RequestStreamIdMinValue(), r.RequestStreamId, r.RequestStreamIdMaxValue())
		}
	}
	if r.RequestIdInActingVersion(actingVersion) {
		if r.RequestId < r.RequestIdMinValue() || r.RequestId > r.RequestIdMaxValue() {
			return fmt.Errorf("Range check failed on r.RequestId (%v < %v > %v)", r.RequestIdMinValue(), r.RequestId, r.RequestIdMaxValue())
		}
	}
	if r.ProtocolVersionInActingVersion(actingVersion) {
		if r.ProtocolVersion < r.ProtocolVersionMinValue() || r.ProtocolVersion > r.ProtocolVersionMaxValue() {
			return fmt.Errorf("Range check failed on r.ProtocolVersion (%v < %v > %v)", r.ProtocolVersionMinValue(), r.ProtocolVersion, r.ProtocolVersionMaxValue())
		}
	}
	if err := r.ValueType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if r.IntentInActingVersion(actingVersion) {
		if r.Intent < r.IntentMinValue() || r.Intent > r.IntentMaxValue() {
			return fmt.Errorf("Range check failed on r.Intent (%v < %v > %v)", r.IntentMinValue(), r.Intent, r.IntentMaxValue())
		}
	}
	if err := r.RejectionType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(r.RejectionReason[:]) {
		return errors.New("r.RejectionReason failed UTF-8 validation")
	}
	return nil
}

func RecordMetadataInit(r *RecordMetadata) {
	return
}

func (*RecordMetadata) SbeBlockLength() (blockLength uint16) {
	return 30
}

func (*RecordMetadata) SbeTemplateId() (templateId uint16) {
	return 200
}

func (*RecordMetadata) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (*RecordMetadata) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*RecordMetadata) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*RecordMetadata) RecordTypeId() uint16 {
	return 1
}

func (*RecordMetadata) RecordTypeSinceVersion() uint16 {
	return 0
}

func (r *RecordMetadata) RecordTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RecordTypeSinceVersion()
}

func (*RecordMetadata) RecordTypeDeprecated() uint16 {
	return 0
}

func (*RecordMetadata) RecordTypeMetaAttribute(meta int) string {
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

func (*RecordMetadata) RequestStreamIdId() uint16 {
	return 2
}

func (*RecordMetadata) RequestStreamIdSinceVersion() uint16 {
	return 0
}

func (r *RecordMetadata) RequestStreamIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RequestStreamIdSinceVersion()
}

func (*RecordMetadata) RequestStreamIdDeprecated() uint16 {
	return 0
}

func (*RecordMetadata) RequestStreamIdMetaAttribute(meta int) string {
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

func (*RecordMetadata) RequestStreamIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*RecordMetadata) RequestStreamIdMaxValue() int32 {
	return math.MaxInt32
}

func (*RecordMetadata) RequestStreamIdNullValue() int32 {
	return math.MinInt32
}

func (*RecordMetadata) RequestIdId() uint16 {
	return 3
}

func (*RecordMetadata) RequestIdSinceVersion() uint16 {
	return 0
}

func (r *RecordMetadata) RequestIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RequestIdSinceVersion()
}

func (*RecordMetadata) RequestIdDeprecated() uint16 {
	return 0
}

func (*RecordMetadata) RequestIdMetaAttribute(meta int) string {
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

func (*RecordMetadata) RequestIdMinValue() uint64 {
	return 0
}

func (*RecordMetadata) RequestIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RecordMetadata) RequestIdNullValue() uint64 {
	return math.MaxUint64
}

func (*RecordMetadata) ProtocolVersionId() uint16 {
	return 4
}

func (*RecordMetadata) ProtocolVersionSinceVersion() uint16 {
	return 0
}

func (r *RecordMetadata) ProtocolVersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ProtocolVersionSinceVersion()
}

func (*RecordMetadata) ProtocolVersionDeprecated() uint16 {
	return 0
}

func (*RecordMetadata) ProtocolVersionMetaAttribute(meta int) string {
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

func (*RecordMetadata) ProtocolVersionMinValue() uint16 {
	return 0
}

func (*RecordMetadata) ProtocolVersionMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*RecordMetadata) ProtocolVersionNullValue() uint16 {
	return math.MaxUint16
}

func (*RecordMetadata) ValueTypeId() uint16 {
	return 5
}

func (*RecordMetadata) ValueTypeSinceVersion() uint16 {
	return 0
}

func (r *RecordMetadata) ValueTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ValueTypeSinceVersion()
}

func (*RecordMetadata) ValueTypeDeprecated() uint16 {
	return 0
}

func (*RecordMetadata) ValueTypeMetaAttribute(meta int) string {
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

func (*RecordMetadata) IntentId() uint16 {
	return 6
}

func (*RecordMetadata) IntentSinceVersion() uint16 {
	return 0
}

func (r *RecordMetadata) IntentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.IntentSinceVersion()
}

func (*RecordMetadata) IntentDeprecated() uint16 {
	return 0
}

func (*RecordMetadata) IntentMetaAttribute(meta int) string {
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

func (*RecordMetadata) IntentMinValue() uint8 {
	return 0
}

func (*RecordMetadata) IntentMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*RecordMetadata) IntentNullValue() uint8 {
	return math.MaxUint8
}

func (*RecordMetadata) RejectionTypeId() uint16 {
	return 7
}

func (*RecordMetadata) RejectionTypeSinceVersion() uint16 {
	return 0
}

func (r *RecordMetadata) RejectionTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RejectionTypeSinceVersion()
}

func (*RecordMetadata) RejectionTypeDeprecated() uint16 {
	return 0
}

func (*RecordMetadata) RejectionTypeMetaAttribute(meta int) string {
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

func (*RecordMetadata) BrokerVersionId() uint16 {
	return 9
}

func (*RecordMetadata) BrokerVersionSinceVersion() uint16 {
	return 2
}

func (r *RecordMetadata) BrokerVersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.BrokerVersionSinceVersion()
}

func (*RecordMetadata) BrokerVersionDeprecated() uint16 {
	return 0
}

func (*RecordMetadata) BrokerVersionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*RecordMetadata) RejectionReasonMetaAttribute(meta int) string {
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

func (*RecordMetadata) RejectionReasonSinceVersion() uint16 {
	return 0
}

func (r *RecordMetadata) RejectionReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RejectionReasonSinceVersion()
}

func (*RecordMetadata) RejectionReasonDeprecated() uint16 {
	return 0
}

func (RecordMetadata) RejectionReasonCharacterEncoding() string {
	return "UTF-8"
}

func (RecordMetadata) RejectionReasonHeaderLength() uint64 {
	return 4
}
