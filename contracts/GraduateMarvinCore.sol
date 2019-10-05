pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./GraduateMarvinBase.sol";

contract GraduateMarvinCore is GraduateMarvinBase {

	string public constant name = "GraduateMarvin";
	string public constant symbol = "42G";

	address public newContractAddress;

	function createGraduate(Graduate calldata newGraduate, string calldata signature) external onlyOwner {
		require(loginToId[newGraduate.login] == 0, 'The student graduate already exists');
		_createGraduate(newGraduate, signature);
	}

	function getGraduate(bytes32 _loginToGet) external view returns (Graduate memory graduate, string memory signature) {
		uint256 id = loginToId[_loginToGet];
		graduate = graduates[id];
		signature = idToSignature[id];
	}
}