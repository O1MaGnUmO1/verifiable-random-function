pragma solidity ^0.8.6;


import "../vrf/VRFConsumerBaseV2.sol";
import "../vrf/VRFCoordinatorV2.sol";
import "hardhat/console.sol";


contract ConsumerMock is VRFConsumerBaseV2 {
    VRFCoordinatorV2 public coordinator;
    uint256[] public random;
    // uint public constant totalSupply = 10**27;
    // string public constant name = "ChainLink Token";
    // uint8 public constant decimals = 18;
    // string public constant symbol = "LINK";

    //   function LinkToken()
    //     public
    //   {
    //     balances[msg.sender] = totalSupply;
    //   }
    constructor(address vrfCoordinator)VRFConsumerBaseV2(vrfCoordinator){
        coordinator = VRFCoordinatorV2(vrfCoordinator);
    }
    function getRandom() public view returns (uint256[] memory){
        return random;
    }
    function generateRandomWords(
        bytes32 keyHash,
        uint64 subId,
        uint16 requestConfirmations,
        uint32 callbackGasLimit,
        uint32 numWords
    ) external {
            coordinator.requestRandomWords(
            keyHash,
            subId,
            requestConfirmations,
            callbackGasLimit,
            numWords
            );
    }
    function fulfillRandomWords(uint256 requestId, uint256[] memory randomWords) internal virtual override {
        console.log("here", randomWords[0]);
        console.log("here", randomWords[1]);
        console.log("here", randomWords[2]);
        console.log("here", randomWords[3]);
        random = randomWords;
    }

    function test(uint256[] memory randomWords) external {
        random = randomWords;
    }
}
