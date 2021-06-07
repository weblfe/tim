package tim

// 消息体
type MsgItem struct {
	MsgType    string      `json:"MsgType"`
	MsgContent interface{} `json:"MsgContent"` // TIMTextElem（文本消息） TIMLocationElem（位置消息） TIMFaceElem（表情消息） TIMCustomElem（自定义消息） TIMSoundElem（语音消息） TIMImageElem（图像消息） TIMFileElem（文件消息）
}

// 推送条件
type Condition struct {
	AttrsAnd *map[string]interface{} `json:"AttrsAnd,omitempty"`
	AttrsOr  *map[string]interface{} `json:"AttrsOr,omitempty"`
	TagsOr   *map[string]interface{} `json:"TagsOr,omitempty"`
	TagsAnd  *map[string]interface{} `json:"TagsAnd,omitempty"`
}

// 离线相关配置
type OfflinePushInfoObj struct {
	PushFlag    int          `json:"PushFlag"`
	Desc        string       `json:"Desc"`
	Ext         string       `json:"Ext"`
	AndroidInfo *AndroidInfo `json:"AndroidInfo,omitempty"`
	ApnsInfo    *ApnsInfo    `json:"ApnsInfo,omitempty"`
}

// 全员推送
type PushMsg struct {
	FromAccount     string              `json:"From_Account"`
	MsgRandom       int                 `json:"MsgRandom"`
	MsgBody         []MsgItem           `json:"MsgBody"`
	OfflinePushInfo *OfflinePushInfoObj `json:"OfflinePushInfo"`
	Condition       *Condition          `json:"Condition,omitempty"`
}

// 请求结果
type PushResp struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
	TaskID       string `json:"TaskId"`
}
