/*
**	Autor: Louise Pieri
*/
// SPDX-License-Identifier: MIT
pragma solidity >=0.5.8 <0.7.0;

contract	FtDiplomaBase {

	string public constant name = "42 Diploma";
	string public constant symbol = "42D";
	string public constant linkOfRepo = "github.com/lpieri/42-Diploma";
	address public constant ftPubAddress = 0x8A21Dc0aeC762cD85de81B2bcd396a9d5676cFD7;

	event Publish42Diploma(address ftPubAddress, string _link);
	event CreateDiploma(bytes32 student, uint256 diplomaId);

	struct	Sign {
		uint8		v;
		bytes32		r;
		bytes32		s;
	}

	struct	Diploma {
		uint64		level;
		uint64[30]	skills;
		Sign		signature;
	}

	Diploma[] private diplomas;
	mapping (bytes32 => uint256) private hashToDiploma;

	constructor () public {
		emit Publish42Diploma(ftPubAddress, linkOfRepo);
	}

	function createDiploma(uint64 _level, uint64[30] memory _skills, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _studentHash) public {
		require(ecrecover(_studentHash, _v, _r, _s) == ftPubAddress, "FtDiplomaBase: Is not 42 sign this diploma");
		diplomas.push(Diploma(_level, _skills, Sign(_v, _r, _s)));
		uint256 newID = diplomas.length - 1;
		hashToDiploma[_studentHash] = newID;
		emit CreateDiploma(_studentHash, newID);
	}

	function getDiploma(bytes32 _studentHash) public view returns (uint64 level, uint64[30] memory skills) {
		Diploma memory _getDiploma = diplomas[hashToDiploma[_studentHash]];
		uint64 levelDiploma = _getDiploma.level;
		uint64[30] memory skillsDiploma = _getDiploma.skills;
		return (levelDiploma, skillsDiploma);
	}
}