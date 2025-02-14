<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Device\V1;

/**
 */
class DeviceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * RegisterLoggedInAccounts registers a set of owner accounts logged for
     * an app install. Currently, a single login is enforced per app install.
     * After using GetLoggedInAccounts to detect stale logins, clients can use
     * this RPC to update the set of accounts with valid login sessions.
     * @param \Code\Device\V1\RegisterLoggedInAccountsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RegisterLoggedInAccounts(\Code\Device\V1\RegisterLoggedInAccountsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.device.v1.Device/RegisterLoggedInAccounts',
        $argument,
        ['\Code\Device\V1\RegisterLoggedInAccountsResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetLoggedInAccounts gets the set of logged in accounts for an app install.
     * Clients can use this RPC to detect stale logins for boot out of the app.
     * @param \Code\Device\V1\GetLoggedInAccountsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetLoggedInAccounts(\Code\Device\V1\GetLoggedInAccountsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.device.v1.Device/GetLoggedInAccounts',
        $argument,
        ['\Code\Device\V1\GetLoggedInAccountsResponse', 'decode'],
        $metadata, $options);
    }

}
