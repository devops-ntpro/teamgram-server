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
	"github.com/devops-ntpro/mtproto/mtproto"
	"github.com/devops-ntpro/teamgram-server/app/messenger/sync/sync"
)

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long auth_key_id:long server_id:string session_id:flags.0?long updates:Updates = Void;
func (c *SyncCore) SyncUpdatesMe(in *sync.TLSyncUpdatesMe) (*mtproto.Void, error) {
	c.pushUpdatesToSession(syncTypeUserMe,
		in.GetUserId(),
		in.GetAuthKeyId(),
		in.GetSessionId().GetValue(),
		in.GetUpdates(),
		in.GetServerId(),
		false)

	return mtproto.EmptyVoid, nil
}
