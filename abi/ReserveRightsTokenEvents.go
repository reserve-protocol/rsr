// This file is auto-generated. Do not edit.

package abi

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *ReserveRightsTokenFilterer) ParseLog(log *types.Log) (fmt.Stringer, error) {
	var event fmt.Stringer
	var eventName string
	switch log.Topics[0].Hex() {
	case "0x78be06d07afe380e04d6deeba0f33c892db454f303fb739d9b768987a5ec6aca": // AccountLocked
		event = new(ReserveRightsTokenAccountLocked)
		eventName = "AccountLocked"
	case "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925": // Approval
		event = new(ReserveRightsTokenApproval)
		eventName = "Approval"
	case "0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258": // Paused
		event = new(ReserveRightsTokenPaused)
		eventName = "Paused"
	case "0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8": // PauserAdded
		event = new(ReserveRightsTokenPauserAdded)
		eventName = "PauserAdded"
	case "0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e": // PauserRemoved
		event = new(ReserveRightsTokenPauserRemoved)
		eventName = "PauserRemoved"
	case "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef": // Transfer
		event = new(ReserveRightsTokenTransfer)
		eventName = "Transfer"
	case "0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa": // Unpaused
		event = new(ReserveRightsTokenUnpaused)
		eventName = "Unpaused"
	default:
		return nil, fmt.Errorf("no such event hash for ReserveRightsToken: %v", log.Topics[0])
	}

	err := c.contract.UnpackLog(event, eventName, *log)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (e ReserveRightsTokenAccountLocked) String() string {
	return fmt.Sprintf("ReserveRightsToken.AccountLocked(%v)", e.LockedAccount.Hex())
}

func (e ReserveRightsTokenApproval) String() string {
	return fmt.Sprintf("ReserveRightsToken.Approval(%v, %v, %v)", e.Owner.Hex(), e.Spender.Hex(), e.Value)
}

func (e ReserveRightsTokenPaused) String() string {
	return fmt.Sprintf("ReserveRightsToken.Paused(%v)", e.Account.Hex())
}

func (e ReserveRightsTokenPauserAdded) String() string {
	return fmt.Sprintf("ReserveRightsToken.PauserAdded(%v)", e.Account.Hex())
}

func (e ReserveRightsTokenPauserRemoved) String() string {
	return fmt.Sprintf("ReserveRightsToken.PauserRemoved(%v)", e.Account.Hex())
}

func (e ReserveRightsTokenTransfer) String() string {
	return fmt.Sprintf("ReserveRightsToken.Transfer(%v, %v, %v)", e.From.Hex(), e.To.Hex(), e.Value)
}

func (e ReserveRightsTokenUnpaused) String() string {
	return fmt.Sprintf("ReserveRightsToken.Unpaused(%v)", e.Account.Hex())
}
