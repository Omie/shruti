#!/bin/bash

export SHRUTI_SERVER_HOST=0.0.0.0
export SHRUTI_SERVER_PORT=9574
export SHRUTI_CONN_STRING='postgres://postgres:@127.0.0.1:5432/postgres?search_path=testshruti&sslmode=disable'

export SHRUTI_PUSHER_APPID=131626
export SHRUTI_PUSHER_KEY=
export SHRUTI_PUSHER_SECRET=
export SHRUTI_PUSHER_CHANNEL=shrutitest
export SHRUTI_PUSHER_EVENT=new-notification

rm shruti
go build
./shruti

