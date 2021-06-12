package tim

const (
	//TimTextElem	文本消息。
	TextElem = "TimTextElem"
	//LocationElem	地理位置消息。
	LocationElem = "TIMLocationElem"
	//FaceElem	表情消息。
	FaceElem = "TIMFaceElem"
	//CustomElem	自定义消息，当接收方为 iOS 系统且应用处在后台时，此消息类型可携带除文本以外的字段到 APNs。一条组合消息中只能包含一个 CustomElem 自定义消息元素
	CustomElem = "TIMCustomElem"
	// 语音消息
	SoundElem = "TIMSoundElem"
	// 图像消息
	ImageElem = "TIMImageElem"
	// 文件消息
	FileElem = "TIMFileElem"
	// 视频消息
	VideoFileElem = "TIMVideoFileElem"
)

// 声音消息
func NewSoundElem(sound SoundContent) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			SoundContent: sound,
		},
		MsgType: SoundElem,
	}
}

// 图片消息
func NewImageElem(images MsgImageContent) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			MsgImageContent: images,
		},
		MsgType: ImageElem,
	}
}

// 文件消息
func NewFileElem(file FileContent) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			FileContent: file,
		},
		MsgType: FileElem,
	}
}

// 视频消息
func NewVideoElem(video VideoContent) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			VideoContent: video,
		},
		MsgType: VideoFileElem,
	}
}

// 文本消息元素
// Text	String	消息内容。当接收方为 iOS 或 Android 后台在线时，作为离线推送的文本展示。
func NewTextElem(text string) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			TextContent: TextContent{Text: text},
		},
		MsgType: TextElem,
	}
}

// 地理位置消息元素
// Desc	String	地理位置描述信息。
// Latitude	Number	纬度。
// Longitude	Number	经度。
func NewLocationElem(desc string, latitude, longitude float64) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			LocationContent: LocationContent{
				Latitude:  latitude,
				Longitude: longitude,
			},
			Desc: desc,
		},
		MsgType: LocationElem,
	}
}

// 自定义消息元素
// Index	Number	表情索引，用户自定义。
// Data	String	额外数据。
func NewFaceElem(index int, data string) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			Data: data,
			FaceContent: FaceContent{

				Index: index,
			},
		},
		MsgType: FaceElem,
	}
}

// 自定义消息元素
// Data	String	自定义消息数据。 不作为 APNs 的 payload 字段下发，故从 payload 中无法获取 Data 字段。
// Desc	String	自定义消息描述信息。当接收方为 iOS 或 Android 后台在线时，做离线推送文本展示。
// 若发送自定义消息的同时设置了 OfflinePushInfo.Desc 字段，此字段会被覆盖，请优先填 OfflinePushInfo.Desc 字段。
// 当消息中只有一个 CustomElem 自定义消息元素时，如果 Desc 字段和 OfflinePushInfo.Desc 字段都不填写，将收不到该条消息的离线推送，需要填写 OfflinePushInfo.Desc 字段才能收到该消息的离线推送。
// Ext	String	扩展字段。当接收方为 iOS 系统且应用处在后台时，此字段作为 APNs 请求包 Payloads 中的 Ext 键值下发，Ext 的协议格式由业务方确定，APNs 只做透传。
// Sound	String	自定义 APNs 推送铃音。
func NewCustomElem(data, desc, ext, sound string) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			Data: data,
			Desc: desc,
			CustomContent: CustomContent{
				Ext:   ext,
				Sound: sound,
			},
		},
		MsgType: CustomElem,
	}
}

//From_Account	String	选填	消息发送方 UserID（用于指定发送消息方帐号）
//To_Account	String	必填	消息接收方 UserID
type SingleChatMsg struct {
	FromAccount string `json:"From_Account"`
	ToAccount   string `json:"To_Account"`
	ChatMsg
}

type BatchChatMsg struct {
	FromAccount string   `json:"From_Account"`
	ToAccount   []string `json:"To_Account"`
	ChatMsg
}

//SyncOtherMachine	Integer
//选填	1：把消息同步到 From_Account 在线终端和漫游上；
//      2：消息不同步至 From_Account；
//若不填写默认情况下会将消息存 From_Account 漫游
//MsgLifeTime	Integer	选填	消息离线保存时长（单位：秒），最长为7天（604800秒）
//若设置该字段为0，则消息只发在线用户，不保存离线
//若设置该字段超过7天（604800秒），仍只保存7天
//若不设置该字段，则默认保存7天
//MsgRandom	Integer	必填	消息随机数，由随机函数产生，用于后台定位问题
//MsgTimeStamp	Integer	选填	消息时间戳，UNIX 时间戳（单位：秒）
//MsgType	String	必填	TIM 消息对象类型，目前支持的消息对象包括：TimTextElem(文本消息)，FaceElem(表情消息)，LocationElem(位置消息)，CustomElem(自定义消息)
//MsgContent	Object	必填	对于每种 MsgType 用不同的 MsgContent 格式，具体可参考 消息格式描述
//OfflinePushInfo	Object	选填	离线推送信息配置，具体可参考 消息格式描述
type ChatMsg struct {
	SyncOtherMachine  int              `json:"SyncOtherMachine,omitempty"`
	SyncFromOldSystem int              `json:"SyncFromOldSystem,omitempty" `
	MsgLifeTime       int              `json:"MsgLifeTime,omitempty"`
	MsgRandom         int              `json:"MsgRandom"`
	MsgTimeStamp      int64            `json:"MsgTimeStamp,omitempty"`
	MsgBody           []MsgBody        `json:"MsgBody"`
	OfflinePushInfo   *OfflinePushInfo `json:"OfflinePushInfo,omitempty"`
}

