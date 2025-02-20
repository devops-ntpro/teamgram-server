// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/devops-ntpro/mtproto/mtproto"
	msgpb "github.com/devops-ntpro/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/devops-ntpro/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/devops-ntpro/teamgram-server/app/service/biz/chat/chat"
)

// MessagesDeleteChat
// messages.deleteChat#5bd0ee50 chat_id:long = Bool;
func (c *ChatsCore) MessagesDeleteChat(in *mtproto.TLMessagesDeleteChat) (*mtproto.Bool, error) {
	// 2. delete chat
	chat, err := c.svcCtx.Dao.ChatClient.ChatDeleteChat(c.ctx, &chatpb.TLChatDeleteChat{
		ChatId: in.ChatId,
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteChat - error: %v", err)
		return nil, err
	}

	pushUpdates := mtproto.MakeUpdatesByUpdatesChats(
		[]*mtproto.Chat{chat.ToChatForbidden()},
		mtproto.MakeUpdateChat(chat.Id()))

	// 1. kicked all
	chat.Walk(func(userId int64, participant *chatpb.ImmutableChatParticipant) error {
		if userId == c.MD.UserId || participant.IsChatMemberStateNormal() {
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
				UserId:  userId,
				Updates: pushUpdates,
			})
		}

		c.svcCtx.Dao.MsgClient.MsgDeleteChatHistory(c.ctx, &msgpb.TLMsgDeleteChatHistory{
			ChatId:       chat.Id(),
			DeleteUserId: userId,
		})
		return nil
	})

	return mtproto.BoolTrue, nil
}
