FROM php:8.1-cli

RUN apt-get update && apt-get install -y \
    git \
    unzip \
    protobuf-compiler \
    cmake \
    build-essential \
    autoconf \
    libtool \
    pkg-config \
    wget

# Install Go for ARM64
RUN wget https://go.dev/dl/go1.21.6.linux-arm64.tar.gz \
    && tar -C /usr/local -xzf go1.21.6.linux-arm64.tar.gz \
    && rm go1.21.6.linux-arm64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
ENV CGO_ENABLED=1

# Install PHP Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# Install protobuf PHP runtime and grpc
RUN pecl install protobuf grpc
RUN docker-php-ext-enable protobuf grpc

# Install protoc PHP plugin using CMake build system
RUN git clone https://github.com/grpc/grpc.git /var/local/git/grpc \
    && cd /var/local/git/grpc \
    && git submodule update --init \
    && mkdir -p cmake/build \
    && cd cmake/build \
    && cmake ../.. \
    && make protoc plugins

# Install protoc-gen-validate
RUN git clone https://github.com/envoyproxy/protoc-gen-validate.git /var/local/git/protoc-gen-validate \
    && cd /var/local/git/protoc-gen-validate \
    && make build

WORKDIR /proto

# Create an entrypoint script that includes validate proto path
RUN echo '#!/bin/sh\n\
find /proto -name "*.proto" -type f | while read -r file; do\n\
    protoc -I/proto \\\n\
        -I/var/local/git/protoc-gen-validate \\\n\
        --php_out=/genproto \\\n\
        --grpc_out=/genproto \\\n\
        --plugin=protoc-gen-grpc=/var/local/git/grpc/cmake/build/grpc_php_plugin \\\n\
        "$file"\n\
done' > /entrypoint.sh && chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
