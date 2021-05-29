package tim

import (
	"encoding/json"
	"fmt"
)

//ActionStatus	String	请求的处理结果，OK 表示处理成功，FAIL 表示失败
//ErrorCode	Integer	错误码，0表示成功，非0表示失败
//ErrorInfo	String	请求处理失败时的错误信息
//ResultItem	Array	单个帐号的结果对象数组
//ResultCode	Integer	单个帐号的错误码，0表示成功，非0表示失败
//ResultInfo	String	单个帐号删除失败时的错误描述信息
//UserID	String	请求删除的帐号的 UserID
type ResultItem struct {
	ResultCode    int      `json:"ResultCode"`
	ResultInfo    string   `json:"ResultInfo"`
	UserID        string   `json:"UserID"`
	AccountStatus string   `json:"AccountStatus"`
	ToAccount     string   `json:"To_Account" `
	State         string   `json:"State" `
	Detail        []Detail `json:"Detail"`
}

type Detail struct {
	Platform string `json:"Platform"`
	Status   string `json:"Status"`
}

// 导入单个帐号
// 本接口适用于将 App 自有帐号导入即时通信 IM，
// 即在即时通信 IM 中为 App 自有帐号创建一个对应的内部
// ID，使未登录过即时通信 IM 的 App 自有帐号能够使用即时通信 IM 服务。
// 例如，App 开发者通过 REST API 给用户 A 发送一条消息，用户 A 如果没有登录过即时通信 IM 服务，
// 由于腾讯内部没有用户 A 对应的内部 ID，那么给用户 A 发送消息将会失败。
// 需要把用户 A 的帐号导入即时通信 IM，系统为用户 A 创建一个内部 ID，才能通过 REST API 给用户 A 发送消息。
//
// 注意：
// 同一个帐号重复导入仅会创建1个内部 ID。
// identifier	String	必填	用户名，长度不超过32字节
// nick			String	选填	用户昵称
// faceUrl		String	选填	用户头像 URL
//
func (s IMServer) AccountImport(identifier, nick, faceUrl string) error {

	if _, err := s.requestWithPath(AccountImportApi, []byte(`{"Identifier":"`+identifier+`",
								"Nick":"`+nick+`",
								"FaceUrl":"`+faceUrl+`"}`)); err != nil {
		return err
	}

	return nil

}

//
// 本接口单次最多支持导入100个帐号，同一个帐号重复导入仅会创建1个内部 ID。
// Accounts		Array	必填	用户名，单个用户名长度不超过32字节，单次最多导入100个用户名
// FailAccounts	Array	导入失败的帐号列表
//
func (s IMServer) MultiAccountImports(account []string) ([]string, error) {
	if len(account) > 100 {
		return nil, fmt.Errorf("导入账号最多支持100个，当前导入%d ", len(account))
	}
	var req struct {
		Accounts []string `json:"Accounts" `
	}

	req.Accounts = account
	r, err := s.requestWithPath(MultiAccountImportApi, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		FailAccounts []string `json:"FailAccounts"`
	}

	if err = json.Unmarshal(r, &resp); err != nil {
		return nil, err
	}

	return resp.FailAccounts, err
}

//
// 仅支持删除体验版帐号。
// 帐号删除时，该用户的关系链、资料等数据也会被删除。
// 帐号删除后，该用户的数据将无法恢复，请谨慎使用该接口。
// accounts	Array	必填	请求删除的帐号的 UserID，单次请求最多支持100个帐号
// DeleteResultItem	Array	单个帐号的结果对象数组
// ResultCode		Integer	单个帐号的错误码，0表示成功，非0表示失败
// ResultInfo		String	单个帐号删除失败时的错误描述信息
// UserID			String	请求删除的帐号的 UserID
//
func (s IMServer) DeleteAccount(accounts []string) ([]ResultItem, error) {
	if len(accounts) > 100 {
		return nil, fmt.Errorf("最多支持删除100个账号 ，当前请求数量%d", len(accounts))
	}
	var req struct {
		Accounts []map[string]string `json:"DeleteItem" `
	}

	for _, ac := range accounts {
		req.Accounts = append(req.Accounts, map[string]string{"UserID": ac})
	}

	var resp struct {
		ResultItem []ResultItem `json:"ResultItem"`
	}

	r, err := s.requestWithPath(AccountDeleteApi, req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(r, &resp); err != nil {
		return nil, err
	}

	return resp.ResultItem, nil
}

//
// 查询帐号 https://cloud.tencent.com/document/product/269/38417
// 用于查询自有帐号是否已导入即时通信 IM，支持批量查询
// accounts 请求检查的帐号对象数组，单次请求最多支持100个帐号
//
func (s IMServer) AccountCheck(accounts []string) ([]ResultItem, error) {
	if len(accounts) > 100 {
		return nil, fmt.Errorf("最多支持查询100个账号 ，当前请求数量%d", len(accounts))
	}
	var req struct {
		Accounts []map[string]string `json:"CheckItem" `
	}

	for _, ac := range accounts {
		req.Accounts = append(req.Accounts, map[string]string{"UserID": ac})
	}

	r, err := s.requestWithPath(AccountCheckApi, req)
	if err != nil {
		return nil, err
	}
	var AutoGenerated struct {
		ResultItem []ResultItem `json:"ResultItem"`
	}
	if err = json.Unmarshal(r, &AutoGenerated); err != nil {
		return nil, err
	}

	return AutoGenerated.ResultItem, nil
}

// https://cloud.tencent.com/document/product/269/3853
// 失效帐号登录态
// 本接口适用于将 App 用户帐号的登录态（例如 UserSig）失效。
// 例如，开发者判断一个用户为恶意帐号后，可以调用本接口将该用户当前的登录态失效，这样用户使用历史 UserSig 登录即时通信 IM 会失败
// 使用该接口将用户登录态失效后，用户如果使用重新生成的 UserSig 可以成功登录即时通信 IM，接口支持一次失效一个帐号。
// userId	String	必填	用户名
func (s IMServer) Kick(userId string) error {
	_, err := s.requestWithPath(KickApi, []byte(`{"Identifier":"`+userId+`"}`))
	return err
}

// https://cloud.tencent.com/document/product/269/2566
// 查询帐号在线状态
// 获取用户当前的登录状态
func (s IMServer) QueryState(accounts []string, isNeedDetail int) ([]ResultItem, error) {

	if len(accounts) <= 0 {
		return nil, fmt.Errorf("用户列表不能为空")
	}

	if len(accounts) > 100 {
		return nil, fmt.Errorf("查询数量不能超过100")
	}

	var req struct {
		Accounts     []string `json:"To_Account" `
		IsNeedDetail int      `json:"IsNeedDetail" `
	}
	req.Accounts = accounts
	req.IsNeedDetail = isNeedDetail

	r, err := s.requestWithPath(QueryStateApi, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		QueryResult []ResultItem `json:"QueryResult"`
	}

	if err = json.Unmarshal(r, &resp); err != nil {
		return nil, err
	}

	return resp.QueryResult, nil
}
