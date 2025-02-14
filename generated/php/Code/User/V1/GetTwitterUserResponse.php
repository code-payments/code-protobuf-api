<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user/v1/identity_service.proto

namespace Code\User\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.user.v1.GetTwitterUserResponse</code>
 */
class GetTwitterUserResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.user.v1.GetTwitterUserResponse.Result result = 1;</code>
     */
    protected $result = 0;
    /**
     * Generated from protobuf field <code>.code.user.v1.TwitterUser twitter_user = 2;</code>
     */
    protected $twitter_user = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $result
     *     @type \Code\User\V1\TwitterUser $twitter_user
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\User\V1\IdentityService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.user.v1.GetTwitterUserResponse.Result result = 1;</code>
     * @return int
     */
    public function getResult()
    {
        return $this->result;
    }

    /**
     * Generated from protobuf field <code>.code.user.v1.GetTwitterUserResponse.Result result = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setResult($var)
    {
        GPBUtil::checkEnum($var, \Code\User\V1\GetTwitterUserResponse\Result::class);
        $this->result = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.user.v1.TwitterUser twitter_user = 2;</code>
     * @return \Code\User\V1\TwitterUser|null
     */
    public function getTwitterUser()
    {
        return $this->twitter_user;
    }

    public function hasTwitterUser()
    {
        return isset($this->twitter_user);
    }

    public function clearTwitterUser()
    {
        unset($this->twitter_user);
    }

    /**
     * Generated from protobuf field <code>.code.user.v1.TwitterUser twitter_user = 2;</code>
     * @param \Code\User\V1\TwitterUser $var
     * @return $this
     */
    public function setTwitterUser($var)
    {
        GPBUtil::checkMessage($var, \Code\User\V1\TwitterUser::class);
        $this->twitter_user = $var;

        return $this;
    }

}

