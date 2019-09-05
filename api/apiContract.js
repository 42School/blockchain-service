var		express = require('express');
var		bodyParser = require('body-parser');
var		request = require('request');
var		router = express.Router();
var		helpers = require('./helpers');
var		Web3 = require('web3');
var		web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:7545'));
var		models = require('./models')

const	GraduateMarvinCore = require("../build/contracts/GraduateMarvinCore.json");
const	contractInstance = new web3.eth.Contract(GraduateMarvinCore.abi, process.env.CONTRACT_ADD);
const	contractOwner = process.env.CONTRACT_OWNER;
const	gasLimit = process.env.GAS_LIMIT;

const	acc1 = {
	address: process.env.AD1,
	privateKey: process.env.PK1
};
const	acc2 = {
	address: process.env.AD2,
	privateKey: process.env.PK2
};
const	acc3 = {
	address: process.env.AD3,
	privateKey: process.env.PK3
};
const	accs = [acc1, acc2, acc3];

router.use(bodyParser.json());

router.post('/create/:login', (req, res) => {
	const	login = req.params.login;
	request({url: `https://api.intra.42.fr/v2/users/${login}`, auth: {'bearer': process.env.ACCESS_TOKEN}}, async (err, res2) => {
		const	apiData = JSON.parse(res2.body);
		let		data = {
			intraLevel: `${apiData.cursus_users[0].level}`,
			promoYears: apiData.pool_year,
			graduateYears: `${new Date().getFullYear()}`,
			flags: '0'
		};
		data = helpers.str2bytes32(data);
		const signature = helpers.multiSig(login, accs);
		contractInstance.methods.createGraduate(data, signature)
		.send({from: contractOwner, gas: gasLimit})
		.then((tx) => {
			const	gasUsed = tx.gasUsed;
			if (tx) {
				let newIntraLink = new models.intraLink();
				newIntraLink.login = apiData.login;
				newIntraLink.graduateId = tx.events.CreateGraduate.returnValues.graduateId;
				newIntraLink.isForgot = false;
				newIntraLink.save((err) => {
					if (err) {
						return res.status(500).json({message: `Failure: The diploma has been well created, but we have error in links !`});
					}
					return res.status(201).json({message: `The diploma has been well created, it cost ${gasUsed} gas !!`});
				})
			}
		})
		.catch((err) => {
			console.log(err);
			return res.status(500).json({message: `Failure: The diploma was not created !`});
		});
	});
})

router.get('/get/:login', async (req, res) => {
	const	login = req.params.login;
	request({url: `https://api.intra.42.fr/v2/users/${login}`, auth: {'bearer': process.env.ACCESS_TOKEN}}, async (err, res2) => {
		const apiData = JSON.parse(res2.body);
		const findDb = await models.intraLink.find({ login: login });
		const graduateId = findDb[0].graduateId;
		const graduate = await contractInstance.methods.getGraduate(graduateId).call({from: contractOwner, gas: gasLimit});
		if (graduate.signature === '') {
			return res.status(200).json({message: `Failure: This person is not certify by 42 !`});
		} else {
			helpers.createPdf(graduate, apiData.displayname, apiData.login);
			graduate.graduate.splice(0, 4);
			let graduateData = Object.assign({}, graduate.graduate);
			graduateData.intraLevel = web3.utils.toUtf8(graduateData.intraLevel);
			return res.status(200).json({message: 'Success', signature: graduate.signature, data: graduateData})
		}
	});
})

module.exports = router