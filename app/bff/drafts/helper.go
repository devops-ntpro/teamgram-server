/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package drafts_helper

import (
	"github.com/devops-ntpro/teamgram-server/app/bff/drafts/internal/config"
	"github.com/devops-ntpro/teamgram-server/app/bff/drafts/internal/server/grpc/service"
	"github.com/devops-ntpro/teamgram-server/app/bff/drafts/internal/svc"
	"github.com/devops-ntpro/teamgram-server/app/bff/drafts/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.DraftsPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
