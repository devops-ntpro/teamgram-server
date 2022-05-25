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

// AuthsessionQueryAuthKey
// authsession.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (c *AuthsessionCore) AuthsessionQueryAuthKey(in *authsession.TLAuthsessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	rValue, err := c.svcCtx.Dao.QueryAuthKey(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("authsession.queryAuthKey - error: %v", err.Error())
		return nil, err
	}

	return rValue, nil
}
