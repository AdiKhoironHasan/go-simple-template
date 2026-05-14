package dto

import "net/http"

func RestResponse(code int, data any, err error) *Response {
	response := &Response{
		Meta: Meta{
			Code:    code,
			Message: http.StatusText(code),
		},
		Data: data,
	}

	if err != nil {
		response.Meta.Error = ParseValidationErrors(err)
	}

	return response
}
