pragma solidity ^0.5.8;

/**
 * @title The standard ERC20 interface
 * @dev see https://eips.ethereum.org/EIPS/eip-20
 */
interface IERC20 {
    function transfer(address, uint256) external returns (bool);
    function approve(address, uint256) external returns (bool);
    function transferFrom(address, address, uint256) external returns (bool);
    function totalSupply() external view returns (uint256);
    function balanceOf(address) external view returns (uint256);
    function allowance(address, address) external view returns (uint256);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed holder, address indexed spender, uint256 value);
}

/// @title Time-delayed ERC-20 wallet contract.
/// Can only transfer tokens after publicly recording the intention to do so
/// at least two weeks in advance.
contract SlowWallet {

    // TYPES

    struct TransferProposal {
        address destination;
        uint256 value;
        uint256 time;
        string notes;
        bool closed;
    }

    // DATA

    IERC20 public token;
    uint256 public constant delay = 2 weeks;
    address public owner;

    TransferProposal[] public proposals;

    // EVENTS

    event TransferProposed(
        uint256 index,
        address indexed destination,
        uint256 value,
        uint256 delayUntil,
        string notes
    );
    event TransferConfirmed(uint256 index, address indexed destination, uint256 value);
    event TransferCancelled(uint256 index, address indexed destination, uint256 value);
    event AllTransfersCancelled();

    // FUNCTIONALITY

    constructor(address tokenAddress) public {
        token = IERC20(tokenAddress);
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "must be owner");
        _;
    }

    /// The earliest that funds could possibly move from this account,
    /// expressed in seconds after the current block.
    function earliestPossibleTransfer() external view returns (uint256) {
        uint256 minimumDelay = delay;
        for (uint i = 0; i < proposals.length; i++) {
            if (proposals[i].closed) {
                continue;
            }
            if (proposals[i].time < now) {
                return 0;
            }
            // solium-disable-next-line security/no-block-members
            uint256 proposalDelay = proposals[i].time - now;
            if (proposalDelay < minimumDelay) {
                minimumDelay = proposalDelay;
            }
        }
        return minimumDelay;
    }

    /// Propose a new transfer, which can be confirmed after two weeks.
    function propose(address destination, uint256 value, string calldata notes) external onlyOwner {
        // Delay by at least two weeks.
        // We are relying on block.timestamp for this, and aware of the possibility of its
        // manipulation by miners. But we are working at a timescale that is already much
        // longer than the variance in timestamps we have observed and expect in the future,
        // so we are satisfied with this choice.
        // solium-disable-next-line security/no-block-members
        uint256 delayUntil = now + delay;

        proposals.push(TransferProposal({
            destination: destination,
            value: value,
            time: delayUntil,
            notes: notes,
            closed: false
        }));

        emit TransferProposed(proposals.length - 1, destination, value, delayUntil, notes);
    }

    /// Cancel a proposed transfer.
    function cancel(uint256 index, address addr, uint256 value) external onlyOwner {
        // Check authorization.
        requireMatchingOpenProposal(index, addr, value);

        // Cancel transfer.
        proposals[index].closed = true;
        emit TransferCancelled(index, addr, value);
    }

    /// Cancel all transfer proposals.
    function cancelAll() external onlyOwner {
        proposals.length = 0;
        emit AllTransfersCancelled();
    }

    /// Confirm and execute a proposed transfer, if enough time has passed since it was proposed.
    function confirm(uint256 index, address destination, uint256 value) external onlyOwner {
        // Check authorization.
        requireMatchingOpenProposal(index, destination, value);

        // See commentary above about using `now`.
        // solium-disable-next-line security/no-block-members
        require(proposals[index].time < now, "too early");

        // Record execution of transfer.
        proposals[index].closed = true;
        emit TransferConfirmed(index, destination, value);

        // Proceed with execution of transfer.
        token.transfer(destination, value);
    }

    /// Throw unless the given transfer proposal exists and matches `destination` and `value`.
    function requireMatchingOpenProposal(uint256 index, address destination, uint256 value) private view {
        require(!proposals[index].closed, "transfer already closed");

        // Slither reports "dangerous strict equality" for each of these, but it's OK.
        // These equalities are to confirm that the transfer entered is equal to the
        // matching previous transfer. We're vetting data entry; strict equality is appropriate.
        require(proposals[index].destination == destination, "destination mismatched");
        require(proposals[index].value == value, "value mismatched");
    }
}
