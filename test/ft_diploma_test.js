const FtDiploma = artifacts.require("FtDiploma");

let skills = [
	857, 542, 62, 942, 661, 416, 902, 902, 902, 36, 222, 55, 145,
	435, 267, 1122, 73, 206, 103, 817, 116, 2828, 920, 157, 1112,
	2126, 328, 423, 203, 416
]

const keystore = "Contents of keystore file";
const account = web3.eth.accounts.decrypt(keystore, 'PASSWORD');

web3.eth.handleRevert = true;

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

	// it("Testing writing not valid diploma", async () => {
	// 	let instance = await FtDiploma.deployed();
	// 	instance.handleRevert = true;
	// 	let dataToHash = "FirstName, LastName, AAAA-MM-JJ, AAAA-MM-JJ";
	// 	let hash = web3.utils.sha3(dataToHash);
	// 	let sign = account.sign(hash);
	// 	const tx = await instance.createDiploma(1517, skills, sign.v, sign.r, sign.s, hash).handleRevert;
	// 	await expectThrow(tx);
	// 	assert.equal(tx, undefined, `Error: It is possible to write a diploma not signed by 42.`);
	// })

	it("Testing double same writing of new diploma", async () => {
		let instance = await FtDiploma.deployed();
		instance.handleRevert = true;
		let dataToHash = "Prenom, Nom, AAAA-MM-JJ, AAAA-MM-JJ";
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash);
		const firstTx = await instance.createDiploma(1517, skills, sign.v, sign.r, sign.s, sign.messageHash);
		const secondeTx = instance.createDiploma(1517, skills, sign.v, sign.r, sign.s, sign.messageHash).handleRevert;
		assert.equal(secondeTx, undefined, `Error: The double insertion of a diploma is possible.`);
	})
})