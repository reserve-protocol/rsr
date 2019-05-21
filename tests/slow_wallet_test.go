package tests

// This test was based off of our MintAndBurnAdmin tests. The comments, therefore, should not be trusted.

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	rsrABI "github.com/reserve-protocol/rsr/abi"
)

var delayInSeconds = bigInt(60 * 60 * 24 * 28)

func TestSlowWallet(t *testing.T) {
	suite.Run(t, new(SlowWalletTestingSuite))
}

type SlowWalletTestingSuite struct {
	TestSuite

	walletContract        *rsrABI.SlowWallet
	walletContractAddress common.Address
	adminAccount          account
	adminSigner           *bind.TransactOpts

	amount *big.Int
}

var (
	// Compile-time check that SlowWalletTestingSuite implements the interfaces we think it does.
	// If it does not implement these interfaces, then the corresponding setup and teardown
	// functions will not actually run.
	_ suite.BeforeTest       = &SlowWalletTestingSuite{}
	_ suite.SetupAllSuite    = &SlowWalletTestingSuite{}
	_ suite.TearDownAllSuite = &SlowWalletTestingSuite{}
)

// SetupSuite runs once, before all of the tests in the suite.
func (s *SlowWalletTestingSuite) SetupSuite() {
	s.setup()
	s.amount = bigInt(1000)
}

// TearDownSuite runs once, after all of the tests in the suite.
func (s *SlowWalletTestingSuite) TearDownSuite() {
}

// After this function is done running, vesting is set to start 10s after current timestamp
func (s *SlowWalletTestingSuite) BeforeTest(suiteName, testName string) {
	s.adminAccount = s.account[len(s.account)-1]
	s.adminSigner = signer(s.adminAccount)

	walletAddress, _, walletContract, err := rsrABI.DeploySlowWallet(s.adminSigner, s.node, s.rsrAddress)
	s.Require().NoError(err)
	s.walletContractAddress = walletAddress
	s.walletContract = walletContract

	// Transfer rsr from account[0] to wallet address
	s.requireTx(s.rsr.Transfer(s.signer, s.walletContractAddress, s.amount))

	// Double check that token vesting contract is holding the balance
	s.assertBalance(s.walletContractAddress, s.amount)

	s.logParsers = map[common.Address]logParser{
		s.walletContractAddress: s.walletContract,
		s.rsrAddress:            s.rsr,
	}
}

func (s *SlowWalletTestingSuite) TestAdminCanTransfer() {
	recipient := common.BigToAddress(bigInt(1))
	amount := bigInt(100)

	// Propose a new mint.
	s.requireTx(s.walletContract.Propose(s.adminSigner, recipient, amount, "notes"))(
		s.transferProposed(0, recipient, amount, "notes"),
	)

	// Trying to confirm it immediately should fail.
	s.requireTxFails(s.walletContract.Confirm(s.adminSigner, bigInt(0), recipient, amount))

	// Advance time.
	s.Require().NoError(s.node.(backend).AdjustTime(672 * time.Hour)) // 4 weeks

	// Trying to confirm it should now succeed.
	s.requireTx(s.walletContract.Confirm(s.adminSigner, bigInt(0), recipient, amount))(
		s.transferConfirmed(0, recipient, amount, "notes"),
		transfer(s.walletContractAddress, recipient, amount),
	)

	// The mint should have happened.
	s.assertBalance(recipient, amount)

	// Trying to confirm a second time should fail.
	s.requireTxFails(s.walletContract.Confirm(s.adminSigner, bigInt(0), recipient, amount))
}

func (s *SlowWalletTestingSuite) TestAdminCanCancelTransfer() {
	recipient := common.BigToAddress(bigInt(2))
	amount := bigInt(100)

	// Propose a new mint.
	s.requireTx(s.walletContract.Propose(s.adminSigner, recipient, amount, "notes"))(
		s.transferProposed(0, recipient, amount, "notes"),
	)

	// And then cancel that minting.
	s.requireTx(s.walletContract.Cancel(s.adminSigner, bigInt(0), recipient, amount))(
		s.transferCancelled(0, recipient, amount, "notes"),
	)

	// Advance time.
	s.Require().NoError(s.node.(backend).AdjustTime(672 * time.Hour)) // 4 weeks

	// Trying to confirm it should now fail, even though time has advanced.
	s.requireTxFails(s.walletContract.Confirm(s.adminSigner, bigInt(0), recipient, amount))

	// The mint should not have happened.
	s.assertBalance(recipient, bigInt(0))

	// Trying to confirm a second time should fail.
	s.requireTxFails(s.walletContract.Confirm(s.adminSigner, bigInt(0), recipient, amount))
}

