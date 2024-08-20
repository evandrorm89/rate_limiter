FROM golang:1.22.5 as build

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd
RUN go test -v ./...

FROM scratch
WORKDIR /app
COPY --from=build /app/app .
COPY --from=build /app/.env .

EXPOSE 8080

CMD ["./app"]
