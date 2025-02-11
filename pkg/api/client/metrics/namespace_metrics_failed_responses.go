// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NamespaceMetricsFailedReader is a Reader for the NamespaceMetricsFailed structure.
type NamespaceMetricsFailedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamespaceMetricsFailedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNamespaceMetricsFailedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNamespaceMetricsFailedOK creates a NamespaceMetricsFailedOK with default headers values
func NewNamespaceMetricsFailedOK() *NamespaceMetricsFailedOK {
	return &NamespaceMetricsFailedOK{}
}

/* NamespaceMetricsFailedOK describes a response with status code 200, with default header values.

successfully got namespace metrics
*/
type NamespaceMetricsFailedOK struct {
}

func (o *NamespaceMetricsFailedOK) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/metrics/failed][%d] namespaceMetricsFailedOK ", 200)
}

func (o *NamespaceMetricsFailedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
