// Copyright 2014, Joel Scoble (github.com/mohae): All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file
// 
package customjson

// MarshalIndentToString wraps MarshalIndent, converting the []byte to a string
// before returning the result, if it didn't error. Errors are thrown away and
// an empty string is returned.
//
// Not ideal to ignore errors but since this function is designed to create a
// readable printout, i.e. MarshalIndent'd, version of an interface, in JSON.
// This makes it useful for debugging, logging, etc.
//
// If error check is necessary, call MarshalIndent first.
func MarshalIndentToString(v interface{}, prefix, indent string) string {
	json, err := MarshalIndent(v, prefix, indent)
	if err != nil {
		return ""
	}
	
	return string(json)
}

// MarshalIndentToString wraps MarshalIndent, converting the []byte to a string
// before returning the result, if it didn't error. Errors are thrown away and
// an empty string is returned.
//
// Not ideal to ignore errors but since this function is designed to create a
// readable printout, i.e. MarshalIndent'd, version of an interface, in JSON.
// This makes it useful for debugging, logging, etc.
//
// If error check is necessary, call MarshalIndent first.
func MarshalToString(v interface{}) string {
	json, err := Marshal(v)
	if err != nil {
		return ""
	}
	
	return string(json)
}


// MarshalString is a struct to wrap the MarshalToString functions. This is
// mainly to simplify the use of MarshalIndentToString as all that needs to
// be passed is the interface to be marshalled to a string.
//
// The defaults for MarshalToString are the most commonly used, imo:
//       prefix: ""
//       indent: "        "
//
// Each of these settings can be overridden individually by calling its
// respective public method.
type MarshalString  struct {
	prefix, indent string 
}

func NewMarshalString() *MarshalString {
	return &MarshalString{indent: "    "}
}


func (m *MarshalString) Prefix(s string) {
	m.prefix = s
}

func (m *MarshalString) Indent(s string) {
	m.indent = s
}

func (m *MarshalString) GetIndented(v interface{}) string {
	return MarshalIndentToString(v, m.prefix, m.indent)
}

func (m *MarshalString) Get(v interface {}) string {
	return MarshalToString(v)
}
