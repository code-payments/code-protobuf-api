// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: push/v1/push_service.proto

package push

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

// Validate checks the field values on AddTokenRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddTokenRequestMultiError, or nil if none found.
func (m *AddTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOwnerAccountId() == nil {
		err := AddTokenRequestValidationError{
			field:  "OwnerAccountId",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetOwnerAccountId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddTokenRequestValidationError{
					field:  "OwnerAccountId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddTokenRequestValidationError{
					field:  "OwnerAccountId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOwnerAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddTokenRequestValidationError{
				field:  "OwnerAccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		err := AddTokenRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSignature()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddTokenRequestValidationError{
					field:  "Signature",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddTokenRequestValidationError{
					field:  "Signature",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddTokenRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetContainerId() == nil {
		err := AddTokenRequestValidationError{
			field:  "ContainerId",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetContainerId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddTokenRequestValidationError{
					field:  "ContainerId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddTokenRequestValidationError{
					field:  "ContainerId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContainerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddTokenRequestValidationError{
				field:  "ContainerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetPushToken()); l < 1 || l > 4096 {
		err := AddTokenRequestValidationError{
			field:  "PushToken",
			reason: "value length must be between 1 and 4096 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _AddTokenRequest_TokenType_InLookup[m.GetTokenType()]; !ok {
		err := AddTokenRequestValidationError{
			field:  "TokenType",
			reason: "value must be in list [FCM_ANDROID FCM_APNS]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAppInstall()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddTokenRequestValidationError{
					field:  "AppInstall",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddTokenRequestValidationError{
					field:  "AppInstall",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAppInstall()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddTokenRequestValidationError{
				field:  "AppInstall",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddTokenRequestMultiError(errors)
	}

	return nil
}

// AddTokenRequestMultiError is an error wrapping multiple validation errors
// returned by AddTokenRequest.ValidateAll() if the designated constraints
// aren't met.
type AddTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddTokenRequestMultiError) AllErrors() []error { return m }

// AddTokenRequestValidationError is the validation error returned by
// AddTokenRequest.Validate if the designated constraints aren't met.
type AddTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddTokenRequestValidationError) ErrorName() string { return "AddTokenRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddTokenRequestValidationError{}

var _AddTokenRequest_TokenType_InLookup = map[TokenType]struct{}{
	1: {},
	2: {},
}

// Validate checks the field values on AddTokenResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddTokenResponseMultiError, or nil if none found.
func (m *AddTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return AddTokenResponseMultiError(errors)
	}

	return nil
}

// AddTokenResponseMultiError is an error wrapping multiple validation errors
// returned by AddTokenResponse.ValidateAll() if the designated constraints
// aren't met.
type AddTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddTokenResponseMultiError) AllErrors() []error { return m }

// AddTokenResponseValidationError is the validation error returned by
// AddTokenResponse.Validate if the designated constraints aren't met.
type AddTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddTokenResponseValidationError) ErrorName() string { return "AddTokenResponseValidationError" }

// Error satisfies the builtin error interface
func (e AddTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddTokenResponseValidationError{}

// Validate checks the field values on RemoveTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveTokenRequestMultiError, or nil if none found.
func (m *RemoveTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOwnerAccountId() == nil {
		err := RemoveTokenRequestValidationError{
			field:  "OwnerAccountId",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetOwnerAccountId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemoveTokenRequestValidationError{
					field:  "OwnerAccountId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveTokenRequestValidationError{
					field:  "OwnerAccountId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOwnerAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveTokenRequestValidationError{
				field:  "OwnerAccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		err := RemoveTokenRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSignature()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemoveTokenRequestValidationError{
					field:  "Signature",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveTokenRequestValidationError{
					field:  "Signature",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveTokenRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetContainerId() == nil {
		err := RemoveTokenRequestValidationError{
			field:  "ContainerId",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetContainerId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemoveTokenRequestValidationError{
					field:  "ContainerId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveTokenRequestValidationError{
					field:  "ContainerId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContainerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveTokenRequestValidationError{
				field:  "ContainerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetPushToken()); l < 1 || l > 4096 {
		err := RemoveTokenRequestValidationError{
			field:  "PushToken",
			reason: "value length must be between 1 and 4096 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _RemoveTokenRequest_TokenType_InLookup[m.GetTokenType()]; !ok {
		err := RemoveTokenRequestValidationError{
			field:  "TokenType",
			reason: "value must be in list [FCM_ANDROID FCM_APNS]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RemoveTokenRequestMultiError(errors)
	}

	return nil
}

// RemoveTokenRequestMultiError is an error wrapping multiple validation errors
// returned by RemoveTokenRequest.ValidateAll() if the designated constraints
// aren't met.
type RemoveTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveTokenRequestMultiError) AllErrors() []error { return m }

// RemoveTokenRequestValidationError is the validation error returned by
// RemoveTokenRequest.Validate if the designated constraints aren't met.
type RemoveTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveTokenRequestValidationError) ErrorName() string {
	return "RemoveTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveTokenRequestValidationError{}

var _RemoveTokenRequest_TokenType_InLookup = map[TokenType]struct{}{
	1: {},
	2: {},
}

// Validate checks the field values on RemoveTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveTokenResponseMultiError, or nil if none found.
func (m *RemoveTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return RemoveTokenResponseMultiError(errors)
	}

	return nil
}

// RemoveTokenResponseMultiError is an error wrapping multiple validation
// errors returned by RemoveTokenResponse.ValidateAll() if the designated
// constraints aren't met.
type RemoveTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveTokenResponseMultiError) AllErrors() []error { return m }

// RemoveTokenResponseValidationError is the validation error returned by
// RemoveTokenResponse.Validate if the designated constraints aren't met.
type RemoveTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveTokenResponseValidationError) ErrorName() string {
	return "RemoveTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveTokenResponseValidationError{}
