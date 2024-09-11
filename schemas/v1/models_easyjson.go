// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	easyjson "github.com/zerodha/easyjson"
	jlexer "github.com/zerodha/easyjson/jlexer"
	jwriter "github.com/zerodha/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV1(in *jlexer.Lexer, out *URL) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "url":
			out.URL = string(in.String())
		case "wellKnown":
			out.WellKnown = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV1(out *jwriter.Writer, in URL) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix[1:])
		out.String(string(in.URL))
	}
	if in.WellKnown != "" {
		const prefix string = ",\"wellKnown\":"
		out.RawString(prefix)
		out.String(string(in.WellKnown))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v URL) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v URL) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *URL) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *URL) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV1(l, v)
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV11(in *jlexer.Lexer, out *Projects) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Projects, 0, 0)
			} else {
				*out = Projects{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Project
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV11(out *jwriter.Writer, in Projects) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Projects) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Projects) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Projects) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Projects) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV11(l, v)
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV12(in *jlexer.Lexer, out *Project) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "webpageUrl":
			(out.WebpageURL).UnmarshalEasyJSON(in)
		case "repositoryUrl":
			(out.RepositoryUrl).UnmarshalEasyJSON(in)
		case "licenses":
			if in.IsNull() {
				in.Skip()
				out.Licenses = nil
			} else {
				in.Delim('[')
				if out.Licenses == nil {
					if !in.IsDelim(']') {
						out.Licenses = make([]string, 0, 4)
					} else {
						out.Licenses = []string{}
					}
				} else {
					out.Licenses = (out.Licenses)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Licenses = append(out.Licenses, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "frameworks":
			if in.IsNull() {
				in.Skip()
				out.Frameworks = nil
			} else {
				in.Delim('[')
				if out.Frameworks == nil {
					if !in.IsDelim(']') {
						out.Frameworks = make([]string, 0, 4)
					} else {
						out.Frameworks = []string{}
					}
				} else {
					out.Frameworks = (out.Frameworks)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.Frameworks = append(out.Frameworks, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.Tags = append(out.Tags, v6)
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
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV12(out *jwriter.Writer, in Project) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"webpageUrl\":"
		out.RawString(prefix)
		(in.WebpageURL).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"repositoryUrl\":"
		out.RawString(prefix)
		(in.RepositoryUrl).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"licenses\":"
		out.RawString(prefix)
		if in.Licenses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Licenses {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"frameworks\":"
		out.RawString(prefix)
		if in.Frameworks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Frameworks {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Tags {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Project) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Project) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Project) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Project) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV12(l, v)
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV13(in *jlexer.Lexer, out *Plans) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Plans, 0, 0)
			} else {
				*out = Plans{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v13 Plan
			(v13).UnmarshalEasyJSON(in)
			*out = append(*out, v13)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV13(out *jwriter.Writer, in Plans) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v14, v15 := range in {
			if v14 > 0 {
				out.RawByte(',')
			}
			(v15).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Plans) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Plans) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Plans) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Plans) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV13(l, v)
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV14(in *jlexer.Lexer, out *Plan) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "amount":
			out.Amount = float64(in.Float64())
		case "currency":
			out.Currency = string(in.String())
		case "frequency":
			out.Frequency = string(in.String())
		case "channels":
			if in.IsNull() {
				in.Skip()
				out.Channels = nil
			} else {
				in.Delim('[')
				if out.Channels == nil {
					if !in.IsDelim(']') {
						out.Channels = make([]string, 0, 4)
					} else {
						out.Channels = []string{}
					}
				} else {
					out.Channels = (out.Channels)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.Channels = append(out.Channels, v16)
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
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV14(out *jwriter.Writer, in Plan) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"amount\":"
		out.RawString(prefix)
		out.Float64(float64(in.Amount))
	}
	{
		const prefix string = ",\"currency\":"
		out.RawString(prefix)
		out.String(string(in.Currency))
	}
	{
		const prefix string = ",\"frequency\":"
		out.RawString(prefix)
		out.String(string(in.Frequency))
	}
	{
		const prefix string = ",\"channels\":"
		out.RawString(prefix)
		if in.Channels == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Channels {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Plan) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Plan) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Plan) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Plan) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV14(l, v)
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV15(in *jlexer.Lexer, out *Manifest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "version":
			out.Version = string(in.String())
		case "entity":
			(out.Entity).UnmarshalEasyJSON(in)
		case "projects":
			(out.Projects).UnmarshalEasyJSON(in)
		case "funding":
			easyjsonD2b7633eDecode(in, &out.Funding)
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
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV15(out *jwriter.Writer, in Manifest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"entity\":"
		out.RawString(prefix)
		(in.Entity).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"projects\":"
		out.RawString(prefix)
		(in.Projects).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"funding\":"
		out.RawString(prefix)
		easyjsonD2b7633eEncode(out, in.Funding)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Manifest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV15(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Manifest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV15(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Manifest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV15(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Manifest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV15(l, v)
}
func easyjsonD2b7633eDecode(in *jlexer.Lexer, out *struct {
	Channels Channels `json:"channels"`
	Plans    Plans    `json:"plans"`
	History  History  `json:"history"`
}) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "channels":
			(out.Channels).UnmarshalEasyJSON(in)
		case "plans":
			(out.Plans).UnmarshalEasyJSON(in)
		case "history":
			(out.History).UnmarshalEasyJSON(in)
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
func easyjsonD2b7633eEncode(out *jwriter.Writer, in struct {
	Channels Channels `json:"channels"`
	Plans    Plans    `json:"plans"`
	History  History  `json:"history"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"channels\":"
		out.RawString(prefix[1:])
		(in.Channels).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"plans\":"
		out.RawString(prefix)
		(in.Plans).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"history\":"
		out.RawString(prefix)
		(in.History).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV16(in *jlexer.Lexer, out *HistoryItem) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "year":
			out.Year = int(in.Int())
		case "income":
			out.Income = float64(in.Float64())
		case "expenses":
			out.Expenses = float64(in.Float64())
		case "description":
			out.Description = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV16(out *jwriter.Writer, in HistoryItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"year\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Year))
	}
	{
		const prefix string = ",\"income\":"
		out.RawString(prefix)
		out.Float64(float64(in.Income))
	}
	{
		const prefix string = ",\"expenses\":"
		out.RawString(prefix)
		out.Float64(float64(in.Expenses))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HistoryItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV16(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HistoryItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV16(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HistoryItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV16(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HistoryItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV16(l, v)
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV17(in *jlexer.Lexer, out *History) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(History, 0, 1)
			} else {
				*out = History{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v19 HistoryItem
			(v19).UnmarshalEasyJSON(in)
			*out = append(*out, v19)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV17(out *jwriter.Writer, in History) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v20, v21 := range in {
			if v20 > 0 {
				out.RawByte(',')
			}
			(v21).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v History) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV17(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v History) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV17(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *History) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV17(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *History) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV17(l, v)
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV18(in *jlexer.Lexer, out *Entity) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "role":
			out.Role = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "telephone":
			out.Telephone = string(in.String())
		case "webpageUrl":
			(out.WebpageURL).UnmarshalEasyJSON(in)
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
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV18(out *jwriter.Writer, in Entity) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		out.String(string(in.Role))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"telephone\":"
		out.RawString(prefix)
		out.String(string(in.Telephone))
	}
	{
		const prefix string = ",\"webpageUrl\":"
		out.RawString(prefix)
		(in.WebpageURL).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Entity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV18(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Entity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV18(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Entity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV18(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Entity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV18(l, v)
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV19(in *jlexer.Lexer, out *Channels) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Channels, 0, 1)
			} else {
				*out = Channels{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v22 Channel
			(v22).UnmarshalEasyJSON(in)
			*out = append(*out, v22)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV19(out *jwriter.Writer, in Channels) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v23, v24 := range in {
			if v23 > 0 {
				out.RawByte(',')
			}
			(v24).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Channels) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV19(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Channels) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV19(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Channels) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV19(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Channels) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV19(l, v)
}
func easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV110(in *jlexer.Lexer, out *Channel) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "address":
			out.Address = string(in.String())
		case "description":
			out.Description = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV110(out *jwriter.Writer, in Channel) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Channel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{FloatFmt: ""}
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV110(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Channel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComFlossFundGoFundingJsonSchemasV110(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Channel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV110(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Channel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComFlossFundGoFundingJsonSchemasV110(l, v)
}
