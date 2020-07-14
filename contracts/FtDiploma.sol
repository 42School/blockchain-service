/*
**	Autor: Louise Pieri
*/
// SPDX-License-Identifier: MIT
pragma solidity >=0.5.8 <0.7.0;

contract	FtDiploma {

	string public constant name = "42 Alumni";
	string public constant symbol = "42A";
	string public constant linkOfRepo = "github.com/lpieri/42-Alumni";
	address public constant ftPubAddress = 0x8A21Dc0aeC762cD85de81B2bcd396a9d5676cFD7;

	event Publish42Diploma(address ftPubAddress, string _link);
	event CreateDiploma(bytes32 student);

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

	mapping (bytes32 => Diploma) hashToDiploma;

	constructor () public {
		emit Publish42Diploma(ftPubAddress, linkOfRepo);
	}

	function createDiploma(uint64 _level, uint64[30] memory _skills, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _studentHash) public {
		require(ecrecover(_studentHash, _v, _r, _s) == ftPubAddress, "FtDiploma: Is not 42 sign this diploma");
		require(hashToDiploma[_studentHash].level == 0, "FtDiploma: The diploma already exists.");
		hashToDiploma[_studentHash] = Diploma(_level, _skills, Sign(_v, _r, _s));
		emit CreateDiploma(_studentHash);
	}

	function getDiploma(bytes32 _studentHash) public view returns (uint64 level, uint64[30] memory skills) {
		Diploma memory _getDiploma = hashToDiploma[_studentHash];
		uint64 levelDiploma = _getDiploma.level;
		uint64[30] memory skillsDiploma = _getDiploma.skills;
		return (levelDiploma, skillsDiploma);
	}
}