package key

import "encoding/binary"

func job() {

}

func Int64Byte(i uint64) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, i)
	return key
}
