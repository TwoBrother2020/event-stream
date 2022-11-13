package event_stream

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_zipDir(t *testing.T) {

	data, err := zipDir("/home/xioahei/workspace/event-stream/src")
	assert.Nil(t, err)
	tempDir := t.TempDir()

	err = unzip(tempDir, bytes.NewReader(data.Bytes()))
	assert.Nil(t, err)
}
