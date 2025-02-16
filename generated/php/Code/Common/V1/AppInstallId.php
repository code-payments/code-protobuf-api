<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/v1/model.proto

namespace Code\Common\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * AppInstallId is a unque ID tied to a client app installation. It does not
 * identify a device. Value should remain private and not be shared across
 * installs.
 *
 * Generated from protobuf message <code>code.common.v1.AppInstallId</code>
 */
class AppInstallId extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string value = 1 [(.validate.rules) = {</code>
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
        \GPBMetadata\Common\V1\Model::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string value = 1 [(.validate.rules) = {</code>
     * @return string
     */
    public function getValue()
    {
        return $this->value;
    }

    /**
     * Generated from protobuf field <code>string value = 1 [(.validate.rules) = {</code>
     * @param string $var
     * @return $this
     */
    public function setValue($var)
    {
        GPBUtil::checkString($var, True);
        $this->value = $var;

        return $this;
    }

}

