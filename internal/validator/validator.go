package validator

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type CustomValidator interface {
	StructCtx(ctx context.Context, s interface{}) (err error)
}

func NewValidator() CustomValidator {
	return validator.New()
}
