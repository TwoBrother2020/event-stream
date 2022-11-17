// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"fmt"
	"io"
	"reflect"
)

type ValueTypeEnum uint8
type ValueTypeValues struct {
	JOB                              ValueTypeEnum
	DEPLOYMENT                       ValueTypeEnum
	PROCESS_INSTANCE                 ValueTypeEnum
	INCIDENT                         ValueTypeEnum
	MESSAGE                          ValueTypeEnum
	MESSAGE_SUBSCRIPTION             ValueTypeEnum
	PROCESS_MESSAGE_SUBSCRIPTION     ValueTypeEnum
	JOB_BATCH                        ValueTypeEnum
	TIMER                            ValueTypeEnum
	MESSAGE_START_EVENT_SUBSCRIPTION ValueTypeEnum
	VARIABLE                         ValueTypeEnum
	VARIABLE_DOCUMENT                ValueTypeEnum
	PROCESS_INSTANCE_CREATION        ValueTypeEnum
	ERROR                            ValueTypeEnum
	PROCESS_INSTANCE_RESULT          ValueTypeEnum
	PROCESS                          ValueTypeEnum
	DEPLOYMENT_DISTRIBUTION          ValueTypeEnum
	PROCESS_EVENT                    ValueTypeEnum
	DECISION                         ValueTypeEnum
	DECISION_REQUIREMENTS            ValueTypeEnum
	DECISION_EVALUATION              ValueTypeEnum
	PROCESS_INSTANCE_MODIFICATION    ValueTypeEnum
	CHECKPOINT                       ValueTypeEnum
	NullValue                        ValueTypeEnum
}

var ValueType = ValueTypeValues{0, 4, 5, 6, 10, 11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 254, 255}

func (v ValueTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(v)); err != nil {
		return err
	}
	return nil
}

func (v *ValueTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(v)); err != nil {
		return err
	}
	return nil
}

func (v ValueTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ValueType)
	for idx := 0; idx < value.NumField(); idx++ {
		if v == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ValueType, unknown enumeration value %d", v)
}

func (*ValueTypeEnum) EncodedLength() int64 {
	return 1
}

func (*ValueTypeEnum) JOBSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) JOBInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.JOBSinceVersion()
}

func (*ValueTypeEnum) JOBDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) DEPLOYMENTSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) DEPLOYMENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.DEPLOYMENTSinceVersion()
}

func (*ValueTypeEnum) DEPLOYMENTDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) PROCESS_INSTANCESinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) PROCESS_INSTANCEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.PROCESS_INSTANCESinceVersion()
}

func (*ValueTypeEnum) PROCESS_INSTANCEDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) INCIDENTSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) INCIDENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.INCIDENTSinceVersion()
}

func (*ValueTypeEnum) INCIDENTDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) MESSAGESinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) MESSAGEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.MESSAGESinceVersion()
}

func (*ValueTypeEnum) MESSAGEDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) MESSAGE_SUBSCRIPTIONSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) MESSAGE_SUBSCRIPTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.MESSAGE_SUBSCRIPTIONSinceVersion()
}

func (*ValueTypeEnum) MESSAGE_SUBSCRIPTIONDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) PROCESS_MESSAGE_SUBSCRIPTIONSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) PROCESS_MESSAGE_SUBSCRIPTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.PROCESS_MESSAGE_SUBSCRIPTIONSinceVersion()
}

func (*ValueTypeEnum) PROCESS_MESSAGE_SUBSCRIPTIONDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) JOB_BATCHSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) JOB_BATCHInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.JOB_BATCHSinceVersion()
}

func (*ValueTypeEnum) JOB_BATCHDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) TIMERSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) TIMERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.TIMERSinceVersion()
}

func (*ValueTypeEnum) TIMERDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) MESSAGE_START_EVENT_SUBSCRIPTIONSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) MESSAGE_START_EVENT_SUBSCRIPTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.MESSAGE_START_EVENT_SUBSCRIPTIONSinceVersion()
}

func (*ValueTypeEnum) MESSAGE_START_EVENT_SUBSCRIPTIONDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) VARIABLESinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) VARIABLEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.VARIABLESinceVersion()
}

func (*ValueTypeEnum) VARIABLEDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) VARIABLE_DOCUMENTSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) VARIABLE_DOCUMENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.VARIABLE_DOCUMENTSinceVersion()
}

func (*ValueTypeEnum) VARIABLE_DOCUMENTDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) PROCESS_INSTANCE_CREATIONSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) PROCESS_INSTANCE_CREATIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.PROCESS_INSTANCE_CREATIONSinceVersion()
}

func (*ValueTypeEnum) PROCESS_INSTANCE_CREATIONDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) ERRORSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) ERRORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.ERRORSinceVersion()
}

func (*ValueTypeEnum) ERRORDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) PROCESS_INSTANCE_RESULTSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) PROCESS_INSTANCE_RESULTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.PROCESS_INSTANCE_RESULTSinceVersion()
}

func (*ValueTypeEnum) PROCESS_INSTANCE_RESULTDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) PROCESSSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) PROCESSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.PROCESSSinceVersion()
}

func (*ValueTypeEnum) PROCESSDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) DEPLOYMENT_DISTRIBUTIONSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) DEPLOYMENT_DISTRIBUTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.DEPLOYMENT_DISTRIBUTIONSinceVersion()
}

func (*ValueTypeEnum) DEPLOYMENT_DISTRIBUTIONDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) PROCESS_EVENTSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) PROCESS_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.PROCESS_EVENTSinceVersion()
}

func (*ValueTypeEnum) PROCESS_EVENTDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) DECISIONSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) DECISIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.DECISIONSinceVersion()
}

func (*ValueTypeEnum) DECISIONDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) DECISION_REQUIREMENTSSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) DECISION_REQUIREMENTSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.DECISION_REQUIREMENTSSinceVersion()
}

func (*ValueTypeEnum) DECISION_REQUIREMENTSDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) DECISION_EVALUATIONSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) DECISION_EVALUATIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.DECISION_EVALUATIONSinceVersion()
}

func (*ValueTypeEnum) DECISION_EVALUATIONDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) PROCESS_INSTANCE_MODIFICATIONSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) PROCESS_INSTANCE_MODIFICATIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.PROCESS_INSTANCE_MODIFICATIONSinceVersion()
}

func (*ValueTypeEnum) PROCESS_INSTANCE_MODIFICATIONDeprecated() uint16 {
	return 0
}

func (*ValueTypeEnum) CHECKPOINTSinceVersion() uint16 {
	return 0
}

func (v *ValueTypeEnum) CHECKPOINTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.CHECKPOINTSinceVersion()
}

func (*ValueTypeEnum) CHECKPOINTDeprecated() uint16 {
	return 0
}
