FROM golang:1.18 as build

WORKDIR /authorization_service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine

WORKDIR /app

COPY --from=build /authorization_service ./

EXPOSE 80

CMD [ "./app" ]
