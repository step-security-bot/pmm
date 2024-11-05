// Code generated by go-swagger; DO NOT EDIT.

package qan_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExplainFingerprintByQueryIDReader is a Reader for the ExplainFingerprintByQueryID structure.
type ExplainFingerprintByQueryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExplainFingerprintByQueryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExplainFingerprintByQueryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExplainFingerprintByQueryIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExplainFingerprintByQueryIDOK creates a ExplainFingerprintByQueryIDOK with default headers values
func NewExplainFingerprintByQueryIDOK() *ExplainFingerprintByQueryIDOK {
	return &ExplainFingerprintByQueryIDOK{}
}

/*
ExplainFingerprintByQueryIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type ExplainFingerprintByQueryIDOK struct {
	Payload *ExplainFingerprintByQueryIDOKBody
}

// IsSuccess returns true when this explain fingerprint by query Id Ok response has a 2xx status code
func (o *ExplainFingerprintByQueryIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this explain fingerprint by query Id Ok response has a 3xx status code
func (o *ExplainFingerprintByQueryIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this explain fingerprint by query Id Ok response has a 4xx status code
func (o *ExplainFingerprintByQueryIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this explain fingerprint by query Id Ok response has a 5xx status code
func (o *ExplainFingerprintByQueryIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this explain fingerprint by query Id Ok response a status code equal to that given
func (o *ExplainFingerprintByQueryIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the explain fingerprint by query Id Ok response
func (o *ExplainFingerprintByQueryIDOK) Code() int {
	return 200
}

func (o *ExplainFingerprintByQueryIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan:explainFingerprint][%d] explainFingerprintByQueryIdOk %s", 200, payload)
}

func (o *ExplainFingerprintByQueryIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan:explainFingerprint][%d] explainFingerprintByQueryIdOk %s", 200, payload)
}

func (o *ExplainFingerprintByQueryIDOK) GetPayload() *ExplainFingerprintByQueryIDOKBody {
	return o.Payload
}

func (o *ExplainFingerprintByQueryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ExplainFingerprintByQueryIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExplainFingerprintByQueryIDDefault creates a ExplainFingerprintByQueryIDDefault with default headers values
func NewExplainFingerprintByQueryIDDefault(code int) *ExplainFingerprintByQueryIDDefault {
	return &ExplainFingerprintByQueryIDDefault{
		_statusCode: code,
	}
}

/*
ExplainFingerprintByQueryIDDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ExplainFingerprintByQueryIDDefault struct {
	_statusCode int

	Payload *ExplainFingerprintByQueryIDDefaultBody
}

// IsSuccess returns true when this explain fingerprint by query ID default response has a 2xx status code
func (o *ExplainFingerprintByQueryIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this explain fingerprint by query ID default response has a 3xx status code
func (o *ExplainFingerprintByQueryIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this explain fingerprint by query ID default response has a 4xx status code
func (o *ExplainFingerprintByQueryIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this explain fingerprint by query ID default response has a 5xx status code
func (o *ExplainFingerprintByQueryIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this explain fingerprint by query ID default response a status code equal to that given
func (o *ExplainFingerprintByQueryIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the explain fingerprint by query ID default response
func (o *ExplainFingerprintByQueryIDDefault) Code() int {
	return o._statusCode
}

func (o *ExplainFingerprintByQueryIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan:explainFingerprint][%d] ExplainFingerprintByQueryID default %s", o._statusCode, payload)
}

func (o *ExplainFingerprintByQueryIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan:explainFingerprint][%d] ExplainFingerprintByQueryID default %s", o._statusCode, payload)
}

func (o *ExplainFingerprintByQueryIDDefault) GetPayload() *ExplainFingerprintByQueryIDDefaultBody {
	return o.Payload
}

func (o *ExplainFingerprintByQueryIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ExplainFingerprintByQueryIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExplainFingerprintByQueryIDBody ExplainFingerprintByQueryIDRequest get explain fingerprint for given query ID.
swagger:model ExplainFingerprintByQueryIDBody
*/
type ExplainFingerprintByQueryIDBody struct {
	// serviceid
	Serviceid string `json:"serviceid,omitempty"`

	// query id
	QueryID string `json:"query_id,omitempty"`
}

// Validate validates this explain fingerprint by query ID body
func (o *ExplainFingerprintByQueryIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this explain fingerprint by query ID body based on context it is used
func (o *ExplainFingerprintByQueryIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDBody) UnmarshalBinary(b []byte) error {
	var res ExplainFingerprintByQueryIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ExplainFingerprintByQueryIDDefaultBody explain fingerprint by query ID default body
swagger:model ExplainFingerprintByQueryIDDefaultBody
*/
type ExplainFingerprintByQueryIDDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ExplainFingerprintByQueryIDDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this explain fingerprint by query ID default body
func (o *ExplainFingerprintByQueryIDDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExplainFingerprintByQueryIDDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExplainFingerprintByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExplainFingerprintByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this explain fingerprint by query ID default body based on the context it is used
func (o *ExplainFingerprintByQueryIDDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExplainFingerprintByQueryIDDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExplainFingerprintByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExplainFingerprintByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDDefaultBody) UnmarshalBinary(b []byte) error {
	var res ExplainFingerprintByQueryIDDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ExplainFingerprintByQueryIDDefaultBodyDetailsItems0 explain fingerprint by query ID default body details items0
swagger:model ExplainFingerprintByQueryIDDefaultBodyDetailsItems0
*/
type ExplainFingerprintByQueryIDDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// explain fingerprint by query ID default body details items0
	ExplainFingerprintByQueryIDDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ExplainFingerprintByQueryIDDefaultBodyDetailsItems0

	rcv.AtType = stage1.AtType
	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "@type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		o.ExplainFingerprintByQueryIDDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}

	stage1.AtType = o.AtType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.ExplainFingerprintByQueryIDDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this explain fingerprint by query ID default body details items0
func (o *ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this explain fingerprint by query ID default body details items0 based on context it is used
func (o *ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ExplainFingerprintByQueryIDDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ExplainFingerprintByQueryIDOKBody ExplainFingerprintByQueryIDResponse is explain fingerprint and placeholders count for given query ID.
swagger:model ExplainFingerprintByQueryIDOKBody
*/
type ExplainFingerprintByQueryIDOKBody struct {
	// explain fingerprint
	ExplainFingerprint string `json:"explain_fingerprint,omitempty"`

	// placeholders count
	PlaceholdersCount int64 `json:"placeholders_count,omitempty"`
}

// Validate validates this explain fingerprint by query ID OK body
func (o *ExplainFingerprintByQueryIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this explain fingerprint by query ID OK body based on context it is used
func (o *ExplainFingerprintByQueryIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDOKBody) UnmarshalBinary(b []byte) error {
	var res ExplainFingerprintByQueryIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}