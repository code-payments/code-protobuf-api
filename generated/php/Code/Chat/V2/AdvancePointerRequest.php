<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat/v2/chat_service.proto

namespace Code\Chat\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.chat.v2.AdvancePointerRequest</code>
 */
class AdvancePointerRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.common.v1.ChatId chat_id = 1 [(.validate.rules) = {</code>
     */
    protected $chat_id = null;
    /**
     * Generated from protobuf field <code>.code.chat.v2.Pointer pointer = 2 [(.validate.rules) = {</code>
     */
    protected $pointer = null;
    /**
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 3 [(.validate.rules) = {</code>
     */
    protected $owner = null;
    /**
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 4 [(.validate.rules) = {</code>
     */
    protected $signature = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\ChatId $chat_id
     *     @type \Code\Chat\V2\Pointer $pointer
     *     @type \Code\Common\V1\SolanaAccountId $owner
     *     @type \Code\Common\V1\Signature $signature
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Chat\V2\ChatService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.ChatId chat_id = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\ChatId|null
     */
    public function getChatId()
    {
        return $this->chat_id;
    }

    public function hasChatId()
    {
        return isset($this->chat_id);
    }

    public function clearChatId()
    {
        unset($this->chat_id);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.ChatId chat_id = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\ChatId $var
     * @return $this
     */
    public function setChatId($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\ChatId::class);
        $this->chat_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.chat.v2.Pointer pointer = 2 [(.validate.rules) = {</code>
     * @return \Code\Chat\V2\Pointer|null
     */
    public function getPointer()
    {
        return $this->pointer;
    }

    public function hasPointer()
    {
        return isset($this->pointer);
    }

    public function clearPointer()
    {
        unset($this->pointer);
    }

    /**
     * Generated from protobuf field <code>.code.chat.v2.Pointer pointer = 2 [(.validate.rules) = {</code>
     * @param \Code\Chat\V2\Pointer $var
     * @return $this
     */
    public function setPointer($var)
    {
        GPBUtil::checkMessage($var, \Code\Chat\V2\Pointer::class);
        $this->pointer = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 3 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getOwner()
    {
        return $this->owner;
    }

    public function hasOwner()
    {
        return isset($this->owner);
    }

    public function clearOwner()
    {
        unset($this->owner);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 3 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setOwner($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->owner = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 4 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Signature|null
     */
    public function getSignature()
    {
        return $this->signature;
    }

    public function hasSignature()
    {
        return isset($this->signature);
    }

    public function clearSignature()
    {
        unset($this->signature);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 4 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Signature $var
     * @return $this
     */
    public function setSignature($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Signature::class);
        $this->signature = $var;

        return $this;
    }

}

