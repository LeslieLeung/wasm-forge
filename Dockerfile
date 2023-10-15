FROM golang:1.21 AS builder
WORKDIR $GOPATH/src/packages/wasm-forge
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/wasm-forge cmd/wasm-forge/main.go

FROM golang:1.21
WORKDIR /app
COPY --from=builder /go/wasm-forge /app/wasm-forge
ENV PORT 9000
ENV PATH $PATH:/usr/local/tinygo/bin
RUN mkdir runtime
EXPOSE 9000

RUN wget https://github.com/tinygo-org/tinygo/releases/download/v0.30.0/tinygo0.30.0.linux-amd64.tar.gz && \
    tar -xzf tinygo0.30.0.linux-amd64.tar.gz && \
    mv tinygo /usr/local/tinygo && \
    rm -rf tinygo && \
    rm tinygo0.30.0.linux-amd64.tar.gz


CMD ["/app/wasm-forge", "-m", "api"]