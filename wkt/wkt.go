package wkt

// msgp doesn't write useful tests for primitive types
//go:generate msgp -tests=0

// go primitives

type Bool bool
type String string
type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Uint uint
type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64
type Byte byte

// Bytes is []byte because that's a very common idiom
type Bytes []byte

// Uintptr is excluded because msgp doesn't know how to handle it
// type Uintptr uintptr

type Rune rune
type Float32 float32
type Float64 float64
type Complex64 complex64
type Complex128 complex128
