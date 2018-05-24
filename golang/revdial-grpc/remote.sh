#!/bin/bash
# Execute the cicd-agent in a docker container and dial to the mothership
CID=`docker create --rm debian:latest /usr/local/bin/cicd-agent -host 0.tcp.ngrok.io:17210`
docker cp ./client $CID:/usr/local/bin/cicd-agent
docker start -i $CID
