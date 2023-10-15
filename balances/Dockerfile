FROM golang:1.20.10-alpine3.17 as build
RUN apk update
WORKDIR /app
RUN apk add librdkafka-dev build-base
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -a -ldflags '-linkmode external -extldflags "-static" -s -w' -tags musl -v -o walletbalance ./cmd/walletbalance

FROM scratch as prod
COPY --from=build /app/walletbalance .
EXPOSE 3003
CMD ["./walletbalance"]