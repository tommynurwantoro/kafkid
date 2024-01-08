package validator

import "context"

type Validator interface {
	Validate(ctx context.Context, data interface{}) error
}
