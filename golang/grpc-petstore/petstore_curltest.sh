#!/bin/sh
curl -s --request POST http://127.0.0.1:8080/v1/pets \
     --header "Content-Type: application/json" \
     --data '{"name":"sheep"}' > /dev/null

curl -s --request POST http://127.0.0.1:8080/v1/pets \
     --header "Content-Type: application/json" \
     --data '{"name":"cow"}' > /dev/null

curl -s --request GET http://127.0.0.1:8080/v1/pets
