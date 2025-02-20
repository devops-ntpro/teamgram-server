// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/devops-ntpro/mtproto/mtproto"
	"github.com/devops-ntpro/teamgram-server/app/service/authsession/authsession"
)

// AuthLogOut5717DA40
// auth.logOut#5717da40 = Bool;
func (c *AuthorizationCore) AuthLogOut5717DA40(in *mtproto.TLAuthLogOut5717DA40) (*mtproto.Bool, error) {
	// unbind auth_key and user_id
	_, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionUnbindAuthKeyUser(c.ctx, &authsession.TLAuthsessionUnbindAuthKeyUser{
		AuthKeyId: c.MD.AuthId,
		UserId:    c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("auth.logOut - error: %v", err)
	} else {
		if c.svcCtx.Plugin != nil {
			c.svcCtx.Plugin.OnAuthLogout(c.ctx, c.MD.UserId, c.MD.AuthId)
		}
	}

	return mtproto.BoolTrue, nil
}
