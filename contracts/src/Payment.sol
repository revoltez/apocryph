// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol";

using SafeERC20 for IERC20;

contract Payment {
    error AlreadyExists();
    error DoesNotExist();
    error AmountRequired();
    error ChannelLocked();
    error InsufficientFunds();

    event UnlockTimerStarted(
        address indexed publisher, address indexed provider, IERC20 indexed token, uint256 unlockedAt
    );
    event Deposited(address indexed publisher, address indexed provider, IERC20 indexed token, uint256 depositAmount);
    event Unlocked(address indexed publisher, address indexed provider, IERC20 indexed token, uint256 unlockedAmount);
    event Withdrawn(address indexed publisher, address indexed provider, IERC20 indexed token, uint256 withdrawnAmount);
    event ChannelClosed(address indexed publisher, address indexed provider, IERC20 indexed token);

    struct Channel {
        uint256 investedByPublisher;
        uint256 withdrawnByProvider;
        uint256 unlockTime; // minimum time in seconds needed to unlock the funds
        uint256 unlockedAt; // time @ unlock + unlockTime
    }

    // publisher => provider => token => funds
    mapping(address => mapping(address => mapping(IERC20 => Channel))) public channels;

    // called by publisher to create a new payment channel; must approve a withdraw by this contract's address
    function createChannel(address provider, IERC20 token, uint256 unlockTime, uint256 initialAmount) public {
        if (initialAmount == 0) revert AmountRequired();
        address publisher = msg.sender;
        Channel storage channel = channels[publisher][provider][token];
        if (channel.investedByPublisher != 0) revert AlreadyExists();
        assert(channel.withdrawnByProvider == 0);
        channel.investedByPublisher = initialAmount;
        channel.unlockTime = unlockTime;

        emit Deposited(publisher, provider, token, initialAmount);

        token.safeTransferFrom(msg.sender, address(this), initialAmount);
    }

    // add more funds to the payment channel
    function deposit(address provider, IERC20 token, uint256 amount) public {
        if (amount == 0) revert AmountRequired();
        address publisher = msg.sender;
        Channel storage channel = channels[publisher][provider][token];

        channel.investedByPublisher = channel.investedByPublisher + amount;
        channel.unlockedAt = 0;

        emit Deposited(publisher, provider, token, amount);

        token.safeTransferFrom(msg.sender, address(this), amount);
    }

    // initiate the process of unlocking the funds stored in the contract
    function unlock(address provider, IERC20 token) public {
        address publisher = msg.sender;
        Channel storage channel = channels[publisher][provider][token];
        if (channel.investedByPublisher == 0) revert DoesNotExist();

        uint256 newUnlockedAt = block.timestamp + channel.unlockTime;
        if (channel.unlockedAt == 0 || channel.unlockedAt < newUnlockedAt) {
            channel.unlockedAt = newUnlockedAt;
            emit UnlockTimerStarted(publisher, provider, token, newUnlockedAt);
        }
    }

    // transfer the now-unlocked funds back to the publisher
    function withdrawUnlocked(address provider, IERC20 token) public {
        address publisher = msg.sender;
        Channel storage channel = channels[publisher][provider][token];
        if (channel.unlockedAt == 0 || block.timestamp < channel.unlockedAt) revert ChannelLocked();

        uint256 leftoverFunds = channel.investedByPublisher - channel.withdrawnByProvider;
        if (leftoverFunds == 0) revert AmountRequired();

        channel.investedByPublisher = channel.withdrawnByProvider;

        emit Unlocked(publisher, provider, token, leftoverFunds);

        token.safeTransfer(publisher, leftoverFunds);
    }

    // withdrawUnlockedFunds and destroy all previous traces of the channel's existence
    function closeChannel(address provider, IERC20 token) public {
        address publisher = msg.sender;
        Channel storage channel = channels[publisher][provider][token];
        if (channel.unlockedAt == 0 || block.timestamp < channel.unlockedAt) revert ChannelLocked();

        uint256 leftoverFunds = channel.investedByPublisher - channel.withdrawnByProvider;
        delete channels[publisher][provider][token];

        if (leftoverFunds != 0) emit Unlocked(publisher, provider, token, leftoverFunds);
        emit ChannelClosed(publisher, provider, token);

        if (leftoverFunds != 0) token.safeTransfer(publisher, leftoverFunds);
    }

    // allows the provider to withdraw as many tokens as would be needed to reach totalWithdrawlAmount since the opening of the channel
    function withdrawUpTo(address publisher, IERC20 token, uint256 totalWithdrawlAmount, address transferAddress)
        public
    {
        if (transferAddress == address(0)) {
            transferAddress = msg.sender;
        }

        address provider = msg.sender;
        Channel storage channel = channels[publisher][provider][token];
        if (totalWithdrawlAmount > channel.investedByPublisher) revert InsufficientFunds();
        if (totalWithdrawlAmount <= channel.withdrawnByProvider) revert AmountRequired();

        uint256 transferAmonut = totalWithdrawlAmount - channel.withdrawnByProvider;
        channel.withdrawnByProvider = totalWithdrawlAmount;

        emit Withdrawn(publisher, provider, token, transferAmonut);

        if (channel.unlockedAt != 0) {
            channel.unlockedAt = block.timestamp;
        }

        token.safeTransfer(transferAddress, transferAmonut);
    }

    // allows the provider to withdraw amount more tokens
    function withdraw(address publisher, IERC20 token, uint256 amount, address transferAddress) public {
        withdrawUpTo(
            publisher, token, channels[publisher][msg.sender][token].withdrawnByProvider + amount, transferAddress
        );
    }

    // allows one to check the amount of as-of-yet unclaimed tokens
    function available(address publisher, address provider, IERC20 token) public view returns (uint256) {
        Channel storage channel = channels[publisher][provider][token];
        return channel.investedByPublisher - channel.withdrawnByProvider;
    }
}
