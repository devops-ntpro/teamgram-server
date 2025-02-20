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
	"github.com/devops-ntpro/teamgram-server/app/service/authsession/authsession"
)

// AuthsessionGetPermAuthKeyId
// authsession.getPermAuthKeyId auth_key_id:long= Int64;
func (c *AuthsessionCore) AuthsessionGetPermAuthKeyId(in *authsession.TLAuthsessionGetPermAuthKeyId) (*mtproto.Int64, error) {
	v := c.svcCtx.Dao.GetPermAuthKeyId(c.ctx, in.AuthKeyId)

	return mtproto.MakeTLInt64(&mtproto.Int64{
		V: v,
	}).To_Int64(), nil
}
