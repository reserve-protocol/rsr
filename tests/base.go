package tests

// This is a base test suite that is extended by slow_wallet_test.go and vesting_test.go

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"

	rsrABI "github.com/reserve-protocol/rsr/abi"
)

type logParser interface {
	ParseLog(*types.Log) (fmt.Stringer, error)
}

// TestSuite holds functionality common between our two test suites.
//
// It knows how to create a connection to an Ethereum node, it holds a list of accounts
// to use with that node, and it implements common test assertions.
type TestSuite struct {
	suite.Suite

	account []account
	signer  *bind.TransactOpts
	node    interface {
		bind.ContractBackend
		TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	}
	rsr        *rsrABI.ReserveRightsToken
	rsrAddress common.Address

	utilContract *bind.BoundContract

	logParsers map[common.Address]logParser
}

// requireTx requires that a transaction is successfully mined and does
// not revert. It also takes an extra error argument, and checks that the
// error is nil. This signature allows the function to directly wrap our
// abigen'd mutator calls.
//
// requireTx returns a closure that can be used to assert the list of events
// that were emitted during the transaction. This API is a bit weird -- it would
// be more natural to pass the events in to the `requireTx` call itself -- but
// this is the cleanest way that is compatible with directly wrapping the abigen'd
// calls, without using intermediate placeholder variables in calling code.
func (s *TestSuite) requireTx(tx *types.Transaction, err error) func(assertEvent ...fmt.Stringer) {
	receipt := s._requireTxStatus(tx, err, types.ReceiptStatusSuccessful)

	// return a closure that can take a varargs list of events,
	// and assert that the transaction generates those events.
	return func(assertEvent ...fmt.Stringer) {
		if s.Equal(len(assertEvent), len(receipt.Logs), "did not get the expected number of events") {
			for i, wantEvent := range assertEvent {
				parser := s.logParsers[receipt.Logs[i].Address]
				if s.NotNil(parser, "got an event from an unexpected contract address: "+receipt.Logs[i].Address.Hex()) {
					gotEvent, err := parser.ParseLog(receipt.Logs[i])
					if s.NoErrorf(err, "parsing event %v", i) {
						s.Equal(wantEvent.String(), gotEvent.String())
					}
				}
			}
		}
	}
}

// requireTxFails is like requireTx, but it requires that the transaction either
// reverts or is not successfully made in the first place due to gas estimation
// failing.
func (s *TestSuite) requireTxFails(tx *types.Transaction, err error) {
	if err != nil && err.Error() ==
		"failed to estimate gas needed: gas required exceeds allowance or always failing transaction" {
		return
	}

	receipt := s._requireTxStatus(tx, err, types.ReceiptStatusFailed)
	s.Equal(0, len(receipt.Logs), "Zero logs should be generated for a failed transaction")
}

func (s *TestSuite) _requireTxStatus(tx *types.Transaction, err error, status uint64) *types.Receipt {
	s.Require().NoError(err)
	s.Require().NotNil(tx)
	receipt, err := bind.WaitMined(context.Background(), s.node, tx)
	s.Require().NoError(err)
	s.Require().Equal(status, receipt.Status)
	return receipt
}

// assertBalance asserts that the RSR balance of `address` is `amount`.
func (s *TestSuite) assertBalance(address common.Address, amount *big.Int) {
	s.Equal(amount.String(), s.currentBalance(address).String()) // assert.Equal can mis-compare big.Ints, so compare strings instead
}

// currentBalance returns the current balance of an address.
func (s *TestSuite) currentBalance(address common.Address) *big.Int {
	balance, err := s.rsr.BalanceOf(nil, address)
	s.NoError(err)
	return balance
}

// assertTotalSupply asserts that the total supply of RSR is `amount`.
func (s *TestSuite) assertTotalSupply(amount *big.Int) {
	totalSupply, err := s.rsr.TotalSupply(nil)
	s.NoError(err)
	s.Equal(amount.String(), totalSupply.String())
}

func (s *TestSuite) currentTimestamp() *big.Int {
	result := new(big.Int)
	s.NoError(s.utilContract.Call(nil, &result, "time"))
	return result
}

// createFastNode creates a fast in-process Ethereum node. It is then available as `s.node`.
func (s *TestSuite) createFastNode() {
	genesisAlloc := core.GenesisAlloc{}
	for _, account := range s.account {
		genesisAlloc[account.address()] = core.GenesisAccount{
			Balance: big.NewInt(math.MaxInt64),
		}
	}
	s.node = backend{
		backends.NewSimulatedBackend(
			genesisAlloc,
			// Block gas limit. Needs to be more than 7e6, which is about the cost
			// of the ReserveDollarV2 constructor. But we still want it about the
			// same order of magnitude as mainnet.
			//
			// The ReserveDollar constructor is edging close to the mainnet block limit.
			// We'll probably stay under it without any problem. If not, we can split
			// the Eternal Storage contract deployment into a different transaction.
			8e6,
		),
	}
}

