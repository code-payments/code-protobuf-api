// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file micropayment/v1/micro_payment_service.proto (package code.micropayment.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetStatusRequest, GetStatusResponse, RegisterWebhookRequest, RegisterWebhookResponse } from "./micro_payment_service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * todo: Migrate this to a generic "request" service
 *
 * @generated from service code.micropayment.v1.MicroPayment
 */
export const MicroPayment = {
  typeName: "code.micropayment.v1.MicroPayment",
  methods: {
    /**
     * GetStatus gets basic request status
     *
     * @generated from rpc code.micropayment.v1.MicroPayment.GetStatus
     */
    getStatus: {
      name: "GetStatus",
      I: GetStatusRequest,
      O: GetStatusResponse,
      kind: MethodKind.Unary,
    },
    /**
     * RegisterWebhook registers a webhook for a request
     *
     * todo: Once Kik codes can encode the entire payment request details, we can
     *       remove the messaging service component and have a Create RPC that
     *       reserves the intent ID with payment details, plus registers the webhook
     *       at the same time. Until that's possible, we're stuck with two RPC calls.
     *
     * @generated from rpc code.micropayment.v1.MicroPayment.RegisterWebhook
     */
    registerWebhook: {
      name: "RegisterWebhook",
      I: RegisterWebhookRequest,
      O: RegisterWebhookResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

