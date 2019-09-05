pragma solidity ^0.5.0;

import "./lib/tools/Ownable.sol";

contract GraduateMarvinBase is Ownable {

	event CreateGraduate(bytes32 intraLevel, uint256 graduateId);

	struct Graduate {
		bytes32 intraLevel;
		uint256 promoYears;
		uint256 graduateYears;
		uint64	flags;
	}

	Graduate[] graduates;
	mapping(uint256 => bytes32) public idToSignature;

	function _createGraduate(Graduate memory _newGraduate, bytes32 _signature) internal onlyOwner {
		uint256 newGraduateId = graduates.push(_newGraduate) - 1;
		idToSignature[newGraduateId] = _signature;
		emit CreateGraduate(_newGraduate.intraLevel, newGraduateId);
	}

}