// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file currency/v1/currency_service.proto (package code.currency.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message code.currency.v1.GetAllRatesRequest
 */
export class GetAllRatesRequest extends Message<GetAllRatesRequest> {
  /**
   * If timestamp is included, the returned rate will be the most recent available
   * exchange rate prior to the provided timestamp within the same day. Otherwise,
   * the latest rates will be returned.
   *
   * @generated from field: google.protobuf.Timestamp timestamp = 1;
   */
  timestamp?: Timestamp;

  constructor(data?: PartialMessage<GetAllRatesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.currency.v1.GetAllRatesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "timestamp", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAllRatesRequest {
    return new GetAllRatesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAllRatesRequest {
    return new GetAllRatesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAllRatesRequest {
    return new GetAllRatesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetAllRatesRequest | PlainMessage<GetAllRatesRequest> | undefined, b: GetAllRatesRequest | PlainMessage<GetAllRatesRequest> | undefined): boolean {
    return proto3.util.equals(GetAllRatesRequest, a, b);
  }
}

/**
 * @generated from message code.currency.v1.GetAllRatesResponse
 */
export class GetAllRatesResponse extends Message<GetAllRatesResponse> {
  /**
   * @generated from field: code.currency.v1.GetAllRatesResponse.Result result = 1;
   */
  result = GetAllRatesResponse_Result.OK;

  /**
   * The time the exchange rates were observed
   *
   * @generated from field: google.protobuf.Timestamp as_of = 2;
   */
  asOf?: Timestamp;

  /**
   * The price of 1 core mint token in different currencies, keyed on 3- or 4-
   * letter lowercase currency code.
   *
   * @generated from field: map<string, double> rates = 3;
   */
  rates: { [key: string]: number } = {};

  constructor(data?: PartialMessage<GetAllRatesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.currency.v1.GetAllRatesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(GetAllRatesResponse_Result) },
    { no: 2, name: "as_of", kind: "message", T: Timestamp },
    { no: 3, name: "rates", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 1 /* ScalarType.DOUBLE */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAllRatesResponse {
    return new GetAllRatesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAllRatesResponse {
    return new GetAllRatesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAllRatesResponse {
    return new GetAllRatesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetAllRatesResponse | PlainMessage<GetAllRatesResponse> | undefined, b: GetAllRatesResponse | PlainMessage<GetAllRatesResponse> | undefined): boolean {
    return proto3.util.equals(GetAllRatesResponse, a, b);
  }
}

/**
 * @generated from enum code.currency.v1.GetAllRatesResponse.Result
 */
export enum GetAllRatesResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * No currency data is available for the requested timestamp.
   *
   * @generated from enum value: MISSING_DATA = 1;
   */
  MISSING_DATA = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(GetAllRatesResponse_Result)
proto3.util.setEnumType(GetAllRatesResponse_Result, "code.currency.v1.GetAllRatesResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "MISSING_DATA" },
]);

