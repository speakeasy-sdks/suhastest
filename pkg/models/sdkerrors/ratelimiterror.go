// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/speakeasy-sdks/suhastest/pkg/models/shared"
)

// RateLimitError - Rate Limit Error
type RateLimitError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	// rate_limit_error
	Type *shared.RateLimitErrorType `json:"type,omitempty"`
}

var _ error = &RateLimitError{}

func (e *RateLimitError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}