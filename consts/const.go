package consts

import "time"

const (
	TokenExpireTime = 30 * 24 * time.Hour
)

const (
	TraceIDHeader       = "X-Trace-ID"
	AuthorizationHeader = "Authorization"
)

const (
	NotLoginErrorCode     = 420
	NoAuthErrorCode       = 450
	InvalidTokenErrorCode = 480
	OperationErrorCode    = 500
)

const EnumNotFound = "<UNSET>"
