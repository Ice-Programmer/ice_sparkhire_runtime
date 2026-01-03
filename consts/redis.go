package consts

import "time"

const (
	ServerName        = "ice.sparkhire.runtime"
	ServerRedisPrefix = "ice:sparkhire:"

	RecordValue = "record"
)

// 登录相关
const (
	HasSendCodePrefixKey = ServerRedisPrefix + "has_send:"
	VerifyCodePrefixKey  = ServerRedisPrefix + "verify_code:"

	VerifyCodeDuration        = 5 * time.Minute
	HasSendVerifyCodeDuration = 1 * time.Minute
)

// information
const (
	MajorListKey    = ServerRedisPrefix + "major_list"
	IndustryListKey = ServerRedisPrefix + "industry_list"
	SchoolListKey   = ServerRedisPrefix + "school_list"
	GeoListKey      = ServerRedisPrefix + "geo_level"
	CareerListKey   = ServerRedisPrefix + "career_list"

	InformationListDuration = 30 * 24 * time.Hour
)
