/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TrusthubV1SupportingDocument struct for TrusthubV1SupportingDocument
type TrusthubV1SupportingDocument struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The set of parameters that compose the Supporting Documents resource
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The image type of the file
	MimeType *string `json:"mime_type,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The verification status of the Supporting Document resource
	Status *string `json:"status,omitempty"`
	// The type of the Supporting Document
	Type *string `json:"type,omitempty"`
	// The absolute URL of the Supporting Document resource
	Url *string `json:"url,omitempty"`
}
