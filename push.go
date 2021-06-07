package tim

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"time"
)

// https://cloud.tencent.com/document/product/269/45934
// 支持全员推送。
//支持按用户属性推送。
//支持按用户标签推送。
//管理员推送消息，接收方看到消息发送者是管理员。
//管理员指定某一帐号向其他帐号推送消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
//支持消息离线存储，不支持漫游。
//由于全员推送需要下发的帐号数量巨大，下发完全部帐号需要一定时间（根据帐号总数而定，一般在一分钟内）。
// Condition	Object	选填	Condition 共有4种条件类型，分别是：
// 属性的或条件 AttrsOr
// 属性的与条件 AttrsAnd
// 标签的或条件 TagsOr
// 标签的与条件 TagsAnd
//AttrsOr 和 AttrsAnd 可以并存，TagsOr 和 TagsAnd 也可以并存。但是标签和属性条件不能并存。如果没有 Condition，则推送给全部用户
//MsgRandom	Integer	必填	消息随机数，由随机函数产生。用于推送任务去重。对于不同的推送请求，MsgRandom7 天之内不能重复，否则视为相同的推送任务（调用推送 API 返回失败的时候可以用相同的 MsgRandom 进行重试）
//MsgBody	Object	必填	消息内容，具体格式请参考 MsgBody 消息内容说明（一条消息可包括多种消息元素，所以 MsgBody 为 Array 类型）
//MsgType	String	必填	TIM 消息对象类型，目前支持的消息对象包括：
//TIMTextElem（文本消息）
//TIMLocationElem（位置消息）
//TIMFaceElem（表情消息）
//TIMCustomElem（自定义消息）
//TIMSoundElem（语音消息）
//TIMImageElem（图像消息）
//TIMFileElem（文件消息）
//TIMVideoFileElem（视频消息）
//MsgContent	Object	必填	对于每种 MsgType，用不同的 MsgContent 格式，具体可参考 TIMMsgElement 对象 的定义
//MsgLifeTime	Integer	选填	消息离线存储时间，单位秒，最多保存7天（604800秒）。默认为0，表示不离线存储
//From_Account	String	选填	消息推送方帐号
//AttrsOr	Object	选填	属性条件的并集。注意属性推送和标签推送不可同时作为推送条件
//AttrsAnd	Object	选填	属性条件的交集。注意属性推送和标签推送不可同时作为推送条件
//TagsOr	Object	选填	标签条件的并集。标签是一个不超过50字节的字符串。注意属性推送和标签推送不可同时作为推送条件。TagsOr 条件中的标签个数不能超过10个
//TagsAnd	Object	选填	标签条件的交集。标签是一个不超过50字节的字符串。注意属性推送和标签推送不可同时作为推送条件。TagsAnd 条件中的标签个数不能超过10个
//OfflinePushInfo	Object	选填	离线推送信息配置，具体可参考 消息格式描述
func (s IMServer) PostPushAll(m PushMsg) (*PushResp, error) {
	var v, err = s.requestWithPath(PostPushAll, m)
	if err != nil {
		return nil, err
	}
	var resp = &PushResp{}
	err = json.Unmarshal(v, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 创建推送消息
func CreatePushMsg(msgType string, data interface{}) PushMsg {
	var n, _ = rand.Int(rand.Reader, big.NewInt(time.Now().Unix()))
	return PushMsg{
		MsgRandom: int(n.Int64()),
		MsgBody: []MsgItem{
			{MsgType: msgType, MsgContent: data},
		},
	}
}
