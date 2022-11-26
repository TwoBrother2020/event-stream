package key

import "testing"

func TestCreate(t *testing.T) {
	kv := NewDiskKV(1, 2)
	_, err := kv.Open(nil)
	if err != nil {
		t.Fatal(err)
	}

}
