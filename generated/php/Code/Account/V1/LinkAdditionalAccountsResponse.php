<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: account/v1/account_service.proto

namespace Code\Account\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.account.v1.LinkAdditionalAccountsResponse</code>
 */
class LinkAdditionalAccountsResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.account.v1.LinkAdditionalAccountsResponse.Result result = 1;</code>
     */
    protected $result = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $result
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Account\V1\AccountService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.account.v1.LinkAdditionalAccountsResponse.Result result = 1;</code>
     * @return int
     */
    public function getResult()
    {
        return $this->result;
    }

    /**
     * Generated from protobuf field <code>.code.account.v1.LinkAdditionalAccountsResponse.Result result = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setResult($var)
    {
        GPBUtil::checkEnum($var, \Code\Account\V1\LinkAdditionalAccountsResponse\Result::class);
        $this->result = $var;

        return $this;
    }

}

