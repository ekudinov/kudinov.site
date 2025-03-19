FROM golang:alpine AS builder

COPY . .

RUN mkdir /build && \
    GOARCH=wasm GOOS=js go build -o /build/web/app.wasm && \
    go build -o /build/server && \
    cp -Rf assets /build/web

FROM alpine

WORKDIR /app

COPY --from=builder /build /app

CMD ["./server"]

