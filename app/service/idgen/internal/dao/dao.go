/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dao

import (
	"log"

	"github.com/bwmarrin/snowflake"
	"github.com/devops-ntpro/teamgram-server/app/service/idgen/internal/config"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

type Dao struct {
	*snowflake.Node
	kv.Store
}

func New(c config.Config) *Dao {
	var (
		err error
		d   = new(Dao)
	)

	d.Node, err = snowflake.NewNode(c.NodeId)
	if err != nil {
		log.Fatal("new snowflake node error: ", err)
	}
	d.Store = kv.NewStore(c.SeqIDGen)

	return d
}
