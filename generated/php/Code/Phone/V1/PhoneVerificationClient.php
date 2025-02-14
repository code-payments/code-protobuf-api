<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Phone\V1;

/**
 */
class PhoneVerificationClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * SendVerificationCode sends a verification code to the provided phone number
     * over SMS. If an active verification is already taking place, the existing code
     * will be resent.
     * @param \Code\Phone\V1\SendVerificationCodeRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SendVerificationCode(\Code\Phone\V1\SendVerificationCodeRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.phone.v1.PhoneVerification/SendVerificationCode',
        $argument,
        ['\Code\Phone\V1\SendVerificationCodeResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * CheckVerificationCode validates a verification code. On success, a one-time use
     * token to link an owner account is provided. 
     * @param \Code\Phone\V1\CheckVerificationCodeRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CheckVerificationCode(\Code\Phone\V1\CheckVerificationCodeRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.phone.v1.PhoneVerification/CheckVerificationCode',
        $argument,
        ['\Code\Phone\V1\CheckVerificationCodeResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetAssociatedPhoneNumber gets the latest verified phone number linked to an owner account.
     * @param \Code\Phone\V1\GetAssociatedPhoneNumberRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetAssociatedPhoneNumber(\Code\Phone\V1\GetAssociatedPhoneNumberRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.phone.v1.PhoneVerification/GetAssociatedPhoneNumber',
        $argument,
        ['\Code\Phone\V1\GetAssociatedPhoneNumberResponse', 'decode'],
        $metadata, $options);
    }

}
