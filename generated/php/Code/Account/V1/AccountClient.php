<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Account\V1;

/**
 */
class AccountClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * IsCodeAccount returns whether an owner account is a Code account. This hints
     * to the client whether the account can be logged in, used for making payments,
     * etc.
     * @param \Code\Account\V1\IsCodeAccountRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function IsCodeAccount(\Code\Account\V1\IsCodeAccountRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.account.v1.Account/IsCodeAccount',
        $argument,
        ['\Code\Account\V1\IsCodeAccountResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetTokenAccountInfos returns token account metadata relevant to the Code owner
     * account.
     * @param \Code\Account\V1\GetTokenAccountInfosRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetTokenAccountInfos(\Code\Account\V1\GetTokenAccountInfosRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.account.v1.Account/GetTokenAccountInfos',
        $argument,
        ['\Code\Account\V1\GetTokenAccountInfosResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * LinkAdditionalAccounts allows a client to declare additional accounts to
     * be tracked and used within Code. The accounts declared in this RPC are not
     * managed by Code (ie. not a Timelock account), created externally and cannot
     * be linked automatically (ie. authority derived off user 12 words).
     * @param \Code\Account\V1\LinkAdditionalAccountsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function LinkAdditionalAccounts(\Code\Account\V1\LinkAdditionalAccountsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.account.v1.Account/LinkAdditionalAccounts',
        $argument,
        ['\Code\Account\V1\LinkAdditionalAccountsResponse', 'decode'],
        $metadata, $options);
    }

}
