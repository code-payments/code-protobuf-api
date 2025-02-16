<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: invite/v2/invite_service.proto

namespace Code\Invite\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.invite.v2.InvitePhoneNumberRequest</code>
 */
class InvitePhoneNumberRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * The phone number receiving the invite.
     *
     * Generated from protobuf field <code>.code.common.v1.PhoneNumber receiver = 2 [(.validate.rules) = {</code>
     */
    protected $receiver = null;
    protected $source;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\UserId $user
     *     @type \Code\Invite\V2\InviteCode $invite_code
     *     @type \Code\Common\V1\PhoneNumber $receiver
     *           The phone number receiving the invite.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Invite\V2\InviteService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.UserId user = 1;</code>
     * @return \Code\Common\V1\UserId|null
     */
    public function getUser()
    {
        return $this->readOneof(1);
    }

    public function hasUser()
    {
        return $this->hasOneof(1);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.UserId user = 1;</code>
     * @param \Code\Common\V1\UserId $var
     * @return $this
     */
    public function setUser($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\UserId::class);
        $this->writeOneof(1, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.invite.v2.InviteCode invite_code = 3;</code>
     * @return \Code\Invite\V2\InviteCode|null
     */
    public function getInviteCode()
    {
        return $this->readOneof(3);
    }

    public function hasInviteCode()
    {
        return $this->hasOneof(3);
    }

    /**
     * Generated from protobuf field <code>.code.invite.v2.InviteCode invite_code = 3;</code>
     * @param \Code\Invite\V2\InviteCode $var
     * @return $this
     */
    public function setInviteCode($var)
    {
        GPBUtil::checkMessage($var, \Code\Invite\V2\InviteCode::class);
        $this->writeOneof(3, $var);

        return $this;
    }

    /**
     * The phone number receiving the invite.
     *
     * Generated from protobuf field <code>.code.common.v1.PhoneNumber receiver = 2 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\PhoneNumber|null
     */
    public function getReceiver()
    {
        return $this->receiver;
    }

    public function hasReceiver()
    {
        return isset($this->receiver);
    }

    public function clearReceiver()
    {
        unset($this->receiver);
    }

    /**
     * The phone number receiving the invite.
     *
     * Generated from protobuf field <code>.code.common.v1.PhoneNumber receiver = 2 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\PhoneNumber $var
     * @return $this
     */
    public function setReceiver($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\PhoneNumber::class);
        $this->receiver = $var;

        return $this;
    }

    /**
     * @return string
     */
    public function getSource()
    {
        return $this->whichOneof("source");
    }

}

