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

package core

import (
	"github.com/devops-ntpro/mtproto/mtproto"
)

// MessagesGetDocumentByHash
// messages.getDocumentByHash#338e2464 sha256:bytes size:int mime_type:string = Document;
func (c *FilesCore) MessagesGetDocumentByHash(in *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.getDocumentByHash blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return mtproto.MakeTLDocumentEmpty(nil).To_Document(), nil
}
