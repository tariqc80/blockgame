package util

import "strconv"

// Uint64ToHexBytes returns a base 16 byte slice of the given integer
func Uint64ToHexBytes(i uint64) []byte {
	return []byte(strconv.FormatUint(i, 16))
}

// Int64ToHexBytes returns a base 16 byte slice of the given integer
func Int64ToHexBytes(i int64) []byte {
	return []byte(strconv.FormatInt(i, 16))
}
