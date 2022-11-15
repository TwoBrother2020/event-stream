package uk_co_real_logic_sbe_benchmarks

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	var car Car
	var buf = new(bytes.Buffer)
	m := NewSbeGoMarshaller()

	header := SbeGoMessageHeader{car.SbeBlockLength(), car.SbeTemplateId(), car.SbeSchemaId(), car.SbeSchemaVersion()}
	header.Encode(m, buf)
}
