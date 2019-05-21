package tests

// These tests MUST be run serially, because they rely on the block timestamp advancing predictably.

// These tests don't look at events.

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	rsrABI "github.com/reserve-protocol/rsr/abi"
)

func TestVesting(t *testing.T) {
	suite.Run(t, new(VestingTestSuite))
}

type VestingTestSuite struct {
	TestSuite

	vesting        *rsrABI.TokenVesting
	vestingAddress common.Address
	beneficiary    account
	amount         *big.Int
	vestTime       *big.Int
	vestStart      *big.Int
	currentTime    *big.Int
}

var (
	// Compile-time check that VestingTestSuite implements the interfaces we think it does.
	// If it does not implement these interfaces, then the corresponding setup and teardown
	// functions will not actually run.
	_ suite.BeforeTest       = &VestingTestSuite{}
	_ suite.SetupAllSuite    = &VestingTestSuite{}
	_ suite.TearDownAllSuite = &VestingTestSuite{}
)

// SetupSuite runs once, before all of the tests in the suite.
func (s *VestingTestSuite) SetupSuite() {
	s.setup()
	s.beneficiary = s.account[1]
	s.amount = big.NewInt(1000)
	s.vestTime = big.NewInt(100)
}

// TearDownSuite runs once, after all of the tests in the suite.
func (s *VestingTestSuite) TearDownSuite() {
}

// After this function is done running, vesting is set to start 10s after current timestamp
func (s *VestingTestSuite) BeforeTest(suiteName, testName string) {
	s.vestStart = big.NewInt(0).Add(s.currentTimestamp(), big.NewInt(30))
	vestingAddress, _, vesting, err := rsrABI.DeployTokenVesting(s.signer, s.node, s.beneficiary.address(), s.vestStart, big.NewInt(0), s.vestTime)
	s.Require().NoError(err)
	s.vestingAddress = vestingAddress
	s.vesting = vesting

	// Transfer rsr from account[0] to vesting address
	s.requireTx(s.rsr.Transfer(s.signer, s.vestingAddress, s.amount))
	s.currentTime = s.currentTimestamp()

	// Double check that token vesting contract is holding the balance
	s.assertBalance(s.vestingAddress, s.amount)
}

// Each test is going to assume there was a TokenVesting contract deployed 20s before the current timestamp.

func (s *VestingTestSuite) TestBeneficiary() {
	beneficiary, err := s.vesting.Beneficiary(nil)
	s.Nil(err)

	s.Equal(s.beneficiary.address(), beneficiary)
}

func (s *VestingTestSuite) TestCliff() {
	cliff, err := s.vesting.Cliff(nil)
	s.Nil(err)

	s.Equal(cliff.String(), s.vestStart.String())
}

func (s *VestingTestSuite) TestStart() {
	start, err := s.vesting.Start(nil)
	s.Nil(err)

	s.Equal(start.String(), s.vestStart.String())
}

func (s *VestingTestSuite) TestDuration() {
	duration, err := s.vesting.Duration(nil)
	s.Nil(err)

	s.Equal(duration.String(), s.vestTime.String())
}

func (s *VestingTestSuite) TestReleased() {
	released, err := s.vesting.Released(nil, s.rsrAddress)
	s.Nil(err)

	s.Equal(released.String(), big.NewInt(0).String())
}

func (s *VestingTestSuite) TestReleasable() {
	releasable, err := s.vesting.Releasable(nil, s.rsrAddress)
	s.Nil(err)

	s.Equal(releasable.String(), big.NewInt(0).String())
}

func (s *VestingTestSuite) TestReleasableAdvancingTime() {
	releasable, err := s.vesting.Releasable(nil, s.rsrAddress)
	s.Nil(err)

	// We are 10s before vesting starts, so releasable should be 0
	s.Equal(releasable.String(), big.NewInt(0).String())

	s.node.(backend).AdjustTime(0 * time.Second) // Advances 10s

	releasable, err = s.vesting.Releasable(nil, s.rsrAddress)
	s.Nil(err)

	// We are at the moment vesting starts, so releasable should be 0
	s.Equal(releasable.String(), big.NewInt(0).String())

	s.node.(backend).AdjustTime(0 * time.Second) // Advances 10s
	releasable, err = s.vesting.Releasable(nil, s.rsrAddress)
	s.Nil(err)

	// We are 10s after vesting starts, so releasable amount should be 1/10 of 1000
	s.Equal(releasable.String(), big.NewInt(100).String())

	s.node.(backend).AdjustTime(70 * time.Second) // Advances 80s
	releasable, err = s.vesting.Releasable(nil, s.rsrAddress)
	s.Nil(err)

	// We are 90s through the 100s vesting period, so releasable should be 9/10 of 1000
	s.Equal(releasable.String(), big.NewInt(900).String())

	s.node.(backend).AdjustTime(0 * time.Second) // Advances 10s
	releasable, err = s.vesting.Releasable(nil, s.rsrAddress)
	s.Nil(err)

	// We are at the exact end of the vesting period, so releasable should be the full amount
	s.Equal(releasable.String(), big.NewInt(1000).String())

	s.node.(backend).AdjustTime(1000000 * time.Second) // Advances way beyond the end of the vesting period
	releasable, err = s.vesting.Releasable(nil, s.rsrAddress)
	s.Nil(err)

	// We are way beyond the vesting period, so releasable should be the full amount
	s.Equal(releasable.String(), big.NewInt(1000).String())
}

func (s *VestingTestSuite) TestRelease() {
	// We are before vesting has started
	// The signer shouldn't matter, so we make it someone other than the beneficiary
	s.requireTxFails(s.vesting.Release(signer(s.account[1]), s.rsrAddress))

	s.node.(backend).AdjustTime(0 * time.Second) // Advances 10s

	// We are at the point vesting has started, which means the next block mined successfully will contain a timestamp 10s greater.
	s.requireTx(s.vesting.Release(s.signer, s.rsrAddress))
	s.assertBalance(s.beneficiary.address(), big.NewInt(100))

	s.node.(backend).AdjustTime(10 * time.Second) // Advance 20s

	// We should be able to release 3/10th now, leading to 4/10th total
	s.requireTx(s.vesting.Release(s.signer, s.rsrAddress))
	s.assertBalance(s.beneficiary.address(), big.NewInt(400))

	s.node.(backend).AdjustTime(1000000 * time.Second) // Advance way beyond the end

	// We should be able to release the rest
	s.requireTx(s.vesting.Release(s.signer, s.rsrAddress))
	s.assertBalance(s.beneficiary.address(), big.NewInt(1000))
}
