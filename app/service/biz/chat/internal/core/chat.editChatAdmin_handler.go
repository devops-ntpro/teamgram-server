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
	"github.com/devops-ntpro/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
	"time"

	"github.com/devops-ntpro/mtproto/mtproto"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/chat/chat"
)

// ChatEditChatAdmin
// chat.editChatAdmin chat_id:long operator_id:long edit_chat_admin_id:long is_admin:Bool = MutableChat;
func (c *ChatCore) ChatEditChatAdmin(in *chat.TLChatEditChatAdmin) (*chat.MutableChat, error) {
	var (
		now           = time.Now().Unix()
		chat2         *chat.MutableChat
		me, editAdmin *chat.ImmutableChatParticipant
		err           error
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId, in.OperatorId, in.EditChatAdminId)
	if err != nil {
		c.Logger.Errorf("chat.editChatAdmin - error: %v", err)
		err = mtproto.ErrChatIdInvalid
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(in.OperatorId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrUserNotParticipant
		c.Logger.Errorf("chat.editChatAdmin - error: %v", err)
		return nil, err
	}

	editAdmin, _ = chat2.GetImmutableChatParticipant(in.EditChatAdminId)
	if editAdmin != nil && editAdmin.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("chat.editChatAdmin - error: %v", err)
		return nil, err
	}

	if !me.CanAdminAddAdmins() {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.editChatAdmin - error: %v", err)
		return nil, err
	}

	tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if mtproto.FromBool(in.IsAdmin) {
			_, result.Err = c.svcCtx.Dao.ChatParticipantsDAO.UpdateParticipantTypeTx(tx, mtproto.ChatMemberAdmin, editAdmin.Id)
			if result.Err != nil {
				c.Logger.Errorf("chat.editChatAdmin - error: %v", result.Err)
				return
			}

			if editAdmin.Link == "" {
				editAdmin.Link = chat.GenChatInviteHash()
				c.svcCtx.Dao.ChatParticipantsDAO.UpdateLinkTx(tx, editAdmin.Link, in.ChatId, in.EditChatAdminId)
				c.svcCtx.Dao.ChatInvitesDAO.InsertTx(tx, &dataobject.ChatInvitesDO{
					ChatId:    in.ChatId,
					AdminId:   in.EditChatAdminId,
					Link:      editAdmin.Link,
					Permanent: true,
					Date2:     now,
				})
			}

			editAdmin.AdminRights = mtproto.MakeDefaultChatAdminRights()
			editAdmin.ParticipantType = mtproto.ChatMemberAdmin
		} else {
			_, result.Err = c.svcCtx.Dao.ChatParticipantsDAO.UpdateParticipantType(c.ctx, mtproto.ChatMemberNormal, editAdmin.Id)
			if result.Err != nil {
				c.Logger.Errorf("chat.editChatAdmin - error: %v", result.Err)
				return
			}
			editAdmin.AdminRights = nil
			editAdmin.ParticipantType = mtproto.ChatMemberNormal
			editAdmin.Link = ""
		}
	})
	if tR.Err != nil {
		return nil, tR.Err
	}

	chat2.Chat.Version += 1
	chat2.Chat.Date = now
	return chat2, nil
}
