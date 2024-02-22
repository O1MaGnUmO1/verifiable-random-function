pragma solidity ^0.8.6;


interface ERC677Receiver {
  function onTokenTransfer(address _sender, uint _value, bytes memory _data) external;
}
