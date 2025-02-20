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
	"github.com/devops-ntpro/teamgram-server/app/service/biz/auth/auth"
)

// AuthImportLoginToken
// auth.importLoginToken token:bytes = auth.LoginToken;
func (c *AuthCore) AuthImportLoginToken(in *auth.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error) {
	// TODO: not impl
	c.Logger.Errorf("auth.importLoginToken - error: method AuthImportLoginToken not impl")

	return nil, mtproto.ErrMethodNotImpl
}
