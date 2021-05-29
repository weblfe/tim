/*
 *
 * server_list.go
 * server
 *
 * Created by lintao on 2020/4/16 2:24 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package tim

type ApiPath string

const (

	VERSION = "v4"
	BaseUrl = "https://console.tim.qq.com"

	//帐号管理
	MultiAccountImportApi ApiPath = "/im_open_login_svc/multiaccount_import" //导入多个帐号
	AccountImportApi      ApiPath = "/im_open_login_svc/account_import"      //导入单个帐号
	AccountDeleteApi      ApiPath = "/im_open_login_svc/account_delete"      //删除帐号	v4/im_open_login_svc/account_delete
	AccountCheckApi       ApiPath = "/im_open_login_svc/account_check"       //查询帐号	v4/im_open_login_svc/account_check
	KickApi               ApiPath = "/im_open_login_svc/kick"                //失效帐号登录态	v4/im_open_login_svc/kick

	//单聊消息
	QueryStateApi    ApiPath = "/openim/querystate"        //查询帐号在线状态	v4/openim/querystate
	SendMsg          ApiPath = "/openim/sendmsg"           //单发单聊消息	v4/openim/sendmsg
	BatchSendMsg     ApiPath = "/openim/batchsendmsg"      //批量发单聊消息	v4/openim/batchsendmsg
	ImportMsg        ApiPath = "/openim/importmsg"         //导入单聊消息	v4/openim/importmsg
	AdminGetRoamMsg  ApiPath = "/openim/admin_getroammsg"  //查询单聊消息	v4/openim/admin_getroammsg
	AdminMsgWithDraw ApiPath = "/openim/admin_msgwithdraw" //撤回单聊消息
	//资料管理
	PortraitGet     ApiPath = "/profile/portrait_get" //拉取资料	v4/profile/portrait_get
	PortraitPostSet ApiPath = "/profile/portrait_set" //设置资料	v4/profile/portrait_set

	//关系链管理
	FriendAddApi       ApiPath = "/sns/friend_add"        //添加好友	v4/sns/friend_add
	FriendImport       ApiPath = "/sns/friend_import"     //导入好友	v4/sns/friend_import
	FriendDeleteApi    ApiPath = "/sns/friend_delete"     //删除好友	v4/sns/friend_delete
	FriendUpdate       ApiPath = "/sns/friend_update"     //更新好友
	FriendDeleteAll    ApiPath = "/sns/friend_delete_all" //删除所有好友	v4/sns/friend_delete_all
	FriendCheck        ApiPath = "/sns/friend_check"      //校验好友	v4/sns/friend_check
	FriendGetApi       ApiPath = "/sns/friend_get"        //拉取好友	v4/sns/friend_get
	FriendGetList      ApiPath = "/sns/friend_get_list"   //拉取指定好友	v4/sns/friend_get_list
	BlackListAddApi    ApiPath = "/sns/black_list_add"    //添加黑名单	v4/sns/black_list_add
	BlackListDeleteApi ApiPath = "/sns/black_list_delete" //删除黑名单	v4/sns/black_list_delete
	BlackListGetApi    ApiPath = "/sns/black_list_get"    //拉取黑名单	v4/sns/black_list_get
	BlackListCheckApi  ApiPath = "/sns/black_list_check"  //校验黑名单	v4/sns/black_list_check
	GroupAddApi        ApiPath = "/sns/group_add"         //添加分组	v4/sns/group_add
	GroupDeleteApi     ApiPath = "/sns/group_delete"      //删除分组	v4/sns/group_delete

	//群组管理

	//获取 App 中的所有群组	v4/group_open_http_svc/get_appid_group_list
	GetAppidGroupListApi        ApiPath = "/group_open_http_svc/get_appid_group_list"
	CreateGroupApi              ApiPath = "/group_open_http_svc/create_group"                   //创建群组	v4/group_open_http_svc/create_group
	GetGroupInfoApi             ApiPath = "/group_open_http_svc/get_group_info"                 //获取群组详细资料	v4/group_open_http_svc/get_group_info
	GetGroupMemberInfoApi       ApiPath = "/group_open_http_svc/get_group_member_info"          //获取群成员详细资料	v4/group_open_http_svc/get_group_member_info
	ModifyGroupBaseInfoApi      ApiPath = "/group_open_http_svc/modify_group_base_info"         //修改群组基础资料	v4/group_open_http_svc/modify_group_base_info
	AddGroupMemberApi           ApiPath = "/group_open_http_svc/add_group_member"               //增加群组成员	v4/group_open_http_svc/add_group_member
	DeleteGroupMemberApi        ApiPath = "/group_open_http_svc/delete_group_member"            //删除群组成员	v4/group_open_http_svc/delete_group_member
	ModifyGroupMemberInfoApi    ApiPath = "/group_open_http_svc/modify_group_member_info"       //修改群组成员资料	v4/group_open_http_svc/modify_group_member_info
	DestroyGroupApi             ApiPath = "/group_open_http_svc/destroy_group"                  //解散群组	v4/group_open_http_svc/destroy_group
	GetJoinedGroupListApi       ApiPath = "/group_open_http_svc/get_joined_group_list"          //获取用户所加入的群组	v4/group_open_http_svc/get_joined_group_list
	GetRoleInGroup              ApiPath = "/group_open_http_svc/get_role_in_group"              //查询用户在群组中的身份	v4/group_open_http_svc/get_role_in_group
	ForbidSendMsg               ApiPath = "/group_open_http_svc/forbid_send_msg"                //批量禁言和取消禁言	v4/group_open_http_svc/forbid_send_msg
	GetGroupShuttedUin          ApiPath = "/group_open_http_svc/get_group_shutted_uin"          //获取群组被禁言用户列表	v4/group_open_http_svc/get_group_shutted_uin
	SendGroupMsg                ApiPath = "/group_open_http_svc/send_group_msg"                 //在群组中发送普通消息	v4/group_open_http_svc/send_group_msg
	SendGroupSystemNotification ApiPath = "/group_open_http_svc/send_group_system_notification" //在群组中发送系统通知	v4/group_open_http_svc/send_group_system_notification
	GroupMsgRecall              ApiPath = "/group_open_http_svc/group_msg_recall"               //群组消息撤回	v4/group_open_http_svc/group_msg_recall
	ChangeGroupOwner            ApiPath = "/group_open_http_svc/change_group_owner"             //转让群组	v4/group_open_http_svc/change_group_owner
	ImportGroupApi              ApiPath = "/group_open_http_svc/import_group"                   //导入群基础资料	v4/group_open_http_svc/import_group
	ImportGroupMsgApi           ApiPath = "/group_open_http_svc/import_group_msg"               //导入群消息	v4/group_open_http_svc/import_group_msg
	ImportGroupMemberApi        ApiPath = "/group_open_http_svc/import_group_member"            //导入群成员	v4/group_open_http_svc/import_group_member
	SetUnreadMsgNumApi          ApiPath = "/group_open_http_svc/set_unread_msg_num"             //设置成员未读消息计数	v4/group_open_http_svc/set_unread_msg_num
	DeleteGroupMsgBySenderApi   ApiPath = "/group_open_http_svc/delete_group_msg_by_sender"     //删除指定用户发送的消息	v4/group_open_http_svc/delete_group_msg_by_sender
	GroupMsgGetSimpleApi        ApiPath = "/group_open_http_svc/group_msg_get_simple"           //拉取群漫游消息	v4/group_open_http_svc/group_msg_get_simple

	//全局禁言管理
	SetNoSpeakingApi ApiPath = "/openconfigsvr/setnospeaking" //设置全局禁言	v4/openconfigsvr/setnospeaking
	GetNoSpeakingApi ApiPath = "/openconfigsvr/getnospeaking" //查询全局禁言	v4/openconfigsvr/getnospeaking

	//运营管理
	GetHistoryApi ApiPath = "/open_msg_svc/get_history" //下载消息记录	v4/open_msg_svc/get_history
	GetAppInfoApi ApiPath = "/openconfigsvr/getappinfo" //拉取运营数据	v4/openconfigsvr/getappinfo
)
