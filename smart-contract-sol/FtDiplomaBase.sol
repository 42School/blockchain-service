pragma solidity >=0.6.1;

contract	FtDiplomaBase {

	string public constant name = "42 Diploma";
	string public constant symbol = "42D";

	event CreateDiploma(address student, uint256 diplomaId);
	event Transfer(address sender, address recipient, uint256 diplomaId);

	struct	Diploma {
		string	data;
		bytes32	signature;
	}

	Diploma[] private diplomas;
	mapping (address => uint256) private addressToDiploma;

	function _createDiploma(string memory _data, bytes32 _signature, address _studentAddress) private {
		require(_studentAddress != address(0), "42-Diploma: The student address is equal at 0.");
		diplomas.push(Diploma(_data, _signature));
		uint256 newID = diplomas.length - 1;
		addressToDiploma[_studentAddress] = newID;
		emit CreateDiploma(_studentAddress, newID);
	}

	function _getDiploma(address _studentAddress) private view returns (string memory data) {
		require(_studentAddress != address(0), "42-Diploma: The student address is equal at 0.");
		Diploma memory getDiploma = diplomas[addressToDiploma[_studentAddress]];
		string memory dataDiploma = getDiploma.data;
		return (dataDiploma);
	}

	function _transferDiploma(address sender, address recipient, uint256 diplomaId) private {
		require(sender != address(0), "42-Diploma: Transfer sender the zero address.");
		require(recipient != address(0), "42-Diploma: Transfer recipient the zero address.");
		require(addressToDiploma[sender] == diplomaId, "42-Diploma: The address 'sender' is not the owner of the diploma.");
		require(addressToDiploma[recipient] == 0, "42-Diploma: The 'recipient' address is the owner of a other diploma.");

		addressToDiploma[sender] = 0;
		addressToDiploma[recipient] = diplomaId;
		emit Transfer(sender, recipient, diplomaId);
	}

}