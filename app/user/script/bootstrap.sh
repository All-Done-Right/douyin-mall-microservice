#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/UserService"
exec "$CURDIR/bin/UserService"
