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
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/devops-ntpro/teamgram-server/app/bff/usernames/internal/config"
	sync_client "github.com/devops-ntpro/teamgram-server/app/messenger/sync/client"
	chat_client "github.com/devops-ntpro/teamgram-server/app/service/biz/chat/client"
	user_client "github.com/devops-ntpro/teamgram-server/app/service/biz/user/client"
	username_client "github.com/devops-ntpro/teamgram-server/app/service/biz/username/client"
)

type Dao struct {
	user_client.UserClient
	chat_client.ChatClient
	username_client.UsernameClient
	sync_client.SyncClient
}

func New(c config.Config) *Dao {
	return &Dao{
		UserClient:     user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:     chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		UsernameClient: username_client.NewUsernameClient(rpcx.GetCachedRpcClient(c.UsernameClient)),
		SyncClient:     sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
	}
}
