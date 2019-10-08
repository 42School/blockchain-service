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

const	acc1 = web3.eth.accounts.create();
const	acc2 = web3.eth.accounts.create();
const	acc3 = web3.eth.accounts.create();
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
		const signature = helpers.multiSig(apiData.login, accs);
		console.log(signature);
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
						return res.status(500).json({message: `The diploma has been well created, but we have error in links !`});
					}
					return res.status(201).json({message: `The diploma has been well created, it cost ${gasUsed} gas !!`});
				})
			}
		})
		.catch((err) => {
			console.log(err);
			return res.status(500).json({message: `The diploma was not created !`});
		});
	});
})

router.get('/get/:login', async (req, res) => {
	const	login = req.params.login;
	const findDb = await models.intraLink.find({ login: login });
	const graduateId = findDb[0].graduateId;
	const graduate = await contractInstance.methods.getGraduate(graduateId).call({from: contractOwner, gas: gasLimit});
	if (graduate.signature === '') {
		return res.status(200).json({message: `This person is not certify by 42 !`});
	} else {
		helpers.createPdf(graduate);
		return res.status(200).json({message: 'Succes'})
	}
})

module.exports = router