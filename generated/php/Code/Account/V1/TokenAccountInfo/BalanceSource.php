<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: account/v1/account_service.proto

namespace Code\Account\V1\TokenAccountInfo;

use UnexpectedValueException;

/**
 * Protobuf type <code>code.account.v1.TokenAccountInfo.BalanceSource</code>
 */
class BalanceSource
{
    /**
     * The account's balance could not be determined. This may be returned when
     * the data source is unstable and a reliable balance cannot be determined.
     *
     * Generated from protobuf enum <code>BALANCE_SOURCE_UNKNOWN = 0;</code>
     */
    const BALANCE_SOURCE_UNKNOWN = 0;
    /**
     * The account's balance was fetched directly from a finalized state on the
     * blockchain.
     *
     * Generated from protobuf enum <code>BALANCE_SOURCE_BLOCKCHAIN = 1;</code>
     */
    const BALANCE_SOURCE_BLOCKCHAIN = 1;
    /**
     * The account's balance was calculated using cached values in Code. Accuracy
     * is only guaranteed when management_state is LOCKED.
     *
     * Generated from protobuf enum <code>BALANCE_SOURCE_CACHE = 2;</code>
     */
    const BALANCE_SOURCE_CACHE = 2;

    private static $valueToName = [
        self::BALANCE_SOURCE_UNKNOWN => 'BALANCE_SOURCE_UNKNOWN',
        self::BALANCE_SOURCE_BLOCKCHAIN => 'BALANCE_SOURCE_BLOCKCHAIN',
        self::BALANCE_SOURCE_CACHE => 'BALANCE_SOURCE_CACHE',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(BalanceSource::class, \Code\Account\V1\TokenAccountInfo_BalanceSource::class);

