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

// UserPredefinedBindRegisteredUserId
// user.predefinedBindRegisteredUserId phone:string registered_userId:int = Bool;
func (c *UserCore) UserPredefinedBindRegisteredUserId(in *user.TLUserPredefinedBindRegisteredUserId) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("user.predefinedBindRegisteredUserId - error: method UserPredefinedBindRegisteredUserId not impl")

	return nil, mtproto.ErrMethodNotImpl
}
