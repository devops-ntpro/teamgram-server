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

package dao

import (
	"github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/chat/plugin"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/chat/internal/config"
	media_client "github.com/devops-ntpro/teamgram-server/app/service/media/client"
)

// Dao dao.
type Dao struct {
	*Mysql
	sqlc.CachedConn
	media_client.MediaClient
	Plugin plugin.ChatPlugin
}

// New new a dao and return.
func New(c config.Config, plugin plugin.ChatPlugin) (dao *Dao) {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:       newMysqlDao(db),
		CachedConn:  sqlc.NewConn(db, c.Cache),
		MediaClient: media_client.NewMediaClient(rpcx.GetCachedRpcClient(c.MediaClient)),
		Plugin:      plugin,
	}
}
