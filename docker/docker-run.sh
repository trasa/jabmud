#!/bin/bash

docker run -d \
    -p 5222:5222 \
    -p 5269:5269 \
    -p 5275:5275 \
    -p 5280:5280 \
    --name jabmud-ejabberd \
    jabmud/ejabberd
