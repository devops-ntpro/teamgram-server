/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package user_helper

import (
	"github.com/devops-ntpro/teamgram-server/app/service/biz/user/internal/config"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/user/internal/server/grpc/service"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/user/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
