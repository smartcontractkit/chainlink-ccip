// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package filter

import (
	"fmt"
	"strings"
)

const (
	// FieldPlugin is a Field of type Plugin.
	FieldPlugin Field = "Plugin"
	// FieldComponent is a Field of type Component.
	FieldComponent Field = "Component"
	// FieldLogLevel is a Field of type LogLevel.
	FieldLogLevel Field = "LogLevel"
	// FieldMessage is a Field of type Message.
	FieldMessage Field = "Message"
	// FieldCaller is a Field of type Caller.
	FieldCaller Field = "Caller"
	// FieldLoggerName is a Field of type LoggerName.
	FieldLoggerName Field = "LoggerName"
	// FieldDONID is a Field of type DONID.
	FieldDONID Field = "DONID"
)

var ErrInvalidField = fmt.Errorf("not a valid Field, try [%s]", strings.Join(_FieldNames, ", "))

var _FieldNames = []string{
	string(FieldPlugin),
	string(FieldComponent),
	string(FieldLogLevel),
	string(FieldMessage),
	string(FieldCaller),
	string(FieldLoggerName),
	string(FieldDONID),
}

// FieldNames returns a list of possible string values of Field.
func FieldNames() []string {
	tmp := make([]string, len(_FieldNames))
	copy(tmp, _FieldNames)
	return tmp
}

// String implements the Stringer interface.
func (x Field) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Field) IsValid() bool {
	_, err := ParseField(string(x))
	return err == nil
}

var _FieldValue = map[string]Field{
	"Plugin":     FieldPlugin,
	"plugin":     FieldPlugin,
	"Component":  FieldComponent,
	"component":  FieldComponent,
	"LogLevel":   FieldLogLevel,
	"loglevel":   FieldLogLevel,
	"Message":    FieldMessage,
	"message":    FieldMessage,
	"Caller":     FieldCaller,
	"caller":     FieldCaller,
	"LoggerName": FieldLoggerName,
	"loggername": FieldLoggerName,
	"DONID":      FieldDONID,
	"donid":      FieldDONID,
}

// ParseField attempts to convert a string to a Field.
func ParseField(name string) (Field, error) {
	if x, ok := _FieldValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _FieldValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Field(""), fmt.Errorf("%s is %w", name, ErrInvalidField)
}

const (
	// FilterOPAND is a FilterOP of type AND.
	FilterOPAND FilterOP = "AND"
	// FilterOPOR is a FilterOP of type OR.
	FilterOPOR FilterOP = "OR"
)

var ErrInvalidFilterOP = fmt.Errorf("not a valid FilterOP, try [%s]", strings.Join(_FilterOPNames, ", "))

var _FilterOPNames = []string{
	string(FilterOPAND),
	string(FilterOPOR),
}

// FilterOPNames returns a list of possible string values of FilterOP.
func FilterOPNames() []string {
	tmp := make([]string, len(_FilterOPNames))
	copy(tmp, _FilterOPNames)
	return tmp
}

// String implements the Stringer interface.
func (x FilterOP) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x FilterOP) IsValid() bool {
	_, err := ParseFilterOP(string(x))
	return err == nil
}

var _FilterOPValue = map[string]FilterOP{
	"AND": FilterOPAND,
	"and": FilterOPAND,
	"OR":  FilterOPOR,
	"or":  FilterOPOR,
}

// ParseFilterOP attempts to convert a string to a FilterOP.
func ParseFilterOP(name string) (FilterOP, error) {
	if x, ok := _FilterOPValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _FilterOPValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return FilterOP(""), fmt.Errorf("%s is %w", name, ErrInvalidFilterOP)
}
