package httphdl

import (
	"daveslist/internal/core/port"
	errors "daveslist/pkg/go-errors"
	"daveslist/pkg/validator"
)

type Config struct {
	Validator       validator.Validator
	CategoryService port.CategoryService
	ListingService  port.ListingService
	MessageService  port.MessageService
}

type HTTPHandler struct {
	validator       validator.Validator
	categoryService port.CategoryService
	listingService  port.ListingService
	messageService  port.MessageService
}

func NewHTTP(cfg Config) *HTTPHandler {
	return &HTTPHandler{
		validator:       cfg.Validator,
		categoryService: cfg.CategoryService,
		listingService:  cfg.ListingService,
		messageService:  cfg.MessageService,
	}
}

func (hdl *HTTPHandler) validate(request interface{}) error {
	errs := hdl.validator.StrcutWithTranslateError(request)
	if len(errs) != 0 {
		validateErr := errors.ErrValidation
		for _, err := range errs {
			validateErr.SetError(err)
		}
		return validateErr
	}

	return nil

}
