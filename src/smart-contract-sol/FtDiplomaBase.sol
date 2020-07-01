/*
**	Autor: Louise Pieri
*/
// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.6.8;

contract	FtDiplomaBase {

	string public constant name = "42 Diploma";
	string public constant symbol = "42D";
	address public ftPubAddress;
	string public linkOfRepo;

	event Publish42Diploma(address ftPubAddress, string linkOfRepo);
	// event CreateDiploma(address student, uint256 diplomaId);
	// event Transfer(address sender, address recipient, uint256 diplomaId);

	struct	Diploma {
		bytes32		level;
		bytes32[30]	skills;
		bytes32		signature;
	}

	Diploma[] private diplomas;
	mapping (address => uint256) private hashToDiploma;

	constructor (address _ftAddress, string memory _linkOfRepo) internal {
		require(_ftAddress != address(0), "42-Diploma: The _ftAddress is equal to zero.");

		ftPubAddress = _ftAddress;
		linkOfRepo = _linkOfRepo;
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