// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: device/v1/device_service.proto

package device

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on RegisterLoggedInAccountsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegisterLoggedInAccountsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAppInstall() == nil {
		return RegisterLoggedInAccountsRequestValidationError{
			field:  "AppInstall",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAppInstall()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterLoggedInAccountsRequestValidationError{
				field:  "AppInstall",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetOwners()) > 1 {
		return RegisterLoggedInAccountsRequestValidationError{
			field:  "Owners",
			reason: "value must contain no more than 1 item(s)",
		}
	}

	for idx, item := range m.GetOwners() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterLoggedInAccountsRequestValidationError{
					field:  fmt.Sprintf("Owners[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(m.GetSignatures()) > 1 {
		return RegisterLoggedInAccountsRequestValidationError{
			field:  "Signatures",
			reason: "value must contain no more than 1 item(s)",
		}
	}

	for idx, item := range m.GetSignatures() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterLoggedInAccountsRequestValidationError{
					field:  fmt.Sprintf("Signatures[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RegisterLoggedInAccountsRequestValidationError is the validation error
// returned by RegisterLoggedInAccountsRequest.Validate if the designated
// constraints aren't met.
type RegisterLoggedInAccountsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterLoggedInAccountsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterLoggedInAccountsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterLoggedInAccountsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterLoggedInAccountsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterLoggedInAccountsRequestValidationError) ErrorName() string {
	return "RegisterLoggedInAccountsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterLoggedInAccountsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterLoggedInAccountsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterLoggedInAccountsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterLoggedInAccountsRequestValidationError{}

// Validate checks the field values on RegisterLoggedInAccountsResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *RegisterLoggedInAccountsResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if len(m.GetInvalidOwners()) > 1 {
		return RegisterLoggedInAccountsResponseValidationError{
			field:  "InvalidOwners",
			reason: "value must contain no more than 1 item(s)",
		}
	}

	for idx, item := range m.GetInvalidOwners() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterLoggedInAccountsResponseValidationError{
					field:  fmt.Sprintf("InvalidOwners[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RegisterLoggedInAccountsResponseValidationError is the validation error
// returned by RegisterLoggedInAccountsResponse.Validate if the designated
// constraints aren't met.
type RegisterLoggedInAccountsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterLoggedInAccountsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterLoggedInAccountsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterLoggedInAccountsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterLoggedInAccountsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterLoggedInAccountsResponseValidationError) ErrorName() string {
	return "RegisterLoggedInAccountsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterLoggedInAccountsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterLoggedInAccountsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterLoggedInAccountsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterLoggedInAccountsResponseValidationError{}

// Validate checks the field values on GetLoggedInAccountsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetLoggedInAccountsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAppInstall() == nil {
		return GetLoggedInAccountsRequestValidationError{
			field:  "AppInstall",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAppInstall()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetLoggedInAccountsRequestValidationError{
				field:  "AppInstall",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetLoggedInAccountsRequestValidationError is the validation error returned
// by GetLoggedInAccountsRequest.Validate if the designated constraints aren't met.
type GetLoggedInAccountsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLoggedInAccountsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLoggedInAccountsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLoggedInAccountsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLoggedInAccountsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLoggedInAccountsRequestValidationError) ErrorName() string {
	return "GetLoggedInAccountsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetLoggedInAccountsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLoggedInAccountsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLoggedInAccountsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLoggedInAccountsRequestValidationError{}

// Validate checks the field values on GetLoggedInAccountsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetLoggedInAccountsResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if len(m.GetOwners()) > 1 {
		return GetLoggedInAccountsResponseValidationError{
			field:  "Owners",
			reason: "value must contain no more than 1 item(s)",
		}
	}

	for idx, item := range m.GetOwners() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetLoggedInAccountsResponseValidationError{
					field:  fmt.Sprintf("Owners[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetLoggedInAccountsResponseValidationError is the validation error returned
// by GetLoggedInAccountsResponse.Validate if the designated constraints
// aren't met.
type GetLoggedInAccountsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLoggedInAccountsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLoggedInAccountsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLoggedInAccountsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLoggedInAccountsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLoggedInAccountsResponseValidationError) ErrorName() string {
	return "GetLoggedInAccountsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetLoggedInAccountsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLoggedInAccountsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLoggedInAccountsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLoggedInAccountsResponseValidationError{}
