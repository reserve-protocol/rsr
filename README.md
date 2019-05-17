# RSR
This repo contains the code for the Reserve Rights token (RSR) and some associated contracts: TokenVesting.sol, and SlowWallet.sol. 

## ReserveRights.sol

ReserveRights.sol is an ERC20 token that represents the cryptographic right to purchase the excess Reserve tokens produced by the Reserve protocol. Alongside ERC-20 and full-token pausing, it supports specifying addresses deliberately to be locked forever. See https://reserve.org for more information.

[Etherscan](https://etherscan.io/token/0x8762db106B2c2A0bccB3A80d1Ed41273552616E8)

### What does the code do and why?

We've extended ERC20Pausable, which makes at a pausable ERC20 token. 

We made it pausable because we want to be able to pause the contract if we do a hard fork. A common scam is to sell somebody tokens of a deprecated address. 


#### Constructor 

Its constructor copies over all balances from our "trove" contract, which was used to track team member allocations and early investor allocations. It also locks all these tokens permanently, until the next hard fork. This is to prevent team members and investors from selling their tokens into the market until mainnet launch. It also copies over the balance of our genesis wallet. For this address, _it does not lock the account_. This is because we need to transfer tokens out of the genesis address before the mainnet launch. 

#### lockMyTokensForever

Can be called by anyone to lock msg.sender. It's not clear why you would want to do this, but we wanted to have the option of locking more accounts in the future. This could happen either with future investors, or with current team members who haven't received their tokens yet. 

#### Transfer + TransferFrom

The only other thing in this contract is a wrapper around the ERC20Pausable `transfer` and `transferFrom` functions. The only thing that happens is that we mandate that the account must not be locked in order for the transfer to occur, and we also return the bool result of the transfer so it can be used with SafeERC20. 


## TokenVesting.sol

TokenVesting.sol is a contract we use to distribute tokens according to a linear vesting schedule. It is a minor variant of the OpenZeppelin contract from their version 2.0.0 August 2018 audit, taken at commit 4115686b4f8c1abf29f1f855eb15308076159959. 

### What does this code do and why?

This code provides the ability for tokens to be transferred to a beneficiary's address according to a linear vesting schedule. 

#### release

At any point in time anyone on the blockchain can call `release` in order to transfer all available vested balance to the beneficiary's address. You could for example call this every day, in order to do day-by-day vesting. 

#### releasable

We added a public view that returns the total amount releasable by the contract. This is just a nicety, since there isn't otherwise a good way to determine how many tokens can potentially be released at some arbitrary time. 

#### revoke

We don't invoke the constructor with revocable equal to true, so this function isn't used. 

## SlowWallet.sol

SlowWallet.sol is a contract we use to prevent Reserve from transferring tokens out of the Reserve main wallet unexpectedly. More precisely, it delays transfers out of the wallet by 2 weeks. It is based off a contract that was part of our audit with Trail of Bits. 

### What does this code do and why?

This code tracks a list of proposals, where each proposal is a time-gated `transfer` function. The owner of the contract can add proposals to the list by calling `propose`, and then after 2 weeks has passed, can call `confirm` in order to carry out the transfer. They can also `cancel` the propsal, and they can also `cancelAll`. This feature is important because it means if an attacker gains a copy of the key, while they can propose transfers, we can easily cancel any number of them at a low cost. 

