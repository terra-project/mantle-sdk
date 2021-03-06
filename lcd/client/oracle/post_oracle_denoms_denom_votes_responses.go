// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// PostOracleDenomsDenomVotesReader is a Reader for the PostOracleDenomsDenomVotes structure.
type PostOracleDenomsDenomVotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOracleDenomsDenomVotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOracleDenomsDenomVotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOracleDenomsDenomVotesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOracleDenomsDenomVotesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOracleDenomsDenomVotesOK creates a PostOracleDenomsDenomVotesOK with default headers values
func NewPostOracleDenomsDenomVotesOK() *PostOracleDenomsDenomVotesOK {
	return &PostOracleDenomsDenomVotesOK{}
}

/*PostOracleDenomsDenomVotesOK handles this case with default header values.

OK
*/
type PostOracleDenomsDenomVotesOK struct {
	Payload *models.StdTx
}

func (o *PostOracleDenomsDenomVotesOK) Error() string {
	return fmt.Sprintf("[POST /oracle/denoms/{denom}/votes][%d] postOracleDenomsDenomVotesOK  %+v", 200, o.Payload)
}

func (o *PostOracleDenomsDenomVotesOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *PostOracleDenomsDenomVotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOracleDenomsDenomVotesBadRequest creates a PostOracleDenomsDenomVotesBadRequest with default headers values
func NewPostOracleDenomsDenomVotesBadRequest() *PostOracleDenomsDenomVotesBadRequest {
	return &PostOracleDenomsDenomVotesBadRequest{}
}

/*PostOracleDenomsDenomVotesBadRequest handles this case with default header values.

Bad request
*/
type PostOracleDenomsDenomVotesBadRequest struct {
}

func (o *PostOracleDenomsDenomVotesBadRequest) Error() string {
	return fmt.Sprintf("[POST /oracle/denoms/{denom}/votes][%d] postOracleDenomsDenomVotesBadRequest ", 400)
}

func (o *PostOracleDenomsDenomVotesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOracleDenomsDenomVotesInternalServerError creates a PostOracleDenomsDenomVotesInternalServerError with default headers values
func NewPostOracleDenomsDenomVotesInternalServerError() *PostOracleDenomsDenomVotesInternalServerError {
	return &PostOracleDenomsDenomVotesInternalServerError{}
}

/*PostOracleDenomsDenomVotesInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostOracleDenomsDenomVotesInternalServerError struct {
}

func (o *PostOracleDenomsDenomVotesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oracle/denoms/{denom}/votes][%d] postOracleDenomsDenomVotesInternalServerError ", 500)
}

func (o *PostOracleDenomsDenomVotesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
