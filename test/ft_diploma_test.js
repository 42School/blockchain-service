const truffleConfig = require("../truffle-config");

const FtDiploma = artifacts.require("FtDiploma");

let skillsLevel = [
	857, 542, 62, 942, 661, 416, 902, 902, 902, 36, 222, 55, 145,
	435, 267, 1122, 73, 206, 103, 817, 116, 2828, 920, 157, 1112,
	2126, 328, 423, 203, 416
]

let skillsSlug = [
	"Written Communication", "Web", "Unix", "Technology integration", "Shell",
	"Security", "Ruby", "Rigor", "Python", "Parallel computing", "Organization",
	"Object-oriented programming", "Network & system administration", "Motion Design",
	"Imperative programming", "Group & interpersonal", "Graphics", "Functional programming",
	"DB & Data", "Company experience", "Basics", "Algorithms & AI", "Adaptation & creativity",
	"157", "1112", "2126", "328", "423", "203", "416"
]

const keystore = `{"address":"7e12234e994384a757e2689addb2a463ccd3b47d","crypto":{"cipher":"aes-128-ctr","ciphertext":"6d9f85fc277cbbd66f01f15358a62fff3a1cef9c7a96a8024043d8b3c9adeaf6","cipherparams":{"iv":"39516f9f12fe7cccfbdaa11dab569cb9"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"56bb6dac014acbb558e2cb6c533d73d25dc3fe7fa2b1653e664ce32fa72347d5"},"mac":"7d69e74114b8529df420adb173d62a97ef12f631272253c8ce69416c2234245d"},"id":"3ae16ea0-ef89-4b2c-b527-ad1259a53b4d","version":3}`;
const account = web3.eth.accounts.decrypt(keystore, 'password');

contract("FtDiploma", async (accounts) => {

	it("Testing the creation of new diploma", async () => {
		let instance = await FtDiploma.deployed();
		let dataToHash = "Louise, Pieri, 1998-12-27, 2020-06-25";
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash);
		const tx = await instance.createDiploma(1517, skillsLevel, skillsSlug, sign.v, sign.r, sign.s, sign.messageHash);
		txStudent = tx.logs[0].args.student;
		assert.equal(txStudent, sign.messageHash, 'Error: The writing is not valid!');
	})

	it("Testing the get of diploma", async () => {
		let instance = await FtDiploma.deployed();
		let dataToHash = "Louise, Pieri, 1998-12-27, 2020-06-25";
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash);
		const tx = await instance.getDiploma(sign.messageHash);
		level = web3.utils.BN(tx[0]).toNumber();
		assert.equal(level, 1517, `Error: The diploma couldn't be gotten, or doesn't exist.`);
	})

	it("Testing writing not valid diploma", async () => {
		let instance = await FtDiploma.deployed();
		let dataToHash = "FirstName, LastName, AAAA-MM-JJ, AAAA-MM-JJ";
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash);
		try {
			await instance.createDiploma(1517, skillsLevel, skillsSlug, sign.v, sign.r, sign.s, hash);
		} catch (error) {
			assert.equal(error.reason, 'FtDiploma: Is not 42 sign this diploma.', `Error: It is possible to write a diploma not signed by 42.`);
		}
	})

	it("Testing double same writing of new diploma", async () => {
		let instance = await FtDiploma.deployed();
		let dataToHash = "Prenom, Nom, AAAA-MM-JJ, AAAA-MM-JJ";
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash);
		const firstTx = await instance.createDiploma(1517, skillsLevel, skillsSlug, sign.v, sign.r, sign.s, sign.messageHash);
		try {
			await instance.createDiploma(1517, skillsLevel, skillsSlug, sign.v, sign.r, sign.s, sign.messageHash);
		} catch (error) {
			assert.equal(error.reason, "FtDiploma: The diploma already exists.", `Error: The double insertion of a diploma is possible.`);
		}
	})

	it("Testing the gets all data if not 42", async () => {
		let instance = await FtDiploma.deployed();
		try {
			await instance.getAllDiploma();
		} catch (error) {
			assert.ok(error, "FtDiploma: Is not 42.", `Error: The gets of all diplomas is not authorize if is not 42.`);
		}
	})

	it("Testing the gets all data", async () => {
		let instance = await FtDiploma.deployed();
		let datas = await instance.getAllDiploma({from: "0x7e12234e994384a757e2689addb2a463ccd3b47d"})
		console.log(datas);
		assert.isNotEmpty(datas);
	})
})