package business

const (
	//普通错误
	Unknown            = 1000
	InvalidBodyParse   = 1001
	NotFound           = 1002
	Forbidden          = 1003
	ServiceUnavailable = 1004
	MethodNowAllowed   = 1005
	PathNotFound       = 1006

	//资源错误
	MemberNotFound = 1100

	//验证错误
	TokenRequired = 1200
	TokenInvalid  = 1201
	TokenExpired  = 1202
)