func (s *SlowWalletTestingSuite) TestConstructor() {
	addr, err := s.walletContract.Owner(nil)
	s.Nil(err)

	s.Equal(s.adminAccount.address(), addr, "admin should be set to the walletContractAddress")

	addr, err = s.walletContract.Token(nil)
	s.Nil(err)

	s.Equal(s.rsrAddress, addr, "reserve should be set to the Reserve contract address")
}

func (s *SlowWalletTestingSuite) TestPropose() {
	recipient := common.BigToAddress(bigInt(3))
	amount := bigInt(100)
	futureAmount := bigInt(1000)

	// Confirm that the length is 0.
	length, err := s.walletContract.ProposalsLength(nil)
	s.Nil(err)
	s.Equal(length.String(), bigInt(0).String())

	// Trying to propose as someone other than the admin signer should fail.
	s.requireTxFails(s.walletContract.Propose(s.signer, recipient, amount, "notes"))

	// Propose a new mint.
	s.requireTx(s.walletContract.Propose(s.adminSigner, recipient, amount, "notes"))(
		s.transferProposed(0, recipient, amount, "notes"),
	)

	// Retrieving the 0th proposal should now work.
	_, err = s.walletContract.Proposals(nil, bigInt(0))
	s.NoError(err)

	// Advance time by 12 hours.
	s.Require().NoError(s.node.(backend).AdjustTime(12 * time.Hour))

	// Propose a second mint.
	s.requireTx(s.walletContract.Propose(s.adminSigner, recipient, futureAmount, "notes"))(
		s.transferProposed(1, recipient, futureAmount, "notes"),
	)

	// Proposals should now contain the first mint proposal at index 0
	proposalOne, err := s.walletContract.Proposals(nil, bigInt(0))
	s.Nil(err)
	proposalTwo, err := s.walletContract.Proposals(nil, bigInt(1))
	s.Nil(err)

	// Check proposal one is correct
	s.Equal(recipient, proposalOne.Destination, "0th proposal address should be recipient")
	s.Equal(amount.String(), proposalOne.Value.String(), "0th proposal value should be amount")

	// Check proposal two is correct
	s.Equal(recipient, proposalTwo.Destination, "1th proposal address should be recipient")
	s.Equal(futureAmount.String(), proposalTwo.Value.String(), "1th proposal value should be amount")

	// Confirm times are separated by 12 hours plus the blocktime (20 s)
	diff := bigInt(0).Sub(proposalTwo.Time, proposalOne.Time)
	s.Equal(bigInt(12*3600+20).String(), diff.String(), "proposals should be separated by exactly 12 hours")
}

func (s *SlowWalletTestingSuite) TestCancel() {
	recipient := common.BigToAddress(bigInt(4))
	amount := bigInt(100)
	index := bigInt(0)

	// Create a proposal.
	s.requireTx(s.walletContract.Propose(s.adminSigner, recipient, amount, "notes"))(
		s.transferProposed(0, recipient, amount, "notes"),
	)

	// Trying to cancel as someone other than the admin should fail.
	s.requireTxFails(s.walletContract.Cancel(s.signer, index, recipient, amount))

	// Trying to cancel with mismatching recipient should fail.
	s.requireTxFails(s.walletContract.Cancel(s.adminSigner, index, common.BigToAddress(bigInt(2)), amount))

	// Trying to cancel with mismatching amount should fail.
	s.requireTxFails(s.walletContract.Cancel(s.adminSigner, index, recipient, bigInt(1)))

	// Trying to cancel with mismatching index should fail.
	s.requireTxFails(s.walletContract.Cancel(s.adminSigner, bigInt(1), recipient, amount))

	// Should be able to cancel proposal when supplied properly.
	s.requireTx(s.walletContract.Cancel(s.adminSigner, index, recipient, amount))(
		s.transferCancelled(0, recipient, amount, "notes"),
	)

	// Should be marked as closed.
	proposal, err := s.walletContract.Proposals(nil, index)
	s.NoError(err)
	s.True(proposal.Closed)

	// Should not be able to cancel a second time.
	s.requireTxFails(s.walletContract.Cancel(s.adminSigner, index, recipient, amount))
}

