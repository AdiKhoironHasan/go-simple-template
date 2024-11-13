package rest

import "net/http"

// ApiResponse is a function to define the basic api response
// Default response code is 200 with message "OK"
func ApiResponse() *Response {
	return &Response{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}
}

// CursorPaginationResponse is a function to define the cursor pagination response
// It returns the pagination response with type "cursor"
// It requires page, perPage, total, pageSize, and nextCursor as the parameters
func CursorPaginationResponse(pageSize int, nextCursor string) *Pagination {
	return &Pagination{
		Type:       "cursor",
		PageSize:   pageSize,
		NextCursor: nextCursor,
	}
}

// PagePaginationResponse is a function to define the page pagination response
// It returns the pagination response with type "page"
// It requires page, perPage, total, nextPage, and prevPage as the parameters
func PagePaginationResponse(page, perPage, total, nextPage, prevPage int) *Pagination {
	return &Pagination{
		Type:     "page",
		Page:     page,
		PerPage:  perPage,
		Total:    total,
		NextPage: nextPage,
		PrevPage: prevPage,
	}
}

func ErrorResponse(field string, message string) Error {
	return Error{
		Field:   field,
		Message: message,
	}
}

func (resp *Response) WithData(data interface{}) *Response {
	resp.Data = data

	return resp
}

func (resp *Response) WithCode(code int) *Response {
	resp.Meta.Code = code
	resp.Meta.Message = http.StatusText(code)

	return resp
}

func (resp *Response) WithPagination(pagination *Pagination) *Response {
	resp.Meta.Pagination = pagination

	return resp
}

func (resp *Response) WithErrors(errors ...Error) *Response {
	resp.Meta.Errors = append(resp.Meta.Errors, errors...)

	return resp
}

func (resp *Response) WithTraceId(traceId string) *Response {
	resp.Meta.Trace = &Trace{
		TraceId: traceId,
	}

	return resp
}

func (resp *Response) WithMessage(message string) *Response {
	resp.Meta.Message = message

	return resp
}
