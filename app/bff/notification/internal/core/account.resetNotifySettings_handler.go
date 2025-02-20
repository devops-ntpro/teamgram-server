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
	"github.com/gogo/protobuf/types"
	"github.com/devops-ntpro/mtproto/mtproto"
	"github.com/devops-ntpro/teamgram-server/app/messenger/sync/sync"
	userpb "github.com/devops-ntpro/teamgram-server/app/service/biz/user/user"
)

// AccountResetNotifySettings
// account.resetNotifySettings#db7e1747 = Bool;
func (c *NotificationCore) AccountResetNotifySettings(in *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.UserClient.UserResetNotifySettings(c.ctx, &userpb.TLUserResetNotifySettings{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("getNotifySettings error - %v", err)
		// We ignore error
		return mtproto.BoolFalse, nil
	}

	pushNotifySettingsFunc := func(peerType int32) {
		peer := &mtproto.PeerUtil{
			PeerType: peerType,
			PeerId:   0,
		}
		syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateNotifySettings(&mtproto.Update{
			Peer_NOTIFYPEER: peer.ToNotifyPeer(),
			NotifySettings: mtproto.MakeTLPeerNotifySettings(&mtproto.PeerNotifySettings{
				ShowPreviews: mtproto.BoolTrue,
				Silent:       mtproto.BoolFalse,
				MuteUntil:    &types.Int32Value{Value: 0},
				Sound:        &types.StringValue{Value: "default"},
			}).To_PeerNotifySettings(),
		}).To_Update())
		c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
			Constructor: 0,
			UserId:      c.MD.UserId,
			AuthKeyId:   c.MD.AuthId,
			Updates:     syncUpdates,
		})
	}

	pushNotifySettingsFunc(mtproto.PEER_USERS)
	pushNotifySettingsFunc(mtproto.PEER_CHATS)
	pushNotifySettingsFunc(mtproto.PEER_BROADCASTS)

	return mtproto.BoolTrue, nil
}
