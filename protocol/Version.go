// Generated SBE (Simple Binary Encoding) message codec

package protocol

import (
	"fmt"
	"io"
	"math"
)

type Version struct {
	MajorVersion int32
	MinorVersion int32
	PatchVersion int32
}

func (v *Version) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, v.MajorVersion); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, v.MinorVersion); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, v.PatchVersion); err != nil {
		return err
	}
	return nil
}

func (v *Version) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !v.MajorVersionInActingVersion(actingVersion) {
		v.MajorVersion = v.MajorVersionNullValue()
	} else {
		if err := _m.ReadInt32(_r, &v.MajorVersion); err != nil {
			return err
		}
	}
	if !v.MinorVersionInActingVersion(actingVersion) {
		v.MinorVersion = v.MinorVersionNullValue()
	} else {
		if err := _m.ReadInt32(_r, &v.MinorVersion); err != nil {
			return err
		}
	}
	if !v.PatchVersionInActingVersion(actingVersion) {
		v.PatchVersion = v.PatchVersionNullValue()
	} else {
		if err := _m.ReadInt32(_r, &v.PatchVersion); err != nil {
			return err
		}
	}
	return nil
}

func (v *Version) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if v.MajorVersionInActingVersion(actingVersion) {
		if v.MajorVersion < v.MajorVersionMinValue() || v.MajorVersion > v.MajorVersionMaxValue() {
			return fmt.Errorf("Range check failed on v.MajorVersion (%v < %v > %v)", v.MajorVersionMinValue(), v.MajorVersion, v.MajorVersionMaxValue())
		}
	}
	if v.MinorVersionInActingVersion(actingVersion) {
		if v.MinorVersion < v.MinorVersionMinValue() || v.MinorVersion > v.MinorVersionMaxValue() {
			return fmt.Errorf("Range check failed on v.MinorVersion (%v < %v > %v)", v.MinorVersionMinValue(), v.MinorVersion, v.MinorVersionMaxValue())
		}
	}
	if v.PatchVersionInActingVersion(actingVersion) {
		if v.PatchVersion < v.PatchVersionMinValue() || v.PatchVersion > v.PatchVersionMaxValue() {
			return fmt.Errorf("Range check failed on v.PatchVersion (%v < %v > %v)", v.PatchVersionMinValue(), v.PatchVersion, v.PatchVersionMaxValue())
		}
	}
	return nil
}

func VersionInit(v *Version) {
	return
}

func (*Version) EncodedLength() int64 {
	return 12
}

func (*Version) MajorVersionMinValue() int32 {
	return math.MinInt32 + 1
}

func (*Version) MajorVersionMaxValue() int32 {
	return math.MaxInt32
}

func (*Version) MajorVersionNullValue() int32 {
	return math.MinInt32
}

func (*Version) MajorVersionSinceVersion() uint16 {
	return 2
}

func (v *Version) MajorVersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.MajorVersionSinceVersion()
}

func (*Version) MajorVersionDeprecated() uint16 {
	return 0
}

func (*Version) MinorVersionMinValue() int32 {
	return math.MinInt32 + 1
}

func (*Version) MinorVersionMaxValue() int32 {
	return math.MaxInt32
}

func (*Version) MinorVersionNullValue() int32 {
	return math.MinInt32
}

func (*Version) MinorVersionSinceVersion() uint16 {
	return 2
}

func (v *Version) MinorVersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.MinorVersionSinceVersion()
}

func (*Version) MinorVersionDeprecated() uint16 {
	return 0
}

func (*Version) PatchVersionMinValue() int32 {
	return math.MinInt32 + 1
}

func (*Version) PatchVersionMaxValue() int32 {
	return math.MaxInt32
}

func (*Version) PatchVersionNullValue() int32 {
	return math.MinInt32
}

func (*Version) PatchVersionSinceVersion() uint16 {
	return 2
}

func (v *Version) PatchVersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.PatchVersionSinceVersion()
}

func (*Version) PatchVersionDeprecated() uint16 {
	return 0
}
