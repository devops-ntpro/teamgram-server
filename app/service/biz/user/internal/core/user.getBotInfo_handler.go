/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/devops-ntpro/mtproto/mtproto"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/devops-ntpro/teamgram-server/app/service/biz/user/user"
)

// UserGetBotInfo
// user.getBotInfo bot_id:long = BotInfo;
func (c *UserCore) UserGetBotInfo(in *user.TLUserGetBotInfo) (*mtproto.BotInfo, error) {
	botsDO, err := c.svcCtx.BotsDAO.Select(c.ctx, in.BotId)
	if err != nil {
		c.Logger.Errorf("user.getBotInfo - error: %v", err)
		return nil, err
	} else if botsDO == nil {
		return nil, mtproto.ErrUserIdInvalid
	}

	botInfo := mtproto.MakeTLBotInfo(&mtproto.BotInfo{
		UserId:      in.BotId,
		Description: botsDO.Description,
		Commands:    []*mtproto.BotCommand{},
	}).To_BotInfo()

	c.svcCtx.Dao.BotCommandsDAO.SelectListWithCB(
		c.ctx,
		in.BotId,
		func(i int, v *dataobject.BotCommandsDO) {
			botInfo.Commands = append(botInfo.Commands, mtproto.MakeTLBotCommand(&mtproto.BotCommand{
				Command:     v.Command,
				Description: v.Description,
			}).To_BotCommand())
		})

	return botInfo, nil
}
