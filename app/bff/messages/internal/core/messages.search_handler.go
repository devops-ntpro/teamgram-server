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
	"math"

	"github.com/devops-ntpro/mtproto/mtproto"
	chatpb "github.com/devops-ntpro/teamgram-server/app/service/biz/chat/chat"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/message/message"
	userpb "github.com/devops-ntpro/teamgram-server/app/service/biz/user/user"
)

// MessagesSearch
// messages.search#a0fda762 flags:# peer:InputPeer q:string from_id:flags.0?InputPeer top_msg_id:flags.1?int filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (c *MessagesCore) MessagesSearch(in *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if c.MD.IsBot {
		err := mtproto.ErrBotMethodInvalid
		c.Logger.Errorf("messages.search - error: %v", err)
		return nil, err
	}

	var (
		rValues  *mtproto.Messages_Messages
		offsetId = in.OffsetId
		limit    = in.Limit
		boxList  *message.Vector_MessageBox
		err      error
	)

	if offsetId == 0 {
		offsetId = math.MaxInt32
	}

	if limit > 50 {
		limit = 50
	}

	peer := mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
	if peer.IsChannel() {
		// TODO: not impl
		c.Logger.Errorf("messages.search blocked, License key from https://teamgram.net required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	rValues = mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages()

	filterType := mtproto.FromMessagesFilter(in.Filter)
	switch filterType {
	case mtproto.FilterPhotos:
		// TODO
		// c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterVideo:
		// TODO
		// c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterPhotoVideo:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_PHOTOVIDEO,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterDocument:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_FILE,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterUrl:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_URL,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterGif:
		// TODO
		// c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterVoice:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterMusic:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_MUSIC,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterChatPhotos:
		// TODO
	case mtproto.FilterPhoneCalls:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_PHONE_CALL,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterRoundVoice:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByMediaType(c.ctx, &message.TLMessageSearchByMediaType{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			MediaType: mtproto.MEDIA_AUDIO,
			Offset:    offsetId,
			Limit:     limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterRoundVideo:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterMyMentions:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterGeo:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterContacts:
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	case mtproto.FilterPinned:
		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearchByPinned(c.ctx, &message.TLMessageSearchByPinned{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	case mtproto.FilterEmpty:
		if in.Q == "" {
			err = mtproto.ErrSearchQueryEmpty
			c.Logger.Errorf("messages.search - error: %v", err)
			return nil, err
		}

		boxList, err = c.svcCtx.Dao.MessageClient.MessageSearch(c.ctx, &message.TLMessageSearch{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
			Q:        in.Q,
			Offset:   offsetId,
			Limit:    limit,
		})
		if err != nil {
			c.Logger.Errorf("messages.search - error: %v", err)
			return rValues, nil
		}
	default:
		// TODO
		c.Logger.Errorf("messages.search - invalid filter: %s", in.DebugString())
		return rValues, nil
	}

	//
	if peer.PeerType == mtproto.PEER_CHANNEL {
		rValues.Count = boxList.Length()
		//channelLogic, err := s.ChannelCore.NewChannelLogicById(ctx, peer.PeerId)
		//if err != nil {
		//	messages.Pts = channelLogic.Pts
		//}
	} else {

	}

	boxList.Visit(c.MD.UserId,
		func(messageList []*mtproto.Message) {
			rValues.Messages = messageList
		},
		func(userIdList []int64) {
			mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			rValues.Users = append(rValues.Users, mUsers.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			mChats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			rValues.Chats = append(rValues.Chats, mChats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			//mChannels, _ := c.svcCtx.Dao.ChannelClient.ChannelGetChannelListByIdList(c.ctx,
			//	&channelpb.TLChannelGetChannelListByIdList{
			//		SelfUserId: c.MD.UserId,
			//		Id:         channelIdList,
			//	})
			//if len(mChannels.GetDatas()) > 0 {
			//	rValues.Chats = append(rValues.Chats, mChannels.GetDatas()...)
			//}
		})

	return rValues, nil
}
