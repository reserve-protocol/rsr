// This file is auto-generated. Do not edit.

package abi

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *SlowWalletFilterer) ParseLog(log *types.Log) (fmt.Stringer, error) {
	var event fmt.Stringer
	var eventName string
	switch log.Topics[0].Hex() {
	case "0xf4a9bad09d9916720a4c78733bb1b637bccdc7b56ae0f78e6bb63c99934b49b9": // AllTransfersCancelled
		event = new(SlowWalletAllTransfersCancelled)
		eventName = "AllTransfersCancelled"
	case "0xa99797fde63ee29177c0bcd12725053c1fcecc39957a12d910dad426ddefafcf": // TransferCancelled
		event = new(SlowWalletTransferCancelled)
		eventName = "TransferCancelled"
	case "0x0725491da501611b9ba3843fe633719becb32a870a63c7008fab061512c83417": // TransferConfirmed
		event = new(SlowWalletTransferConfirmed)
		eventName = "TransferConfirmed"
	case "0xc21d9f71ad43be26dd663dc2ae89d8a5f2e2e627520624eeff74bf1004d80b3d": // TransferProposed
		event = new(SlowWalletTransferProposed)
		eventName = "TransferProposed"
	default:
		return nil, fmt.Errorf("no such event hash for SlowWallet: %v", log.Topics[0])
	}

	err := c.contract.UnpackLog(event, eventName, *log)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (e SlowWalletAllTransfersCancelled) String() string {
	return fmt.Sprintf("SlowWallet.AllTransfersCancelled()")
}

func (e SlowWalletTransferCancelled) String() string {
	return fmt.Sprintf("SlowWallet.TransferCancelled(%v, %v, %v, %q)", e.Index, e.Destination.Hex(), e.Value, e.Notes)
}

func (e SlowWalletTransferConfirmed) String() string {
	return fmt.Sprintf("SlowWallet.TransferConfirmed(%v, %v, %v, %q)", e.Index, e.Destination.Hex(), e.Value, e.Notes)
}

func (e SlowWalletTransferProposed) String() string {
	return fmt.Sprintf("SlowWallet.TransferProposed(%v, %v, %v, %v, %q)", e.Index, e.Destination.Hex(), e.Value, e.DelayUntil, e.Notes)
}
