// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// SchemasRateLimitErrorType - rate_limit_error
type SchemasRateLimitErrorType string

const (
	SchemasRateLimitErrorTypeRateLimitError SchemasRateLimitErrorType = "rate_limit_error"
)

func (e SchemasRateLimitErrorType) ToPointer() *SchemasRateLimitErrorType {
	return &e
}

func (e *SchemasRateLimitErrorType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rate_limit_error":
		*e = SchemasRateLimitErrorType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasRateLimitErrorType: %v", v)
	}
}

type RateLimitError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	// rate_limit_error
	Type *SchemasRateLimitErrorType `json:"type,omitempty"`
}

var _ error = &RateLimitError{}

func (e *RateLimitError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
