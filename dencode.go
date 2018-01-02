package dencode

import (
	"encoding/json"
	"io"
)

type FormatType int32

const (
	TomlFormat FormatType = iota
	YamlFormat
	JsonFormat
)

// EncodeProvider interface
type EncodeProvider interface {
	Encode(interface{}) (err error)
}

// DecodeProvider interface
type DecodeProvider interface {
	Decode(interface{}) (err error)
}

// Encoder object
type Encoder struct {
	Provider EncodeProvider
}

// Decoder object
type Decoder struct {
	Provider DecodeProvider
}

// NewEncoder from io.Writer with specified format
func NewEncoder(f FormatType, w io.Writer) (enc Encoder) {
	var e EncodeProvider = nil

	switch f {
	case TomlFormat:
		e = NewTomlEncoder(w)
	case YamlFormat:
		e = NewYamlEncoder(w)
	case JsonFormat:
		e = json.NewEncoder(w)
	default:
	}

	return Encoder{Provider: e}
}

// Encode specified item
func (e Encoder) Encode(in interface{}) (err error) {
	return e.Provider.Encode(in)
}

// NewDecoder from io.Reader with specified format
func NewDecoder(f FormatType, r io.Reader) (dec Decoder) {
	var d DecodeProvider = nil

	switch f {
	case TomlFormat:
		d = NewTomlDecoder(r)
	case YamlFormat:
		d = NewYamlDecoder(r)
	case JsonFormat:
		d = json.NewDecoder(r)
	default:
	}

	return Decoder{Provider: d}
}

// Decode into specified object
func (d Decoder) Decode(out interface{}) (err error) {
	return d.Provider.Decode(out)
}
