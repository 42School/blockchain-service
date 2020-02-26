pragma solidity >=0.6.1;

contract	FtDiplomaBase {

	string public constant name = "42 Diploma";
	string public constant symbol = "42D";

	event CreateDiploma(address student, uint256 diplomaId);

	struct	Diploma {
		string	data;
	}

	Diploma[] diplomas;
	mapping (address => uint256) AddressToDiploma;

	function _createDiploma(string memory _data, address _studentAddress) private {
		diplomas.push(Diploma(_data));
		uint256 newID = diplomas.length - 1;
		AddressToDiploma[_studentAddress] = newID;
		emit CreateDiploma(_studentAddress, newID);
	}

}