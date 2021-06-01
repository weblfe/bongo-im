package bootstrap

import "bongo-im/internal"

// 初始化路由
func Initialize() {
	createRouter(internal.GetAppType())
}

// 创建路由
func createRouter(appType internal.AppType){
	if appType.Empty() {
		InitApiHandler()
		InitWebHandler()
		return
	}
	// 按应用类型初始化路由
	switch appType {
	case internal.AppTypeWeb:
		InitWebHandler()
	case internal.AppTypeGrpc:
		InitGrpcHandler()
	case internal.AppTypeApi:
		InitApiHandler()
	case internal.AppTypeAll:
		InitWebHandler()
		InitGrpcHandler()
		InitApiHandler()
	case internal.AppTypeApiGrpc:
		InitGrpcHandler()
		InitApiHandler()
	case internal.AppTypeWebGrpc:
		InitGrpcHandler()
		InitWebHandler()
	case internal.AppTypeDefault:
		fallthrough
	default:
		InitApiHandler()
		InitWebHandler()
	}
}
