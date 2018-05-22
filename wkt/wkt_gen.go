package wkt

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Bool) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 bool
		zb0001, err = dc.ReadBool()
		if err != nil {
			return
		}
		(*z) = Bool(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Bool) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBool(bool(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Bool) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBool(o, bool(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Bool) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 bool
		zb0001, bts, err = msgp.ReadBoolBytes(bts)
		if err != nil {
			return
		}
		(*z) = Bool(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Bool) Msgsize() (s int) {
	s = msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Byte) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 byte
		zb0001, err = dc.ReadByte()
		if err != nil {
			return
		}
		(*z) = Byte(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Byte) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteByte(byte(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Byte) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendByte(o, byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Byte) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 byte
		zb0001, bts, err = msgp.ReadByteBytes(bts)
		if err != nil {
			return
		}
		(*z) = Byte(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Byte) Msgsize() (s int) {
	s = msgp.ByteSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Complex128) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 complex128
		zb0001, err = dc.ReadComplex128()
		if err != nil {
			return
		}
		(*z) = Complex128(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Complex128) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteComplex128(complex128(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Complex128) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendComplex128(o, complex128(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Complex128) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 complex128
		zb0001, bts, err = msgp.ReadComplex128Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Complex128(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Complex128) Msgsize() (s int) {
	s = msgp.Complex128Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Complex64) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 complex64
		zb0001, err = dc.ReadComplex64()
		if err != nil {
			return
		}
		(*z) = Complex64(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Complex64) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteComplex64(complex64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Complex64) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendComplex64(o, complex64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Complex64) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 complex64
		zb0001, bts, err = msgp.ReadComplex64Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Complex64(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Complex64) Msgsize() (s int) {
	s = msgp.Complex64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Float32) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 float32
		zb0001, err = dc.ReadFloat32()
		if err != nil {
			return
		}
		(*z) = Float32(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Float32) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteFloat32(float32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Float32) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendFloat32(o, float32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Float32) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 float32
		zb0001, bts, err = msgp.ReadFloat32Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Float32(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Float32) Msgsize() (s int) {
	s = msgp.Float32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Float64) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 float64
		zb0001, err = dc.ReadFloat64()
		if err != nil {
			return
		}
		(*z) = Float64(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Float64) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteFloat64(float64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Float64) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendFloat64(o, float64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Float64) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 float64
		zb0001, bts, err = msgp.ReadFloat64Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Float64(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Float64) Msgsize() (s int) {
	s = msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Int) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int
		zb0001, err = dc.ReadInt()
		if err != nil {
			return
		}
		(*z) = Int(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Int) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt(int(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Int) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Int) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			return
		}
		(*z) = Int(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Int) Msgsize() (s int) {
	s = msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Int16) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int16
		zb0001, err = dc.ReadInt16()
		if err != nil {
			return
		}
		(*z) = Int16(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Int16) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt16(int16(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Int16) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt16(o, int16(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Int16) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int16
		zb0001, bts, err = msgp.ReadInt16Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Int16(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Int16) Msgsize() (s int) {
	s = msgp.Int16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Int32) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int32
		zb0001, err = dc.ReadInt32()
		if err != nil {
			return
		}
		(*z) = Int32(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Int32) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Int32) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Int32) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int32
		zb0001, bts, err = msgp.ReadInt32Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Int32(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Int32) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Int64) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int64
		zb0001, err = dc.ReadInt64()
		if err != nil {
			return
		}
		(*z) = Int64(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Int64) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt64(int64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Int64) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt64(o, int64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Int64) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int64
		zb0001, bts, err = msgp.ReadInt64Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Int64(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Int64) Msgsize() (s int) {
	s = msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Int8) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int8
		zb0001, err = dc.ReadInt8()
		if err != nil {
			return
		}
		(*z) = Int8(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Int8) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt8(int8(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Int8) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt8(o, int8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Int8) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int8
		zb0001, bts, err = msgp.ReadInt8Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Int8(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Int8) Msgsize() (s int) {
	s = msgp.Int8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Rune) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int32
		zb0001, err = dc.ReadInt32()
		if err != nil {
			return
		}
		(*z) = Rune(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Rune) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Rune) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Rune) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int32
		zb0001, bts, err = msgp.ReadInt32Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Rune(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Rune) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *String) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 string
		zb0001, err = dc.ReadString()
		if err != nil {
			return
		}
		(*z) = String(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z String) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z String) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *String) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		(*z) = String(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z String) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Uint) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint
		zb0001, err = dc.ReadUint()
		if err != nil {
			return
		}
		(*z) = Uint(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Uint) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint(uint(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Uint) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint(o, uint(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Uint) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint
		zb0001, bts, err = msgp.ReadUintBytes(bts)
		if err != nil {
			return
		}
		(*z) = Uint(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Uint) Msgsize() (s int) {
	s = msgp.UintSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Uint16) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint16
		zb0001, err = dc.ReadUint16()
		if err != nil {
			return
		}
		(*z) = Uint16(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Uint16) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint16(uint16(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Uint16) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint16(o, uint16(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Uint16) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint16
		zb0001, bts, err = msgp.ReadUint16Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Uint16(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Uint16) Msgsize() (s int) {
	s = msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Uint32) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint32
		zb0001, err = dc.ReadUint32()
		if err != nil {
			return
		}
		(*z) = Uint32(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Uint32) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint32(uint32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Uint32) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint32(o, uint32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Uint32) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint32
		zb0001, bts, err = msgp.ReadUint32Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Uint32(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Uint32) Msgsize() (s int) {
	s = msgp.Uint32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Uint64) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint64
		zb0001, err = dc.ReadUint64()
		if err != nil {
			return
		}
		(*z) = Uint64(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Uint64) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint64(uint64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Uint64) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint64(o, uint64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Uint64) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint64
		zb0001, bts, err = msgp.ReadUint64Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Uint64(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Uint64) Msgsize() (s int) {
	s = msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Uint8) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint8
		zb0001, err = dc.ReadUint8()
		if err != nil {
			return
		}
		(*z) = Uint8(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Uint8) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint8(uint8(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Uint8) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint8(o, uint8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Uint8) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint8
		zb0001, bts, err = msgp.ReadUint8Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Uint8(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Uint8) Msgsize() (s int) {
	s = msgp.Uint8Size
	return
}
