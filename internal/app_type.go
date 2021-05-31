package internal

const (
	AppTypeWeb     AppType = "web"
	AppTypeApi     AppType = "api"
	AppTypeGrpc    AppType = "grpc"
	AppTypeApiGrpc AppType = "api+grpc"
	AppTypeWebGrpc AppType = "web+grpc"
	AppTypeAll     AppType = "*"
	AppTypeDefault AppType = "default"
	AppTypeUnknown AppType = ""
)

type AppType string

func GetAppType() AppType {
	var di = GetDi()
	v := di.Get(GAppType)
	if v == ErrorNilDi {
		return AppTypeUnknown
	}
	switch v.(type) {
	case AppType:
		return v.(AppType)
	case string:
		return ParseAppType(v.(string))
	}
	return AppTypeUnknown
}

func IsGrpcApp() bool {
	return false
}

func RegisterAppType(name AppType) bool {
	var appType = GetAppType()
	if appType.Empty() {
		return Register(GAppType, name)
	}
	return false
}

func (s AppType) String() string {
	return string(s)
}

func (s AppType) IsGrpcApp() bool {
	switch s {
	case AppTypeGrpc:
		return true
	case AppTypeApiGrpc:
		return true
	case AppTypeWebGrpc:
		return true
	case AppTypeAll:
		return true
	}
	return false
}

func (s AppType) Empty() bool {
	return s.String() == ""
}

func ParseAppType(str string) AppType {
	return AppType(str)
}
