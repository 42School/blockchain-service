const	GraduateMarvin = artifacts.require("GraduateMarvinCore");
const	helpers = require('./helpers/helpers');
const	assert = require('assert');
const	fs = require("fs");
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

	it("Create 1 random Russe graduate", async () => {
		const data = await helpers.getRandomUser('Russia');
		const dataHex = web3.utils.toHex(data);
		const signature = web3.eth.accounts.sign(dataHex, acc42.privateKey);
		let instance = await GraduateMarvin.deployed();
		let new_graduate = await instance.createGraduate(data, signature.signature);
		let ret = await instance.getGraduate(data.login);
		let graduate = ret.graduate;
		graduate.splice(0, 9);
		graduate = Object.assign({}, graduate);
		const message = web3.utils.toHex(graduate);
		assert.equal(web3.utils.sha3(message), web3.utils.sha3(dataHex), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random Japan graduate", async () => {
		const data = await helpers.getRandomUser('Japan');
		const dataHex = web3.utils.toHex(data);
		const signature = web3.eth.accounts.sign(dataHex, acc42.privateKey);
		let instance = await GraduateMarvin.deployed();
		let new_graduate = await instance.createGraduate(data, signature.signature);
		let ret = await instance.getGraduate(data.login);
		let graduate = ret.graduate;
		graduate.splice(0, 9);
		graduate = Object.assign({}, graduate);
		const message = web3.utils.toHex(graduate);
		assert.equal(web3.utils.sha3(message), web3.utils.sha3(dataHex), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random Korea graduate", async () => {
		const data = await helpers.getRandomUser('Korea');
		const dataHex = web3.utils.toHex(data);
		const signature = web3.eth.accounts.sign(dataHex, acc42.privateKey);
		let instance = await GraduateMarvin.deployed();
		let new_graduate = await instance.createGraduate(data, signature.signature);
		let ret = await instance.getGraduate(data.login);
		let graduate = ret.graduate;
		graduate.splice(0, 9);
		graduate = Object.assign({}, graduate);
		const message = web3.utils.toHex(graduate);
		assert.equal(web3.utils.sha3(message), web3.utils.sha3(dataHex), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random Greece graduate", async () => {
		const data = await helpers.getRandomUser('Greece');
		const dataHex = web3.utils.toHex(data);
		const signature = web3.eth.accounts.sign(dataHex, acc42.privateKey);
		let instance = await GraduateMarvin.deployed();
		let new_graduate = await instance.createGraduate(data, signature.signature);
		let ret = await instance.getGraduate(data.login);
		let graduate = ret.graduate;
		graduate.splice(0, 9);
		graduate = Object.assign({}, graduate);
		const message = web3.utils.toHex(graduate);
		assert.equal(web3.utils.sha3(message), web3.utils.sha3(dataHex), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random Egypt graduate", async () => {
		const data = await helpers.getRandomUser('Egypt');
		const dataHex = web3.utils.toHex(data);
		const signature = web3.eth.accounts.sign(dataHex, acc42.privateKey);
		let instance = await GraduateMarvin.deployed();
		let new_graduate = await instance.createGraduate(data, signature.signature);
		let ret = await instance.getGraduate(data.login);
		let graduate = ret.graduate;
		graduate.splice(0, 9);
		graduate = Object.assign({}, graduate);
		const message = web3.utils.toHex(graduate);
		assert.equal(web3.utils.sha3(message), web3.utils.sha3(dataHex), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random China graduate", async () => {
		const data = await helpers.getRandomUser('China');
		const dataHex = web3.utils.toHex(data);
		const signature = web3.eth.accounts.sign(dataHex, acc42.privateKey);
		let instance = await GraduateMarvin.deployed();
		let new_graduate = await instance.createGraduate(data, signature.signature);
		let ret = await instance.getGraduate(data.login);
		let graduate = ret.graduate;
		graduate.splice(0, 9);
		graduate = Object.assign({}, graduate);
		const message = web3.utils.toHex(graduate);
		assert.equal(web3.utils.sha3(message), web3.utils.sha3(dataHex), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 500 random graduates", async () => {
		let gasUsed = {};
		users = helpers.getRandomUsers();
		for (let i = 0; i < users.length; i++) {
			const data = users[i];
			const dataHex = web3.utils.toHex(data);
			const signature = web3.eth.accounts.sign(dataHex, acc42.privateKey);
			let instance = await GraduateMarvin.deployed();
			let new_graduate = await instance.createGraduate(data, signature.signature);
			gasUsed[i] = new_graduate.receipt.gasUsed;
			let ret = await instance.getGraduate(data.login);
			let graduate = ret.graduate;
			graduate.splice(0, 9);
			graduate = Object.assign({}, graduate);
			const message = web3.utils.toHex(graduate);
			assert.equal(web3.utils.sha3(message), web3.utils.sha3(dataHex), 'The data insert into the blockchain it\'s not the same than input');
		};
		fs.writeFileSync('UsersCreationGasUsed.json', JSON.stringify(gasUsed));
	})

})