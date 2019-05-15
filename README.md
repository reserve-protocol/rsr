# RSR
This repo contains the code for the Reserve Rights token (RSR) and some associated contracts: TokenVesting.sol, and SlowWallet.sol. 

## ReserveRights.sol

ReserveRights.sol is an ERC20 token that represents the cryptographic right to purchase the excess Reserve tokens produced by the Reserve protocol. See https://reserve.org for more information.

## TokenVesting.sol

TokenVesting.sol is a contract we use to distribute tokens according to a linear vesting schedule. It is a variant of the OpenZeppelin contract from their version 2.0.0 August 2018 audit, taken at commit 4115686b4f8c1abf29f1f855eb15308076159959. 

## SlowWallet.sol

SlowWallet.sol is a contract we use to prevent Reserve from transferring tokens out of the Reserve main wallet unexpectedly. More precisely, it delays transfers out of the wallet by 2 weeks. 
