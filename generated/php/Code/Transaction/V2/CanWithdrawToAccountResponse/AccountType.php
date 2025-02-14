<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2\CanWithdrawToAccountResponse;

use UnexpectedValueException;

/**
 * Protobuf type <code>code.transaction.v2.CanWithdrawToAccountResponse.AccountType</code>
 */
class AccountType
{
    /**
     * Server cannot determine
     *
     * Generated from protobuf enum <code>Unknown = 0;</code>
     */
    const Unknown = 0;
    /**
     * Client uses the address as is in SubmitIntent
     *
     * Generated from protobuf enum <code>TokenAccount = 1;</code>
     */
    const TokenAccount = 1;
    /**
     * Client locally derives the ATA to use in SubmitIntent
     *
     * Generated from protobuf enum <code>OwnerAccount = 2;</code>
     */
    const OwnerAccount = 2;

    private static $valueToName = [
        self::Unknown => 'Unknown',
        self::TokenAccount => 'TokenAccount',
        self::OwnerAccount => 'OwnerAccount',
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
class_alias(AccountType::class, \Code\Transaction\V2\CanWithdrawToAccountResponse_AccountType::class);

