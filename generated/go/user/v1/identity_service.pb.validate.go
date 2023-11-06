// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/v1/identity_service.proto

package user

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

// Validate checks the field values on LinkAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LinkAccountRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerAccountId() == nil {
		return LinkAccountRequestValidationError{
			field:  "OwnerAccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOwnerAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LinkAccountRequestValidationError{
				field:  "OwnerAccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		return LinkAccountRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LinkAccountRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Token.(type) {

	case *LinkAccountRequest_Phone:

		if m.GetPhone() == nil {
			return LinkAccountRequestValidationError{
				field:  "Phone",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetPhone()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LinkAccountRequestValidationError{
					field:  "Phone",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LinkAccountRequestValidationError is the validation error returned by
// LinkAccountRequest.Validate if the designated constraints aren't met.
type LinkAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LinkAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LinkAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LinkAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LinkAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LinkAccountRequestValidationError) ErrorName() string {
	return "LinkAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LinkAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLinkAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LinkAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LinkAccountRequestValidationError{}

// Validate checks the field values on LinkAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LinkAccountResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LinkAccountResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDataContainerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LinkAccountResponseValidationError{
				field:  "DataContainerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Metadata.(type) {

	case *LinkAccountResponse_Phone:

		if m.GetPhone() == nil {
			return LinkAccountResponseValidationError{
				field:  "Phone",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetPhone()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LinkAccountResponseValidationError{
					field:  "Phone",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LinkAccountResponseValidationError is the validation error returned by
// LinkAccountResponse.Validate if the designated constraints aren't met.
type LinkAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LinkAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LinkAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LinkAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LinkAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LinkAccountResponseValidationError) ErrorName() string {
	return "LinkAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LinkAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLinkAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LinkAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LinkAccountResponseValidationError{}

// Validate checks the field values on UnlinkAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UnlinkAccountRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerAccountId() == nil {
		return UnlinkAccountRequestValidationError{
			field:  "OwnerAccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOwnerAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UnlinkAccountRequestValidationError{
				field:  "OwnerAccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		return UnlinkAccountRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UnlinkAccountRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.IdentifyingFeature.(type) {

	case *UnlinkAccountRequest_PhoneNumber:

		if m.GetPhoneNumber() == nil {
			return UnlinkAccountRequestValidationError{
				field:  "PhoneNumber",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetPhoneNumber()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UnlinkAccountRequestValidationError{
					field:  "PhoneNumber",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// UnlinkAccountRequestValidationError is the validation error returned by
// UnlinkAccountRequest.Validate if the designated constraints aren't met.
type UnlinkAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnlinkAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnlinkAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnlinkAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnlinkAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnlinkAccountRequestValidationError) ErrorName() string {
	return "UnlinkAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UnlinkAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnlinkAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnlinkAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnlinkAccountRequestValidationError{}

// Validate checks the field values on UnlinkAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UnlinkAccountResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	return nil
}

// UnlinkAccountResponseValidationError is the validation error returned by
// UnlinkAccountResponse.Validate if the designated constraints aren't met.
type UnlinkAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnlinkAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnlinkAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnlinkAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnlinkAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnlinkAccountResponseValidationError) ErrorName() string {
	return "UnlinkAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UnlinkAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnlinkAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnlinkAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnlinkAccountResponseValidationError{}

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerAccountId() == nil {
		return GetUserRequestValidationError{
			field:  "OwnerAccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOwnerAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserRequestValidationError{
				field:  "OwnerAccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		return GetUserRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.IdentifyingFeature.(type) {

	case *GetUserRequest_PhoneNumber:

		if m.GetPhoneNumber() == nil {
			return GetUserRequestValidationError{
				field:  "PhoneNumber",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetPhoneNumber()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUserRequestValidationError{
					field:  "PhoneNumber",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDataContainerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserResponseValidationError{
				field:  "DataContainerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for EnableInternalFlags

	switch m.Metadata.(type) {

	case *GetUserResponse_Phone:

		if m.GetPhone() == nil {
			return GetUserResponseValidationError{
				field:  "Phone",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetPhone()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUserResponseValidationError{
					field:  "Phone",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetUserResponseValidationError is the validation error returned by
// GetUserResponse.Validate if the designated constraints aren't met.
type GetUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserResponseValidationError) ErrorName() string { return "GetUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserResponseValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() == nil {
		return UserValidationError{
			field:  "Id",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetView() == nil {
		return UserValidationError{
			field:  "View",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetView()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "View",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on View with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *View) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPhoneNumber() == nil {
		return ViewValidationError{
			field:  "PhoneNumber",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPhoneNumber()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ViewValidationError{
				field:  "PhoneNumber",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ViewValidationError is the validation error returned by View.Validate if the
// designated constraints aren't met.
type ViewValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ViewValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ViewValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ViewValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ViewValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ViewValidationError) ErrorName() string { return "ViewValidationError" }

// Error satisfies the builtin error interface
func (e ViewValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sView.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ViewValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ViewValidationError{}

// Validate checks the field values on PhoneMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PhoneMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IsLinked

	return nil
}

// PhoneMetadataValidationError is the validation error returned by
// PhoneMetadata.Validate if the designated constraints aren't met.
type PhoneMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PhoneMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PhoneMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PhoneMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PhoneMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PhoneMetadataValidationError) ErrorName() string { return "PhoneMetadataValidationError" }

// Error satisfies the builtin error interface
func (e PhoneMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPhoneMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PhoneMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PhoneMetadataValidationError{}