//https://cloud.tencent.com/document/product/269/2720#.E7.A6.BB.E7.BA.BF.E6.8E.A8.E9.80.81-offlinepushinfo-.E8.AF.B4.E6.98.8E
//PushFlag	Integer	选填	0表示推送，1表示不离线推送。
//Title	String	选填	离线推送标题。该字段为 iOS 和 Android 共用。
//Desc	String	选填	离线推送内容。该字段会覆盖上面各种消息元素 TIMMsgElement 的离线推送展示文本。
//若发送的消息只有一个 CustomElem 自定义消息元素，该 Desc 字段会覆盖 CustomElem 中的 Desc 字段。如果两个 Desc 字段都不填，将收不到该自定义消息的离线推送。
//Ext	String	选填	离线推送透传内容。
//AndroidInfo.Sound	String	选填	Android 离线推送声音文件路径。
//AndroidInfo.OPPOChannelID	String	选填	OPPO 手机 Android 8.0 以上的 NotificationChannel 通知适配字段。
//ApnsInfo.BadgeMode	Integer	选填	这个字段缺省或者为0表示需要计数，为1表示本条消息不需要计数，即右上角图标数字不增加。
//ApnsInfo.Title	String	选填	该字段用于标识 APNs 推送的标题，若填写则会覆盖最上层 Title。
//ApnsInfo.SubTitle	String	选填	该字段用于标识 APNs 推送的子标题。
//ApnsInfo.Image	String	选填	该字段用于标识 APNs 携带的图片地址，当客户端拿到该字段时，可以通过下载图片资源的方式将图片展示在弹窗上。
type OfflinePushInfo struct {
	PushFlag    int          `json:"PushFlag"`
	Desc        string       `json:"Desc"`
	Ext         string       `json:"Ext"`
	AndroidInfo *AndroidInfo `json:"AndroidInfo,omitempty"`
	ApnsInfo    *ApnsInfo    `json:"ApnsInfo,omitempty"`
}

type ErrorList struct {
	ToAccount string `json:"To_Account"`
	ErrorCode int    `json:"ErrorCode"`
}

type AndroidInfo struct {
	Sound string `json:"Sound"`
}

type ApnsInfo struct {
	Sound     string `json:"Sound"`
	BadgeMode int    `json:"BadgeMode"`
	Title     string `json:"Title"`
	SubTitle  string `json:"SubTitle"`
	Image     string `json:"Image"`
}

type RoamMsgReq struct {
	FromAccount string `json:"From_Account"`
	ToAccount   string `json:"To_Account"`
	MaxCnt      int    `json:"MaxCnt"`
	MinTime     int    `json:"MinTime"`
	MaxTime     int    `json:"MaxTime"`
	LastMsgKey  string `json:"LastMsgKey"`
}

type RoamMsg struct {
	FromAccount  string            `json:"From_Account"`
	ToAccount    string            `json:"To_Account"`
	MsgSeq       int               `json:"MsgSeq"`
	MsgRandom    int               `json:"MsgRandom"`
	MsgTimeStamp int               `json:"MsgTimeStamp"`
	MsgFlagBits  int               `json:"MsgFlagBits"`
	MsgKey       string            `json:"MsgKey"`
	MsgBody      []RoamMsgBodyItem `json:"MsgBody"`
}

type RoamMsgBodyItem struct {
	MsgType    string              `json:"MsgType"`
	MsgContent *RoamMsgBodyContent `json:"MsgContent,omitempty"`
}

type RoamMsgBodyContent struct {
	Text string `json:"Text"`
}

type RoamMsgResp struct {
	ActionStatus string    `json:"ActionStatus"`
	ErrorInfo    string    `json:"ErrorInfo"`
	ErrorCode    int       `json:"ErrorCode"`
	Complete     int       `json:"Complete"`
	MsgCnt       int       `json:"MsgCnt"`
	LastMsgTime  int       `json:"LastMsgTime"`
	LastMsgKey   string    `json:"LastMsgKey"`
	MsgList      []RoamMsg `json:"MsgList"`
}

