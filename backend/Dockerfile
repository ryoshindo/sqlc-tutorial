FROM golang:1.20

RUN apt-get -y update && apt-get -y install postgresql-client

WORKDIR /go/src/github.com/ryoshindo/sqlc-tutorial/backend

ADD go.mod go.sum ./

RUN go install \
    github.com/99designs/gqlgen \
    github.com/cosmtrek/air
RUN go install github.com/rubenv/sql-migrate/sql-migrate@v1.2.0
