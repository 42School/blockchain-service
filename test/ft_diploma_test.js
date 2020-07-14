const FtDiploma = artifacts.require("FtDiploma")

let skills = [
	857, 542, 62, 942, 661, 416, 902, 902, 902, 36, 222, 55, 145,
	435, 267, 1122, 73, 206, 103, 817, 116, 2828, 920, 157, 1112,
	2126, 328, 423, 203, 416
]

const keystore = "Contents of keystore file";
const account = web3.eth.accounts.decrypt(keystore, 'PASSWORD');

contract("FtDiploma", async (accounts) => {

	it("Testing the creation of new diploma", async () => {
		let instance = await FtDiploma.deployed();
		let dataToHash = "Louise, Pieri, 1998-12-27, 2020-06-25"
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash)
		const tx = await instance.createDiploma(1517, skills, sign.v, sign.r, sign.s, sign.messageHash);
		console.log(tx)
	})

	it("Testing the get of diploma", async () => {
		let instance = await FtDiploma.deployed();
		let dataToHash = "Louise, Pieri, 1998-12-27, 2020-06-25"
		let hash = web3.utils.sha3(dataToHash);
		let sign = account.sign(hash)
		const tx = await instance.getDiploma(sign.messageHash);
		console.log(tx)
	})
})