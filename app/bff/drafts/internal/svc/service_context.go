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

package svc

import (
	"github.com/devops-ntpro/teamgram-server/app/bff/drafts/internal/config"
	"github.com/devops-ntpro/teamgram-server/app/bff/drafts/internal/dao"
	"github.com/devops-ntpro/teamgram-server/app/bff/drafts/plugin"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	Plugin plugin.DraftsPlugin
}

func NewServiceContext(c config.Config, plugin plugin.DraftsPlugin) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
		Plugin: plugin,
	}
}
