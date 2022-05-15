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

package server

import (
	//	"net"
	"github.com/teamgram/marmota/pkg/timer2"
	"github.com/zeromicro/go-zero/core/logx"
	
	"github.com/devops-ntpro/teamgram-server/app/interface/ntproxy/internal/config"
)

var (
	//etcdPrefix is a etcd globe key prefix
	endpoints string
)

type Server struct {
	c      *config.Config
	timer          *timer2.Timer // 32 * 2048
}

func New(c config.Config) *Server {
	var (
		s   = new(Server)
	)

	s.timer = timer2.NewTimer(1024)
	logx.Infof("New: %v", c)

	return s
}

func (_ *Server) Serve() error {

	logx.Infof("Serve")
	return nil
}

func (_ *Server) Close() {
	logx.Infof("Close")
}
