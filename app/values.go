package app

//This package provides helpers to derive value from string, like
//strings input by a user (through dialog or command line options).
//It is mainly derived from go standard flag library.

import (
	"fmt"
	"strconv"
)

type value interface {
	Set(string) error
	String() string
}

type boolValue bool

func newBoolValue(p *bool) *boolValue {
	return (*boolValue)(p)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

func (b *boolValue) String() string {
	return strconv.FormatBool(bool(*b))
}

type int64Value int64

func newInt64Value(p *int64) *int64Value {
	return (*int64Value)(p)
}

func (i *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = int64Value(v)
	return err
}

func (i *int64Value) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

type stringValue string

func newStringValue(p *string) *stringValue {
	return (*stringValue)(p)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

func (s *stringValue) String() string {
	return string(*s)
}

//stringsValue accumulates strings
type stringsValue []string

func newStringsValue(p *[]string) *stringsValue {
	return (*stringsValue)(p)
}

func (s *stringsValue) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func (s *stringsValue) String() string {
	return fmt.Sprintf("%v", *s)
}

//enumValue is a value whose value should be in a given set of strings
type enumValue struct {
	value   *string
	options []string
}

func newEnumValue(p *string, options ...string) *enumValue {
	return &enumValue{
		value:   p,
		options: options,
	}
}

func (e *enumValue) Set(value string) error {
	for _, v := range e.options {
		if v == value {
			*e.value = value
			return nil
		}
	}

	return fmt.Errorf("enum value must be one of %v, got '%s'", e.options, value)
}

func (e *enumValue) String() string {
	return *e.value
}
