/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/devops-ntpro/mtproto/mtproto"
	"github.com/devops-ntpro/teamgram-server/app/bff/drafts/internal/core"
)

// MessagesSaveDraft
// messages.saveDraft#bc39e14b flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> = Bool;
func (s *Service) MessagesSaveDraft(ctx context.Context, request *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.saveDraft - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSaveDraft(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.saveDraft - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetAllDrafts
// messages.getAllDrafts#6a3f8d65 = Updates;
func (s *Service) MessagesGetAllDrafts(ctx context.Context, request *mtproto.TLMessagesGetAllDrafts) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getAllDrafts - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetAllDrafts(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getAllDrafts - reply: %s", r.DebugString())
	return r, err
}

// MessagesClearAllDrafts
// messages.clearAllDrafts#7e58ee9c = Bool;
func (s *Service) MessagesClearAllDrafts(ctx context.Context, request *mtproto.TLMessagesClearAllDrafts) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.clearAllDrafts - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesClearAllDrafts(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.clearAllDrafts - reply: %s", r.DebugString())
	return r, err
}
