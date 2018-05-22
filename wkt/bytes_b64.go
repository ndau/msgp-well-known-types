package wkt

import (
	"bytes"
	"encoding/base64"
)

// The Bytes type should know how to serialize itself as text

// Bytes converts Bytes to a primitive
func (b *Bytes) Bytes() []byte {
	return []byte(*b)
}

// UnmarshalText satisfies the encoding.TextUnmarshaler interface
func (b *Bytes) UnmarshalText(text []byte) error {
	bytes, err := base64.StdEncoding.DecodeString(string(text))
	if err == nil {
		*b = bytes
	}
	return err
}

// MarshalText satisfies the encoding.TextMarshaler interface
func (b Bytes) MarshalText() (text []byte, err error) {
	text = []byte(base64.StdEncoding.EncodeToString(b.Bytes()))
	return
}

// Equal returns true when `b` and `other` are equal
func (b *Bytes) Equal(other *Bytes) bool {
	if b == nil && other == nil {
		return true
	}
	if b == nil || other == nil {
		return false
	}
	return bytes.Equal(b.Bytes(), other.Bytes())
}
