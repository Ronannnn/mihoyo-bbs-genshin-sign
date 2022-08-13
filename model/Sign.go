package model

// SignInfo Data struct from /event/bbs_sign_reward/info
type SignInfo struct {
	TotalSignDay  int    `json:"total_sign_day"`
	Today         string `json:"today"`
	IsSign        bool   `json:"is_sign"`
	FirstBind     bool   `json:"first_bind"`
	IsSub         bool   `json:"is_sub"`
	MonthFirst    bool   `json:"month_first"`
	SignCntMissed int    `json:"sign_cnt_missed"`
}

// SignUrlParam Sign info request url parameters
type SignUrlParam struct {
	ActId  string `param:"act_id"`
	Uid    string `param:"uid"`
	Region string `param:"region"`
}

// SignAward Sign Award Item Details
type SignAward struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
	Cnt  int    `json:"cnt"`
}

// SignAwardList Sign award List
type SignAwardList struct {
	Month  int         `json:"month"`
	Awards []SignAward `json:"awards"`
	Resign bool        `json:"resign"`
}

// SignAwardsInfoReqParam Sign award list request url parameters
type SignAwardsInfoReqParam struct {
	ActId string `param:"act_id"`
}
