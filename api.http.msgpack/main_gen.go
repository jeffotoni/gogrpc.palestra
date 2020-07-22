package main

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Login) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "lastname":
			z.Lastname, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Lastname")
				return
			}
		case "age":
			z.Age, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Age")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Login) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "name"
	err = en.Append(0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "lastname"
	err = en.Append(0xa8, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Lastname)
	if err != nil {
		err = msgp.WrapError(err, "Lastname")
		return
	}
	// write "age"
	err = en.Append(0xa3, 0x61, 0x67, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Age)
	if err != nil {
		err = msgp.WrapError(err, "Age")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Login) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "name"
	o = append(o, 0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "lastname"
	o = append(o, 0xa8, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Lastname)
	// string "age"
	o = append(o, 0xa3, 0x61, 0x67, 0x65)
	o = msgp.AppendInt(o, z.Age)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Login) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "lastname":
			z.Lastname, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Lastname")
				return
			}
		case "age":
			z.Age, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Age")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Login) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.StringPrefixSize + len(z.Lastname) + 4 + msgp.IntSize
	return
}