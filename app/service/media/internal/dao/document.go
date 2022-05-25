// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/zeromicro/go-zero/core/mr"
	"time"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/devops-ntpro/mtproto/mtproto"
	"github.com/devops-ntpro/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	cacheDocumentPrefix = "document"
)

func genCacheDocumentKey(id int64) string {
	return fmt.Sprintf("%s_%d", cacheDocumentPrefix, id)
}

/*
document#1e87342b flags:#
	id:long
	access_hash:long
	file_reference:bytes
	date:int
	mime_type:string
	size:int
	thumbs:flags.0?Vector<PhotoSize>
	video_thumbs:flags.1?Vector<VideoSize>
	dc_id:int
	attributes:Vector<DocumentAttribute> = Document;
*/
func (m *Dao) makeDocumentByDO(
	ctx context.Context,
	document *mtproto.Document,
	id int64,
	do *dataobject.DocumentsDO,
	thumbs []*mtproto.PhotoSize,
	videoThumbs []*mtproto.VideoSize) {
	document.Id = do.DocumentId

	if do == nil {
		document.Id = id
		mtproto.MakeTLDocumentEmpty(&mtproto.Document{
			Id: id,
		})
		return
	}

	mtproto.MakeTLDocument(document)
	document.AccessHash = do.AccessHash
	document.FileReference = []byte{}
	document.Date = int32(do.Date2)
	if document.Date == 0 {
		document.Date = int32(time.Now().Unix())
	}
	document.MimeType = do.MimeType
	document.Size2 = do.FileSize

	if do.ThumbId != 0 && do.VideoThumbId != 0 {
		if len(thumbs) > 0 && len(videoThumbs) > 0 {
			document.Thumbs = thumbs
			document.VideoThumbs = videoThumbs
		} else {
			mr.FinishVoid(
				func() {
					document.Thumbs = m.GetPhotoSizeListV2(ctx, do.ThumbId)
				},
				func() {
					document.VideoThumbs = m.GetVideoSizeList(ctx, do.VideoThumbId)
				})
		}
	} else {
		// thumbs
		if do.ThumbId != 0 {
			if len(thumbs) > 0 {
				document.Thumbs = thumbs
			} else {
				document.Thumbs = m.GetPhotoSizeListV2(ctx, do.ThumbId)
			}
		}

		// video_thumbs
		if do.VideoThumbId != 0 {
			if len(videoThumbs) > 0 {
				document.VideoThumbs = videoThumbs
			} else {
				document.VideoThumbs = m.GetVideoSizeList(ctx, do.VideoThumbId)
			}
		}
	}

	document.DcId = 1
	err := json.Unmarshal([]byte(do.Attributes), &document.Attributes)
	if err != nil {
		document.Attributes = []*mtproto.DocumentAttribute{}
	}
}

func (m *Dao) GetDocumentById(ctx context.Context, id int64) *mtproto.Document {
	var (
		key      = genCacheDocumentKey(id)
		document = new(mtproto.Document)
	)

	m.CachedConn.QueryRow(ctx, document, key, func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
		do, _ := m.DocumentsDAO.SelectByDocumentId(ctx, id)
		if do == nil {
			logx.WithContext(ctx).Infof("not found document by id: %d", id)
		}
		m.makeDocumentByDO(ctx, v.(*mtproto.Document), id, do, nil, nil)
		return nil
	})

	return document
}

