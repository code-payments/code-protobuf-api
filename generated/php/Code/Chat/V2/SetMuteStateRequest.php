<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat/v2/chat_service.proto

namespace Code\Chat\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.chat.v2.SetMuteStateRequest</code>
 */
class SetMuteStateRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.common.v1.ChatId chat_id = 1 [(.validate.rules) = {</code>
     */
    protected $chat_id = null;
    /**
     * Generated from protobuf field <code>.code.chat.v2.ChatMemberId member_id = 2 [(.validate.rules) = {</code>
     */
    protected $member_id = null;
    /**
     * Generated from protobuf field <code>bool is_muted = 3;</code>
     */
    protected $is_muted = false;
    /**
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 4 [(.validate.rules) = {</code>
     */
    protected $owner = null;
    /**
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 5 [(.validate.rules) = {</code>
     */
    protected $signature = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\ChatId $chat_id
     *     @type \Code\Chat\V2\ChatMemberId $member_id
     *     @type bool $is_muted
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
     * Generated from protobuf field <code>.code.chat.v2.ChatMemberId member_id = 2 [(.validate.rules) = {</code>
     * @return \Code\Chat\V2\ChatMemberId|null
     */
    public function getMemberId()
    {
        return $this->member_id;
    }

    public function hasMemberId()
    {
        return isset($this->member_id);
    }

    public function clearMemberId()
    {
        unset($this->member_id);
    }

    /**
     * Generated from protobuf field <code>.code.chat.v2.ChatMemberId member_id = 2 [(.validate.rules) = {</code>
     * @param \Code\Chat\V2\ChatMemberId $var
     * @return $this
     */
    public function setMemberId($var)
    {
        GPBUtil::checkMessage($var, \Code\Chat\V2\ChatMemberId::class);
        $this->member_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool is_muted = 3;</code>
     * @return bool
     */
    public function getIsMuted()
    {
        return $this->is_muted;
    }

    /**
     * Generated from protobuf field <code>bool is_muted = 3;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsMuted($var)
    {
        GPBUtil::checkBool($var);
        $this->is_muted = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 4 [(.validate.rules) = {</code>
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
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 4 [(.validate.rules) = {</code>
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
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 5 [(.validate.rules) = {</code>
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
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 5 [(.validate.rules) = {</code>
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

