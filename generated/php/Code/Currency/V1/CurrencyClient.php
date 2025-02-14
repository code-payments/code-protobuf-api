<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Currency\V1;

/**
 */
class CurrencyClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * GetAllRates returns the exchange rates for Kin against all available currencies
     * @param \Code\Currency\V1\GetAllRatesRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetAllRates(\Code\Currency\V1\GetAllRatesRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.currency.v1.Currency/GetAllRates',
        $argument,
        ['\Code\Currency\V1\GetAllRatesResponse', 'decode'],
        $metadata, $options);
    }

}
