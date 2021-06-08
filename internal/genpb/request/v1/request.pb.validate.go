// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: request/v1/request.proto

package requestv1

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
)

// Validate checks the field values on CheckResourceSetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckResourceSetRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestId

	if l := len(m.GetActions()); l < 1 || l > 10 {
		return CheckResourceSetRequestValidationError{
			field:  "Actions",
			reason: "value must contain between 1 and 10 items, inclusive",
		}
	}

	_CheckResourceSetRequest_Actions_Unique := make(map[string]struct{}, len(m.GetActions()))

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if _, exists := _CheckResourceSetRequest_Actions_Unique[item]; exists {
			return CheckResourceSetRequestValidationError{
				field:  fmt.Sprintf("Actions[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_CheckResourceSetRequest_Actions_Unique[item] = struct{}{}
		}

		if utf8.RuneCountInString(item) < 1 {
			return CheckResourceSetRequestValidationError{
				field:  fmt.Sprintf("Actions[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	if m.GetPrincipal() == nil {
		return CheckResourceSetRequestValidationError{
			field:  "Principal",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPrincipal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResourceSetRequestValidationError{
				field:  "Principal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetResource() == nil {
		return CheckResourceSetRequestValidationError{
			field:  "Resource",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResourceSetRequestValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IncludeMeta

	return nil
}

// CheckResourceSetRequestValidationError is the validation error returned by
// CheckResourceSetRequest.Validate if the designated constraints aren't met.
type CheckResourceSetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetRequestValidationError) ErrorName() string {
	return "CheckResourceSetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetRequestValidationError{}

// Validate checks the field values on ResourceSet with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ResourceSet) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetKind()) < 1 {
		return ResourceSetValidationError{
			field:  "Kind",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_ResourceSet_Kind_Pattern.MatchString(m.GetKind()) {
		return ResourceSetValidationError{
			field:  "Kind",
			reason: "value does not match regex pattern \"^[[:alpha:]][[:word:]\\\\@\\\\.\\\\-/]*(\\\\:[[:alpha:]][[:word:]\\\\@\\\\.\\\\-/]*)*$\"",
		}
	}

	if !_ResourceSet_PolicyVersion_Pattern.MatchString(m.GetPolicyVersion()) {
		return ResourceSetValidationError{
			field:  "PolicyVersion",
			reason: "value does not match regex pattern \"^[[:word:]]*$\"",
		}
	}

	if l := len(m.GetInstances()); l < 1 || l > 20 {
		return ResourceSetValidationError{
			field:  "Instances",
			reason: "value must contain between 1 and 20 pairs, inclusive",
		}
	}

	for key, val := range m.GetInstances() {
		_ = val

		if val == nil {
			return ResourceSetValidationError{
				field:  fmt.Sprintf("Instances[%v]", key),
				reason: "value cannot be sparse, all pairs must be non-nil",
			}
		}

		// no validation rules for Instances[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceSetValidationError{
					field:  fmt.Sprintf("Instances[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ResourceSetValidationError is the validation error returned by
// ResourceSet.Validate if the designated constraints aren't met.
type ResourceSetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceSetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceSetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceSetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceSetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceSetValidationError) ErrorName() string { return "ResourceSetValidationError" }

// Error satisfies the builtin error interface
func (e ResourceSetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourceSet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceSetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceSetValidationError{}

var _ResourceSet_Kind_Pattern = regexp.MustCompile("^[[:alpha:]][[:word:]\\@\\.\\-/]*(\\:[[:alpha:]][[:word:]\\@\\.\\-/]*)*$")

var _ResourceSet_PolicyVersion_Pattern = regexp.MustCompile("^[[:word:]]*$")

// Validate checks the field values on AttributesMap with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AttributesMap) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetAttr() {
		_ = val

		if val == nil {
			return AttributesMapValidationError{
				field:  fmt.Sprintf("Attr[%v]", key),
				reason: "value cannot be sparse, all pairs must be non-nil",
			}
		}

		// no validation rules for Attr[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AttributesMapValidationError{
					field:  fmt.Sprintf("Attr[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AttributesMapValidationError is the validation error returned by
// AttributesMap.Validate if the designated constraints aren't met.
type AttributesMapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributesMapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributesMapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributesMapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributesMapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributesMapValidationError) ErrorName() string { return "AttributesMapValidationError" }

// Error satisfies the builtin error interface
func (e AttributesMapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributesMap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributesMapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributesMapValidationError{}

// Validate checks the field values on CheckResourceBatchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckResourceBatchRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestId

	if m.GetPrincipal() == nil {
		return CheckResourceBatchRequestValidationError{
			field:  "Principal",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPrincipal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResourceBatchRequestValidationError{
				field:  "Principal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := len(m.GetResources()); l < 1 || l > 20 {
		return CheckResourceBatchRequestValidationError{
			field:  "Resources",
			reason: "value must contain between 1 and 20 items, inclusive",
		}
	}

	for idx, item := range m.GetResources() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResourceBatchRequestValidationError{
					field:  fmt.Sprintf("Resources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CheckResourceBatchRequestValidationError is the validation error returned by
// CheckResourceBatchRequest.Validate if the designated constraints aren't met.
type CheckResourceBatchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceBatchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceBatchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceBatchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceBatchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceBatchRequestValidationError) ErrorName() string {
	return "CheckResourceBatchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceBatchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceBatchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceBatchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceBatchRequestValidationError{}

// Validate checks the field values on PolicyFile with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PolicyFile) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetFileName()) < 1 {
		return PolicyFileValidationError{
			field:  "FileName",
			reason: "value length must be at least 1 runes",
		}
	}

	if l := len(m.GetContents()); l < 1 || l > 1048576 {
		return PolicyFileValidationError{
			field:  "Contents",
			reason: "value length must be between 1 and 1048576 bytes, inclusive",
		}
	}

	return nil
}

// PolicyFileValidationError is the validation error returned by
// PolicyFile.Validate if the designated constraints aren't met.
type PolicyFileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolicyFileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolicyFileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolicyFileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolicyFileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolicyFileValidationError) ErrorName() string { return "PolicyFileValidationError" }

// Error satisfies the builtin error interface
func (e PolicyFileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicyFile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolicyFileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolicyFileValidationError{}

// Validate checks the field values on PlaygroundValidateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PlaygroundValidateRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PlaygroundId

	if l := len(m.GetPolicyFiles()); l < 1 || l > 10 {
		return PlaygroundValidateRequestValidationError{
			field:  "PolicyFiles",
			reason: "value must contain between 1 and 10 items, inclusive",
		}
	}

	for idx, item := range m.GetPolicyFiles() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundValidateRequestValidationError{
					field:  fmt.Sprintf("PolicyFiles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PlaygroundValidateRequestValidationError is the validation error returned by
// PlaygroundValidateRequest.Validate if the designated constraints aren't met.
type PlaygroundValidateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundValidateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundValidateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundValidateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundValidateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundValidateRequestValidationError) ErrorName() string {
	return "PlaygroundValidateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundValidateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundValidateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundValidateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundValidateRequestValidationError{}

// Validate checks the field values on PlaygroundEvaluateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PlaygroundEvaluateRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PlaygroundId

	if l := len(m.GetPolicyFiles()); l < 1 || l > 10 {
		return PlaygroundEvaluateRequestValidationError{
			field:  "PolicyFiles",
			reason: "value must contain between 1 and 10 items, inclusive",
		}
	}

	for idx, item := range m.GetPolicyFiles() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundEvaluateRequestValidationError{
					field:  fmt.Sprintf("PolicyFiles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetPrincipal() == nil {
		return PlaygroundEvaluateRequestValidationError{
			field:  "Principal",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPrincipal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlaygroundEvaluateRequestValidationError{
				field:  "Principal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetResource() == nil {
		return PlaygroundEvaluateRequestValidationError{
			field:  "Resource",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlaygroundEvaluateRequestValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := len(m.GetActions()); l < 1 || l > 10 {
		return PlaygroundEvaluateRequestValidationError{
			field:  "Actions",
			reason: "value must contain between 1 and 10 items, inclusive",
		}
	}

	_PlaygroundEvaluateRequest_Actions_Unique := make(map[string]struct{}, len(m.GetActions()))

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if _, exists := _PlaygroundEvaluateRequest_Actions_Unique[item]; exists {
			return PlaygroundEvaluateRequestValidationError{
				field:  fmt.Sprintf("Actions[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_PlaygroundEvaluateRequest_Actions_Unique[item] = struct{}{}
		}

		if utf8.RuneCountInString(item) < 1 {
			return PlaygroundEvaluateRequestValidationError{
				field:  fmt.Sprintf("Actions[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	return nil
}

// PlaygroundEvaluateRequestValidationError is the validation error returned by
// PlaygroundEvaluateRequest.Validate if the designated constraints aren't met.
type PlaygroundEvaluateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundEvaluateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundEvaluateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundEvaluateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundEvaluateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundEvaluateRequestValidationError) ErrorName() string {
	return "PlaygroundEvaluateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundEvaluateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundEvaluateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundEvaluateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundEvaluateRequestValidationError{}

// Validate checks the field values on AddOrUpdatePolicyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddOrUpdatePolicyRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetPolicies()); l < 1 || l > 10 {
		return AddOrUpdatePolicyRequestValidationError{
			field:  "Policies",
			reason: "value must contain between 1 and 10 items, inclusive",
		}
	}

	for idx, item := range m.GetPolicies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddOrUpdatePolicyRequestValidationError{
					field:  fmt.Sprintf("Policies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AddOrUpdatePolicyRequestValidationError is the validation error returned by
// AddOrUpdatePolicyRequest.Validate if the designated constraints aren't met.
type AddOrUpdatePolicyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddOrUpdatePolicyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddOrUpdatePolicyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddOrUpdatePolicyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddOrUpdatePolicyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddOrUpdatePolicyRequestValidationError) ErrorName() string {
	return "AddOrUpdatePolicyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddOrUpdatePolicyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddOrUpdatePolicyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddOrUpdatePolicyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddOrUpdatePolicyRequestValidationError{}

// Validate checks the field values on CheckResourceBatchRequest_BatchEntry
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *CheckResourceBatchRequest_BatchEntry) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetActions()); l < 1 || l > 10 {
		return CheckResourceBatchRequest_BatchEntryValidationError{
			field:  "Actions",
			reason: "value must contain between 1 and 10 items, inclusive",
		}
	}

	_CheckResourceBatchRequest_BatchEntry_Actions_Unique := make(map[string]struct{}, len(m.GetActions()))

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if _, exists := _CheckResourceBatchRequest_BatchEntry_Actions_Unique[item]; exists {
			return CheckResourceBatchRequest_BatchEntryValidationError{
				field:  fmt.Sprintf("Actions[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_CheckResourceBatchRequest_BatchEntry_Actions_Unique[item] = struct{}{}
		}

		if utf8.RuneCountInString(item) < 1 {
			return CheckResourceBatchRequest_BatchEntryValidationError{
				field:  fmt.Sprintf("Actions[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	if m.GetResource() == nil {
		return CheckResourceBatchRequest_BatchEntryValidationError{
			field:  "Resource",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResourceBatchRequest_BatchEntryValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CheckResourceBatchRequest_BatchEntryValidationError is the validation error
// returned by CheckResourceBatchRequest_BatchEntry.Validate if the designated
// constraints aren't met.
type CheckResourceBatchRequest_BatchEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceBatchRequest_BatchEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceBatchRequest_BatchEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceBatchRequest_BatchEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceBatchRequest_BatchEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceBatchRequest_BatchEntryValidationError) ErrorName() string {
	return "CheckResourceBatchRequest_BatchEntryValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceBatchRequest_BatchEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceBatchRequest_BatchEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceBatchRequest_BatchEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceBatchRequest_BatchEntryValidationError{}
