// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"path"
	"strings"

	"github.com/devops-ntpro/mtproto/mtproto"
)

var (
	emptyPhoto = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
)

func getFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	if ext == "" {
		ext = "partial"
	}
	return strings.ToLower(ext)
}
