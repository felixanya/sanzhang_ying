// Code generated by protoc-gen-go.
// source: pb/client_msg.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/client_msg.proto
	pb/config.proto
	pb/game.proto
	pb/messageId.proto
	pb/rpc.proto
	pb/server_msg.proto
	pb/server_msgId.proto

It has these top-level messages:
	ClientMsg
	Config
	PrizeDef
	UserMatchRecordDef
	ShareInfoDef
	UserDef
	ActiveStatusDef
	Msg_GetAciveStatusReq
	Msg_GetAciveStatusAck
	BoxLogsDef
	SafeBoxDef
	UserVipTaskDef
	UserMagicItemDef
	UserPrizeTaskDef
	WanRenWaitQueueUserDef
	UserPokerDeskDef
	PokerDeskDef
	MsgJoinWaitQueue
	MsgLeaveWaitQueue
	MsgUpdateGold
	MsgUpdateGoldInGame
	MsgLoginReq
	MsgLoginRes
	MsgUpdateUserInfoReq
	MsgUpdateUserInfoRes
	MsgGetUserInfoReq
	MsgGetUserInfoRes
	MsgEnterPokerDeskReq
	MsgEnterPokerDeskRes
	MsgEnterPokerDeskBro
	MsgLeavePokerDeskReq
	MsgLeavePokerDesk
	WanRenLookupBetItem
	MsgKickPlayerReq
	MsgKickPlayerRes
	MsgShowTips
	MsgGoodsTips
	MsgOpenFishTankReq
	MsgOpenFishTankRes
	MsgUpdateVipTask
	Msg_ReConnectReq
	Msg_ReConnectRes
	PrizeOnline
	Msg_UpdateOnlinePrize
	Msg_LoginSafeBoxReq
	Msg_LoginSafeBoxRes
	Msg_ChangePwdSafeBoxReq
	Msg_ChangePwdSafeBoxRes
	Msg_ResetPwdSafeBoxReq
	Msg_ResetPwdSafeBoxRes
	Msg_ResetPwdUserReq
	Msg_ResetPwdUserRes
	MsgRegisterReq
	MsgRegisterRes
	Msg_SaveGoldSafeBoxReq
	Msg_SaveGoldSafeBoxRes
	Msg_GetGoldSafeBoxReq
	Msg_GetGoldSafeBoxRes
	Msg_GiftGoldSafeBoxReq
	Msg_GiftGoldSafeBoxRes
	Msg_GetPayTokenReq
	Msg_GetPayTokenRes
	Msg_GetVerifyCodeReq
	Msg_GetVerifyCodeRes
	MsgGetPokerDeskInfoRes
	MsgGameStart
	MsgPlayerRoundBeginRes
	MsgOpCardReq
	MsgOpCardRes
	MsgUseMagicItemReq
	MsgUseMagicItemRes
	MsgUseMagicItemBro
	MsgPokerDeskGameEndRes
	MsgSitDownPokerDeskRes
	MsgStandUpPokerDeskRes
	MsgExchangeGoldReq
	MsgExchangeGoldRes
	MsgExchangeGameGoodsReq
	MsgExchangeGameGoodsRes
	MsgExchangeScoreGoodsReq
	MsgExchangeScoreGoodsRes
	RankingItem
	RankingList
	MsgGetRankingListReq
	MsgGetRankingListRes
	MsgChat
	MsgGetShopLogRes
	MsgPlayLuckWheelReq
	MsgPlayLuckWheelRes
	MsgGetLuckWheelLogRes
	MsgGainVipPrizeReq
	MsgGainVipPrizeRes
	MsgGainOnlinePrizeReq
	MsgGainOnlinePrizeRes
	MsgGainTaskPrizeReq
	MsgGainTaskPrizeRes
	MsgUpdateTaskRes
	MsgReadySNGBro
	MsgLookupBetGold
	MsgLookupBetGoldRes
	MsgLookupBetGoldBro
	MsgLookupBetWin
	Msg_SlotMachinesPlayReq
	Msg_SlotMachinesPlayRes
	Msg_SlotMachinesReplaceReq
	Msg_SlotMachinesReplaceRes
	Msg_SlotMachinesPrizeRes
	MsgGetMatchRecordRes
	MsgGetPrizeAddressRes
	MsgBindPrizeAddressReq
	MsgBindPrizeAddressRes
	MsgSubsidyPrizeRes
	MsgRewardInGame
	MsgExchangeCodePrizeReq
	MsgExchangeCodePrizeRes
	MsgGetOnlineStatusRes
	PrizeMailDef
	MsgGetPrizeMailListRes
	MsgGainMailPrizeReq
	MsgGainMailPrizeRes
	MsgRobotMaxCardUser
	MsgRobotCards
	MsgRobotSetGold
	MsgReportUserHeadIllegalReq
	MsgSignInRecord
	MsgSignInRes
	MsgGetRechargeInfoRes
	MsgChangeGameTypeNotify
	MsgUseThreePartyPayRes
	ExchangeGoodsInfo
	MsgGetExchangeGoodsRes
	MsgDeskInfo
	MsgGetActiveContentRes
	MsgGetOnlineConfig
	MsgGetOnlineConfigRes
	NewbeTask
	MsgGetNewbeTaskListRes
	MsgGetNewbeTaskReward
	MsgGetNewbeTaskRewardRes
	MsgNewbeTaskCompletedNotify
	MsgExchangeNewbeTaskHf
	MsgExchangeNewbeTaskHfRes
	ActiveDataDef
	MsgGetActiveDataReq
	MsgGetActiveDataAck
	GoodsListDef
	MsgGetGoodsListReq
	MsgGetGoodsListAck
	MatchConfigDef
	MsgGetMatchConfigReq
	MsgGetMatchConfigAck
	MsgGetActiveTokenReq
	MsgGetActiveTokenRes
	LigameChipDef
	MsgGetLittleGameConfigReq
	MsgGetLittleGameConfigRes
	MsgLittleGameReq
	MsgLittleGameRes
	RpcRequest
	RpcResponse
	ServerMsg
	MsgMatchResult
	MsgLockUser
	MsgUpdateRechargeDiamond
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ClientMsg struct {
	MsgId            *int32 `protobuf:"varint,1,req,name=msgId" json:"msgId,omitempty"`
	MsgBody          []byte `protobuf:"bytes,2,opt,name=msgBody" json:"msgBody,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClientMsg) Reset()         { *m = ClientMsg{} }
func (m *ClientMsg) String() string { return proto.CompactTextString(m) }
func (*ClientMsg) ProtoMessage()    {}

func (m *ClientMsg) GetMsgId() int32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *ClientMsg) GetMsgBody() []byte {
	if m != nil {
		return m.MsgBody
	}
	return nil
}