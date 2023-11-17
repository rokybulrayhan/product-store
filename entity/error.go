package entity

import "errors"

var (
	ErrSomethingWentWrong   = errors.New("something went wrong")
	ErrNoDataMatched        = errors.New("no record found")
	ErrInvalidRefreshToken  = errors.New("invalid refresh token")
	ErrSlugGenerationFailed = errors.New("custom_error!!!slug generation failed")
	ErrSlugExists           = errors.New("custom_error!!!slug already exits")
)
