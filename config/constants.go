package config

import "fmt"

// url-related
var (
	SignHost         = "api-takumi.mihoyo.com"
	SignBaseUrl      = fmt.Sprintf("https://%s/event/bbs_sign_reward", SignHost)
	SignAwardHomeUri = "/home"
	SignAwardSignUri = "/sign"
	SignAwardInfoUri = "/info"
)

// request-related
const (
	ActId             = "e202009291139501"
	XRpcClientType    = "2"
	XRpcClientVersion = "2.28.1"
	DsSalt            = "dWCcD2FsOUXEstC5f9xubswZxEeoBOTc"
)

// Region-related
const (
	RegionCnGf = "cn_gf01" // 中国官服
	RegionCnQd = "cn_qd01" // 中国渠道服
)

const (
	HttpQueryTagName = "param"
)

// const for gin
const (
	DefaultConfigName     = "config"
	DefaultConfigType     = "yaml"
	DefaultConfigFilename = DefaultConfigName + "." + DefaultConfigType
	DefaultConfigPath     = "./"

	Version      = "0.0.1"
	Project      = "mihoyo-bbs-genshin-sign-backend"
	CtxKeyUserId = "userId"
	CtxKeyUserCn = "userCn"
)

const (
	LogTimeFormat = "2006-01-02 15:04:05.000"
)
