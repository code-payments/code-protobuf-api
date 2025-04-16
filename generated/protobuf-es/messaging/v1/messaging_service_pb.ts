// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file messaging/v1/messaging_service.proto (package code.messaging.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message as Message$1, proto3, Timestamp } from "@bufbuild/protobuf";
import { ClientPong, Domain, IntentId, ServerPing, Signature, SolanaAccountId } from "../../common/v1/model_pb";
import { AdditionalFeePayment, ExchangeData, ExchangeDataWithoutRate, Metadata } from "../../transaction/v2/transaction_service_pb";

/**
 * @generated from message code.messaging.v1.OpenMessageStreamRequest
 */
export class OpenMessageStreamRequest extends Message$1<OpenMessageStreamRequest> {
  /**
   * @generated from field: code.messaging.v1.RendezvousKey rendezvous_key = 1;
   */
  rendezvousKey?: RendezvousKey;

  /**
   * The signature is of serialize(OpenMessageStreamRequest) using rendezvous_key.
   *
   * todo: Make required once clients migrate
   *
   * @generated from field: code.common.v1.Signature signature = 2;
   */
  signature?: Signature;

  constructor(data?: PartialMessage<OpenMessageStreamRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.OpenMessageStreamRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "rendezvous_key", kind: "message", T: RendezvousKey },
    { no: 2, name: "signature", kind: "message", T: Signature },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OpenMessageStreamRequest {
    return new OpenMessageStreamRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OpenMessageStreamRequest {
    return new OpenMessageStreamRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OpenMessageStreamRequest {
    return new OpenMessageStreamRequest().fromJsonString(jsonString, options);
  }

  static equals(a: OpenMessageStreamRequest | PlainMessage<OpenMessageStreamRequest> | undefined, b: OpenMessageStreamRequest | PlainMessage<OpenMessageStreamRequest> | undefined): boolean {
    return proto3.util.equals(OpenMessageStreamRequest, a, b);
  }
}

/**
 * @generated from message code.messaging.v1.OpenMessageStreamResponse
 */
export class OpenMessageStreamResponse extends Message$1<OpenMessageStreamResponse> {
  /**
   * @generated from field: repeated code.messaging.v1.Message messages = 1;
   */
  messages: Message[] = [];

  constructor(data?: PartialMessage<OpenMessageStreamResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.OpenMessageStreamResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "messages", kind: "message", T: Message, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OpenMessageStreamResponse {
    return new OpenMessageStreamResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OpenMessageStreamResponse {
    return new OpenMessageStreamResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OpenMessageStreamResponse {
    return new OpenMessageStreamResponse().fromJsonString(jsonString, options);
  }

  static equals(a: OpenMessageStreamResponse | PlainMessage<OpenMessageStreamResponse> | undefined, b: OpenMessageStreamResponse | PlainMessage<OpenMessageStreamResponse> | undefined): boolean {
    return proto3.util.equals(OpenMessageStreamResponse, a, b);
  }
}

/**
 * @generated from message code.messaging.v1.OpenMessageStreamWithKeepAliveRequest
 */
export class OpenMessageStreamWithKeepAliveRequest extends Message$1<OpenMessageStreamWithKeepAliveRequest> {
  /**
   * @generated from oneof code.messaging.v1.OpenMessageStreamWithKeepAliveRequest.request_or_pong
   */
  requestOrPong: {
    /**
     * @generated from field: code.messaging.v1.OpenMessageStreamRequest request = 1;
     */
    value: OpenMessageStreamRequest;
    case: "request";
  } | {
    /**
     * @generated from field: code.common.v1.ClientPong pong = 2;
     */
    value: ClientPong;
    case: "pong";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<OpenMessageStreamWithKeepAliveRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.OpenMessageStreamWithKeepAliveRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "request", kind: "message", T: OpenMessageStreamRequest, oneof: "request_or_pong" },
    { no: 2, name: "pong", kind: "message", T: ClientPong, oneof: "request_or_pong" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OpenMessageStreamWithKeepAliveRequest {
    return new OpenMessageStreamWithKeepAliveRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OpenMessageStreamWithKeepAliveRequest {
    return new OpenMessageStreamWithKeepAliveRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OpenMessageStreamWithKeepAliveRequest {
    return new OpenMessageStreamWithKeepAliveRequest().fromJsonString(jsonString, options);
  }

  static equals(a: OpenMessageStreamWithKeepAliveRequest | PlainMessage<OpenMessageStreamWithKeepAliveRequest> | undefined, b: OpenMessageStreamWithKeepAliveRequest | PlainMessage<OpenMessageStreamWithKeepAliveRequest> | undefined): boolean {
    return proto3.util.equals(OpenMessageStreamWithKeepAliveRequest, a, b);
  }
}

/**
 * @generated from message code.messaging.v1.OpenMessageStreamWithKeepAliveResponse
 */
export class OpenMessageStreamWithKeepAliveResponse extends Message$1<OpenMessageStreamWithKeepAliveResponse> {
  /**
   * @generated from oneof code.messaging.v1.OpenMessageStreamWithKeepAliveResponse.response_or_ping
   */
  responseOrPing: {
    /**
     * @generated from field: code.messaging.v1.OpenMessageStreamResponse response = 1;
     */
    value: OpenMessageStreamResponse;
    case: "response";
  } | {
    /**
     * @generated from field: code.common.v1.ServerPing ping = 2;
     */
    value: ServerPing;
    case: "ping";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<OpenMessageStreamWithKeepAliveResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.OpenMessageStreamWithKeepAliveResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "response", kind: "message", T: OpenMessageStreamResponse, oneof: "response_or_ping" },
    { no: 2, name: "ping", kind: "message", T: ServerPing, oneof: "response_or_ping" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OpenMessageStreamWithKeepAliveResponse {
    return new OpenMessageStreamWithKeepAliveResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OpenMessageStreamWithKeepAliveResponse {
    return new OpenMessageStreamWithKeepAliveResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OpenMessageStreamWithKeepAliveResponse {
    return new OpenMessageStreamWithKeepAliveResponse().fromJsonString(jsonString, options);
  }

  static equals(a: OpenMessageStreamWithKeepAliveResponse | PlainMessage<OpenMessageStreamWithKeepAliveResponse> | undefined, b: OpenMessageStreamWithKeepAliveResponse | PlainMessage<OpenMessageStreamWithKeepAliveResponse> | undefined): boolean {
    return proto3.util.equals(OpenMessageStreamWithKeepAliveResponse, a, b);
  }
}

/**
 * @generated from message code.messaging.v1.PollMessagesRequest
 */
export class PollMessagesRequest extends Message$1<PollMessagesRequest> {
  /**
   * @generated from field: code.messaging.v1.RendezvousKey rendezvous_key = 1;
   */
  rendezvousKey?: RendezvousKey;

  /**
   * The signature is of serialize(PollMessagesRequest) using rendezvous_key.
   *
   * @generated from field: code.common.v1.Signature signature = 2;
   */
  signature?: Signature;

  constructor(data?: PartialMessage<PollMessagesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.PollMessagesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "rendezvous_key", kind: "message", T: RendezvousKey },
    { no: 2, name: "signature", kind: "message", T: Signature },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PollMessagesRequest {
    return new PollMessagesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PollMessagesRequest {
    return new PollMessagesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PollMessagesRequest {
    return new PollMessagesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: PollMessagesRequest | PlainMessage<PollMessagesRequest> | undefined, b: PollMessagesRequest | PlainMessage<PollMessagesRequest> | undefined): boolean {
    return proto3.util.equals(PollMessagesRequest, a, b);
  }
}

/**
 * @generated from message code.messaging.v1.PollMessagesResponse
 */
export class PollMessagesResponse extends Message$1<PollMessagesResponse> {
  /**
   * @generated from field: repeated code.messaging.v1.Message messages = 1;
   */
  messages: Message[] = [];

  constructor(data?: PartialMessage<PollMessagesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.PollMessagesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "messages", kind: "message", T: Message, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PollMessagesResponse {
    return new PollMessagesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PollMessagesResponse {
    return new PollMessagesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PollMessagesResponse {
    return new PollMessagesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: PollMessagesResponse | PlainMessage<PollMessagesResponse> | undefined, b: PollMessagesResponse | PlainMessage<PollMessagesResponse> | undefined): boolean {
    return proto3.util.equals(PollMessagesResponse, a, b);
  }
}

/**
 * @generated from message code.messaging.v1.AckMessagesRequest
 */
export class AckMessagesRequest extends Message$1<AckMessagesRequest> {
  /**
   * @generated from field: code.messaging.v1.RendezvousKey rendezvous_key = 1;
   */
  rendezvousKey?: RendezvousKey;

  /**
   * @generated from field: repeated code.messaging.v1.MessageId message_ids = 2;
   */
  messageIds: MessageId[] = [];

  constructor(data?: PartialMessage<AckMessagesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.AckMessagesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "rendezvous_key", kind: "message", T: RendezvousKey },
    { no: 2, name: "message_ids", kind: "message", T: MessageId, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AckMessagesRequest {
    return new AckMessagesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AckMessagesRequest {
    return new AckMessagesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AckMessagesRequest {
    return new AckMessagesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AckMessagesRequest | PlainMessage<AckMessagesRequest> | undefined, b: AckMessagesRequest | PlainMessage<AckMessagesRequest> | undefined): boolean {
    return proto3.util.equals(AckMessagesRequest, a, b);
  }
}

/**
 * @generated from message code.messaging.v1.AckMesssagesResponse
 */
export class AckMesssagesResponse extends Message$1<AckMesssagesResponse> {
  /**
   * @generated from field: code.messaging.v1.AckMesssagesResponse.Result result = 1;
   */
  result = AckMesssagesResponse_Result.OK;

  constructor(data?: PartialMessage<AckMesssagesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.AckMesssagesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(AckMesssagesResponse_Result) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AckMesssagesResponse {
    return new AckMesssagesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AckMesssagesResponse {
    return new AckMesssagesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AckMesssagesResponse {
    return new AckMesssagesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: AckMesssagesResponse | PlainMessage<AckMesssagesResponse> | undefined, b: AckMesssagesResponse | PlainMessage<AckMesssagesResponse> | undefined): boolean {
    return proto3.util.equals(AckMesssagesResponse, a, b);
  }
}

/**
 * @generated from enum code.messaging.v1.AckMesssagesResponse.Result
 */
export enum AckMesssagesResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,
}
// Retrieve enum metadata with: proto3.getEnumType(AckMesssagesResponse_Result)
proto3.util.setEnumType(AckMesssagesResponse_Result, "code.messaging.v1.AckMesssagesResponse.Result", [
  { no: 0, name: "OK" },
]);

/**
 * @generated from message code.messaging.v1.SendMessageRequest
 */
export class SendMessageRequest extends Message$1<SendMessageRequest> {
  /**
   * The message to send. Types of messages clients can send are restricted.
   *
   * @generated from field: code.messaging.v1.Message message = 1;
   */
  message?: Message;

  /**
   * The rendezvous key that the message should be routed to.
   *
   * @generated from field: code.messaging.v1.RendezvousKey rendezvous_key = 2;
   */
  rendezvousKey?: RendezvousKey;

  /**
   * The signature is of serialize(Message) using the PrivateKey of the keypair.
   *
   * @generated from field: code.common.v1.Signature signature = 3;
   */
  signature?: Signature;

  constructor(data?: PartialMessage<SendMessageRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.SendMessageRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "message", T: Message },
    { no: 2, name: "rendezvous_key", kind: "message", T: RendezvousKey },
    { no: 3, name: "signature", kind: "message", T: Signature },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendMessageRequest {
    return new SendMessageRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendMessageRequest {
    return new SendMessageRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendMessageRequest {
    return new SendMessageRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SendMessageRequest | PlainMessage<SendMessageRequest> | undefined, b: SendMessageRequest | PlainMessage<SendMessageRequest> | undefined): boolean {
    return proto3.util.equals(SendMessageRequest, a, b);
  }
}

/**
 * @generated from message code.messaging.v1.SendMessageResponse
 */
export class SendMessageResponse extends Message$1<SendMessageResponse> {
  /**
   * @generated from field: code.messaging.v1.SendMessageResponse.Result result = 1;
   */
  result = SendMessageResponse_Result.OK;

  /**
   * Set if result == OK.
   *
   * @generated from field: code.messaging.v1.MessageId message_id = 2;
   */
  messageId?: MessageId;

  constructor(data?: PartialMessage<SendMessageResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.SendMessageResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(SendMessageResponse_Result) },
    { no: 2, name: "message_id", kind: "message", T: MessageId },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendMessageResponse {
    return new SendMessageResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendMessageResponse {
    return new SendMessageResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendMessageResponse {
    return new SendMessageResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SendMessageResponse | PlainMessage<SendMessageResponse> | undefined, b: SendMessageResponse | PlainMessage<SendMessageResponse> | undefined): boolean {
    return proto3.util.equals(SendMessageResponse, a, b);
  }
}

/**
 * @generated from enum code.messaging.v1.SendMessageResponse.Result
 */
export enum SendMessageResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: NO_ACTIVE_STREAM = 1;
   */
  NO_ACTIVE_STREAM = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(SendMessageResponse_Result)
proto3.util.setEnumType(SendMessageResponse_Result, "code.messaging.v1.SendMessageResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "NO_ACTIVE_STREAM" },
]);

/**
 * RendezvousKey is a unique key pair, typically derived from a scan code payload,
 * which is used to establish a secure communication channel anonymously to coordinate
 * a flow using messages.
 *
 * @generated from message code.messaging.v1.RendezvousKey
 */
export class RendezvousKey extends Message$1<RendezvousKey> {
  /**
   * @generated from field: bytes value = 1;
   */
  value = new Uint8Array(0);

  constructor(data?: PartialMessage<RendezvousKey>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.RendezvousKey";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RendezvousKey {
    return new RendezvousKey().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RendezvousKey {
    return new RendezvousKey().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RendezvousKey {
    return new RendezvousKey().fromJsonString(jsonString, options);
  }

  static equals(a: RendezvousKey | PlainMessage<RendezvousKey> | undefined, b: RendezvousKey | PlainMessage<RendezvousKey> | undefined): boolean {
    return proto3.util.equals(RendezvousKey, a, b);
  }
}

/**
 * MessageId identifies a message. It is only guaranteed to be unique when
 * paired with a destination (i.e. the rendezvous public key).
 *
 * @generated from message code.messaging.v1.MessageId
 */
export class MessageId extends Message$1<MessageId> {
  /**
   * @generated from field: bytes value = 1;
   */
  value = new Uint8Array(0);

  constructor(data?: PartialMessage<MessageId>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.MessageId";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MessageId {
    return new MessageId().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MessageId {
    return new MessageId().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MessageId {
    return new MessageId().fromJsonString(jsonString, options);
  }

  static equals(a: MessageId | PlainMessage<MessageId> | undefined, b: MessageId | PlainMessage<MessageId> | undefined): boolean {
    return proto3.util.equals(MessageId, a, b);
  }
}

/**
 * Request that a pulled out bill be sent to the requested address.
 *
 * This message type is only initiated by clients.
 *
 * @generated from message code.messaging.v1.RequestToGrabBill
 */
export class RequestToGrabBill extends Message$1<RequestToGrabBill> {
  /**
   * Requestor is the virtual core mint token account on the VM to which a
   * payment should be sent.
   *
   * @generated from field: code.common.v1.SolanaAccountId requestor_account = 1;
   */
  requestorAccount?: SolanaAccountId;

  constructor(data?: PartialMessage<RequestToGrabBill>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.RequestToGrabBill";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "requestor_account", kind: "message", T: SolanaAccountId },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequestToGrabBill {
    return new RequestToGrabBill().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequestToGrabBill {
    return new RequestToGrabBill().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequestToGrabBill {
    return new RequestToGrabBill().fromJsonString(jsonString, options);
  }

  static equals(a: RequestToGrabBill | PlainMessage<RequestToGrabBill> | undefined, b: RequestToGrabBill | PlainMessage<RequestToGrabBill> | undefined): boolean {
    return proto3.util.equals(RequestToGrabBill, a, b);
  }
}

/**
 * Request that a bill of a requested value is created and sent to the requested
 * address.
 *
 * This message type is only initiated by clients.
 *
 * @generated from message code.messaging.v1.RequestToReceiveBill
 */
export class RequestToReceiveBill extends Message$1<RequestToReceiveBill> {
  /**
   * Requestor is the virtual core mint token account on the VM to which a
   * payment should be sent.
   *
   * @generated from field: code.common.v1.SolanaAccountId requestor_account = 1;
   */
  requestorAccount?: SolanaAccountId;

  /**
   * The exchange data for the requested bill value.
   *
   * @generated from oneof code.messaging.v1.RequestToReceiveBill.exchange_data
   */
  exchangeData: {
    /**
     * An exact amount of core mint tokens. Payment is guaranteed to transfer the specified
     * quarks in the requested currency and exchange rate.
     *
     * Only supports the core mint account. Use exchange_data.partial for fiat amounts.
     *
     * @generated from field: code.transaction.v2.ExchangeData exact = 2;
     */
    value: ExchangeData;
    case: "exact";
  } | {
    /**
     * Fiat amount request. The amount of core mint tokens are determined at
     * time of payment with a recent exchange rate provided by the paying client
     * and validated by server.
     *
     * Only supports fiat amounts. Use exchange_data.exact for the core mint account.
     *
     * @generated from field: code.transaction.v2.ExchangeDataWithoutRate partial = 3;
     */
    value: ExchangeDataWithoutRate;
    case: "partial";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * The third-party's domain name, which is its primary identifier. Server
   * guarantees to perform domain verification against the verifier account.
   *
   * @generated from field: code.common.v1.Domain domain = 4;
   */
  domain?: Domain;

  /**
   * Owner account owned by the third party used in domain verification.
   *
   * @generated from field: code.common.v1.SolanaAccountId verifier = 5;
   */
  verifier?: SolanaAccountId;

  /**
   * Signature of this message using the verifier private key, which in addition
   * to domain verification, authenticates the third party.
   *
   * @generated from field: code.common.v1.Signature signature = 6;
   */
  signature?: Signature;

  /**
   * Rendezvous key to avoid replay attacks
   *
   * @generated from field: code.messaging.v1.RendezvousKey rendezvous_key = 7;
   */
  rendezvousKey?: RendezvousKey;

  /**
   * Additional fee payments splitting the requested amount. This is in addition
   * to the hard-coded Code $0.01 USD fee.
   *
   * @generated from field: repeated code.transaction.v2.AdditionalFeePayment additional_fees = 8;
   */
  additionalFees: AdditionalFeePayment[] = [];

  constructor(data?: PartialMessage<RequestToReceiveBill>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.RequestToReceiveBill";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "requestor_account", kind: "message", T: SolanaAccountId },
    { no: 2, name: "exact", kind: "message", T: ExchangeData, oneof: "exchange_data" },
    { no: 3, name: "partial", kind: "message", T: ExchangeDataWithoutRate, oneof: "exchange_data" },
    { no: 4, name: "domain", kind: "message", T: Domain },
    { no: 5, name: "verifier", kind: "message", T: SolanaAccountId },
    { no: 6, name: "signature", kind: "message", T: Signature },
    { no: 7, name: "rendezvous_key", kind: "message", T: RendezvousKey },
    { no: 8, name: "additional_fees", kind: "message", T: AdditionalFeePayment, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequestToReceiveBill {
    return new RequestToReceiveBill().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequestToReceiveBill {
    return new RequestToReceiveBill().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequestToReceiveBill {
    return new RequestToReceiveBill().fromJsonString(jsonString, options);
  }

  static equals(a: RequestToReceiveBill | PlainMessage<RequestToReceiveBill> | undefined, b: RequestToReceiveBill | PlainMessage<RequestToReceiveBill> | undefined): boolean {
    return proto3.util.equals(RequestToReceiveBill, a, b);
  }
}

/**
 * A status update on a stream to indicate a scan code was scanned. This can appear
 * multiple times for the same stream.
 *
 * This message type is only initiated by client
 *
 * @generated from message code.messaging.v1.CodeScanned
 */
export class CodeScanned extends Message$1<CodeScanned> {
  /**
   * Timestamp the client scanned the code
   *
   * @generated from field: google.protobuf.Timestamp timestamp = 1;
   */
  timestamp?: Timestamp;

  constructor(data?: PartialMessage<CodeScanned>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.CodeScanned";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "timestamp", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CodeScanned {
    return new CodeScanned().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CodeScanned {
    return new CodeScanned().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CodeScanned {
    return new CodeScanned().fromJsonString(jsonString, options);
  }

  static equals(a: CodeScanned | PlainMessage<CodeScanned> | undefined, b: CodeScanned | PlainMessage<CodeScanned> | undefined): boolean {
    return proto3.util.equals(CodeScanned, a, b);
  }
}

/**
 * Payment is rejected by the client
 *
 * This message type is only initiated by clients
 *
 * @generated from message code.messaging.v1.ClientRejectedPayment
 */
export class ClientRejectedPayment extends Message$1<ClientRejectedPayment> {
  /**
   * @generated from field: code.common.v1.IntentId intent_id = 1;
   */
  intentId?: IntentId;

  constructor(data?: PartialMessage<ClientRejectedPayment>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.ClientRejectedPayment";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "intent_id", kind: "message", T: IntentId },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClientRejectedPayment {
    return new ClientRejectedPayment().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClientRejectedPayment {
    return new ClientRejectedPayment().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClientRejectedPayment {
    return new ClientRejectedPayment().fromJsonString(jsonString, options);
  }

  static equals(a: ClientRejectedPayment | PlainMessage<ClientRejectedPayment> | undefined, b: ClientRejectedPayment | PlainMessage<ClientRejectedPayment> | undefined): boolean {
    return proto3.util.equals(ClientRejectedPayment, a, b);
  }
}

/**
 * Intent was submitted via SubmitIntent
 *
 * This message type is only initiated by server
 *
 * @generated from message code.messaging.v1.IntentSubmitted
 */
export class IntentSubmitted extends Message$1<IntentSubmitted> {
  /**
   * @generated from field: code.common.v1.IntentId intent_id = 1;
   */
  intentId?: IntentId;

  /**
   * Metadata is available for intents where it can be safely propagated publicly.
   * Anything else requires an additional authenticated RPC call (eg. login).
   *
   * @generated from field: code.transaction.v2.Metadata metadata = 2;
   */
  metadata?: Metadata;

  constructor(data?: PartialMessage<IntentSubmitted>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.IntentSubmitted";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "intent_id", kind: "message", T: IntentId },
    { no: 2, name: "metadata", kind: "message", T: Metadata },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IntentSubmitted {
    return new IntentSubmitted().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IntentSubmitted {
    return new IntentSubmitted().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IntentSubmitted {
    return new IntentSubmitted().fromJsonString(jsonString, options);
  }

  static equals(a: IntentSubmitted | PlainMessage<IntentSubmitted> | undefined, b: IntentSubmitted | PlainMessage<IntentSubmitted> | undefined): boolean {
    return proto3.util.equals(IntentSubmitted, a, b);
  }
}

/**
 * Webhook was successfully called
 *
 * This message type is only initiated by server
 *
 * @generated from message code.messaging.v1.WebhookCalled
 */
export class WebhookCalled extends Message$1<WebhookCalled> {
  /**
   * Estimated time webhook was received
   *
   * @generated from field: google.protobuf.Timestamp timestamp = 1;
   */
  timestamp?: Timestamp;

  constructor(data?: PartialMessage<WebhookCalled>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.WebhookCalled";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "timestamp", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WebhookCalled {
    return new WebhookCalled().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WebhookCalled {
    return new WebhookCalled().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WebhookCalled {
    return new WebhookCalled().fromJsonString(jsonString, options);
  }

  static equals(a: WebhookCalled | PlainMessage<WebhookCalled> | undefined, b: WebhookCalled | PlainMessage<WebhookCalled> | undefined): boolean {
    return proto3.util.equals(WebhookCalled, a, b);
  }
}

/**
 * @generated from message code.messaging.v1.Message
 */
export class Message extends Message$1<Message> {
  /**
   * MessageId is the Id of the message. This ID is generated by the
   * server, and will _always_ be set when receiving a message.
   *
   * Server generates the message to:
   *     1. Reserve the ability for any future ID changes
   *     2. Prevent clients attempting to collide message IDs.
   *
   * @generated from field: code.messaging.v1.MessageId id = 1;
   */
  id?: MessageId;

  /**
   * The signature sent from SendMessageRequest, which will be injected by server.
   * This enables clients to ensure no MITM attacks were performed to hijack contents
   * of the typed message. This is only applicable for messages not generated by server.
   *
   * @generated from field: code.common.v1.Signature send_message_request_signature = 3;
   */
  sendMessageRequestSignature?: Signature;

  /**
   * Next field number is 13
   *
   * @generated from oneof code.messaging.v1.Message.kind
   */
  kind: {
    /**
     * @generated from field: code.messaging.v1.RequestToGrabBill request_to_grab_bill = 2;
     */
    value: RequestToGrabBill;
    case: "requestToGrabBill";
  } | {
    /**
     * @generated from field: code.messaging.v1.RequestToReceiveBill request_to_receive_bill = 5;
     */
    value: RequestToReceiveBill;
    case: "requestToReceiveBill";
  } | {
    /**
     * @generated from field: code.messaging.v1.CodeScanned code_scanned = 6;
     */
    value: CodeScanned;
    case: "codeScanned";
  } | {
    /**
     * @generated from field: code.messaging.v1.ClientRejectedPayment client_rejected_payment = 7;
     */
    value: ClientRejectedPayment;
    case: "clientRejectedPayment";
  } | {
    /**
     * @generated from field: code.messaging.v1.IntentSubmitted intent_submitted = 8;
     */
    value: IntentSubmitted;
    case: "intentSubmitted";
  } | {
    /**
     * @generated from field: code.messaging.v1.WebhookCalled webhook_called = 9;
     */
    value: WebhookCalled;
    case: "webhookCalled";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Message>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "code.messaging.v1.Message";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: MessageId },
    { no: 3, name: "send_message_request_signature", kind: "message", T: Signature },
    { no: 2, name: "request_to_grab_bill", kind: "message", T: RequestToGrabBill, oneof: "kind" },
    { no: 5, name: "request_to_receive_bill", kind: "message", T: RequestToReceiveBill, oneof: "kind" },
    { no: 6, name: "code_scanned", kind: "message", T: CodeScanned, oneof: "kind" },
    { no: 7, name: "client_rejected_payment", kind: "message", T: ClientRejectedPayment, oneof: "kind" },
    { no: 8, name: "intent_submitted", kind: "message", T: IntentSubmitted, oneof: "kind" },
    { no: 9, name: "webhook_called", kind: "message", T: WebhookCalled, oneof: "kind" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Message {
    return new Message().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Message {
    return new Message().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Message {
    return new Message().fromJsonString(jsonString, options);
  }

  static equals(a: Message | PlainMessage<Message> | undefined, b: Message | PlainMessage<Message> | undefined): boolean {
    return proto3.util.equals(Message, a, b);
  }
}

