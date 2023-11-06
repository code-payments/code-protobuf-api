// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: contact/v1/contact_list_service.proto

package contact

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

// Validate checks the field values on AddContactsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddContactsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerAccountId() == nil {
		return AddContactsRequestValidationError{
			field:  "OwnerAccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOwnerAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddContactsRequestValidationError{
				field:  "OwnerAccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		return AddContactsRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddContactsRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetContainerId() == nil {
		return AddContactsRequestValidationError{
			field:  "ContainerId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetContainerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddContactsRequestValidationError{
				field:  "ContainerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := len(m.GetContacts()); l < 1 || l > 1024 {
		return AddContactsRequestValidationError{
			field:  "Contacts",
			reason: "value must contain between 1 and 1024 items, inclusive",
		}
	}

	for idx, item := range m.GetContacts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddContactsRequestValidationError{
					field:  fmt.Sprintf("Contacts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AddContactsRequestValidationError is the validation error returned by
// AddContactsRequest.Validate if the designated constraints aren't met.
type AddContactsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddContactsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddContactsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddContactsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddContactsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddContactsRequestValidationError) ErrorName() string {
	return "AddContactsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddContactsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddContactsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddContactsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddContactsRequestValidationError{}

// Validate checks the field values on AddContactsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddContactsResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	// no validation rules for ContactStatus

	return nil
}

// AddContactsResponseValidationError is the validation error returned by
// AddContactsResponse.Validate if the designated constraints aren't met.
type AddContactsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddContactsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddContactsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddContactsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddContactsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddContactsResponseValidationError) ErrorName() string {
	return "AddContactsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddContactsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddContactsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddContactsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddContactsResponseValidationError{}

// Validate checks the field values on RemoveContactsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveContactsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerAccountId() == nil {
		return RemoveContactsRequestValidationError{
			field:  "OwnerAccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOwnerAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveContactsRequestValidationError{
				field:  "OwnerAccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		return RemoveContactsRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveContactsRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetContainerId() == nil {
		return RemoveContactsRequestValidationError{
			field:  "ContainerId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetContainerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveContactsRequestValidationError{
				field:  "ContainerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := len(m.GetContacts()); l < 1 || l > 1024 {
		return RemoveContactsRequestValidationError{
			field:  "Contacts",
			reason: "value must contain between 1 and 1024 items, inclusive",
		}
	}

	for idx, item := range m.GetContacts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RemoveContactsRequestValidationError{
					field:  fmt.Sprintf("Contacts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RemoveContactsRequestValidationError is the validation error returned by
// RemoveContactsRequest.Validate if the designated constraints aren't met.
type RemoveContactsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveContactsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveContactsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveContactsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveContactsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveContactsRequestValidationError) ErrorName() string {
	return "RemoveContactsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveContactsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveContactsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveContactsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveContactsRequestValidationError{}

// Validate checks the field values on RemoveContactsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveContactsResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	return nil
}

// RemoveContactsResponseValidationError is the validation error returned by
// RemoveContactsResponse.Validate if the designated constraints aren't met.
type RemoveContactsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveContactsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveContactsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveContactsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveContactsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveContactsResponseValidationError) ErrorName() string {
	return "RemoveContactsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveContactsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveContactsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveContactsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveContactsResponseValidationError{}

// Validate checks the field values on GetContactsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetContactsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerAccountId() == nil {
		return GetContactsRequestValidationError{
			field:  "OwnerAccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOwnerAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetContactsRequestValidationError{
				field:  "OwnerAccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		return GetContactsRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetContactsRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetContainerId() == nil {
		return GetContactsRequestValidationError{
			field:  "ContainerId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetContainerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetContactsRequestValidationError{
				field:  "ContainerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPageToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetContactsRequestValidationError{
				field:  "PageToken",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IncludeOnlyInAppContacts

	return nil
}

// GetContactsRequestValidationError is the validation error returned by
// GetContactsRequest.Validate if the designated constraints aren't met.
type GetContactsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetContactsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetContactsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetContactsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetContactsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetContactsRequestValidationError) ErrorName() string {
	return "GetContactsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetContactsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetContactsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetContactsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetContactsRequestValidationError{}

// Validate checks the field values on GetContactsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetContactsResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if len(m.GetContacts()) > 1024 {
		return GetContactsResponseValidationError{
			field:  "Contacts",
			reason: "value must contain no more than 1024 item(s)",
		}
	}

	for idx, item := range m.GetContacts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetContactsResponseValidationError{
					field:  fmt.Sprintf("Contacts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetNextPageToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetContactsResponseValidationError{
				field:  "NextPageToken",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetContactsResponseValidationError is the validation error returned by
// GetContactsResponse.Validate if the designated constraints aren't met.
type GetContactsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetContactsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetContactsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetContactsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetContactsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetContactsResponseValidationError) ErrorName() string {
	return "GetContactsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetContactsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetContactsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetContactsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetContactsResponseValidationError{}

// Validate checks the field values on Contact with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Contact) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPhoneNumber() == nil {
		return ContactValidationError{
			field:  "PhoneNumber",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPhoneNumber()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ContactValidationError{
				field:  "PhoneNumber",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetStatus() == nil {
		return ContactValidationError{
			field:  "Status",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ContactValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ContactValidationError is the validation error returned by Contact.Validate
// if the designated constraints aren't met.
type ContactValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContactValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContactValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContactValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContactValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContactValidationError) ErrorName() string { return "ContactValidationError" }

// Error satisfies the builtin error interface
func (e ContactValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContact.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContactValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContactValidationError{}

// Validate checks the field values on ContactStatus with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ContactStatus) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IsRegistered

	// no validation rules for IsInvited

	// no validation rules for IsInviteRevoked

	return nil
}

// ContactStatusValidationError is the validation error returned by
// ContactStatus.Validate if the designated constraints aren't met.
type ContactStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContactStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContactStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContactStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContactStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContactStatusValidationError) ErrorName() string { return "ContactStatusValidationError" }

// Error satisfies the builtin error interface
func (e ContactStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContactStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContactStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContactStatusValidationError{}

// Validate checks the field values on PageToken with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PageToken) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetValue()); l < 1 || l > 128 {
		return PageTokenValidationError{
			field:  "Value",
			reason: "value length must be between 1 and 128 bytes, inclusive",
		}
	}

	return nil
}

// PageTokenValidationError is the validation error returned by
// PageToken.Validate if the designated constraints aren't met.
type PageTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageTokenValidationError) ErrorName() string { return "PageTokenValidationError" }

// Error satisfies the builtin error interface
func (e PageTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageTokenValidationError{}
