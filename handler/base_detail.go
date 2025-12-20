package handler

import "ice_sparkhire_runtime/kitex_gen/base"

func ConstructSuccessResp() *base.BaseResp {
	return &base.BaseResp{
		StatusCode:    0,
		StatusMessage: "success",
	}
}

func GenErrorBaseResp(message string) *base.BaseResp {
	return &base.BaseResp{
		StatusCode:    -1,
		StatusMessage: message,
	}
}
