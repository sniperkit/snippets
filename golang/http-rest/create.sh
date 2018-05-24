#!/bin/sh
echo "=== api/v1/token/info"; echo
curl	--header "X-Xg-Auth-User: jerry" \
	--header "X-Xg-Api-Token: awesome" \
	http://127.0.0.1:3000/api/v1/token/info
echo

echo "=== api/v1/token/list"; echo
curl	--header "X-Xg-Auth-User: jerry" \
	--header "X-Xg-Api-Token: awesome" \
	http://127.0.0.1:3000/api/v1/token/list
echo

echo "=== api/v1/token/create"; echo
curl	--header "X-Xg-Auth-User: jerry" \
	--header "X-Xg-Api-Token: awesome1337" \
	http://127.0.0.1:3000/api/v1/token/create
echo

echo "=== api/v1/token/delete"; echo
curl    --header "Content-Type: application/json" \
	-X POST -d '{"id":"awesome"}' \
	--header "X-Xg-Auth-User: jerry" \
	--header "X-Xg-Api-Token: awesome1337" \
	http://127.0.0.1:3000/api/v1/token/delete
echo

echo "=== api/v1/token/create (user: xyz)"; echo
curl    --header "Content-Type: application/json" \
	-X POST -d '{"user":"xyz"}' \
	--header "X-Xg-Auth-User: jerry" \
	--header "X-Xg-Api-Token: awesome1337" \
	http://127.0.0.1:3000/api/v1/token/create
echo
