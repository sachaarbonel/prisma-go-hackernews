# go-graphql

## Prerequistes
- Prisma cli version 1.19.0-alpha21 installed
- Go installed
- gqlgen installed

## Getting started

```sh
~ $ git clone https://github.com/Sach97/prisma-go-hackernews/edit/master/README.md 
~ $ cd prisma-go-hackernews
prisma-go-hackernews $ dep ensure
prisma-go-hackernews $ docker-compose up -d # (kill your docker container running with you have errors)
prisma-go-hackernews $ go run server/*.go # Install missing dependencies then rerun this command
```

## Development Workflow

- Modify your datamodel.prisma
- prisma deploy
- Modify your schema.graphql
- gqlgen -v (must be executed in your GOPATH)
- Modify your generated tmp/resolver with the prisma generated types
- run server

## TODOs
- [ ] Auth and permissions
