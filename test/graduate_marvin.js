const	GraduateMarvin = artifacts.require("GraduateMarvinCore");
const	assert = require('assert');
const	crypto = require('crypto');

const	acc42 = web3.eth.accounts.create();

contract("GraduateMarvinCore", async (accounts) => {

	it("Testing function createGraduate", async () => {
		let data_new_graduate = {
			login: web3.utils.padRight(web3.utils.utf8ToHex('Login'), 64),
			firstName: web3.utils.padRight(web3.utils.utf8ToHex('FirstName'), 64),
			lastName: web3.utils.padRight(web3.utils.utf8ToHex('LastName'), 64),
			intraLevel: web3.utils.padRight(web3.utils.utf8ToHex('IntraLevel'), 64),
			birthDate: web3.utils.padRight(web3.utils.utf8ToHex('BirthDate'), 64),
			birthCity: web3.utils.padRight(web3.utils.utf8ToHex('BirthCity'), 64),
			birthCountry: web3.utils.padRight(web3.utils.utf8ToHex('BirthCountry'), 64),
			promoYears: '2017',
			graduateYears: '2020',
		}
		data = web3.utils.toHex(data_new_graduate);
		signature = web3.eth.accounts.sign(data, acc42.privateKey);
		let instance = await GraduateMarvin.deployed();
		let new_graduate = await instance.createGraduate(data_new_graduate, signature.signature);
		let login = new_graduate.logs[0].args.login;
		assert.equal(web3.utils.hexToUtf8(login), 'Login', 'The data is not valid !');
	})

	it("Testing function getGraduate", async () => {
		let loginToGet = web3.utils.utf8ToHex('Login');
		let instance = await GraduateMarvin.deployed();
		let ret = await instance.getGraduate(loginToGet);
		let graduate = ret.graduate;
		graduate.splice(0, 9);
		graduate = Object.assign({}, graduate);
		message = web3.utils.toHex(graduate);
		add42 = web3.eth.accounts.recover(message, ret.signature);
		assert.equal(add42, acc42.address, 'The data return by the smart contract is altered');
	})

	it("Testing function deleteGraduate", async () => {
		let loginToDelete = web3.utils.utf8ToHex('Login');
		let instance = await GraduateMarvin.deployed();
		let deleteGraduate = await instance.deleteGraduate(loginToDelete);
		let loginDeleted = login = deleteGraduate.logs[0].args.login;
		assert.equal(web3.utils.hexToUtf8(loginDeleted), 'Login', 'The data is not valid !');
	})

})