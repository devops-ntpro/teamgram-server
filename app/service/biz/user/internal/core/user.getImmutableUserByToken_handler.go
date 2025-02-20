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

// UserGetImmutableUserByToken
// user.getImmutableUserByToken token:string = ImmutableUser;
func (c *UserCore) UserGetImmutableUserByToken(in *user.TLUserGetImmutableUserByToken) (*user.ImmutableUser, error) {
	botId, err := c.svcCtx.Dao.BotsDAO.SelectByToken(c.ctx, in.Token)
	if err != nil || botId == 0 {
		err = mtproto.ErrTokenInvalid
		return nil, err
	}

	return c.UserGetImmutableUser(&user.TLUserGetImmutableUser{
		Id: botId,
	})
}
