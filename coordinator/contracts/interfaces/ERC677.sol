pragma solidity ^0.8.6;

// import { ERC20 as linkERC20 } from "./ERC20.sol";
// import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


interface ERC677 {
  function transferAndCall(address to, uint value, bytes memory data) external returns (bool success);

  event Transfer(address indexed from, address indexed to, uint value, bytes  data);
}
