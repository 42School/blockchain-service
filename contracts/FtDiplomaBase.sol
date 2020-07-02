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
	event CreateDiploma(bytes32 student, uint256 diplomaId);

	struct	Diploma {
		bytes32		level;
		bytes32[30]	skills;
		bytes32		signature;
	}

	Diploma[] private diplomas;
	mapping (bytes32 => uint256) private hashToDiploma;

	constructor () public {
		emit Publish42Diploma(ftPubAddress, linkOfRepo);
	}

	function createDiploma(bytes32 _level, bytes32[30] memory _skills, bytes32 _signature, bytes32 _studentHash) public {
		diplomas.push(Diploma(_level, _skills, _signature));
		uint256 newID = diplomas.length - 1;
		hashToDiploma[_studentHash] = newID;
		emit CreateDiploma(_studentHash, newID);
	}

	function getDiploma(bytes32 _studentHash) public view returns (bytes32 level, bytes32[30] memory skills) {
		Diploma memory _getDiploma = diplomas[hashToDiploma[_studentHash]];
		bytes32 levelDiploma = _getDiploma.level;
		bytes32[30] memory skillsDiploma = _getDiploma.skills;
		return (levelDiploma, skillsDiploma);
	}
}