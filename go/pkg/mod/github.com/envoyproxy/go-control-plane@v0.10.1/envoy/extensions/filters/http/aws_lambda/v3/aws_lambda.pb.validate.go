// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/aws_lambda/v3/aws_lambda.proto

package envoy_extensions_filters_http_aws_lambda_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConfigMultiError, or nil if none found.
func (m *Config) ValidateAll() error {
	return m.validate(true)
}

func (m *Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetArn()) < 1 {
		err := ConfigValidationError{
			field:  "Arn",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PayloadPassthrough

	if _, ok := Config_InvocationMode_name[int32(m.GetInvocationMode())]; !ok {
		err := ConfigValidationError{
			field:  "InvocationMode",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ConfigMultiError(errors)
	}
	return nil
}

// ConfigMultiError is an error wrapping multiple validation errors returned by
// Config.ValidateAll() if the designated constraints aren't met.
type ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigMultiError) AllErrors() []error { return m }

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

// Validate checks the field values on PerRouteConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PerRouteConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PerRouteConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PerRouteConfigMultiError,
// or nil if none found.
func (m *PerRouteConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *PerRouteConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInvokeConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PerRouteConfigValidationError{
					field:  "InvokeConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PerRouteConfigValidationError{
					field:  "InvokeConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInvokeConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PerRouteConfigValidationError{
				field:  "InvokeConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PerRouteConfigMultiError(errors)
	}
	return nil
}

// PerRouteConfigMultiError is an error wrapping multiple validation errors
// returned by PerRouteConfig.ValidateAll() if the designated constraints
// aren't met.
type PerRouteConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PerRouteConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PerRouteConfigMultiError) AllErrors() []error { return m }

// PerRouteConfigValidationError is the validation error returned by
// PerRouteConfig.Validate if the designated constraints aren't met.
type PerRouteConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PerRouteConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PerRouteConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PerRouteConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PerRouteConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PerRouteConfigValidationError) ErrorName() string { return "PerRouteConfigValidationError" }

// Error satisfies the builtin error interface
func (e PerRouteConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerRouteConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PerRouteConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PerRouteConfigValidationError{}