func (s *SlowWalletTestingSuite) TestConfirm() {
	recipient := common.BigToAddress(bigInt(5))
	amount := bigInt(100)
	index := bigInt(0)

	// Create a proposal.
	s.requireTx(s.walletContract.Propose(s.adminSigner, recipient, amount, "notes"))(
		s.transferProposed(0, recipient, amount, "notes"),
	)

	// Should not be able to confirm until time has passed.
	s.requireTxFails(s.walletContract.Confirm(s.adminSigner, index, recipient, amount))

	// Advance time.
	s.Require().NoError(s.node.(backend).AdjustTime(672 * time.Hour))

	// Trying to confirm as someone other than the admin should fail.
	s.requireTxFails(s.walletContract.Confirm(s.signer, index, recipient, amount))

	// Trying to confirm with mismatching recipient should fail.
	s.requireTxFails(s.walletContract.Confirm(s.adminSigner, index, common.BigToAddress(bigInt(2)), amount))

	// Trying to confirm with mismatching amount should fail.
	s.requireTxFails(s.walletContract.Confirm(s.adminSigner, index, recipient, bigInt(1)))

	// Trying to confirm with mismatching index should fail.
	s.requireTxFails(s.walletContract.Confirm(s.adminSigner, bigInt(1), recipient, amount))

	// Confirm proposal.
	s.requireTx(s.walletContract.Confirm(s.adminSigner, index, recipient, amount))(
		s.transferConfirmed(0, recipient, amount, "notes"),
		transfer(s.walletContractAddress, recipient, amount),
	)

	// Should be marked as closed.
	proposal, err := s.walletContract.Proposals(nil, index)
	s.NoError(err)
	s.True(proposal.Closed)

	// Should not be able to confirm proposal a second time.
	s.requireTxFails(s.walletContract.Confirm(s.adminSigner, index, recipient, amount))

	// Confirm mint went through.
	s.assertBalance(recipient, amount)
}

func (s *SlowWalletTestingSuite) TestVoidAll() {
	// Create several proposals.
	for i := 0; i < 5; i++ {
		recipient := common.BigToAddress(bigInt(uint32(i)))
		value := bigInt(uint32((i + 1) * 100))
		s.requireTx(s.walletContract.Propose(s.adminSigner, recipient, value, "notes"))(
			s.transferProposed(i, recipient, value, "notes"),
		)
	}

	// Confirm that voidAll cannot be called by someone unauthorized.
	s.requireTxFails(s.walletContract.VoidAll(signer(s.account[2])))

	// Cancel all.
	s.requireTx(s.walletContract.VoidAll(s.adminSigner))(
		rsrABI.SlowWalletAllTransfersCancelled{},
	)

	// Confirm that the length is now 0.
	length, err := s.walletContract.ProposalsLength(nil)
	s.Nil(err)
	s.Equal(length.String(), bigInt(0).String())
}

func (s *SlowWalletTestingSuite) transferProposed(
	i int, addr common.Address, value *big.Int, notes string,
) rsrABI.SlowWalletTransferProposed {
	return rsrABI.SlowWalletTransferProposed{
		Index:       bigInt(uint32(i)),
		Destination: addr,
		Value:       value,
		Notes:       notes,
		DelayUntil:  new(big.Int).Add(s.currentTimestamp(), delayInSeconds),
	}
}

func (s *SlowWalletTestingSuite) transferConfirmed(
	i int, addr common.Address, value *big.Int, notes string,
) rsrABI.SlowWalletTransferConfirmed {
	return rsrABI.SlowWalletTransferConfirmed{
		Index:       bigInt(uint32(i)),
		Destination: addr,
		Value:       value,
		Notes:       notes,
	}
}

func (s *SlowWalletTestingSuite) transferCancelled(
	i int, addr common.Address, value *big.Int, notes string,
) rsrABI.SlowWalletTransferCancelled {
	return rsrABI.SlowWalletTransferCancelled{
		Index:       bigInt(uint32(i)),
		Destination: addr,
		Value:       value,
		Notes:       notes,
	}
}
