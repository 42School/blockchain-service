const	GraduateMarvin = artifacts.require("GraduateMarvinCore");
const	crypto = require('crypto');

const	acc42 = web3.eth.accounts.create();

contract("GraduateMarvinCore", async (accounts) => {

	it("Testing function createGraduate", async () => {
		let data_new_graduate = {
			'login': web3.utils.utf8ToHex('Login'),
			'firstName': web3.utils.utf8ToHex('FirstName'),
			'lastName': web3.utils.utf8ToHex('LastName'),
			'intraLevel': web3.utils.utf8ToHex('IntraLevel'),
			'birthDate': web3.utils.utf8ToHex('BirthDate'),
			'birthCity': web3.utils.utf8ToHex('BirthCity'),
			'birthCountry': web3.utils.utf8ToHex('BirthCountry'),
			'promoYears': 2017,
			'graduateYears': 2020,
		}
		data = web3.utils.toHex(data_new_graduate);
		signature = web3.eth.accounts.sign(data, acc42.privateKey);
		let instance = await GraduateMarvin.deployed();
		let new_graduate = await instance.createGraduate(data_new_graduate, signature.signature);
		// console.log(new_graduate);
		let graduate = await instance.getGraduate(web3.utils.utf8ToHex('Login'));
		// console.log(graduate);
	})

})