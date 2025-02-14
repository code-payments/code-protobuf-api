<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2\FeePaymentAction;

use UnexpectedValueException;

/**
 * Protobuf type <code>code.transaction.v2.FeePaymentAction.FeeType</code>
 */
class FeeType
{
    /**
     * Hardcoded $0.01 USD fee to a dynamic fee account specified by server
     *
     * Generated from protobuf enum <code>CODE = 0;</code>
     */
    const CODE = 0;
    /**
     * Third party fee specified at time of payment request
     *
     * Generated from protobuf enum <code>THIRD_PARTY = 1;</code>
     */
    const THIRD_PARTY = 1;

    private static $valueToName = [
        self::CODE => 'CODE',
        self::THIRD_PARTY => 'THIRD_PARTY',
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
class_alias(FeeType::class, \Code\Transaction\V2\FeePaymentAction_FeeType::class);

