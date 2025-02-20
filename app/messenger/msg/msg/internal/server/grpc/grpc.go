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

package grpc

import (
	"github.com/devops-ntpro/teamgram-server/app/messenger/msg/msg/internal/server/grpc/service"
	"github.com/devops-ntpro/teamgram-server/app/messenger/msg/msg/internal/svc"
	msgpb "github.com/devops-ntpro/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c zrpc.RpcServerConf) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		msgpb.RegisterRPCMsgServer(grpcServer, service.New(ctx))
	})
	logx.Must(err)
	return s
}
