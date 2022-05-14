/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package miscellaneous_helper

import (
	"github.com/devops-ntpro/teamgram-server/app/bff/miscellaneous/internal/config"
	"github.com/devops-ntpro/teamgram-server/app/bff/miscellaneous/internal/server/grpc/service"
	"github.com/devops-ntpro/teamgram-server/app/bff/miscellaneous/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
