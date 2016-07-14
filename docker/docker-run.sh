#!/bin/bash

docker run -d \
    -p 4560:4560 \
    -p 5222:5222 \
    -p 5269:5269 \
    -p 5275:5275 \
    -p 5280:5280 \
    -p 5443:5443 \
    -e "EJABBERD_LOGLEVEL=5" \
    -e "EJABBERD_HTTPS=false" \
    -e "EJABBERD_USERS=jabmud_user@localhost:password" \
    --name jabmud-ejabberd \
       jabmud/ejabberd
