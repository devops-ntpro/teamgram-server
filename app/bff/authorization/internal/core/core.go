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
	"context"
	//"fmt"
	//"math/rand"
	//"time"

	"github.com/devops-ntpro/mtproto/mtproto"
	"github.com/devops-ntpro/mtproto/mtproto/rpc/metadata"
	"github.com/devops-ntpro/teamgram-server/app/bff/authorization/internal/svc"
	//msgpb "github.com/devops-ntpro/teamgram-server/app/messenger/msg/msg/msg"
	//"github.com/devops-ntpro/teamgram-server/pkg/env2"
	"github.com/devops-ntpro/teamgram-server/pkg/phonenumber"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizationCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorizationCore {
	return &AuthorizationCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func checkPhoneNumberInvalid(phone string) (string, error) {
	// 3. check number
	// 3.1. empty
	if phone == "" {
		// log.Errorf("check phone_number error - empty")
		return "", mtproto.ErrPhoneNumberInvalid
	}

	// 3.2. check phone_number
	pNumber, err := phonenumber.MakePhoneNumberHelper(phone, "")
	if err != nil {
		// log.Errorf("check phone_number error - %v", err)
		// err = mtproto.ErrPhoneNumberInvalid
		return "", mtproto.ErrPhoneNumberInvalid
	}

	return pNumber.Number, nil
}

const signInMessageTpl = `Login code: %s. Do not give this code to anyone, even if they say they are from %s!

This code can be used to log in to your %s account. We never ask it for anything else.

If you didn't request this code by trying to log in on another device, simply ignore this message.`

func (c *AuthorizationCore) pushSignInMessage(ctx context.Context, signInUserId int64, code string) {
	/*
	time.AfterFunc(2*time.Second, func() {
		message := mtproto.MakeTLMessage(&mtproto.Message{
			Out:     true,
			Date:    int32(time.Now().Unix()),
			FromId:  mtproto.MakePeerUser(777000),
			PeerId:  mtproto.MakeTLPeerUser(&mtproto.Peer{UserId: signInUserId}).To_Peer(),
			Message: fmt.Sprintf(signInMessageTpl, code, env2.MyAppName, env2.MyAppName),
			Entities: []*mtproto.MessageEntity{
				mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
					Offset: 0,
					Length: 11,
				}).To_MessageEntity(),
				mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
					Offset: 22,
					Length: 3,
				}).To_MessageEntity(),
			},
		}).To_Message()

		_ = ctx
		c.svcCtx.Dao.MsgClient.MsgPushUserMessage(
			context.Background(),
			&msgpb.TLMsgPushUserMessage{
				UserId:    777000,
				AuthKeyId: 0,
				PeerType:  mtproto.PEER_USER,
				PeerId:    signInUserId,
				PushType:  1,
				Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
					NoWebpage:    false,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      message,
					ScheduleDate: nil,
				}).To_OutboxMessage(),
			})
	})*/
}
