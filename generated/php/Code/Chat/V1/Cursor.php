<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat/v1/chat_service.proto

namespace Code\Chat\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Opaque cursor used across paged APIs. Underlying bytes may change as paging
 * strategies evolve.
 *
 * Generated from protobuf message <code>code.chat.v1.Cursor</code>
 */
class Cursor extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>bytes value = 1 [(.validate.rules) = {</code>
     */
    protected $value = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $value
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Chat\V1\ChatService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>bytes value = 1 [(.validate.rules) = {</code>
     * @return string
     */
    public function getValue()
    {
        return $this->value;
    }

    /**
     * Generated from protobuf field <code>bytes value = 1 [(.validate.rules) = {</code>
     * @param string $var
     * @return $this
     */
    public function setValue($var)
    {
        GPBUtil::checkString($var, False);
        $this->value = $var;

        return $this;
    }

}

