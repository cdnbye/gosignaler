#!/bin/bash

SERVER="gosignaler"
SERVICE="$SERVER.service"
PREFIX="/lib/systemd/system"
BASE_DIR=$PWD
INTERVAL=2

# 命令行参数，需要手动指定
ARGS=""

function generateServiceFile()
{
echo "[Unit]
Description=Signaler

[Service]
Type=simple
LimitCORE=infinity
LimitNOFILE=1000000
LimitNPROC=1000000


WorkingDirectory=$BASE_DIR

ExecStart=$BASE_DIR/$SERVER

StandardOutput=null
StandardError=null

Restart=on-failure
ExecStop=

[Install]
WantedBy=multi-user.target" > ./$SERVICE
}

function deploy()
{
    echo "Generate Service File"
	generateServiceFile
	sudo cp $SERVICE $PREFIX
	sudo systemctl daemon-reload
	sudo systemctl enable $SERVICE
}

function start()
{
	deploy

	if [ "`pgrep $SERVER`" != "" ];then
		echo "$SERVER already running"
		exit 1
	fi

	sudo systemctl start $SERVICE
	echo "sleeping..." &&  sleep $INTERVAL

	# check status
	if [ "`pgrep $SERVER`" == "" ];then
		echo "$SERVER start failed"
		exit 1
	fi

	echo "$SERVER start succeed"
}

function status()
{
	if [ "`pgrep $SERVER`" != "" ];then
		echo $SERVER is running
	else
		echo $SERVER is not running
	fi
}

function stop()
{
	if [ "`pgrep $SERVER`" != "" ];then
		sudo systemctl stop $SERVICE
	fi

	echo "sleeping..." &&  sleep $INTERVAL

	if [ "`pgrep $SERVER`" != "" ];then
		echo "$SERVER stop failed"
		exit 1
	fi

	echo "$SERVER stop succeed"
}

function restart()
{
	sudo systemctl restart $SERVICE
	echo "sleeping..." &&  sleep $INTERVAL

	if [ "`pgrep $SERVER`" == "" ];then
		echo "$SERVER start failed"
		exit 1
	fi

	echo "$SERVER start succeed"
}

function test()
{
    if [ "`pgrep $SERVER -u $UID`" != "" ];then
        kill -9 `pgrep $SERVER -u $UID`
    fi

    echo "sleeping..." &&  sleep $INTERVAL

    if [ "`pgrep $SERVER -u $UID`" != "" ];then
        echo "$SERVER stop failed"
        exit 1
    fi

    echo "$SERVER stop succeed"

    ulimit -n 1000000

    nohup $BASE_DIR/$SERVER -c conf/test_config.yaml &>/dev/null &

    echo "sleeping..." &&  sleep $INTERVAL

    # check status
    if [ "`pgrep $SERVER -u $UID`" == "" ];then
        echo "$SERVER start failed"
        exit 1
    fi

    echo "$SERVER start succeed"
}

case "$1" in
	'start')
	start
	;;
	'stop')
	stop
	;;
	'status')
	status
	;;
	'restart')
	restart
	;;
	'test')
    test
    ;;
	*)
	echo "usage: $0 {start|stop|restart|status}"
	exit 1
	;;
esac
