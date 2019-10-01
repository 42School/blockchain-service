pragma solidity ^0.5.0;

import "./GraduateMarvinBase.sol";

contract GraduateMarvinCore is GraduateMarvinBase {

	string public constant name = "GraduateMarvin";
	string public constant symbol = "42G";

	address public newContractAddress;

	function createGraduate(
		bytes32 _login,
		bytes32 _firstName,
		bytes32 _lastName,
		bytes32 _intraLevel,
		bytes32 _birthDate,
		bytes32 _birthCity,
		bytes32 _birthCountry,
		uint256 _promoYears,
		uint256 _graduateYears
	) external onlyOwner {
			Graduate memory newGraduate = Graduate({
				login: _login,
				firstName: _firstName,
				lastName: _lastName,
				intraLevel: _intraLevel,
				birthDate: _birthDate,
				birthCity: _birthCity,
				birthCountry: _birthCountry,
				promoYears: _promoYears,
				graduateYears: _graduateYears
			});
			_createGraduate(newGraduate);
		}

	function deleteGraduate(bytes32 _login) external onlyOwner {
		_deleteGraduate(_login);
	}

	function getGraduate(bytes32 _loginToGet) external view returns (
		bytes32 _login,
		bytes32 _firstName,
		bytes32 _lastName,
		bytes32 _intraLevel,
		bytes32 _birthDate,
		bytes32 _birthCity,
		bytes32 _birthCountry,
		uint256 _promoYears,
		uint256 _graduateYears
	) {
		uint256 id = loginToId[_loginToGet];
		Graduate memory graduate = graduates[id];
		return (
			graduate.login, graduate.firstName, graduate.lastName, graduate.intraLevel, graduate.birthDate, graduate.birthCity,
			graduate.birthCountry, graduate.promoYears, graduate.graduateYears
		);
	}
}