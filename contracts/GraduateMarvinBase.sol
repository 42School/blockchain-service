pragma solidity ^0.5.0;

import "./lib/tools/Ownable.sol";

contract GraduateMarvinBase is Ownable {

	event CreateGraduate(string login, string intraLevel, uint256 graduateId);
	event DeleteGraduate(string login);

	struct Graduate {
		string login;
		string firstName;
		string lastName;
		string intraLevel;
		string birthDate;
		string birthCity;
		string birthCountry;
		uint256 promoYears;
		uint256 graduateYears;
	}

	Graduate[] graduates;
	mapping(string => uint256) public loginToId;

	function _createGraduate(Graduate memory _newGraduate) internal onlyOwner {
		uint256 newGraduateId = graduates.push(_newGraduate) - 1;
		loginToId[_newGraduate.login] = newGraduateId;
		emit CreateGraduate(_newGraduate.login, _newGraduate.intraLevel, newGraduateId);
	}

	function _deleteGraduate(string memory _loginToDelete) internal {
		uint256 idToDelete = loginToId[_loginToDelete];
		delete graduates[idToDelete];
		delete loginToId[_loginToDelete];
		emit DeleteGraduate(_loginToDelete);
	}

}