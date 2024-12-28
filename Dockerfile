FROM golang:1.22


# Install buf
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/v1.8.0/buf-Linux-x86_64" -o /usr/local/bin/buf \
    && chmod +x /usr/local/bin/buf

# Install protoc-gen-grpc-gateway plugin
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

ENV PATH="/go/bin:${PATH}"

WORKDIR /workspace

# Copy the current directory contents into the container
COPY . /workspace

# Install Go dependencies if needed
RUN go mod tidy

CMD ["buf", "generate", "proto"]
