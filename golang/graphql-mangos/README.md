# GraphQL over nanomsg (go-mangos)

GraphQL over nanomsg, instead of normal HTTP.
This hosts the Star Wars graphql [testutil](https://github.com/graphql-go/graphql/blob/master/testutil/testutil.go)

## Server

```
go run server/server.go tcp://127.0.0.1:1337
```

## Client

```
go run client/client.go tcp://127.0.0.1:1337 '{hero{name}}'
{"data":{"hero":{"name":"R2-D2"}}}

go run client/client.go tcp://127.0.0.1:1337 '{hero(episode:EMPIRE){name}}'
{"data":{"hero":{"name":"Luke Skywalker"}}}

go run client/client.go tcp://127.0.0.1:1337 '{hero{name friends{name}}}'
{"data":{"hero":{"friends":[{"name":"Luke Skywalker"},{"name":"Han Solo"},{"name":"Leia Organa"}],"name":"R2-D2"}}}
```
