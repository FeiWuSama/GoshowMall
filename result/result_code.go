package result

type Code struct {
	Code int
	Msg  string
}

var (
	OK = Code{
		Code: 20000,
		Msg:  "Ok",
	}
	ServerError = Code{
		Code: 50000,
		Msg:  "Server Error",
	}
	ParamError = Code{
		Code: 40000,
		Msg:  "Param Error",
	}
	NotFound = Code{
		Code: 40400,
		Msg:  "Not Found",
	}
	Unauthorized = Code{
		Code: 40100,
		Msg:  "Unauthorized",
	}
	PermissionDenied = Code{
		Code: 40300,
		Msg:  "Permission Denied",
	}
)
