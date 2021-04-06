/*
**	Autor: Louise Pieri
*/
// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;
pragma abicoder v2;

contract	FtDiploma {

	string public constant name = "42 Alumni";
	string public constant symbol = "42A";
	string public constant linkOfRepo = "github.com/42School/blockchain-service";
	address public constant ftPubAddress = 0x7e12234E994384A757E2689aDdB2A463ccD3B47d;

	event Publish42Diploma(address ftPubAddress, string _link);
	event CreateDiploma(bytes32 student);

	struct	Sign {
		uint8		v;
		bytes32		r;
		bytes32		s;
	}

	struct	Skill {
		string		slug;
		uint64		level;
	}

	struct	Diplomas {
		uint64		level;
		Skill[30]	skills;
		bytes32		hash;
		Sign		signature;
	}

	struct	Diploma {
		uint64		level;
		mapping(uint64 => Skill) skills;
		bytes32		hash;
		Sign		signature;
	}

	mapping (bytes32 => Diploma) hashToDiploma;
	bytes32[] intToHash;

	constructor () public {
		emit Publish42Diploma(ftPubAddress, linkOfRepo);
	}

	function createDiploma(uint64 _level, uint64[30] memory _skillLevel, string[30] memory _skillSlug, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _studentHash) public {
		require(ecrecover(_studentHash, _v, _r, _s) == ftPubAddress, "FtDiploma: Is not 42 sign this diploma.");
		require(hashToDiploma[_studentHash].level == 0, "FtDiploma: The diploma already exists.");
		Diploma storage dp = hashToDiploma[_studentHash];
		dp.level = _level;
		dp.hash = _studentHash;
		dp.signature = Sign(_v, _r, _s);
		for (uint64 i = 0; i < 30; i++) {
			dp.skills[i] = Skill(_skillSlug[i], _skillLevel[i]);
		}
		intToHash.push(_studentHash);
		emit CreateDiploma(_studentHash);
	}

	function getDiploma(bytes32 _studentHash) public view returns (uint64, Skill[] memory) {
		Diploma storage _getDiploma = hashToDiploma[_studentHash];
		uint64 levelDiploma = _getDiploma.level;
		Skill[] memory skills = new Skill[](30);
		for (uint64 i = 0; i < 30; i++) {
			skills[i] = _getDiploma.skills[i];
		}
		return (levelDiploma, skills);
	}

	function getAllDiploma() public view returns (Diplomas[] memory) {
		require(msg.sender == ftPubAddress, "FtDiploma: Is not 42.");
		Diplomas[] memory diplomas = new Diplomas[](intToHash.length);
		for (uint256 i = 0; i < intToHash.length; i++) {
			Diploma storage dp = hashToDiploma[intToHash[i]];
			diplomas[i].level = dp.level;
			for (uint64 y = 0; y < 30; y++) {
				diplomas[i].skills[y] = dp.skills[y];
			}
			diplomas[i].hash = dp.hash;
			diplomas[i].signature = dp.signature;
		}
		return (diplomas);
	}
}
