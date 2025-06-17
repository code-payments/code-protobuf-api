<img width="1280" alt="Code Protobuf API" src="https://github.com/code-payments/code-protobuf-api/assets/5760385/e43106cd-4c50-4dfa-954f-a936290b8b86">

# Code Protobuf API

[![Release](https://img.shields.io/github/v/release/code-payments/code-protobuf-api.svg)](https://github.com/code-payments/code-protobuf-api/releases/latest)
[![GitHub License](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/code-payments/code-protobuf-api/blob/main/LICENSE.md)

The APIs and models for communication between Code clients and server.

## Services

- [Account](https://github.com/code-payments/code-protobuf-api/blob/main/proto/account/v1/account_service.proto)
- [Currency](https://github.com/code-payments/code-protobuf-api/blob/main/proto/currency/v1/currency_service.proto)
- [Messaging](https://github.com/code-payments/code-protobuf-api/blob/main/proto/messaging/v1/messaging_service.proto)
- [Transaction](https://github.com/code-payments/code-protobuf-api/blob/main/proto/transaction/v2/transaction_service.proto)

## Code Generation

Generated code can be found under the `generated/` directory. The following languages are directly supported:
- Go
- TypeScript/JavaScript (via [Protobuf-ES](https://github.com/bufbuild/protobuf-es))

To generate all code, run:

```bash
make
```

Or to generate a specific language, run:

```bash
make {go, protobuf-es}
```

## Getting Help

If you have any questions or need help integrating Code into your website or application, please reach out to us on [Discord](https://discord.gg/T8Tpj8DBFp) or [Twitter](https://twitter.com/getcode).

##  Contributing

For now the best way to contribute is to share feedback on [Discord](https://discord.gg/T8Tpj8DBFp). This will evolve as we continue to build out the platform and open up more ways to contribute.
