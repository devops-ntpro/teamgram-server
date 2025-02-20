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

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (c *UserCore) UserGetImmutableUserByPhone(in *user.TLUserGetImmutableUserByPhone) (*user.ImmutableUser, error) {
	do, err := c.svcCtx.Dao.UsersDAO.SelectByPhoneNumber(c.ctx, in.Phone)
	if err != nil {
		c.Logger.Errorf("user.getImmutableUserByPhone - error: %v", err)
		return nil, err
	} else if do == nil {
		err = mtproto.ErrPhoneNumberUnoccupied
		c.Logger.Errorf("user.getImmutableUserByPhone - error: %v", err)
		return nil, err
	}

	return c.UserGetImmutableUser(&user.TLUserGetImmutableUser{
		Id: do.Id,
	})
}
