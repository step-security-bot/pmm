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

// QueryExistsReader is a Reader for the QueryExists structure.
type QueryExistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryExistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryExistsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQueryExistsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryExistsOK creates a QueryExistsOK with default headers values
func NewQueryExistsOK() *QueryExistsOK {
	return &QueryExistsOK{}
}

/*
QueryExistsOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryExistsOK struct {
	Payload *QueryExistsOKBody
}

// IsSuccess returns true when this query exists Ok response has a 2xx status code
func (o *QueryExistsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query exists Ok response has a 3xx status code
func (o *QueryExistsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query exists Ok response has a 4xx status code
func (o *QueryExistsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query exists Ok response has a 5xx status code
func (o *QueryExistsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query exists Ok response a status code equal to that given
func (o *QueryExistsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query exists Ok response
func (o *QueryExistsOK) Code() int {
	return 200
}

func (o *QueryExistsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan/query:exists][%d] queryExistsOk %s", 200, payload)
}

func (o *QueryExistsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan/query:exists][%d] queryExistsOk %s", 200, payload)
}

func (o *QueryExistsOK) GetPayload() *QueryExistsOKBody {
	return o.Payload
}

func (o *QueryExistsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(QueryExistsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryExistsDefault creates a QueryExistsDefault with default headers values
func NewQueryExistsDefault(code int) *QueryExistsDefault {
	return &QueryExistsDefault{
		_statusCode: code,
	}
}

/*
QueryExistsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type QueryExistsDefault struct {
	_statusCode int

	Payload *QueryExistsDefaultBody
}

// IsSuccess returns true when this query exists default response has a 2xx status code
func (o *QueryExistsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query exists default response has a 3xx status code
func (o *QueryExistsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query exists default response has a 4xx status code
func (o *QueryExistsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query exists default response has a 5xx status code
func (o *QueryExistsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query exists default response a status code equal to that given
func (o *QueryExistsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query exists default response
func (o *QueryExistsDefault) Code() int {
	return o._statusCode
}

func (o *QueryExistsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan/query:exists][%d] QueryExists default %s", o._statusCode, payload)
}

func (o *QueryExistsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan/query:exists][%d] QueryExists default %s", o._statusCode, payload)
}

func (o *QueryExistsDefault) GetPayload() *QueryExistsDefaultBody {
	return o.Payload
}

func (o *QueryExistsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(QueryExistsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
QueryExistsBody QueryExistsRequest check if provided query exists or not.
swagger:model QueryExistsBody
*/
type QueryExistsBody struct {
	// serviceid
	Serviceid string `json:"serviceid,omitempty"`

	// query
	Query string `json:"query,omitempty"`
}

// Validate validates this query exists body
func (o *QueryExistsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query exists body based on context it is used
func (o *QueryExistsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryExistsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryExistsBody) UnmarshalBinary(b []byte) error {
	var res QueryExistsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryExistsDefaultBody query exists default body
swagger:model QueryExistsDefaultBody
*/
type QueryExistsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*QueryExistsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this query exists default body
func (o *QueryExistsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryExistsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("QueryExists default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("QueryExists default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this query exists default body based on the context it is used
func (o *QueryExistsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryExistsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("QueryExists default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("QueryExists default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryExistsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryExistsDefaultBody) UnmarshalBinary(b []byte) error {
	var res QueryExistsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryExistsDefaultBodyDetailsItems0 query exists default body details items0
swagger:model QueryExistsDefaultBodyDetailsItems0
*/
type QueryExistsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// query exists default body details items0
	QueryExistsDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *QueryExistsDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv QueryExistsDefaultBodyDetailsItems0

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
		o.QueryExistsDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o QueryExistsDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.QueryExistsDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.QueryExistsDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this query exists default body details items0
func (o *QueryExistsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query exists default body details items0 based on context it is used
func (o *QueryExistsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryExistsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryExistsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res QueryExistsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryExistsOKBody QueryExistsResponse returns true if query exists.
swagger:model QueryExistsOKBody
*/
type QueryExistsOKBody struct {
	// exists
	Exists bool `json:"exists,omitempty"`
}

// Validate validates this query exists OK body
func (o *QueryExistsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query exists OK body based on context it is used
func (o *QueryExistsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryExistsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryExistsOKBody) UnmarshalBinary(b []byte) error {
	var res QueryExistsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}