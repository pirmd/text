package visual

import (
	"bytes"
)

// PadRight completes a slice of bytes with spaces until its "visual" size
// reaches the provided limit.
func PadRight(s []byte, sz int) []byte {
	in := TrimTrailingSpace(s)
	if freespace := sz - Width(in); freespace > 0 {
		return append(in, bytes.Repeat([]byte{' '}, freespace)...)
	}
	return in
}

// PadLeft prefixes a slice of bytes with spaces until its "visual" size
// reaches the provided limit.
func PadLeft(s []byte, sz int) []byte {
	in := TrimLeadingSpace(s)
	if freespace := sz - Width(in); freespace > 0 {
		return append(bytes.Repeat([]byte{' '}, freespace), in...)
	}
	return in
}

// PadCenter equally prefixes and complete a slice of bytes with spaces until
// its "visual" size reaches the provided limit.
func PadCenter(s []byte, sz int) []byte {
	in := TrimSpace(s)
	if freespace := sz - Width(in); freespace > 0 {
		in = append(bytes.Repeat([]byte{' '}, freespace/2), in...)
		return append(in, bytes.Repeat([]byte{' '}, freespace-(freespace/2))...)
	}
	return in
}
