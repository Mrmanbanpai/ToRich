package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// header
	ErrAppEUINotFound    = &Errno{Code: 10003, Message: "The request header LOLA-ACCESS-APPID not found."}
	ErrNonceNotFound     = &Errno{Code: 10004, Message: "The request header LOLA-ACCESS-NONCE not found."}
	ErrSignatureNotFound = &Errno{Code: 10005, Message: "The request header LOLA-ACCESS-SIGNATURE not found."}
	ErrJwtPayload        = &Errno{Code: 10006, Message: "JWT payload decode error."}
	ErrDataVerify        = &Errno{Code: 10007, Message: "数据完整性校验未通过"}
	ErrNonceTypeWrong    = &Errno{Code: 10008, Message: "LOLA-ACCESS-NONCE 数据类型错误"}
	ErrNonceTimeOut      = &Errno{Code: 10009, Message: "LOLA-ACCESS-NONCE 超时"}

	// system
	ErrValidation    = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase      = &Errno{Code: 20002, Message: "Database error."}
	ErrToken         = &Errno{Code: 20003, Message: "令牌验证失败"}
	ErrReply         = &Errno{Code: 20004, Message: "消息队列回复失败"}
	ErrSmsSend       = &Errno{Code: 20005, Message: "短信发送失败"}
	ErrTokenExpire   = &Errno{Code: 20006, Message: "令牌过期失效"}
	ErrResponseParse = &Errno{Code: 20007, Message: "Response解析失败"}
	ErrJsonParse     = &Errno{Code: 20008, Message: "Json解析失败"}
	ErrProtoParse    = &Errno{Code: 20009, Message: "Proto解析失败"}
	ErrOtherLogin    = &Errno{Code: 20010, Message: "该账户已被他人登录"}

	// user errors
	ErrEncrypt                        = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound                   = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid                   = &Errno{Code: 20103, Message: "该token不具有操作此设备的权限."}
	ErrPwdWrong                       = &Errno{Code: 20104, Message: "密码输入错误"}
	ErrPwdDisaccord                   = &Errno{Code: 20105, Message: "两次密码不一致"}
	ErrWeChatExisted                  = &Errno{Code: 20106, Message: "用户已绑定微信，不可重复绑定"}
	ErrPhone                          = &Errno{Code: 20107, Message: "手机号码错误"}
	ErrUserIsExisted                  = &Errno{Code: 20108, Message: "用户已经存在"}
	ErrQQExisted                      = &Errno{Code: 20109, Message: "用户已绑定QQ，不可重复绑定"}
	ErrInviteCode                     = &Errno{Code: 20110, Message: "邀请码无效"}
	ErrUserDisable                    = &Errno{Code: 20111, Message: "该账号已冻结"}
	ErrUserAppScoreFirstTimeNoEnough  = &Errno{Code: 20112, Message: "用户首次给APP评分不足3天"}
	ErrUserAppScoreSecondTimeNoEnough = &Errno{Code: 20113, Message: "用户第二次给APP评分不足4天"}
	ErrUserAppScoreFinished           = &Errno{Code: 20114, Message: "用户已完成对APP评分,无需继续提示"}
	ErrUserAppScoreAlreadyTwoTimes    = &Errno{Code: 20115, Message: "用户已有2次对APP评分操作,无需继续提示"}
	ErrUserAppScoreConfirmParam       = &Errno{Code: 20116, Message: "用户对APP评分确认,Operate参数错误"}
	ErrCreatePipeDataFail             = &Errno{Code: 20117, Message: "创建pipe数据失败"}

	// parameters
	ErrMustParaIsNull         = &Errno{Code: 20201, Message: "必填参数不可为空"}
	ErrDevEuiOverrideControl  = &Errno{Code: 20202, Message: "DevEUI不在应用权限内"}
	ErrParaTypeWrong          = &Errno{Code: 20203, Message: "参数类型错误"}
	ErrDevInfoNone            = &Errno{Code: 20204, Message: "该设备信息不存在"}
	ErrFportCrossBorder       = &Errno{Code: 20205, Message: "fPort字段超出范围"}
	ErrAppIdNotFound          = &Errno{Code: 20206, Message: "没有找到对应的AppId,请确保LOLA-ACCESS-APPID数据正确"}
	ErrSmsCode                = &Errno{Code: 20207, Message: "短信验证码错误"}
	ErrParas                  = &Errno{Code: 20208, Message: "参数验证失败"}
	ErrQQIsExisted            = &Errno{Code: 20209, Message: "QQ已被使用,同一个QQ只能绑定一个账号"}
	ErrWeChatIsExisted        = &Errno{Code: 20210, Message: "微信已被使用,同一个微信只能绑定一个账号"}
	ErrLike                   = &Errno{Code: 20211, Message: "使用心动失败"}
	ErrWeCharPay              = &Errno{Code: 20212, Message: "微信支付错误"}
	ErrAICashPay              = &Errno{Code: 20213, Message: "AI购买商品错误"}
	ErrIdolStatusCheck        = &Errno{Code: 20214, Message: "public无可领可买状态"}
	ErrIdolStatusDisabled     = &Errno{Code: 20215, Message: "idol未开通"}
	ErrIdolAlreadyHad         = &Errno{Code: 20216, Message: "用户已有该idol不可重复领取或购买"}
	ErrGiftStatusDisabled     = &Errno{Code: 20217, Message: "gift状态disabled"}
	ErrCrowdStatusDisabled    = &Errno{Code: 20218, Message: "crowdfunding状态disabled"}
	ErrApplePayDealWith       = &Errno{Code: 20219, Message: "ApplePay支付处理错误"}
	ErrApplePayDuplicate      = &Errno{Code: 20220, Message: "ApplePay支付结果已处理"}
	ErrApplePayResultVerify   = &Errno{Code: 20221, Message: "ApplePay支付结果验证失败"}
	ErrSetFirstPageFail       = &Errno{Code: 20222, Message: "SetStatusEnabled fail "}
	ErrGetIdolInformationFail = &Errno{Code: 20223, Message: "GetIdolInformation fail"}
	ErrGetUserInformation     = &Errno{Code: 20224, Message: "GetUserInformation"}
	ErrUpdateIdolName         = &Errno{Code: 20225, Message: "update Idol's NickName fail "}
	ErrLotteryLucky           = &Errno{Code: 20226, Message: "抽奖错误"}
	ErrLotteryLuckyList       = &Errno{Code: 20227, Message: "获取奖品列表失败"}
	ErrLotteryUserPrizes      = &Errno{Code: 20228, Message: "抽奖,获取用户中奖列表错误"}
	ErrLotteryMyPrizes        = &Errno{Code: 20229, Message: "抽奖,获取我的中奖列表错误"}
	ErrProductOrderExport     = &Errno{Code: 20230, Message: "导出实物商品订单列表到Excel错误"}
	ErrGetByObjectId          = &Errno{Code: 20231, Message: "获取奖品数据失败"}
	ErrStatusWrong            = &Errno{Code: 20232, Message: "数据状态不是Pending"}
	ErrUpdateFail             = &Errno{Code: 20233, Message: "更新奖品数据失败"}

	// base api errors
	ErrBaseApi        = &Errno{Code: 20301, Message: "基础API错误"}
	ErrNoticeNotFound = &Errno{Code: 20302, Message: "未找到通知信息"}
	ErrOrderService   = &Errno{Code: 20401, Message: "调用OrderService失败"}
	ErrUserService    = &Errno{Code: 20402, Message: "调用UserService失败"}
	ErrIdolService    = &Errno{Code: 20403, Message: "调用IdolService失败"}
	ErrAppStore       = &Errno{Code: 20404, Message: "调用AppStore接口失败"}
	ErrNlpApi         = &Errno{Code: 20405, Message: "NLP接口调用失败"}

	ErrNotImplement = &Errno{Code: 30000, Message: "未实现的接口"}

	ErrRefreshTokenFail = &Errno{Code: 30001, Message: "刷新Token异常"}

	ErrMustNumberType  = &Errno{Code: 30201, Message: "传人的金额必须是数字类型"}
	ErrRedisConnFail   = &Errno{Code: 30202, Message: "链接redis失败!!!"}
	ErrGetUninidError  = &Errno{Code: 30203, Message: "获取唯一订单id异常"}
	ErrSaveRdsError    = &Errno{Code: 30205, Message: "存人Redis缓存map失败"}
	ErrGetRdsDataError = &Errno{Code: 30265, Message: "取Redis缓存数据失败"}

	ErrCheckOrderInfoFail  = &Errno{Code: 30206, Message: "查询订单接口失败"}
	ErrUpdateOrderInfoFail = &Errno{Code: 30207, Message: "更新订单接口失败"}

	ErrGeneralInternalFault = &Errno{Code: 30208, Message: "发生内部服务错误"}
	ErrResourceDuplicate    = &Errno{Code: 30209, Message: "创建的资源和已有的资源冲突"}
	ErrResourceUpdateFail   = &Errno{Code: 30211, Message: "更新该资源不成功"}
	ErrResourceNotFound     = &Errno{Code: 30212, Message: "服务端找不到相应资源"}
	ErrCreateOrderFail      = &Errno{Code: 30213, Message: "生成入库订单信息失败"}

	ErrInitPublicPlatform = &Errno{Code: 30216, Message: "初始化公众平台小程序支付工具失败"}
	ErrInitAppPayProgram  = &Errno{Code: 30217, Message: "初始化公众平台app支付程序失败"}
	ErrTypeMathMiss       = &Errno{Code: 30218, Message: "Type Match Missing"}
	ErrTypeAmountErr      = &Errno{Code: 30219, Message: "传人的金额类型错误"}

	ErrDataIsBothErr = &Errno{Code: 30221, Message: "传人的支付数据完整性校验错误"}

	ErrClickLikeErr = &Errno{Code: 30222, Message: "点赞写入缓存失败"}

	// share 相关
	ErrGetUserInfoByPhone    = &Errno{Code: 40102, Message: "获取用户信息失败"}
	ErrUserInvalidInvitation = &Errno{Code: 40103, Message: "邀请无效"}
	ErrUserMarshal           = &Errno{Code: 40104, Message: "json.Marshal Fail"}
	ErrSaveUserShare         = &Errno{Code: 40105, Message: "保存用户分享失败"}
	ErrSaveUserShareToRedis  = &Errno{Code: 40106, Message: "redis入库失败"}
	ErrUpdateUserInviteCode  = &Errno{Code: 40107, Message: "更新用户邀请码失败"}
	ErrGetUserInfo           = &Errno{Code: 40108, Message: "获取用户信息失败"}
	ErrAtoi                  = &Errno{Code: 40109, Message: "AtoIErr"}
	ErrGetIdolInfo           = &Errno{Code: 40110, Message: "获取idol信息失败"}
	ErrInsertImageFail       = &Errno{Code: 40111, Message: "保存用户上传url失败"}
	ErrGetShareData          = &Errno{Code: 40112, Message: "获取用户分享失败"}
	ErrCheckUserHasIdolFail  = &Errno{Code: 40113, Message: "获取用户是否拥有明星失败"}
	ErrTypeWrong             = &Errno{Code: 40114, Message: "Type Worng"}
	ErrSaveFile              = &Errno{Code: 40115, Message: "保存文件失败，数据可能为空"}
	ErrMgoGetReceive         = &Errno{Code: 40116, Message: "查询数据失败"}

	ErrGetUsersWhoLikeStar      = &Errno{Code: 50101, Message: "获取粉丝数数失败"}
	ErrInternal                 = &Errno{Code: 50102, Message: "internal err "}
	ErrGetUserLikes             = &Errno{Code: 50103, Message: "获取用户守护值错误"}
	ErrGetTopUsers              = &Errno{Code: 50104, Message: "获取排行榜失败"}
	ErrGetUserPosition          = &Errno{Code: 50105, Message: "获取用户排名失败"}
	ErrGetStarList              = &Errno{Code: 50106, Message: "获取明星排行榜失败"}
	ErrGetStarLikes             = &Errno{Code: 50107, Message: "获取明星守护值失败"}
	ErrGetCrowdingFunding       = &Errno{Code: 50108, Message: "获取众筹数据失败"}
	ErrGetGifts                 = &Errno{Code: 50109, Message: "获取礼物数据失败"}
	ErrGetBarrage               = &Errno{Code: 50110, Message: "获取弹幕数据失败"}
	ErrUserCantGetPrize         = &Errno{Code: 50111, Message: "用户无领奖资格"}
	ErrCreateReceive            = &Errno{Code: 50112, Message: "创建用户领奖数据失败"}
	ErrGetClimbValue            = &Errno{Code: 50113, Message: "获取爬墙指数失败"}
	ErrCreateInvestigation      = &Errno{Code: 50114, Message: "保存用户问卷数据失败"}
	ErrNotFound                 = &Errno{Code: 50115, Message: "未查询到用户"}
	ErrGetCount                 = &Errno{Code: 50116, Message: "获取用户总数失败"}
	ErrGiveSomeBonus            = &Errno{Code: 50117, Message: "发放权益失败"}
	ErrGetSmallCountFail        = &Errno{Code: 50118, Message: "获取小于用户消费Count失败"}
	ErrGetThisIdolCount         = &Errno{Code: 50119, Message: "获取喜欢该idol的人数失败"}
	ErrGetIdolWorkByMdw         = &Errno{Code: 50120, Message: "根据mdw获取idolwork失败"}
	ErrGetMdwByMd               = &Errno{Code: 50121, Message: "根据md获取最受欢迎的mdw失败"}
	ErrIncrTotalAndHowMuchCount = &Errno{Code: 50122, Message: "添加人数失败"}
	ErrGetAccessTokenFail       = &Errno{Code: 50123, Message: "get weChat Access Token fail"}
	ErrGetTicketFail            = &Errno{Code: 50124, Message: "get weChat Ticket fail"}
	ErrGetMgoByObjectId         = &Errno{Code: 50125, Message: "get mongo data fail "}

	ErrGetRowsCountErr = &Errno{Code: 30226, Message: "获取数据表计数失败"}

	ErrGetUsersListErr = &Errno{Code: 30227, Message: "获取用户订单列表异常"}

	ErrEncodingMapDataErr = &Errno{Code: 30229, Message: "编码map数据异常"}

	ErrGetZCDataErr = &Errno{Code: 30230, Message: "获取众筹表数据错误"}
	ErrGetDZDataErr = &Errno{Code: 30231, Message: "获取点赞表数据错误"}

	ErrLIMITDZDataErr = &Errno{Code: 30233, Message: "今天点赞次数已达上限"}
	ErrLIMITDZHAVEErr = &Errno{Code: 30234, Message: "获取可点赞数据失败"}

	ErrGetFusionDataErr = &Errno{Code: 30235, Message: "获取融合数据错误"}
	ErrGetIdolStatusErr = &Errno{Code: 30236, Message: "获取idol状态数据错误"}
)
