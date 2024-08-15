#!/bin/bash
set -e

if [ $@ == "npm" ]
then
    npm run build
    exec npm run start -- -p 80
fi
exec $@
