// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package elasticfork

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

func easyjson8092efb6DecodeGithubComOlivereElastic(in *jlexer.Lexer, out *bulkDeleteRequestCommandOp) {
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
		case "_index":
			out.Index = string(in.String())
		case "_type":
			out.Type = string(in.String())
		case "_id":
			out.Id = string(in.String())
		case "parent":
			out.Parent = string(in.String())
		case "routing":
			out.Routing = string(in.String())
		case "version":
			out.Version = int64(in.Int64())
		case "version_type":
			out.VersionType = string(in.String())
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
func easyjson8092efb6EncodeGithubComOlivereElastic(out *jwriter.Writer, in bulkDeleteRequestCommandOp) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Index != "" {
		const prefix string = ",\"_index\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Index))
	}
	if in.Type != "" {
		const prefix string = ",\"_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.Id != "" {
		const prefix string = ",\"_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	if in.Parent != "" {
		const prefix string = ",\"parent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Parent))
	}
	if in.Routing != "" {
		const prefix string = ",\"routing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Routing))
	}
	if in.Version != 0 {
		const prefix string = ",\"version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Version))
	}
	if in.VersionType != "" {
		const prefix string = ",\"version_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.VersionType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v bulkDeleteRequestCommandOp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8092efb6EncodeGithubComOlivereElastic(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v bulkDeleteRequestCommandOp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8092efb6EncodeGithubComOlivereElastic(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *bulkDeleteRequestCommandOp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8092efb6DecodeGithubComOlivereElastic(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *bulkDeleteRequestCommandOp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8092efb6DecodeGithubComOlivereElastic(l, v)
}
func easyjson8092efb6DecodeGithubComOlivereElastic1(in *jlexer.Lexer, out *bulkDeleteRequestCommand) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make(bulkDeleteRequestCommand)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 bulkDeleteRequestCommandOp
			(v1).UnmarshalEasyJSON(in)
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8092efb6EncodeGithubComOlivereElastic1(out *jwriter.Writer, in bulkDeleteRequestCommand) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in {
			if v2First {
				v2First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v2Name))
			out.RawByte(':')
			(v2Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v bulkDeleteRequestCommand) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8092efb6EncodeGithubComOlivereElastic1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v bulkDeleteRequestCommand) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8092efb6EncodeGithubComOlivereElastic1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *bulkDeleteRequestCommand) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8092efb6DecodeGithubComOlivereElastic1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *bulkDeleteRequestCommand) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8092efb6DecodeGithubComOlivereElastic1(l, v)
}