// setup sets up the TestSuite. It must be called before using s.account or s.signer.
func (s *TestSuite) setup() {
	// The first few keys from the following well-known mnemonic used by 0x:
	//	concert load couple harbor equip island argue ramp clarify fence smart topic
	keys := []string{
		"f2f48ee19680706196e2e339e5da3491186e0c4c5030670656b0e0164837257d",
		"5d862464fe9303452126c8bc94274b8c5f9874cbd219789b3eb2128075a76f72",
		"df02719c4df8b9b8ac7f551fcb5d9ef48fa27eef7a66453879f4d8fdc6e78fb1",
		"ff12e391b79415e941a94de3bf3a9aee577aed0731e297d5cfa0b8a1e02fa1d0",
		"752dd9cf65e68cfaba7d60225cbdbc1f4729dd5e5507def72815ed0d8abc6249",
		"efb595a0178eb79a8df953f87c5148402a224cdf725e88c0146727c6aceadccd",
	}
	s.account = make([]account, len(keys))
	for i, key := range keys {
		b, err := hex.DecodeString(key)
		s.Require().NoError(err)
		s.account[i].key, err = crypto.ToECDSA(b)
		s.Require().NoError(err)
	}
	s.signer = signer(s.account[0])
	s.createFastNode()

	// Deploy an instance of Reserve Dollar.
	reserveAddress, tx, reserve, err := rsrABI.DeployReserveDollar(s.signer, s.node)
	s.requireTx(tx, err)()

	// Make the deployment account a minter, pauser, and freezer.
	s.requireTx(reserve.ChangeMinter(s.signer, s.account[0].address()))
	s.requireTx(reserve.Mint(s.signer, s.account[0].address(), big.NewInt(100000000)))

	// Deploy a new version of the Reserve Rights Token
	s.rsrAddress, tx, s.rsr, err = rsrABI.DeployReserveRightsToken(s.signer, s.node, reserveAddress, s.account[0].address())
	s.Require().NoError(err)

	// Deploy utility contract just for reading block time
	bytecode := "0x6080604052348015600f57600080fd5b5060918061001e6000396000f3fe6080604052348015600f57600080fd5b50600436106044577c0100000000000000000000000000000000000000000000000000000000600035046316ada54781146049575b600080fd5b604f6061565b60408051918252519081900360200190f35b429056fea165627a7a723058205524d6a0c4d80ea5535c2ea64615c2619a21518e242cb929275cbd678b04468f0029"
	utilABI, err := ethabi.JSON(strings.NewReader(`
	[{"constant":true,"inputs":[],"name":"time","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]
	`))
	s.Require().NoError(err)

	code, err := hex.DecodeString(strings.TrimPrefix(bytecode, "0x"))
	s.Require().NoError(err)

	_, tx, s.utilContract, err = bind.DeployContract(s.signer, utilABI, code, s.node)
	s.requireTx(tx, err)( /* assert zero events */ )
}

// backend is a wrapper around *backends.SimulatedBackend.
//
// *backends.SimulatedBackend requires blocks to be mined manually -- they are not automatically
// mined on every transaction. We want them to be automatically mined on every transaction, though,
// so we use this wrapper to do so.
type backend struct {
	*backends.SimulatedBackend
}

// SendTransaction overrides the function by the same name in *backends.SimulatedBackend,
// adding auto-mining for each transaction.
func (b backend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	defer b.Commit()
	return b.SimulatedBackend.SendTransaction(ctx, tx)
}

// AdjustTime overrides the function by the same name in *backends.SimulatedBackend,
// adding auto-committing.
func (b backend) AdjustTime(delta time.Duration) error {
	defer b.Commit()
	return b.SimulatedBackend.AdjustTime(delta)
}

// signer returns a *bind.TransactOpts that uses a's private key to sign transactions.
func signer(a account) *bind.TransactOpts {
	return bind.NewKeyedTransactor(a.key)
}

// account is a utility type to make it easier to convert from a private key to an address.
type account struct {
	key *ecdsa.PrivateKey
}

func bigInt(n uint32) *big.Int {
	return big.NewInt(int64(n))
}

// address returns the address corresponding to `a.key`.
func (a account) address() common.Address {
	return crypto.PubkeyToAddress(a.key.PublicKey)
}

func transfer(from common.Address, to common.Address, value *big.Int) rsrABI.ReserveRightsTokenTransfer {
	return rsrABI.ReserveRightsTokenTransfer{
		From:  from,
		To:    to,
		Value: value,
	}
}
