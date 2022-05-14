#!/bin/bash

dalgen3 --xml=$1 --db=teamgram --go2=github.com/devops-ntpro/teamgram-server/app/messenger/msg/internal/dal/dataobject

gofmt -w ../dao/mysql_dao/*.go
gofmt -w ../dataobject/*.go
