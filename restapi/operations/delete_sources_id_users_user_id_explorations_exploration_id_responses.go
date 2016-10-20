package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/chronograf/models"
)

/*DeleteSourcesIDUsersUserIDExplorationsExplorationIDNoContent Exploration session has been removed

swagger:response deleteSourcesIdUsersUserIdExplorationsExplorationIdNoContent
*/
type DeleteSourcesIDUsersUserIDExplorationsExplorationIDNoContent struct {
}

// NewDeleteSourcesIDUsersUserIDExplorationsExplorationIDNoContent creates DeleteSourcesIDUsersUserIDExplorationsExplorationIDNoContent with default headers values
func NewDeleteSourcesIDUsersUserIDExplorationsExplorationIDNoContent() *DeleteSourcesIDUsersUserIDExplorationsExplorationIDNoContent {
	return &DeleteSourcesIDUsersUserIDExplorationsExplorationIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteSourcesIDUsersUserIDExplorationsExplorationIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound Data source id, user, or exploration does not exist.

swagger:response deleteSourcesIdUsersUserIdExplorationsExplorationIdNotFound
*/
type DeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound creates DeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound with default headers values
func NewDeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound() *DeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound {
	return &DeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound{}
}

// WithPayload adds the payload to the delete sources Id users user Id explorations exploration Id not found response
func (o *DeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound) WithPayload(payload *models.Error) *DeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete sources Id users user Id explorations exploration Id not found response
func (o *DeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSourcesIDUsersUserIDExplorationsExplorationIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault Unexpected internal service error

swagger:response deleteSourcesIdUsersUserIdExplorationsExplorationIdDefault
*/
type DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault creates DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault with default headers values
func NewDeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault(code int) *DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete sources ID users user ID explorations exploration ID default response
func (o *DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault) WithStatusCode(code int) *DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete sources ID users user ID explorations exploration ID default response
func (o *DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete sources ID users user ID explorations exploration ID default response
func (o *DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault) WithPayload(payload *models.Error) *DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete sources ID users user ID explorations exploration ID default response
func (o *DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSourcesIDUsersUserIDExplorationsExplorationIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
