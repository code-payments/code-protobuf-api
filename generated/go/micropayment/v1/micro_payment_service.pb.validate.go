// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: micropayment/v1/micro_payment_service.proto

package micropayment

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

// Validate checks the field values on GetStatusRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStatusRequestMultiError, or nil if none found.
func (m *GetStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetIntentId() == nil {
		err := GetStatusRequestValidationError{
			field:  "IntentId",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetIntentId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetStatusRequestValidationError{
					field:  "IntentId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetStatusRequestValidationError{
					field:  "IntentId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIntentId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStatusRequestValidationError{
				field:  "IntentId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetStatusRequestMultiError(errors)
	}

	return nil
}

// GetStatusRequestMultiError is an error wrapping multiple validation errors
// returned by GetStatusRequest.ValidateAll() if the designated constraints
// aren't met.
type GetStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStatusRequestMultiError) AllErrors() []error { return m }

// GetStatusRequestValidationError is the validation error returned by
// GetStatusRequest.Validate if the designated constraints aren't met.
type GetStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStatusRequestValidationError) ErrorName() string { return "GetStatusRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStatusRequestValidationError{}

// Validate checks the field values on GetStatusResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetStatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStatusResponseMultiError, or nil if none found.
func (m *GetStatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Exists

	// no validation rules for CodeScanned

	// no validation rules for IntentSubmitted

	if len(errors) > 0 {
		return GetStatusResponseMultiError(errors)
	}

	return nil
}

// GetStatusResponseMultiError is an error wrapping multiple validation errors
// returned by GetStatusResponse.ValidateAll() if the designated constraints
// aren't met.
type GetStatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStatusResponseMultiError) AllErrors() []error { return m }

// GetStatusResponseValidationError is the validation error returned by
// GetStatusResponse.Validate if the designated constraints aren't met.
type GetStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStatusResponseValidationError) ErrorName() string {
	return "GetStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStatusResponseValidationError{}

// Validate checks the field values on RegisterWebhookRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterWebhookRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterWebhookRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterWebhookRequestMultiError, or nil if none found.
func (m *RegisterWebhookRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterWebhookRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetIntentId() == nil {
		err := RegisterWebhookRequestValidationError{
			field:  "IntentId",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetIntentId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterWebhookRequestValidationError{
					field:  "IntentId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterWebhookRequestValidationError{
					field:  "IntentId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIntentId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterWebhookRequestValidationError{
				field:  "IntentId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetUrl()); l < 1 || l > 1024 {
		err := RegisterWebhookRequestValidationError{
			field:  "Url",
			reason: "value length must be between 1 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !strings.HasPrefix(m.GetUrl(), "http") {
		err := RegisterWebhookRequestValidationError{
			field:  "Url",
			reason: "value does not have prefix \"http\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		err = RegisterWebhookRequestValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := RegisterWebhookRequestValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterWebhookRequestMultiError(errors)
	}

	return nil
}

// RegisterWebhookRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterWebhookRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterWebhookRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterWebhookRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterWebhookRequestMultiError) AllErrors() []error { return m }

// RegisterWebhookRequestValidationError is the validation error returned by
// RegisterWebhookRequest.Validate if the designated constraints aren't met.
type RegisterWebhookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterWebhookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterWebhookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterWebhookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterWebhookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterWebhookRequestValidationError) ErrorName() string {
	return "RegisterWebhookRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterWebhookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterWebhookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterWebhookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterWebhookRequestValidationError{}

// Validate checks the field values on RegisterWebhookResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterWebhookResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterWebhookResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterWebhookResponseMultiError, or nil if none found.
func (m *RegisterWebhookResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterWebhookResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return RegisterWebhookResponseMultiError(errors)
	}

	return nil
}

// RegisterWebhookResponseMultiError is an error wrapping multiple validation
// errors returned by RegisterWebhookResponse.ValidateAll() if the designated
// constraints aren't met.
type RegisterWebhookResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterWebhookResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterWebhookResponseMultiError) AllErrors() []error { return m }

// RegisterWebhookResponseValidationError is the validation error returned by
// RegisterWebhookResponse.Validate if the designated constraints aren't met.
type RegisterWebhookResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterWebhookResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterWebhookResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterWebhookResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterWebhookResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterWebhookResponseValidationError) ErrorName() string {
	return "RegisterWebhookResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterWebhookResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterWebhookResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterWebhookResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterWebhookResponseValidationError{}
