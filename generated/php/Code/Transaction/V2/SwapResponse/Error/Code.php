<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2\SwapResponse\Error;

use UnexpectedValueException;

/**
 * Protobuf type <code>code.transaction.v2.SwapResponse.Error.Code</code>
 */
class Code
{
    /**
     * Denied by a guard (spam, money laundering, etc)
     *
     * Generated from protobuf enum <code>DENIED = 0;</code>
     */
    const DENIED = 0;
    /**
     * There is an issue with the provided signature.
     *
     * Generated from protobuf enum <code>SIGNATURE_ERROR = 2;</code>
     */
    const SIGNATURE_ERROR = 2;
    /**
     * The swap failed server-side validation
     *
     * Generated from protobuf enum <code>INVALID_SWAP = 3;</code>
     */
    const INVALID_SWAP = 3;
    /**
     * The submitted swap transaction failed. Attempt the swap again.
     *
     * Generated from protobuf enum <code>SWAP_FAILED = 4;</code>
     */
    const SWAP_FAILED = 4;

    private static $valueToName = [
        self::DENIED => 'DENIED',
        self::SIGNATURE_ERROR => 'SIGNATURE_ERROR',
        self::INVALID_SWAP => 'INVALID_SWAP',
        self::SWAP_FAILED => 'SWAP_FAILED',
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
class_alias(Code::class, \Code\Transaction\V2\SwapResponse_Error_Code::class);

