<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: phone/v1/phone_verification_service.proto

namespace Code\Phone\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.phone.v1.SendVerificationCodeRequest</code>
 */
class SendVerificationCodeRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * The phone number to send a verification code over SMS to.
     *
     * Generated from protobuf field <code>.code.common.v1.PhoneNumber phone_number = 1 [(.validate.rules) = {</code>
     */
    protected $phone_number = null;
    /**
     * Device token for antispam measures against fake devices
     *
     * Generated from protobuf field <code>.code.common.v1.DeviceToken device_token = 2;</code>
     */
    protected $device_token = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\PhoneNumber $phone_number
     *           The phone number to send a verification code over SMS to.
     *     @type \Code\Common\V1\DeviceToken $device_token
     *           Device token for antispam measures against fake devices
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Phone\V1\PhoneVerificationService::initOnce();
        parent::__construct($data);
    }

    /**
     * The phone number to send a verification code over SMS to.
     *
     * Generated from protobuf field <code>.code.common.v1.PhoneNumber phone_number = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\PhoneNumber|null
     */
    public function getPhoneNumber()
    {
        return $this->phone_number;
    }

    public function hasPhoneNumber()
    {
        return isset($this->phone_number);
    }

    public function clearPhoneNumber()
    {
        unset($this->phone_number);
    }

    /**
     * The phone number to send a verification code over SMS to.
     *
     * Generated from protobuf field <code>.code.common.v1.PhoneNumber phone_number = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\PhoneNumber $var
     * @return $this
     */
    public function setPhoneNumber($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\PhoneNumber::class);
        $this->phone_number = $var;

        return $this;
    }

    /**
     * Device token for antispam measures against fake devices
     *
     * Generated from protobuf field <code>.code.common.v1.DeviceToken device_token = 2;</code>
     * @return \Code\Common\V1\DeviceToken|null
     */
    public function getDeviceToken()
    {
        return $this->device_token;
    }

    public function hasDeviceToken()
    {
        return isset($this->device_token);
    }

    public function clearDeviceToken()
    {
        unset($this->device_token);
    }

    /**
     * Device token for antispam measures against fake devices
     *
     * Generated from protobuf field <code>.code.common.v1.DeviceToken device_token = 2;</code>
     * @param \Code\Common\V1\DeviceToken $var
     * @return $this
     */
    public function setDeviceToken($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\DeviceToken::class);
        $this->device_token = $var;

        return $this;
    }

}

