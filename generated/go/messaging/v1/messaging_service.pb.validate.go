// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: messaging/v1/messaging_service.proto

package messaging

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

// Validate checks the field values on OpenMessageStreamRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OpenMessageStreamRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRendezvousKey() == nil {
		return OpenMessageStreamRequestValidationError{
			field:  "RendezvousKey",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRendezvousKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenMessageStreamRequestValidationError{
				field:  "RendezvousKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenMessageStreamRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OpenMessageStreamRequestValidationError is the validation error returned by
// OpenMessageStreamRequest.Validate if the designated constraints aren't met.
type OpenMessageStreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenMessageStreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenMessageStreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenMessageStreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenMessageStreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenMessageStreamRequestValidationError) ErrorName() string {
	return "OpenMessageStreamRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OpenMessageStreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpenMessageStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenMessageStreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenMessageStreamRequestValidationError{}

// Validate checks the field values on OpenMessageStreamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OpenMessageStreamResponse) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetMessages()); l < 1 || l > 1024 {
		return OpenMessageStreamResponseValidationError{
			field:  "Messages",
			reason: "value must contain between 1 and 1024 items, inclusive",
		}
	}

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpenMessageStreamResponseValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OpenMessageStreamResponseValidationError is the validation error returned by
// OpenMessageStreamResponse.Validate if the designated constraints aren't met.
type OpenMessageStreamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenMessageStreamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenMessageStreamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenMessageStreamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenMessageStreamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenMessageStreamResponseValidationError) ErrorName() string {
	return "OpenMessageStreamResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OpenMessageStreamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpenMessageStreamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenMessageStreamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenMessageStreamResponseValidationError{}

// Validate checks the field values on OpenMessageStreamWithKeepAliveRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *OpenMessageStreamWithKeepAliveRequest) Validate() error {
	if m == nil {
		return nil
	}

	switch m.RequestOrPong.(type) {

	case *OpenMessageStreamWithKeepAliveRequest_Request:

		if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpenMessageStreamWithKeepAliveRequestValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *OpenMessageStreamWithKeepAliveRequest_Pong:

		if v, ok := interface{}(m.GetPong()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpenMessageStreamWithKeepAliveRequestValidationError{
					field:  "Pong",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return OpenMessageStreamWithKeepAliveRequestValidationError{
			field:  "RequestOrPong",
			reason: "value is required",
		}

	}

	return nil
}

// OpenMessageStreamWithKeepAliveRequestValidationError is the validation error
// returned by OpenMessageStreamWithKeepAliveRequest.Validate if the
// designated constraints aren't met.
type OpenMessageStreamWithKeepAliveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenMessageStreamWithKeepAliveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenMessageStreamWithKeepAliveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenMessageStreamWithKeepAliveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenMessageStreamWithKeepAliveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenMessageStreamWithKeepAliveRequestValidationError) ErrorName() string {
	return "OpenMessageStreamWithKeepAliveRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OpenMessageStreamWithKeepAliveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpenMessageStreamWithKeepAliveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenMessageStreamWithKeepAliveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenMessageStreamWithKeepAliveRequestValidationError{}

// Validate checks the field values on OpenMessageStreamWithKeepAliveResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *OpenMessageStreamWithKeepAliveResponse) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ResponseOrPing.(type) {

	case *OpenMessageStreamWithKeepAliveResponse_Response:

		if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpenMessageStreamWithKeepAliveResponseValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *OpenMessageStreamWithKeepAliveResponse_Ping:

		if v, ok := interface{}(m.GetPing()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpenMessageStreamWithKeepAliveResponseValidationError{
					field:  "Ping",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return OpenMessageStreamWithKeepAliveResponseValidationError{
			field:  "ResponseOrPing",
			reason: "value is required",
		}

	}

	return nil
}

// OpenMessageStreamWithKeepAliveResponseValidationError is the validation
// error returned by OpenMessageStreamWithKeepAliveResponse.Validate if the
// designated constraints aren't met.
type OpenMessageStreamWithKeepAliveResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenMessageStreamWithKeepAliveResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenMessageStreamWithKeepAliveResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenMessageStreamWithKeepAliveResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenMessageStreamWithKeepAliveResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenMessageStreamWithKeepAliveResponseValidationError) ErrorName() string {
	return "OpenMessageStreamWithKeepAliveResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OpenMessageStreamWithKeepAliveResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpenMessageStreamWithKeepAliveResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenMessageStreamWithKeepAliveResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenMessageStreamWithKeepAliveResponseValidationError{}

// Validate checks the field values on PollMessagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PollMessagesRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRendezvousKey() == nil {
		return PollMessagesRequestValidationError{
			field:  "RendezvousKey",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRendezvousKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PollMessagesRequestValidationError{
				field:  "RendezvousKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		return PollMessagesRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PollMessagesRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PollMessagesRequestValidationError is the validation error returned by
// PollMessagesRequest.Validate if the designated constraints aren't met.
type PollMessagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PollMessagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PollMessagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PollMessagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PollMessagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PollMessagesRequestValidationError) ErrorName() string {
	return "PollMessagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PollMessagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPollMessagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PollMessagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PollMessagesRequestValidationError{}

// Validate checks the field values on PollMessagesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PollMessagesResponse) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetMessages()) > 1024 {
		return PollMessagesResponseValidationError{
			field:  "Messages",
			reason: "value must contain no more than 1024 item(s)",
		}
	}

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PollMessagesResponseValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PollMessagesResponseValidationError is the validation error returned by
// PollMessagesResponse.Validate if the designated constraints aren't met.
type PollMessagesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PollMessagesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PollMessagesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PollMessagesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PollMessagesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PollMessagesResponseValidationError) ErrorName() string {
	return "PollMessagesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PollMessagesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPollMessagesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PollMessagesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PollMessagesResponseValidationError{}

// Validate checks the field values on AckMessagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AckMessagesRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRendezvousKey() == nil {
		return AckMessagesRequestValidationError{
			field:  "RendezvousKey",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRendezvousKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AckMessagesRequestValidationError{
				field:  "RendezvousKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := len(m.GetMessageIds()); l < 1 || l > 1024 {
		return AckMessagesRequestValidationError{
			field:  "MessageIds",
			reason: "value must contain between 1 and 1024 items, inclusive",
		}
	}

	for idx, item := range m.GetMessageIds() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AckMessagesRequestValidationError{
					field:  fmt.Sprintf("MessageIds[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AckMessagesRequestValidationError is the validation error returned by
// AckMessagesRequest.Validate if the designated constraints aren't met.
type AckMessagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AckMessagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AckMessagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AckMessagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AckMessagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AckMessagesRequestValidationError) ErrorName() string {
	return "AckMessagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AckMessagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAckMessagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AckMessagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AckMessagesRequestValidationError{}

// Validate checks the field values on AckMesssagesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AckMesssagesResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	return nil
}

// AckMesssagesResponseValidationError is the validation error returned by
// AckMesssagesResponse.Validate if the designated constraints aren't met.
type AckMesssagesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AckMesssagesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AckMesssagesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AckMesssagesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AckMesssagesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AckMesssagesResponseValidationError) ErrorName() string {
	return "AckMesssagesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AckMesssagesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAckMesssagesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AckMesssagesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AckMesssagesResponseValidationError{}

// Validate checks the field values on SendMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SendMessageRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMessage() == nil {
		return SendMessageRequestValidationError{
			field:  "Message",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SendMessageRequestValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetRendezvousKey() == nil {
		return SendMessageRequestValidationError{
			field:  "RendezvousKey",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRendezvousKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SendMessageRequestValidationError{
				field:  "RendezvousKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignature() == nil {
		return SendMessageRequestValidationError{
			field:  "Signature",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SendMessageRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SendMessageRequestValidationError is the validation error returned by
// SendMessageRequest.Validate if the designated constraints aren't met.
type SendMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageRequestValidationError) ErrorName() string {
	return "SendMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SendMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageRequestValidationError{}

// Validate checks the field values on SendMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SendMessageResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if v, ok := interface{}(m.GetMessageId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SendMessageResponseValidationError{
				field:  "MessageId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SendMessageResponseValidationError is the validation error returned by
// SendMessageResponse.Validate if the designated constraints aren't met.
type SendMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageResponseValidationError) ErrorName() string {
	return "SendMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SendMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageResponseValidationError{}

// Validate checks the field values on RendezvousKey with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RendezvousKey) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 32 {
		return RendezvousKeyValidationError{
			field:  "Value",
			reason: "value length must be 32 bytes",
		}
	}

	return nil
}

// RendezvousKeyValidationError is the validation error returned by
// RendezvousKey.Validate if the designated constraints aren't met.
type RendezvousKeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RendezvousKeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RendezvousKeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RendezvousKeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RendezvousKeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RendezvousKeyValidationError) ErrorName() string { return "RendezvousKeyValidationError" }

// Error satisfies the builtin error interface
func (e RendezvousKeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRendezvousKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RendezvousKeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RendezvousKeyValidationError{}

// Validate checks the field values on MessageId with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *MessageId) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 16 {
		return MessageIdValidationError{
			field:  "Value",
			reason: "value length must be 16 bytes",
		}
	}

	return nil
}

// MessageIdValidationError is the validation error returned by
// MessageId.Validate if the designated constraints aren't met.
type MessageIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageIdValidationError) ErrorName() string { return "MessageIdValidationError" }

// Error satisfies the builtin error interface
func (e MessageIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageIdValidationError{}

// Validate checks the field values on RequestToGrabBill with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RequestToGrabBill) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRequestorAccount() == nil {
		return RequestToGrabBillValidationError{
			field:  "RequestorAccount",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRequestorAccount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestToGrabBillValidationError{
				field:  "RequestorAccount",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RequestToGrabBillValidationError is the validation error returned by
// RequestToGrabBill.Validate if the designated constraints aren't met.
type RequestToGrabBillValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestToGrabBillValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestToGrabBillValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestToGrabBillValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestToGrabBillValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestToGrabBillValidationError) ErrorName() string {
	return "RequestToGrabBillValidationError"
}

// Error satisfies the builtin error interface
func (e RequestToGrabBillValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestToGrabBill.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestToGrabBillValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestToGrabBillValidationError{}

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Message) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSendMessageRequestSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "SendMessageRequestSignature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Kind.(type) {

	case *Message_RequestToGrabBill:

		if v, ok := interface{}(m.GetRequestToGrabBill()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageValidationError{
					field:  "RequestToGrabBill",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return MessageValidationError{
			field:  "Kind",
			reason: "value is required",
		}

	}

	return nil
}

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}
