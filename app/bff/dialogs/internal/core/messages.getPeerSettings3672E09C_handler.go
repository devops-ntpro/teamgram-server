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
	userpb "github.com/devops-ntpro/teamgram-server/app/service/biz/user/user"
)

// MessagesGetPeerSettings3672E09C
// messages.getPeerSettings#3672e09c peer:InputPeer = PeerSettings;
func (c *DialogsCore) MessagesGetPeerSettings3672E09C(in *mtproto.TLMessagesGetPeerSettings3672E09C) (*mtproto.PeerSettings, error) {
	peer := mtproto.FromInputPeer2(c.MD.UserId, in.Peer)

	peerSettings, err := c.svcCtx.UserClient.UserGetPeerSettings(c.ctx, &userpb.TLUserGetPeerSettings{
		UserId:   c.MD.UserId,
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
	})

	if err != nil {
		c.Logger.Errorf("messages.getPeerSettings - error: %v", err)

		// TODO(@benqi): handle error
		peerSettings = mtproto.MakeTLPeerSettings(nil).To_PeerSettings()
	}

	return peerSettings, nil
}
