package dto

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data,omitempty"`
}

type Meta struct {
	Code             int         `json:"code"`
	Message          string      `json:"message"`
	ValidationErrors []Error     `json:"validation_errors,omitempty"`
	Error            any         `json:"error,omitempty"`
	Pagination       *Pagination `json:"pagination,omitempty"`
}

type Error struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type Pagination struct {
	Page      uint64 `json:"page"`
	Limit     uint64 `json:"limit"`
	Total     uint64 `json:"total,omitempty"`
	TotalPage uint64 `json:"total_page,omitempty"`
}
