/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package messages_helper

import (
	"github.com/devops-ntpro/teamgram-server/app/bff/messages/internal/config"
	"github.com/devops-ntpro/teamgram-server/app/bff/messages/internal/server/grpc/service"
	"github.com/devops-ntpro/teamgram-server/app/bff/messages/internal/svc"
	"github.com/devops-ntpro/teamgram-server/app/bff/messages/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.MessagesPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
