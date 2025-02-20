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
	"github.com/devops-ntpro/teamgram-server/app/service/biz/user/user"
)

// UserUpdateLastSeen
// user.updateLastSeen id:long last_seen_at:long expries:int = Bool;
func (c *UserCore) UserUpdateLastSeen(in *user.TLUserUpdateLastSeen) (*mtproto.Bool, error) {
	if in.GetId() > 0 {
		c.svcCtx.Dao.UserPresencesDAO.UpdateLastSeenAt(c.ctx, in.LastSeenAt, in.Expries, in.Id)
	}

	return mtproto.BoolTrue, nil
}
