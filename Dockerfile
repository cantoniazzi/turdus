FROM golang:alpine AS env-builder

WORKDIR /turdus

COPY . .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags='-w -s -extldflags "-static"' -v -a -o /go/bin/turdus .

FROM alpine:latest  

COPY --from=env-builder /go/bin/turdus /go/bin/turdus

EXPOSE 9000

ENTRYPOINT ["go/bin/turdus"]