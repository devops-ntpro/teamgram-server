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

// UserIsBot
// user.isBot id:long = Bool;
func (c *UserCore) UserIsBot(in *user.TLUserIsBot) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("user.isBot - error: method UserIsBot not impl")

	return nil, mtproto.ErrMethodNotImpl
}
