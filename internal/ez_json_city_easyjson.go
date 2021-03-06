// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package internal

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson458266e0DecodeGithubComTesteradjustSimpleCobraAppInternal(in *jlexer.Lexer, out *EzJsonCity) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "population":
			out.Population = int(in.Int())
		case "people":
			if in.IsNull() {
				in.Skip()
				out.People = nil
			} else {
				in.Delim('[')
				if out.People == nil {
					if !in.IsDelim(']') {
						out.People = make([]EzJsonPerson, 0, 2)
					} else {
						out.People = []EzJsonPerson{}
					}
				} else {
					out.People = (out.People)[:0]
				}
				for !in.IsDelim(']') {
					var v1 EzJsonPerson
					(v1).UnmarshalEasyJSON(in)
					out.People = append(out.People, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson458266e0EncodeGithubComTesteradjustSimpleCobraAppInternal(out *jwriter.Writer, in EzJsonCity) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"population\":"
		out.RawString(prefix)
		out.Int(int(in.Population))
	}
	{
		const prefix string = ",\"people\":"
		out.RawString(prefix)
		if in.People == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.People {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EzJsonCity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson458266e0EncodeGithubComTesteradjustSimpleCobraAppInternal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EzJsonCity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson458266e0EncodeGithubComTesteradjustSimpleCobraAppInternal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EzJsonCity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson458266e0DecodeGithubComTesteradjustSimpleCobraAppInternal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EzJsonCity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson458266e0DecodeGithubComTesteradjustSimpleCobraAppInternal(l, v)
}
