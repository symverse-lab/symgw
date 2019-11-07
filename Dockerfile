FROM golang:1.13
LABEL maintainer="gocheat <itsinil@gmail.com>"

WORKDIR /app/src/github.com/symverse-lab/symgw
COPY . .
ENV GOPATH="/app"
ENV PATH="$PATH:$GOPATH/bin"

# project default package update
RUN go get -u github.com/kardianos/govendor
RUN go get -u github.com/go-redis/redis
RUN go get -u github.com/cpuguy83/go-md2man && go get -u github.com/swaggo/swag/cmd/swag

RUN govendor sync
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./symgw", "--env"]

