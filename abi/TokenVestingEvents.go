// This file is auto-generated. Do not edit.

package abi

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *TokenVestingFilterer) ParseLog(log *types.Log) (fmt.Stringer, error) {
	var event fmt.Stringer
	var eventName string
	switch log.Topics[0].Hex() {
	case "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0": // OwnershipTransferred
		event = new(TokenVestingOwnershipTransferred)
		eventName = "OwnershipTransferred"
	case "0xc7798891864187665ac6dd119286e44ec13f014527aeeb2b8eb3fd413df93179": // TokensReleased
		event = new(TokenVestingTokensReleased)
		eventName = "TokensReleased"
	default:
		return nil, fmt.Errorf("no such event hash for TokenVesting: %v", log.Topics[0])
	}

	err := c.contract.UnpackLog(event, eventName, *log)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (e TokenVestingOwnershipTransferred) String() string {
	return fmt.Sprintf("TokenVesting.OwnershipTransferred(%v, %v)", e.PreviousOwner.Hex(), e.NewOwner.Hex())
}

func (e TokenVestingTokensReleased) String() string {
	return fmt.Sprintf("TokenVesting.TokensReleased(%v, %v)", e.Token.Hex(), e.Amount)
}
