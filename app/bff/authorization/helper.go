/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package authorization_helper

import (
	"github.com/devops-ntpro/teamgram-server/app/bff/authorization/internal/config"
	"github.com/devops-ntpro/teamgram-server/app/bff/authorization/internal/server/grpc/service"
	"github.com/devops-ntpro/teamgram-server/app/bff/authorization/internal/svc"
	"github.com/devops-ntpro/teamgram-server/app/bff/authorization/plugin"
)

type (
	Config               = config.Config
	AuthorizationService = service.Service
)

func New(c Config, plugin plugin.AuthorizationPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
