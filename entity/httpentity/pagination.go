package httpentity

import "github.com/go-contact-service/entity"

const DEFAULT_LIMIT = 10

// TODO: Add validation
type PaginationRequest struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}

func (input *PaginationRequest) Validate() []FieldError {
	return validate(input)
}

func (input *PaginationRequest) GetLimitOffset() entity.Pagination {
	if input.Limit > 50 {
		input.Limit = 50
	} else if input.Limit < 1 {
		input.Limit = DEFAULT_LIMIT
	}
	if input.Page < 1 {
		input.Page = 0
	} else {
		input.Page--
	}
	return entity.Pagination{
		Limit:  input.Limit,
		Offset: input.Page * input.Limit,
	}
}
func (input *PaginationRequest) GetLimit() int {
	if input.Limit > 50 {
		input.Limit = 50
	}
	return input.Limit
}

type PaginationResponse struct {
	Total int `json:"total"`
	Pages int `json:"pages"`
}

func NewPaginationResponse(total, limit int) PaginationResponse {
	return PaginationResponse{
		Total: total,
		Pages: (total + (limit - 1)) / limit,
	}

}