func (m *Dao) GetDocumentListByIdList(ctx context.Context, idList []int64) []*mtproto.Document {
	rList, err := mr.MapReduce(
		func(source chan<- interface{}) {
			for _, id2 := range idList {
				source <- id2
			}
		},
		func(item interface{}, writer mr.Writer, cancel func(error)) {
			id2 := item.(int64)
			document := new(mtproto.Document)
			// since2 := timex.Now()
			err := m.GetCache(ctx, genCacheDocumentKey(id2), document)
			if err != nil {
				if err != sqlc.ErrNotFound {
					cancel(err)
				} else {
					//
				}
			} else if document != nil {
				writer.Write(document)
			}
			// logx.WithDuration(timex.Since(since2)).Infof("getCache: %v", do)
		},
		func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			var documentList2 []*mtproto.Document
			for p := range pipe {
				documentList2 = append(documentList2, p.(*mtproto.Document))
			}
			writer.Write(documentList2)
		})
	if err != nil {
		logx.WithContext(ctx).Errorf("findListByIdList - %v", err)
	}

	var documentList []*mtproto.Document
	if rList != nil {
		documentList = rList.([]*mtproto.Document)
	}
	// logx.Infof("doList: %v", doList)

	if len(documentList) == len(idList) {
		return documentList
	}

	var (
		idList2 []int64
	)

	for _, id2 := range idList {
		for i := 0; i < len(documentList); i++ {
			if documentList[i].Id == id2 {
				goto Line100
			}
		}
		idList2 = append(idList2, id2)
	Line100:
	}

	var (
		thumbSizeIdList      = make([]int64, 0)
		videoThumbSizeIdList = make([]int64, 0)
	)
	missDoList, _ := m.DocumentsDAO.SelectByDocumentIdListWithCB(
		ctx,
		idList2,
		func(i int, v *dataobject.DocumentsDO) {
			if v.ThumbId != 0 {
				thumbSizeIdList = append(thumbSizeIdList, v.ThumbId)
			}
			if v.VideoThumbId != 0 {
				videoThumbSizeIdList = append(videoThumbSizeIdList, v.VideoThumbId)
			}
			v.Id = int64(i)
		})

	if len(missDoList) == 0 {
		return documentList
	}

	var (
		thumbSizeListList      map[int64][]*mtproto.PhotoSize
		videoThumbSizeListList map[int64][]*mtproto.VideoSize
		missDocumentList       = make([]*mtproto.Document, len(missDoList))
	)
	if len(thumbSizeIdList) > 0 && len(videoThumbSizeIdList) > 0 {
		mr.FinishVoid(
			func() {
				thumbSizeListList = m.GetPhotoSizeListList(ctx, thumbSizeIdList)
			},
			func() {
				videoThumbSizeListList = m.GetVideoSizeListList(ctx, videoThumbSizeIdList)
			})
	} else {
		if len(thumbSizeIdList) != 0 {
			thumbSizeListList = m.GetPhotoSizeListList(ctx, thumbSizeIdList)
		}
		if len(videoThumbSizeIdList) != 0 {
			videoThumbSizeListList = m.GetVideoSizeListList(ctx, videoThumbSizeIdList)
		}
	}

	mr.ForEach(
		func(source chan<- interface{}) {
			for i := 0; i < len(missDoList); i++ {
				source <- &missDoList[i]
			}
		},
		func(item interface{}) {
			var (
				do       = item.(*dataobject.DocumentsDO)
				document = new(mtproto.Document)
			)
			m.makeDocumentByDO(ctx, document, do.DocumentId, do, thumbSizeListList[do.ThumbId], videoThumbSizeListList[do.VideoThumbId])
			m.SetCache(ctx, genCacheDocumentKey(do.DocumentId), document)
			missDocumentList[do.Id] = document
		})

	return append(documentList, missDocumentList...)
}

func (m *Dao) SaveDocumentV2(ctx context.Context, fileName string, document *mtproto.Document) {
	var (
		aStr string
	)

	if document.GetAttributes() != nil {
		aBuf, _ := jsonx.Marshal(document.GetAttributes())
		aStr = hack.String(aBuf)
	}

	data := &dataobject.DocumentsDO{
		DocumentId:       document.Id,
		AccessHash:       document.AccessHash,
		DcId:             document.DcId,
		FilePath:         fmt.Sprintf("%d.dat", document.Id),
		FileSize:         document.Size2,
		UploadedFileName: fileName,
		Ext:              getFileExtName(fileName),
		MimeType:         document.MimeType,
		ThumbId:          0,
		VideoThumbId:     0,
		Version:          0,
		Attributes:       aStr,
		Date2:            int64(document.Date),
	}
	if len(document.GetThumbs()) > 0 {
		data.ThumbId = document.Id
	}
	if len(document.GetVideoThumbs()) > 0 {
		data.VideoThumbId = document.Id
	}

	data.Id, _, _ = m.DocumentsDAO.Insert(ctx, data)
	return
}
