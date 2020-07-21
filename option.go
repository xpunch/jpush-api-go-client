package jpushclient

// Option https://docs.jiguang.cn//jpush/server/push/rest_api_v3_push/#options
type Option struct {
	SendNo            int         `json:"sendno,omitempty"`
	TimeLive          int         `json:"time_to_live,omitempty"`
	OverrideMsgID     int64       `json:"override_msg_id,omitempty"`
	ApnsProduction    bool        `json:"apns_production"`
	ApnsCollapseID    string      `json:"apns_collapse_id,omitempty"`
	BigPushDuration   int         `json:"big_push_duration,omitempty"`
	ThirdPartyChannel interface{} `json:"third_party_channel,omitempty"`
}

func (this *Option) SetSendno(no int) {
	this.SendNo = no
}

func (this *Option) SetTimelive(timelive int) {
	this.TimeLive = timelive
}

func (this *Option) SetOverrideMsgId(id int64) {
	this.OverrideMsgID = id
}

func (this *Option) SetApns(apns bool) {
	this.ApnsProduction = apns
}

func (this *Option) SetBigPushDuration(bigPushDuration int) {
	this.BigPushDuration = bigPushDuration
}
