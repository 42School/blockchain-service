const	GraduateMarvin = artifacts.require("GraduateMarvinCore");
const	helpers = require('./helpers/helpers');
const	assert = require('assert');
const	fs = require("fs");
const	crypto = require('crypto');

const	acc1 = web3.eth.accounts.create();
const	acc2 = web3.eth.accounts.create();
const	acc3 = web3.eth.accounts.create();
const	accs = [acc1, acc2, acc3];

contract("GraduateMarvinCore", async (accounts) => {

	it("Testing function createGraduate & getGraduate", async () => {
		let data_new_graduate = {
			intraLevel: web3.utils.padRight(web3.utils.utf8ToHex('IntraLevel'), 64),
			promoYears: '2017',
			graduateYears: '2020',
			flags: '0'
		}
		const signature = helpers.multiSig('Login', accs);
		let instance = await GraduateMarvin.deployed();
		const tx = await instance.createGraduate(data_new_graduate, signature);
		const id = web3.utils.BN(tx.logs[0].args.graduateId).toString();
		const get = await instance.getGraduate(id);
		const getSignature = get.signature;
		assert.equal(getSignature, helpers.multiSig('Login', accs), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random Russe graduate", async () => {
		const data = await helpers.getRandomUser('Russia');
		const signature = helpers.multiSig(data.name, accs);
		let instance = await GraduateMarvin.deployed();
		const tx = await instance.createGraduate(data.data, signature);
		const id = web3.utils.BN(tx.logs[0].args.graduateId).toString();
		const get = await instance.getGraduate(id);
		const getSignature = get.signature;
		assert.equal(getSignature, helpers.multiSig(data.name, accs), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random Japan graduate", async () => {
		const data = await helpers.getRandomUser('Japan');
		const signature = helpers.multiSig(data.name, accs);
		let instance = await GraduateMarvin.deployed();
		const tx = await instance.createGraduate(data.data, signature);
		const id = web3.utils.BN(tx.logs[0].args.graduateId).toString();
		const get = await instance.getGraduate(id);
		const getSignature = get.signature;
		assert.equal(getSignature, helpers.multiSig(data.name, accs), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random Korea graduate", async () => {
		const data = await helpers.getRandomUser('Korea');
		const signature = helpers.multiSig(data.name, accs);
		let instance = await GraduateMarvin.deployed();
		const tx = await instance.createGraduate(data.data, signature);
		const id = web3.utils.BN(tx.logs[0].args.graduateId).toString();
		const get = await instance.getGraduate(id);
		const getSignature = get.signature;
		assert.equal(getSignature, helpers.multiSig(data.name, accs), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random Greece graduate", async () => {
		const data = await helpers.getRandomUser('Greece');
		const signature = helpers.multiSig(data.name, accs);
		let instance = await GraduateMarvin.deployed();
		const tx = await instance.createGraduate(data.data, signature);
		const id = web3.utils.BN(tx.logs[0].args.graduateId).toString();
		const get = await instance.getGraduate(id);
		const getSignature = get.signature;
		assert.equal(getSignature, helpers.multiSig(data.name, accs), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random Egypt graduate", async () => {
		const data = await helpers.getRandomUser('Egypt');
		const signature = helpers.multiSig(data.name, accs);
		let instance = await GraduateMarvin.deployed();
		const tx = await instance.createGraduate(data.data, signature);
		const id = web3.utils.BN(tx.logs[0].args.graduateId).toString();
		const get = await instance.getGraduate(id);
		const getSignature = get.signature;
		assert.equal(getSignature, helpers.multiSig(data.name, accs), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 1 random China graduate", async () => {
		const data = await helpers.getRandomUser('China');
		const signature = helpers.multiSig(data.name, accs);
		let instance = await GraduateMarvin.deployed();
		const tx = await instance.createGraduate(data.data, signature);
		const id = web3.utils.BN(tx.logs[0].args.graduateId).toString();
		const get = await instance.getGraduate(id);
		const getSignature = get.signature;
		assert.equal(getSignature, helpers.multiSig(data.name, accs), 'The data insert into the blockchain it\'s not the same than input');
	})

	it("Create 500 random graduates", async () => {
		let gasUsed = {};
		users = helpers.getRandomUsers();
		for (let i = 0; i < users.length; i++) {
			const data = users[i];
			const signature = helpers.multiSig(data.name, accs);
			let instance = await GraduateMarvin.deployed();
			const tx = await instance.createGraduate(data.data, signature);
			gasUsed[i] = tx.receipt.gasUsed;
			const id = web3.utils.BN(tx.logs[0].args.graduateId).toString();
			const get = await instance.getGraduate(id);
			const getSignature = get.signature;
			assert.equal(getSignature, helpers.multiSig(data.name, accs), 'The data insert into the blockchain it\'s not the same than input');
		};
		fs.writeFileSync(`./gasCalculator/gasUsed_${helpers.randomInt(0, 42)}.json`, JSON.stringify(gasUsed));
	})

})