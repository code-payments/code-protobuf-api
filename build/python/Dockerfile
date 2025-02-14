FROM python:3.11-slim

RUN apt-get update && apt-get install -y \
    git \
    protobuf-compiler \
    python3-pip

# Install Python gRPC tools and protobuf
RUN pip install grpcio grpcio-tools protobuf

# Clone protoc-gen-validate for proto includes only
RUN git clone --depth 1 https://github.com/envoyproxy/protoc-gen-validate.git /var/local/git/protoc-gen-validate

WORKDIR /proto

RUN echo '#!/bin/sh\n\
find /proto -name "*.proto" -type f | while read -r file; do\n\
    python -m grpc_tools.protoc \
        -I/proto \
        -I/var/local/git/protoc-gen-validate \
        --python_out=/genproto \
        --grpc_python_out=/genproto \
        "$file"\n\
done' > /entrypoint.sh && chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]