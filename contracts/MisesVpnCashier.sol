// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MisesVpnCashier is Ownable {
    constructor(address initialOwner) Ownable(initialOwner) {
        // Your contract code here
    }

    /**
     * @dev 接收代币
     * @param tokenAddress ERC20代币合约地址
     * @param amount 接收的代币数量
     * @param uniqueKey 唯一键
     */
    function receiveTokens(address tokenAddress, uint256 amount, string memory uniqueKey) public {
        require(bytes(uniqueKey).length > 0, "Unique key must not be empty.");
        IERC20 token = IERC20(tokenAddress);
        require(token.transferFrom(msg.sender, address(this), amount), "Transfer failed");
    }

    /**
     * @dev 提取全部代币到所有者账户
     * @param tokenAddress ERC20代币合约地址
     */
    function withdrawToken(address tokenAddress) public onlyOwner {
        IERC20 token = IERC20(tokenAddress);
        uint256 balance = token.balanceOf(address(this));
        require(token.transfer(owner(), balance), "Transfer failed");
    }

    /**
     * @dev 转移所有权
     * @param newOwner 新的所有者地址
     */
    function transferOwnership(address newOwner) public override onlyOwner {
        require(newOwner != address(0), "Invalid address");
        transferOwnership(newOwner);
    }

    /**
     * @dev 查看合约账户下指定token的余额
     * @param tokenAddress ERC20代币合约地址
     * @return 合约账户下指定token的余额
     */
    function getTokenBalance(address tokenAddress) public view returns (uint256) {
        IERC20 token = IERC20(tokenAddress);
        return token.balanceOf(address(this));
    }
}
