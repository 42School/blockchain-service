const truffleConfig = require("../truffle-config");

const FtDiploma = artifacts.require("FtDiploma");

let skills = [
	857, 542, 62, 942, 661, 416, 902, 902, 902, 36, 222, 55, 145,
	435, 267, 1122, 73, 206, 103, 817, 116, 2828, 920, 157, 1112,
	2126, 328, 423, 203, 416
]

const keystore = `{"address":"8a21dc0aec762cd85de81b2bcd396a9d5676cfd7","crypto":{"cipher":"aes-128-ctr","ciphertext":"ce9faf40419c176d70a2da4d3ce9aaeff5d07a0aaf4b22d01ec901fd534cc835","cipherparams":{"iv":"48627618eae022f6015c8526655dfed6"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ebc90ec75f99258f745d56c4e76ac57b87e04e8a98ff61e3940b0cd190da7362"},"mac":"9f5b8782b97557ae98713f94a3d940209d291f0f6cc34cdafdc3a9ca4e0c78a9"},"id":"897b76fc-835e-4f25-b4e8-4eb58e95b510","version":3}`;
const account = web3.eth.accounts.decrypt(keystore, 'password');

contract("FtDiploma", async (accounts) => {

	it("Testing the creation of new diploma", async () => {
		let instance = await FtDiploma.deployed();
		let dataToHash = "Louise, Pieri, 1998-12-27, 2020-06-25";
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash);
		const tx = await instance.createDiploma(1517, skills, sign.v, sign.r, sign.s, sign.messageHash);
		txStudent = tx.logs[0].args.student;
		assert.equal(txStudent, sign.messageHash, 'Error: The writing is not valid!');
	})

	it("Testing the get of diploma", async () => {
		let instance = await FtDiploma.deployed();
		let dataToHash = "Louise, Pieri, 1998-12-27, 2020-06-25";
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash);
		const tx = await instance.getDiploma(sign.messageHash);
		level = web3.utils.BN(tx.level).toNumber();
		assert.equal(level, 1517, `Error: The diploma couldn't be gotten, or doesn't exist.`);
	})

	it("Testing writing not valid diploma", async () => {
		let instance = await FtDiploma.deployed();
		// instance.handleRevert = true;
		let dataToHash = "FirstName, LastName, AAAA-MM-JJ, AAAA-MM-JJ";
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash);
		try {
			await instance.createDiploma(1517, skills, sign.v, sign.r, sign.s, hash);
		} catch (error) {
			assert.equal(error.reason, 'FtDiploma: Is not 42 sign this diploma', `Error: It is possible to write a diploma not signed by 42.`);
		}
	})

	it("Testing double same writing of new diploma", async () => {
		let instance = await FtDiploma.deployed();
		instance.handleRevert = true;
		let dataToHash = "Prenom, Nom, AAAA-MM-JJ, AAAA-MM-JJ";
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash);
		const firstTx = await instance.createDiploma(1517, skills, sign.v, sign.r, sign.s, sign.messageHash);
		try {
			await instance.createDiploma(1517, skills, sign.v, sign.r, sign.s, sign.messageHash);
		} catch (error) {
			assert.equal(error.reason, "FtDiploma: The diploma already exists.", `Error: The double insertion of a diploma is possible.`);
		}
	})
})