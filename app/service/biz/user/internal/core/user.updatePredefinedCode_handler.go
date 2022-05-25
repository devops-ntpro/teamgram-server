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

// UserUpdatePredefinedCode
// user.updatePredefinedCode phone:string code:string = PredefinedUser;
func (c *UserCore) UserUpdatePredefinedCode(in *user.TLUserUpdatePredefinedCode) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.updatePredefinedCode - error: method UserUpdatePredefinedCode not impl")

	return nil, mtproto.ErrMethodNotImpl
}
