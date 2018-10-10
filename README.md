# go-graphql

## Prerequistes
- Prisma cli version 1.19.0-alpha21 installed
- Go installed
- gqlgen installed

## Getting started

```sh
~ $ git clone https://github.com/Sach97/prisma-go-hackernews2/edit/master/README.md 
~ $ ln -s ~/prisma-go-hackernews2 ~/go/src/prisma-go-hackernews2 # (convenient but not required see https://codebasecamp.com/2017/04/25/Project-Structure-Go.1/ for an explanation of symlinks )
~ $ cd prisma-go-hackernews
prisma-go-hackernews2 $ docker-compose up -d # (kill your docker container running with you have errors)
prisma-go-hackernews2 $ cd prisma
prisma $ prisma deploy
prisma $ cd ..
prisma-go-hackernews2 $ go run server/*.go # Install missing dependencies then rerun this command
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
