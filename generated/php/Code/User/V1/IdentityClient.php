<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\User\V1;

/**
 */
class IdentityClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * LinkAccount links an owner account to the user identified and authenticated
     * by a one-time use token.
     *
     * Notably, this RPC has the following side effects:
     *   * A new user is automatically created if one doesn't exist.
     *   * Server will create a new data container for at least every unique
     *     owner account linked to the user.
     * @param \Code\User\V1\LinkAccountRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function LinkAccount(\Code\User\V1\LinkAccountRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.user.v1.Identity/LinkAccount',
        $argument,
        ['\Code\User\V1\LinkAccountResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * UnlinkAccount removes links from an owner account. It will NOT remove
     * existing associations between users, owner accounts and identifying
     * features.
     *
     * The following associations will remain intact to ensure owner accounts
     * can continue to be used with a consistent login experience:
     *   * the user continues to be associated to existing owner accounts and
     *     identifying features
     *
     * Client can continue mainting their current login session. Their current
     * user and data container will remain the same.
     *
     * The call is guaranteed to be idempotent. It will not fail if the link is
     * already removed by either a previous call to this RPC or by a more recent
     * call to LinkAccount. A failure will only occur if the link between a user
     * and the owner accout or identifying feature never existed.
     * @param \Code\User\V1\UnlinkAccountRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function UnlinkAccount(\Code\User\V1\UnlinkAccountRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.user.v1.Identity/UnlinkAccount',
        $argument,
        ['\Code\User\V1\UnlinkAccountResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetUser gets user information given a user identifier and an owner account.
     * @param \Code\User\V1\GetUserRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetUser(\Code\User\V1\GetUserRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.user.v1.Identity/GetUser',
        $argument,
        ['\Code\User\V1\GetUserResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * UpdatePreferences updates user preferences.
     * @param \Code\User\V1\UpdatePreferencesRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function UpdatePreferences(\Code\User\V1\UpdatePreferencesRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.user.v1.Identity/UpdatePreferences',
        $argument,
        ['\Code\User\V1\UpdatePreferencesResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * LoginToThirdPartyApp logs a user into a third party app for a given intent
     * ID. If the original request requires payment, then SubmitIntent must be called.
     * @param \Code\User\V1\LoginToThirdPartyAppRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function LoginToThirdPartyApp(\Code\User\V1\LoginToThirdPartyAppRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.user.v1.Identity/LoginToThirdPartyApp',
        $argument,
        ['\Code\User\V1\LoginToThirdPartyAppResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetLoginForThirdPartyApp gets a login for a third party app from an existing
     * request. This endpoint supports all paths where login is possible (login on payment,
     * raw login, etc.).
     * @param \Code\User\V1\GetLoginForThirdPartyAppRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetLoginForThirdPartyApp(\Code\User\V1\GetLoginForThirdPartyAppRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.user.v1.Identity/GetLoginForThirdPartyApp',
        $argument,
        ['\Code\User\V1\GetLoginForThirdPartyAppResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetTwitterUser gets Twitter user information
     *
     * Note 1: This RPC will only return results for Twitter users that have
     * accounts linked with Code.
     *
     * Note 2: This RPC is heavily cached, and may not reflect real-time Twitter
     * information.
     * @param \Code\User\V1\GetTwitterUserRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetTwitterUser(\Code\User\V1\GetTwitterUserRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.user.v1.Identity/GetTwitterUser',
        $argument,
        ['\Code\User\V1\GetTwitterUserResponse', 'decode'],
        $metadata, $options);
    }

}