type MsgWithdraw struct {
	FromAccount string `json:"From_Account"`
	ToAccount   string `json:"To_Account"`
	MsgKey      string `json:"MsgKey"`
}

type MsgBody struct {
	MsgContent MsgContent `json:"MsgContent"`
	MsgType    string     `json:"MsgType"`
}

// 消息内容体嵌套
type MsgContent struct {
	// 自定义消息 Data + Desc + Ext + Sound
	CustomContent
	// 表情消息 Index + Data
	FaceContent
	// 定位消息
	LocationContent
	// 声音消息 Size + Second + DownloadFlag
	SoundContent
	// 文本消息
	TextContent
	// 图片消息
	MsgImageContent
	// 文件消息
	FileContent
	// 视频消息
	VideoContent
	// 公共字段
	Data         string `json:"Data,omitempty"`
	Desc         string `json:"Desc,omitempty"`
	Size         int    `json:"Size,omitempty"`
	DownloadFlag int    `json:"Download_Flag,omitempty"`
}

// 自定义消息
type CustomContent struct {
	// Data  string `json:"Data,omitempty"`
	// Desc  string `json:"Desc,omitempty"`
	Ext   string `json:"Ext,omitempty"`
	Sound string `json:"Sound,omitempty"`
}

// 表情消息
type FaceContent struct {
	// Data  string `json:"Data,omitempty"`
	Index int `json:"Index,omitempty"`
}

// 文本消息
type TextContent struct {
	Text string `json:"Text,omitempty"`
}

// 定位消息
type LocationContent struct {
	Latitude  float64 `json:"Latitude,omitempty"`
	Longitude float64 `json:"Longitude,omitempty"`
}

// 视频消息
type VideoContent struct {
	VideoUrl          string `json:"VideoUrl,omitempty"`          // 视频下载地址。可通过该 URL 地址直接下载相应视频。
	VideoSize         int    `json:"VideoSize,omitempty"`         // 视频数据大小，单位：字节。
	VideoSecond       int    `json:"VideoSecond,omitempty"`       // 视频时长，单位：秒。
	VideoFormat       string `json:"VideoFormat,omitempty"`       // 视频格式，例如 mp4。
	VideoDownloadFlag int    `json:"VideoDownloadFlag,omitempty"` // 	视频下载方式标记。目前 VideoDownloadFlag 取值只能为2，表示可通过VideoUrl字段值的 URL 地址直接下载视频。
	ThumbUrl          string `json:"ThumbUrl,omitempty"`          // 视频缩略图下载地址。可通过该 URL 地址直接下载相应视频缩略图。
	ThumbSize         int    `json:"ThumbSize,omitempty"`         // 缩略图大小，单位：字节。
	ThumbWidth        int    `json:"ThumbWidth,omitempty"`        // 缩略图宽度。
	ThumbHeight       int    `json:"ThumbHeight,omitempty"`       // 缩略图高度。
	ThumbFormat       string `json:"ThumbFormat,omitempty"`       // 缩略图格式，例如 JPG、BMP 等。
	ThumbDownloadFlag int    `json:"ThumbDownloadFlag,omitempty"` //	视频缩略图下载方式标记。目前 ThumbDownloadFlag 取值只能为2，表示可通过ThumbUrl字段值的 URL 地址直接下载视频缩略图。
}

// 声音消息
type SoundContent struct {
	// Size         int `json:"Size,omitempty"`
	Second int `json:"Second,omitempty"`
	// DownloadFlag int `json:"Download_Flag,omitempty"`
}

// 文件
type FileContent struct {
	Url      string `json:"Url,omitempty"`      // 文件下载地址，可通过该 URL 地址直接下载相应文件。
	FileSize int    `json:"FileSize,omitempty"` // 文件数据大小，单位：字节。
	FileName string `json:"FileName,omitempty"` // 文件名称。
	// DownloadFlag uint   `json:"Download_Flag,omitempty"` // 文件下载方式标记。目前 Download_Flag 取值只能为2，表示可通过Url字段值的 URL 地址直接下载文件。
}

// 图片消息
type MsgImageContent struct {
	UUID           string          `json:"UUID,omitempty"`           // 图片序列号。后台用于索引图片的键值。
	ImageFormat    uint64          `json:"ImageFormat,omitempty"`    // 图片格式。JPG = 1，GIF = 2，PNG = 3，BMP = 4，其他 = 255。
	ImageInfoArray []ImageInfoItem `json:"ImageInfoArray,omitempty"` //	原图、缩略图或者大图下载信息。
}

// 图片信息
type ImageInfoItem struct {
	Type   uint   `json:"Type"`   //	图片类型： 1-原图，2-大图，3-缩略图。
	Size   uint   `json:"Size"`   //   图片数据大小，单位：字节。
	Width  int    `json:"Width"`  //   图片宽度。
	Height int    `json:"Height"` //	图片高度。
	URL    string `json:"Url"`    //   图片下载地址。
}
