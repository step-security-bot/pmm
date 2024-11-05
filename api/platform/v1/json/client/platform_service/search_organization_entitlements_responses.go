// Code generated by go-swagger; DO NOT EDIT.

package platform_service

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
	"github.com/go-openapi/validate"
)

// SearchOrganizationEntitlementsReader is a Reader for the SearchOrganizationEntitlements structure.
type SearchOrganizationEntitlementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchOrganizationEntitlementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOrganizationEntitlementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchOrganizationEntitlementsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchOrganizationEntitlementsOK creates a SearchOrganizationEntitlementsOK with default headers values
func NewSearchOrganizationEntitlementsOK() *SearchOrganizationEntitlementsOK {
	return &SearchOrganizationEntitlementsOK{}
}

/*
SearchOrganizationEntitlementsOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchOrganizationEntitlementsOK struct {
	Payload *SearchOrganizationEntitlementsOKBody
}

// IsSuccess returns true when this search organization entitlements Ok response has a 2xx status code
func (o *SearchOrganizationEntitlementsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search organization entitlements Ok response has a 3xx status code
func (o *SearchOrganizationEntitlementsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search organization entitlements Ok response has a 4xx status code
func (o *SearchOrganizationEntitlementsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search organization entitlements Ok response has a 5xx status code
func (o *SearchOrganizationEntitlementsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search organization entitlements Ok response a status code equal to that given
func (o *SearchOrganizationEntitlementsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search organization entitlements Ok response
func (o *SearchOrganizationEntitlementsOK) Code() int {
	return 200
}

func (o *SearchOrganizationEntitlementsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/platform/organization/entitlements][%d] searchOrganizationEntitlementsOk %s", 200, payload)
}

func (o *SearchOrganizationEntitlementsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/platform/organization/entitlements][%d] searchOrganizationEntitlementsOk %s", 200, payload)
}

func (o *SearchOrganizationEntitlementsOK) GetPayload() *SearchOrganizationEntitlementsOKBody {
	return o.Payload
}

func (o *SearchOrganizationEntitlementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(SearchOrganizationEntitlementsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchOrganizationEntitlementsDefault creates a SearchOrganizationEntitlementsDefault with default headers values
func NewSearchOrganizationEntitlementsDefault(code int) *SearchOrganizationEntitlementsDefault {
	return &SearchOrganizationEntitlementsDefault{
		_statusCode: code,
	}
}

/*
SearchOrganizationEntitlementsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SearchOrganizationEntitlementsDefault struct {
	_statusCode int

	Payload *SearchOrganizationEntitlementsDefaultBody
}

// IsSuccess returns true when this search organization entitlements default response has a 2xx status code
func (o *SearchOrganizationEntitlementsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search organization entitlements default response has a 3xx status code
func (o *SearchOrganizationEntitlementsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search organization entitlements default response has a 4xx status code
func (o *SearchOrganizationEntitlementsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search organization entitlements default response has a 5xx status code
func (o *SearchOrganizationEntitlementsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search organization entitlements default response a status code equal to that given
func (o *SearchOrganizationEntitlementsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search organization entitlements default response
func (o *SearchOrganizationEntitlementsDefault) Code() int {
	return o._statusCode
}

func (o *SearchOrganizationEntitlementsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/platform/organization/entitlements][%d] SearchOrganizationEntitlements default %s", o._statusCode, payload)
}

func (o *SearchOrganizationEntitlementsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/platform/organization/entitlements][%d] SearchOrganizationEntitlements default %s", o._statusCode, payload)
}

func (o *SearchOrganizationEntitlementsDefault) GetPayload() *SearchOrganizationEntitlementsDefaultBody {
	return o.Payload
}

func (o *SearchOrganizationEntitlementsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(SearchOrganizationEntitlementsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SearchOrganizationEntitlementsDefaultBody search organization entitlements default body
swagger:model SearchOrganizationEntitlementsDefaultBody
*/
type SearchOrganizationEntitlementsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*SearchOrganizationEntitlementsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this search organization entitlements default body
func (o *SearchOrganizationEntitlementsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationEntitlementsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("SearchOrganizationEntitlements default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SearchOrganizationEntitlements default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search organization entitlements default body based on the context it is used
func (o *SearchOrganizationEntitlementsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationEntitlementsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SearchOrganizationEntitlements default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SearchOrganizationEntitlements default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsDefaultBody) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationEntitlementsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchOrganizationEntitlementsDefaultBodyDetailsItems0 search organization entitlements default body details items0
swagger:model SearchOrganizationEntitlementsDefaultBodyDetailsItems0
*/
type SearchOrganizationEntitlementsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// search organization entitlements default body details items0
	SearchOrganizationEntitlementsDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *SearchOrganizationEntitlementsDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv SearchOrganizationEntitlementsDefaultBodyDetailsItems0

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
		o.SearchOrganizationEntitlementsDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o SearchOrganizationEntitlementsDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.SearchOrganizationEntitlementsDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.SearchOrganizationEntitlementsDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this search organization entitlements default body details items0
func (o *SearchOrganizationEntitlementsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search organization entitlements default body details items0 based on context it is used
func (o *SearchOrganizationEntitlementsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationEntitlementsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchOrganizationEntitlementsOKBody search organization entitlements OK body
swagger:model SearchOrganizationEntitlementsOKBody
*/
type SearchOrganizationEntitlementsOKBody struct {
	// entitlements
	Entitlements []*SearchOrganizationEntitlementsOKBodyEntitlementsItems0 `json:"entitlements"`
}

// Validate validates this search organization entitlements OK body
func (o *SearchOrganizationEntitlementsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEntitlements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationEntitlementsOKBody) validateEntitlements(formats strfmt.Registry) error {
	if swag.IsZero(o.Entitlements) { // not required
		return nil
	}

	for i := 0; i < len(o.Entitlements); i++ {
		if swag.IsZero(o.Entitlements[i]) { // not required
			continue
		}

		if o.Entitlements[i] != nil {
			if err := o.Entitlements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchOrganizationEntitlementsOk" + "." + "entitlements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchOrganizationEntitlementsOk" + "." + "entitlements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search organization entitlements OK body based on the context it is used
func (o *SearchOrganizationEntitlementsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEntitlements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationEntitlementsOKBody) contextValidateEntitlements(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Entitlements); i++ {
		if o.Entitlements[i] != nil {

			if swag.IsZero(o.Entitlements[i]) { // not required
				return nil
			}

			if err := o.Entitlements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchOrganizationEntitlementsOk" + "." + "entitlements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchOrganizationEntitlementsOk" + "." + "entitlements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsOKBody) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationEntitlementsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchOrganizationEntitlementsOKBodyEntitlementsItems0 OrganizationEntitlement contains information about Organization entitlement.
swagger:model SearchOrganizationEntitlementsOKBodyEntitlementsItems0
*/
type SearchOrganizationEntitlementsOKBodyEntitlementsItems0 struct {
	// Entitlement number.
	Number string `json:"number,omitempty"`

	// Entitlement name.
	Name string `json:"name,omitempty"`

	// Entitlement short summary.
	Summary string `json:"summary,omitempty"`

	// Entitlement tier.
	Tier *string `json:"tier,omitempty"`

	// Total units covered by this entitlement.
	TotalUnits *string `json:"total_units,omitempty"`

	// Flag indicates that unlimited units are covered.
	UnlimitedUnits *bool `json:"unlimited_units,omitempty"`

	// Support level covered by this entitlement.
	SupportLevel *string `json:"support_level,omitempty"`

	// Percona product families covered by this entitlement.
	SoftwareFamilies []string `json:"software_families"`

	// Entitlement start data.
	// Note: only date is used here but not time.
	// Format: date-time
	StartDate strfmt.DateTime `json:"start_date,omitempty"`

	// Entitlement end date.
	// Note: only date is used here but not time.
	// Format: date-time
	EndDate strfmt.DateTime `json:"end_date,omitempty"`

	// platform
	Platform *SearchOrganizationEntitlementsOKBodyEntitlementsItems0Platform `json:"platform,omitempty"`
}

// Validate validates this search organization entitlements OK body entitlements items0
func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(o.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date-time", o.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(o.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("end_date", "body", "date-time", o.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(o.Platform) { // not required
		return nil
	}

	if o.Platform != nil {
		if err := o.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this search organization entitlements OK body entitlements items0 based on the context it is used
func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {
	if o.Platform != nil {

		if swag.IsZero(o.Platform) { // not required
			return nil
		}

		if err := o.Platform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationEntitlementsOKBodyEntitlementsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchOrganizationEntitlementsOKBodyEntitlementsItems0Platform Platform indicates platform specific entitlements.
swagger:model SearchOrganizationEntitlementsOKBodyEntitlementsItems0Platform
*/
type SearchOrganizationEntitlementsOKBodyEntitlementsItems0Platform struct {
	// Flag indicates that security advisors are covered by this entitlement.
	SecurityAdvisor *string `json:"security_advisor,omitempty"`

	// Flag indicates that config advisors are covered by this entitlement.
	ConfigAdvisor *string `json:"config_advisor,omitempty"`
}

// Validate validates this search organization entitlements OK body entitlements items0 platform
func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0Platform) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search organization entitlements OK body entitlements items0 platform based on context it is used
func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0Platform) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0Platform) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsOKBodyEntitlementsItems0Platform) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationEntitlementsOKBodyEntitlementsItems0Platform
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}