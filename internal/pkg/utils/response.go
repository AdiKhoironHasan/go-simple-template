package utils

import (
	"go-simple-template/internal/dto"
	"net/http"
)

// ApiResponse is a function to define the basic api response
// Default response code is 200 with message "OK"
func ApiResponse() *dto.ApiResponse {
	return &dto.ApiResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
}

// CursorPaginationResponse is a function to define the cursor pagination response
// It returns the pagination response with type "cursor"
// It requires page, perPage, total, pageSize, and nextCursor as the parameters
func CursorPaginationResponse(pageSize int, nextCursor string) *dto.Pagination {
	return &dto.Pagination{
		Type:       "cursor",
		PageSize:   pageSize,
		NextCursor: nextCursor,
	}
}

// PagePaginationResponse is a function to define the page pagination response
// It returns the pagination response with type "page"
// It requires page, perPage, total, nextPage, and prevPage as the parameters
func PagePaginationResponse(page, perPage, total, nextPage, prevPage int) *dto.Pagination {
	return &dto.Pagination{
		Type:     "page",
		Page:     page,
		PerPage:  perPage,
		Total:    total,
		NextPage: nextPage,
		PrevPage: prevPage,
	}
}

func ErrorResponse(field string, message string) dto.Error {
	return dto.Error{
		Field:   field,
		Message: message,
	}
}
