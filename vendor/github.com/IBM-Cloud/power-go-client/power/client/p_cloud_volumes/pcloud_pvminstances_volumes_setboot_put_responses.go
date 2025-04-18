// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesVolumesSetbootPutReader is a Reader for the PcloudPvminstancesVolumesSetbootPut structure.
type PcloudPvminstancesVolumesSetbootPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesVolumesSetbootPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesVolumesSetbootPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesVolumesSetbootPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesVolumesSetbootPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesVolumesSetbootPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesVolumesSetbootPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPvminstancesVolumesSetbootPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesVolumesSetbootPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot] pcloud.pvminstances.volumes.setboot.put", response, response.Code())
	}
}

// NewPcloudPvminstancesVolumesSetbootPutOK creates a PcloudPvminstancesVolumesSetbootPutOK with default headers values
func NewPcloudPvminstancesVolumesSetbootPutOK() *PcloudPvminstancesVolumesSetbootPutOK {
	return &PcloudPvminstancesVolumesSetbootPutOK{}
}

/*
PcloudPvminstancesVolumesSetbootPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesVolumesSetbootPutOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put o k response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put o k response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put o k response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put o k response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put o k response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances volumes setboot put o k response
func (o *PcloudPvminstancesVolumesSetbootPutOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutOK %s", 200, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutOK %s", 200, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutBadRequest creates a PcloudPvminstancesVolumesSetbootPutBadRequest with default headers values
func NewPcloudPvminstancesVolumesSetbootPutBadRequest() *PcloudPvminstancesVolumesSetbootPutBadRequest {
	return &PcloudPvminstancesVolumesSetbootPutBadRequest{}
}

/*
PcloudPvminstancesVolumesSetbootPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesVolumesSetbootPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put bad request response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put bad request response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put bad request response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put bad request response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put bad request response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances volumes setboot put bad request response
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutUnauthorized creates a PcloudPvminstancesVolumesSetbootPutUnauthorized with default headers values
func NewPcloudPvminstancesVolumesSetbootPutUnauthorized() *PcloudPvminstancesVolumesSetbootPutUnauthorized {
	return &PcloudPvminstancesVolumesSetbootPutUnauthorized{}
}

/*
PcloudPvminstancesVolumesSetbootPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesVolumesSetbootPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put unauthorized response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put unauthorized response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put unauthorized response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put unauthorized response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put unauthorized response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances volumes setboot put unauthorized response
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutForbidden creates a PcloudPvminstancesVolumesSetbootPutForbidden with default headers values
func NewPcloudPvminstancesVolumesSetbootPutForbidden() *PcloudPvminstancesVolumesSetbootPutForbidden {
	return &PcloudPvminstancesVolumesSetbootPutForbidden{}
}

/*
PcloudPvminstancesVolumesSetbootPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesVolumesSetbootPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put forbidden response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put forbidden response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put forbidden response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put forbidden response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put forbidden response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances volumes setboot put forbidden response
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesVolumesSetbootPutForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutNotFound creates a PcloudPvminstancesVolumesSetbootPutNotFound with default headers values
func NewPcloudPvminstancesVolumesSetbootPutNotFound() *PcloudPvminstancesVolumesSetbootPutNotFound {
	return &PcloudPvminstancesVolumesSetbootPutNotFound{}
}

/*
PcloudPvminstancesVolumesSetbootPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesVolumesSetbootPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put not found response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put not found response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put not found response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put not found response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put not found response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances volumes setboot put not found response
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutConflict creates a PcloudPvminstancesVolumesSetbootPutConflict with default headers values
func NewPcloudPvminstancesVolumesSetbootPutConflict() *PcloudPvminstancesVolumesSetbootPutConflict {
	return &PcloudPvminstancesVolumesSetbootPutConflict{}
}

/*
PcloudPvminstancesVolumesSetbootPutConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPvminstancesVolumesSetbootPutConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put conflict response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put conflict response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put conflict response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put conflict response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put conflict response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud pvminstances volumes setboot put conflict response
func (o *PcloudPvminstancesVolumesSetbootPutConflict) Code() int {
	return 409
}

func (o *PcloudPvminstancesVolumesSetbootPutConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutConflict %s", 409, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutConflict %s", 409, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutInternalServerError creates a PcloudPvminstancesVolumesSetbootPutInternalServerError with default headers values
func NewPcloudPvminstancesVolumesSetbootPutInternalServerError() *PcloudPvminstancesVolumesSetbootPutInternalServerError {
	return &PcloudPvminstancesVolumesSetbootPutInternalServerError{}
}

/*
PcloudPvminstancesVolumesSetbootPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesVolumesSetbootPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put internal server error response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put internal server error response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put internal server error response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put internal server error response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances volumes setboot put internal server error response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances volumes setboot put internal server error response
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
