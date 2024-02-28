### `MessagingService.md`

# Messaging Service Documentation

## Overview
This document aims to provide a comprehensive understanding of the messaging service's operations, including message streaming, payment flows, login procedures, and keepalive mechanisms.

The Messaging Service facilitates secure, real-time communication between clients and servers for payment processing, login procedures, and general messaging within the Code platform. This document outlines the service's capabilities, including message streaming, acknowledging messages, and the specific flows for payment processing and user login.

## Service RPCs

### `OpenMessageStream`

Opens a stream for real-time message communication. Utilised for various operations including payment initiation and receiving, login challenges, and more.

#### Flow:

1. **Message Delivery**: Messages routed using a rendezvous keypair's public key, derived by both sender and recipient.
2. **Message Acknowledgement**: Messages must be acknowledged by the client to be marked as processed and are then eligible for deletion.

#### Payment Processing Flow:

- **Sending a Payment**:
  1. Payment sender creates a cash scan code.
  2. Sender initiates `OpenMessageStream` with the derived rendezvous public key.
  3. Payment recipient scans the code, sends account ID back to the sender.
  4. Sender submits payment intent and closes the stream.

- **Receiving a Payment Request**:
  1. Payment recipient sends account ID and payment amount to sender.
  2. Recipient listens for status messages to monitor payment state.
  3. Process concludes with payment submission or timeout.

#### Login Flow:

1. Third-party sends login challenge via `SendMessage`.
2. Third-party listens for status messages indicating login success or failure.
3. Stream is closed once the login process concludes.

### `OpenMessageStreamWithKeepAlive`

Similar to `OpenMessageStream` but includes a keepalive mechanism to monitor stream health.

#### Keepalive Protocol:

1. Server sends pings; client responds with pongs.
2. Protocol ensures stream health, allowing for real-time updates.

### `PollMessages`

Polls for messages, providing an alternative to real-time streaming. Suitable for environments where long-lived connections are not feasible.

### `AckMessages`

Acknowledges receipt of one or more messages, marking them as processed.

### `SendMessage`

Sends a message through the established stream. Supports various message types including payment requests, login attempts, and more.

## Detailed Message Types

This section documents the specific message types supported by the service, their intended use cases, and any special considerations for their use.

- **`RequestToGrabBill`**: Initiates a request for a bill to be sent to a specified address.
- **`RequestToReceiveBill`**: Requests the creation and sending of a bill to a specific address.
- **`CodeScanned`**: Indicates that a scan code was scanned, supporting multiple occurrences per stream.
- **`ClientRejectedPayment`**: Signifies a payment rejection by the client.
- **`IntentSubmitted`**: Indicates that a payment intent was submitted, accompanied by metadata when available.
- **`WebhookCalled`**: Confirms that a webhook call was successfully made.
- **`RequestToLogin`**: Initiates a login request by a third-party through the SDK.
- **`ClientRejectedLogin`**: Indicates a login rejection by the client.
- **`AirdropReceived`**: Notifies the client of an airdrop receipt, detailing the type and value.

## Additional Information

- **Rendezvous Keys**: Essential for establishing secure, anonymous communication channels.
- **Signature Verification**: Ensures message integrity and authenticity.

For further details on implementing these flows, refer to the SDK documentation and best practices for handling message delivery, acknowledgment, and error scenarios.