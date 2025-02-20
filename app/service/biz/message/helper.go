/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package message_helper

import (
	"github.com/devops-ntpro/teamgram-server/app/service/biz/message/internal/config"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/message/internal/dal/dao/mysql_dao"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/message/internal/server/grpc/service"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/message/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}

type (
	MessagesDAO = mysql_dao.MessagesDAO
	MessagesDO  = dataobject.MessagesDO
	// ChannelMessagesDAO   = mysql_dao.ChannelMessagesDAO
	// ChannelMessagesDO    = dataobject.ChannelMessagesDO
	// ScheduledMessagesDAO = mysql_dao.ScheduledMessagesDAO
	// ScheduledMessagesDO  = dataobject.ScheduledMessagesDO
)

var (
	NewMessagesDAO = mysql_dao.NewMessagesDAO
	// NewChannelMessagesDAO   = mysql_dao.NewChannelMessagesDAO
	// NewScheduledMessagesDAO = mysql_dao.NewScheduledMessagesDAO
)
