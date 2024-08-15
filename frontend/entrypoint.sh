#!/bin/bash
set -e

if [ $@ == "npm" ]
then
    npm run build
    exec npm run start
fi
exec $@
