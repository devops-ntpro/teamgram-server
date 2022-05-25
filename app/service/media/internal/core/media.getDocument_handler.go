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
	"github.com/devops-ntpro/teamgram-server/app/service/media/media"
)

// MediaGetDocument
// media.getDocument id:long = Document;
func (c *MediaCore) MediaGetDocument(in *media.TLMediaGetDocument) (*mtproto.Document, error) {
	document := c.svcCtx.Dao.GetDocumentById(c.ctx, in.GetId())

	return document, nil
}
