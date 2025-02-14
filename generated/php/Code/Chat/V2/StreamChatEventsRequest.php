<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat/v2/chat_service.proto

namespace Code\Chat\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.chat.v2.StreamChatEventsRequest</code>
 */
class StreamChatEventsRequest extends \Google\Protobuf\Internal\Message
{
    protected $type;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Chat\V2\OpenChatEventStream $open_stream
     *     @type \Code\Common\V1\ClientPong $pong
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Chat\V2\ChatService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.chat.v2.OpenChatEventStream open_stream = 1;</code>
     * @return \Code\Chat\V2\OpenChatEventStream|null
     */
    public function getOpenStream()
    {
        return $this->readOneof(1);
    }

    public function hasOpenStream()
    {
        return $this->hasOneof(1);
    }

    /**
     * Generated from protobuf field <code>.code.chat.v2.OpenChatEventStream open_stream = 1;</code>
     * @param \Code\Chat\V2\OpenChatEventStream $var
     * @return $this
     */
    public function setOpenStream($var)
    {
        GPBUtil::checkMessage($var, \Code\Chat\V2\OpenChatEventStream::class);
        $this->writeOneof(1, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.ClientPong pong = 2;</code>
     * @return \Code\Common\V1\ClientPong|null
     */
    public function getPong()
    {
        return $this->readOneof(2);
    }

    public function hasPong()
    {
        return $this->hasOneof(2);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.ClientPong pong = 2;</code>
     * @param \Code\Common\V1\ClientPong $var
     * @return $this
     */
    public function setPong($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\ClientPong::class);
        $this->writeOneof(2, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getType()
    {
        return $this->whichOneof("type");
    }

}

