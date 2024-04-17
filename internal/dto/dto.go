package dto

import "net/http"

/*
ApiResponse is a struct to define the response
Meta is a struct to define meta response
Message is a string to define the message of the response
Data is an interface to define the data of the response
*/
type ApiResponse struct {
	Code    int         `json:"code"`
	Meta    *Meta       `json:"meta,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

/*
Meta is a struct to define meta response
Code is an integer to define the status code of the response
Status is a string to define the status of the response
Pagination is a struct to define pagination response
Errors is a slice of Error struct to define error response
*/
type Meta struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Errors     []Error     `json:"errors,omitempty"`
	Trace      *Trace      `json:"trace,omitempty"`
}

/*
Pagination is a struct to define pagination response
Type is a string to define the type of pagination (cursor or page)
Page is an integer to define the current page
PerPage is an integer to define the number of items per page
NextPage is an integer to define the next page
PrevPage is an integer to define the previous page
Total is an integer to define the total number of items
NextCursor is a string to define the next cursor for cursor pagination
PageSize is an integer to define the page size for cursor pagination
*/
type Pagination struct {
	Type string `json:"type"`

	Page     int `json:"page,omitempty"`
	PerPage  int `json:"per_page,omitempty"`
	NextPage int `json:"next_page,omitempty"`
	PrevPage int `json:"prev_page,omitempty"`
	Total    int `json:"total,omitempty"`

	NextCursor string `json:"next_cursor,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
}

/*
Error is a struct to define error response
Field is a string to define the field of the error
Message is a string to define the message of the error
*/
type Error struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type Trace struct {
	TraceId string `json:"trace_id"`
}

func (resp *ApiResponse) WithData(data interface{}) *ApiResponse {
	resp.Data = data

	return resp
}

func (resp *ApiResponse) WithCode(code int) *ApiResponse {
	resp.Code = code

	if resp.Message == http.StatusText(http.StatusOK) {
		resp.Message = http.StatusText(code)
	}

	return resp
}

func (resp *ApiResponse) WithMessage(message string) *ApiResponse {
	resp.Message = message

	return resp
}

func (resp *ApiResponse) WithPagination(pagination *Pagination) *ApiResponse {
	if resp.Meta == nil {
		resp.Meta = &Meta{
			Pagination: pagination,
		}
	} else {
		resp.Meta.Pagination = pagination
	}

	return resp
}

func (resp *ApiResponse) WithErrors(errors ...Error) *ApiResponse {
	if len(errors) == 0 {
		return resp
	}

	if resp.Meta == nil {
		resp.Meta = &Meta{
			Errors: errors,
		}
	} else {
		resp.Meta.Errors = errors
	}

	return resp
}

func (resp *ApiResponse) WithTraceId(traceId string) *ApiResponse {
	if resp.Meta == nil {
		resp.Meta = &Meta{
			Trace: &Trace{
				TraceId: traceId,
			},
		}
	} else {
		resp.Meta.Trace = &Trace{
			TraceId: traceId,
		}
	}

	return resp
}
