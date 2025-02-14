<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Chat\V2;

/**
 */
class ChatClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * GetChats gets the set of chats for an owner account using a paged API.
     * This RPC is aware of all identities tied to the owner account.
     * @param \Code\Chat\V2\GetChatsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetChats(\Code\Chat\V2\GetChatsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v2.Chat/GetChats',
        $argument,
        ['\Code\Chat\V2\GetChatsResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetMessages gets the set of messages for a chat member using a paged API
     * @param \Code\Chat\V2\GetMessagesRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetMessages(\Code\Chat\V2\GetMessagesRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v2.Chat/GetMessages',
        $argument,
        ['\Code\Chat\V2\GetMessagesResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * StreamChatEvents streams chat events in real-time. Chat events include
     * messages, pointer updates, etc.
     *
     * The streaming protocol is follows:
     *  1. Client initiates a stream by sending an OpenChatEventStream message.
     *  2. If an error is encoutered, a ChatStreamEventError message will be
     *     returned by server and the stream will be closed.
     *  3. Server will immediately flush initial chat state.
     *  4. New chat events will be pushed to the stream in real time as they
     *     are received.
     *
     * This RPC supports a keepalive protocol as follows:
     *   1. Client initiates a stream by sending an OpenChatEventStream message.
     *   2. Upon stream initialization, server begins the keepalive protocol.
     *   3. Server sends a ping to the client.
     *   4. Client responds with a pong as fast as possible, making note of
     *      the delay for when to expect the next ping.
     *   5. Steps 3 and 4 are repeated until the stream is explicitly terminated
     *      or is deemed to be unhealthy.
     *
     * Client notes:
     * * Client should be careful to process events async, so any responses to pings are
     *   not delayed.
     * * Clients should implement a reasonable backoff strategy upon continued timeout
     *   failures.
     * * Clients that abuse pong messages may have their streams terminated by server.
     *
     * At any point in the stream, server will respond with events in real time as
     * they are observed. Events sent over the stream should not affect the ping/pong
     * protocol timings.
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function StreamChatEvents($metadata = [], $options = []) {
        return $this->_bidiRequest('/code.chat.v2.Chat/StreamChatEvents',
        ['\Code\Chat\V2\StreamChatEventsResponse','decode'],
        $metadata, $options);
    }

    /**
     * StartChat starts a chat. The RPC call is idempotent and will use existing
     * chats whenever applicable within the context of message routing.
     * @param \Code\Chat\V2\StartChatRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function StartChat(\Code\Chat\V2\StartChatRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v2.Chat/StartChat',
        $argument,
        ['\Code\Chat\V2\StartChatResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * SendMessage sends a message to a chat
     * @param \Code\Chat\V2\SendMessageRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SendMessage(\Code\Chat\V2\SendMessageRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v2.Chat/SendMessage',
        $argument,
        ['\Code\Chat\V2\SendMessageResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * AdvancePointer advances a pointer in message history for a chat member
     * @param \Code\Chat\V2\AdvancePointerRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AdvancePointer(\Code\Chat\V2\AdvancePointerRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v2.Chat/AdvancePointer',
        $argument,
        ['\Code\Chat\V2\AdvancePointerResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * RevealIdentity reveals a chat member's identity if it is anonymous. A chat
     * message will be inserted on success.
     * @param \Code\Chat\V2\RevealIdentityRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RevealIdentity(\Code\Chat\V2\RevealIdentityRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v2.Chat/RevealIdentity',
        $argument,
        ['\Code\Chat\V2\RevealIdentityResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * SetMuteState configures a chat member's mute state
     * @param \Code\Chat\V2\SetMuteStateRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SetMuteState(\Code\Chat\V2\SetMuteStateRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v2.Chat/SetMuteState',
        $argument,
        ['\Code\Chat\V2\SetMuteStateResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * SetSubscriptionState configures a chat member's susbscription state
     * @param \Code\Chat\V2\SetSubscriptionStateRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SetSubscriptionState(\Code\Chat\V2\SetSubscriptionStateRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v2.Chat/SetSubscriptionState',
        $argument,
        ['\Code\Chat\V2\SetSubscriptionStateResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * NotifyIsTypingRequest notifies a chat that the sending member is typing.
     *
     * These requests are transient, and may be dropped at any point.
     * @param \Code\Chat\V2\NotifyIsTypingRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function NotifyIsTyping(\Code\Chat\V2\NotifyIsTypingRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v2.Chat/NotifyIsTyping',
        $argument,
        ['\Code\Chat\V2\NotifyIsTypingResponse', 'decode'],
        $metadata, $options);
    }

}
