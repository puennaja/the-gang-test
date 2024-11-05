package goerrors

const (
	BadRequestStatus          = 400
	UnauthorizedStatus        = 401
	ForbiddenStatus           = 403
	NotFoundStatus            = 404
	InternalServerErrorStatus = 500

	DefaultCode          int = 1100
	ValidationCode       int = 1101
	MongoCode            int = 1102
	UnauthorizedCode     int = 1103
	PermissionDeniedCode int = 1104

	CategoryNotFoundCode      int = 1210
	CategoryAlreadyExistsCode int = 1211

	ListingNotFoundCode      int = 1220
	ListingAlreadyExistsCode int = 1221

	MessageNotFoundCode int = 1230
)

var (
	messages = map[int]string{
		DefaultCode:               "something went wrong",
		ValidationCode:            "validation error",
		MongoCode:                 "mongo error",
		UnauthorizedCode:          "unauthorized",
		PermissionDeniedCode:      "permission denied",
		CategoryNotFoundCode:      "category not found",
		CategoryAlreadyExistsCode: "category already exists",
		ListingNotFoundCode:       "listing not found",
		ListingAlreadyExistsCode:  "listing already exists",
		MessageNotFoundCode:       "message not found",
	}
)

var (
	ErrDefault               = NewInternalErr(InternalServerErrorStatus, DefaultCode, messages[DefaultCode])
	ErrValidation            = NewInternalErr(BadRequestStatus, ValidationCode, messages[ValidationCode])
	ErrMongo                 = NewInternalErr(InternalServerErrorStatus, MongoCode, messages[MongoCode])
	ErrUnauthorized          = NewInternalErr(UnauthorizedStatus, UnauthorizedCode, messages[UnauthorizedCode])
	ErrPermissionDenied      = NewInternalErr(ForbiddenStatus, PermissionDeniedCode, messages[PermissionDeniedCode])
	ErrCategoryNotFound      = NewInternalErr(NotFoundStatus, CategoryNotFoundCode, messages[CategoryNotFoundCode])
	ErrCategoryAlreadyExists = NewInternalErr(InternalServerErrorStatus, CategoryAlreadyExistsCode, messages[CategoryAlreadyExistsCode])
	ErrListingNotFound       = NewInternalErr(NotFoundStatus, ListingNotFoundCode, messages[ListingNotFoundCode])
	ErrListingAlreadyExists  = NewInternalErr(InternalServerErrorStatus, ListingAlreadyExistsCode, messages[ListingAlreadyExistsCode])
	ErrMessageNotFound       = NewInternalErr(NotFoundStatus, MessageNotFoundCode, messages[MessageNotFoundCode])
)
