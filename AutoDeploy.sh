#!/bin/sh

#### File PATH
PWD=`pwd`
BIN_FILE=./gosignaler
PROCESS=gosignaler

#### REMOTE HOST INFO
USER=root
#HOST=120.78.168.126
HOST=211.159.220.166
REMOTE_PATH=/var/www/html/signaler
#REMOTE_FILE=$PROCESS.`date +%s`

#### Stop & Start Command
STOP_COMMAND="killall $PROCESS"
START_COMMAND="nohup $REMOTE_PATH/$PROCESS 1 > $REMOTE_PATH/logs.out 2>&1 &"

#CGO_ENABLED=0 GOOS=linux GOARCH=386 go build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

echo "ssh $USER@$HOST $STOP_COMMAND"
ssh $USER@$HOST $STOP_COMMAND

echo "$USER@$HOST:$REMOTE_PATH/"
scp $BIN_FILE $USER@$HOST:$REMOTE_PATH/ 

#echo "ssh $USER@$HOST $START_COMMAND"
ssh $USER@$HOST $START_COMMAND
