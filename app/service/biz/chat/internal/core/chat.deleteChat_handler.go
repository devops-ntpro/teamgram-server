/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/devops-ntpro/mtproto/mtproto"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/chat/chat"
)

// ChatDeleteChat
// chat.deleteChat chat_id:long operator_id:long = MutableChat;
func (c *ChatCore) ChatDeleteChat(in *chat.TLChatDeleteChat) (*chat.MutableChat, error) {
	mChat, err := c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.deleteChat - error: %v", err)
		return nil, err
	}
	if mChat.Creator() != in.OperatorId {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.deleteChat - error: %v", err)
		return nil, err
	}

	tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		// kicked
		c.svcCtx.Dao.ChatParticipantsDAO.UpdateStateByChatIdTx(tx, mtproto.ChatMemberStateKicked, in.ChatId)
		c.svcCtx.Dao.ChatsDAO.UpdateParticipantCountTx(tx, 0, in.ChatId)
		c.svcCtx.Dao.ChatsDAO.UpdateDeactivatedTx(tx, false, in.ChatId)
	})
	if tR != nil {
		c.Logger.Errorf("chat.deleteChat - error: %v", err)
		return nil, err
	}

	return mChat, nil
}
