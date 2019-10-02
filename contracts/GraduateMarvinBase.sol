pragma solidity ^0.5.0;

import "./lib/tools/Ownable.sol";

contract GraduateMarvinBase is Ownable {

	event CreateGraduate(bytes32 login, bytes32 intraLevel, uint256 graduateId);
	event DeleteGraduate(bytes32 login);

	struct Graduate {
		bytes32 login;
		bytes32 firstName;
		bytes32 lastName;
		bytes32 intraLevel;
		bytes32 birthDate;
		bytes32 birthCity;
		bytes32 birthCountry;
		uint256 promoYears;
		uint256 graduateYears;
	}

	Graduate[] graduates;
	mapping(bytes32 => uint256) public loginToId;
	mapping(uint256 => string) public idToSignature;

	function _createGraduate(Graduate memory _newGraduate, string memory _signature) internal onlyOwner {
		uint256 newGraduateId = graduates.push(_newGraduate) - 1;
		loginToId[_newGraduate.login] = newGraduateId;
		idToSignature[newGraduateId] = _signature;
		emit CreateGraduate(_newGraduate.login, _newGraduate.intraLevel, newGraduateId);
	}

	function _deleteGraduate(bytes32 _loginToDelete) internal {
		uint256 idToDelete = loginToId[_loginToDelete];
		delete graduates[idToDelete];
		delete loginToId[_loginToDelete];
		emit DeleteGraduate(_loginToDelete);
	}

}