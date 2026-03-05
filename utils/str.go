package utils

import "unsafe"

// BytesToString converts a byte slice to a string without allocation.
func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// StringToBytes converts a string to a byte slice without allocation.
//
// WARNING: The returned byte slice must not be modified. Modifying it
// will result in a runtime panic or undefined behavior since strings
// are immutable.
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
