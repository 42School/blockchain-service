pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./GraduateMarvinBase.sol";

contract GraduateMarvinCore is GraduateMarvinBase {

	string public constant name = "GraduateMarvin";
	string public constant symbol = "42G";

	address public newContractAddress;

	function createGraduate(Graduate calldata newGraduate, bytes32 signature) external onlyOwner {
		_createGraduate(newGraduate, signature);
	}

	function getGraduate(uint256 _idToGet) external view returns (Graduate memory graduate, bytes32 signature) {
		graduate = graduates[_idToGet];
		signature = idToSignature[_idToGet];
	}
}