<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Chat\V1;

/**
 * Deprecated: Use the v2 service
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
     * GetChats gets the set of chats for an owner account
     * @param \Code\Chat\V1\GetChatsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetChats(\Code\Chat\V1\GetChatsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v1.Chat/GetChats',
        $argument,
        ['\Code\Chat\V1\GetChatsResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetMessages gets the set of messages for a chat
     * @param \Code\Chat\V1\GetMessagesRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetMessages(\Code\Chat\V1\GetMessagesRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v1.Chat/GetMessages',
        $argument,
        ['\Code\Chat\V1\GetMessagesResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * AdvancePointer advances a pointer in chat history
     * @param \Code\Chat\V1\AdvancePointerRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AdvancePointer(\Code\Chat\V1\AdvancePointerRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v1.Chat/AdvancePointer',
        $argument,
        ['\Code\Chat\V1\AdvancePointerResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * SetMuteState configures the mute state of a chat
     * @param \Code\Chat\V1\SetMuteStateRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SetMuteState(\Code\Chat\V1\SetMuteStateRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v1.Chat/SetMuteState',
        $argument,
        ['\Code\Chat\V1\SetMuteStateResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * SetSubscriptionState configures the susbscription state of a chat
     * @param \Code\Chat\V1\SetSubscriptionStateRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SetSubscriptionState(\Code\Chat\V1\SetSubscriptionStateRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v1.Chat/SetSubscriptionState',
        $argument,
        ['\Code\Chat\V1\SetSubscriptionStateResponse', 'decode'],
        $metadata, $options);
    }

    /**
     *
     * Experimental PoC two-way chat APIs below
     *
     *
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function StreamChatEvents($metadata = [], $options = []) {
        return $this->_bidiRequest('/code.chat.v1.Chat/StreamChatEvents',
        ['\Code\Chat\V1\StreamChatEventsResponse','decode'],
        $metadata, $options);
    }

    /**
     * @param \Code\Chat\V1\SendMessageRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SendMessage(\Code\Chat\V1\SendMessageRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.chat.v1.Chat/SendMessage',
        $argument,
        ['\Code\Chat\V1\SendMessageResponse', 'decode'],
        $metadata, $options);
    }

}
