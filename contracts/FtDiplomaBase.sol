/*
**	Autor: Louise Pieri
*/
// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.7.0;

contract	FtDiplomaBase {

	string public constant name = "42 Diploma";
	string public constant symbol = "42D";
	string public constant linkOfRepo = "github.com/lpieri/42-Diploma";
	address public constant ftPubAddress = 0x8A21Dc0aeC762cD85de81B2bcd396a9d5676cFD7;

	event Publish42Diploma(address ftPubAddress, string _link);
	// event CreateDiploma(address student, uint256 diplomaId);
	// event Transfer(address sender, address recipient, uint256 diplomaId);

	struct	Diploma {
		bytes32		level;
		bytes32[30]	skills;
		bytes32		signature;
	}

	Diploma[] private diplomas;
	mapping (address => uint256) private hashToDiploma;

	constructor () public {
		emit Publish42Diploma(ftPubAddress, linkOfRepo);
	}

	// function changeFtAddress(address sender, address recipient) public {
	// 	require(sender == ftPubAddress, "42-Diploma: The sender is not equal to the address of 42.");
	// 	require(sender != address(0), "42-Diploma: The sender is equal to zero.");
	// 	require(recipient != address(0), "42-Diploma: The sender is equal to zero.");

	// 	ftPubAddress = recipient;
	// 	emit PublishFtAddress(recipient);
	// }

	// function createDiploma(string memory _data, bytes32 _signature, address _studentAddress) public {
	// 	require(_studentAddress != address(0), "42-Diploma: The student address is equal at 0.");

	// 	diplomas.push(Diploma(_data, _signature));
	// 	uint256 newID = diplomas.length - 1;
	// 	addressToDiploma[_studentAddress] = newID;
	// 	emit CreateDiploma(_studentAddress, newID);
	// }

	// function transferDiploma(address sender, address recipient, uint256 diplomaId) public {
	// 	require(sender != address(0), "42-Diploma: Transfer sender the zero address.");
	// 	require(recipient != address(0), "42-Diploma: Transfer recipient the zero address.");
	// 	require(addressToDiploma[sender] == diplomaId, "42-Diploma: The address 'sender' is not the owner of the diploma.");
	// 	require(addressToDiploma[recipient] == 0, "42-Diploma: The 'recipient' address is the owner of a other diploma.");

	// 	addressToDiploma[sender] = 0;
	// 	addressToDiploma[recipient] = diplomaId;
	// 	emit Transfer(sender, recipient, diplomaId);
	// }

	// function getDiploma(address _studentAddress) public view returns (string memory data) {
	// 	require(_studentAddress != address(0), "42-Diploma: The student address is equal at 0.");

	// 	Diploma memory _getDiploma = diplomas[addressToDiploma[_studentAddress]];
	// 	string memory dataDiploma = _getDiploma.data;
	// 	return (dataDiploma);
	// }

}