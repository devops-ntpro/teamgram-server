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

// UserChangePhone
// user.changePhone user_id:int phone:string = Bool;
func (c *UserCore) UserChangePhone(in *user.TLUserChangePhone) (*mtproto.Bool, error) {
	c.svcCtx.Dao.UsersDAO.UpdateUser(c.ctx, map[string]interface{}{
		"phone": in.Phone, // TODO(@benqi): country_code
	}, in.UserId)

	c.svcCtx.Dao.UserContactsDAO.UpdatePhoneByContactId(c.ctx, in.Phone, in.UserId)

	return mtproto.BoolTrue, nil
}
