var		express = require('express');
var		bodyParser = require('body-parser');
var		request = require('request');
var		router = express.Router();
var		helpers = require('./helpers');
var		Web3 = require('web3');
var		web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:7545'));

const	GraduateMarvinCore = require("../build/contracts/GraduateMarvinCore.json");
const	contractInstance = new web3.eth.Contract(GraduateMarvinCore.abi, process.env.CONTRACT_ADD);
const	contractOwner = process.env.CONTRACT_OWNER;
const	gasLimit = process.env.GAS_LIMIT;
const	account42 = {
	address: '0x2CC70d29a7F00C0e04A3B0E78074AB523b7056af',
	privateKey: '0x808874e9b286e5020632a7bcd583ffc9f3e9dd1ae86198a6d2a215ee5e182add'
};

router.use(bodyParser.json());

router.post('/create/:login', (req, res) => {
	const	login = req.params.login;
	request({url: `https://api.intra.42.fr/v2/users/${login}`, auth: {'bearer': process.env.ACCESS_TOKEN}}, async (err, res2) => {
		const	apiData = JSON.parse(res2.body);
		let		data = {
			login: login,
			firstName: apiData.first_name,
			lastName: apiData.last_name,
			intraLevel: `${apiData.cursus_users[0].level}`,
			promoYears: apiData.pool_year,
			graduateYears: `${new Date().getFullYear()}`,
		};
		data = helpers.str2bytes32(data);
		dataBytes = web3.utils.toHex(data);
		signature = web3.eth.accounts.sign(dataBytes, account42.privateKey);
		contractInstance.methods.createGraduate(data, signature.signature)
		.send({from: contractOwner, gas: gasLimit})
		.then((event) => {
			const	gasUsed = event.gasUsed;
			if (event) {
				return res.status(201).json({message: `The diploma has been well created, it cost ${gasUsed} gas !!`});
			}
		})
		.catch(() => {
			return res.status(500).json({message: `The diploma was not created !`});
		});
	});
})

router.get('/get/:login', async (req, res) => {
	const	login = req.params.login;
	const	loginBytes32 = web3.utils.padRight(web3.utils.utf8ToHex(login), 64);
	const	data = await contractInstance.methods.getGraduate(loginBytes32).call({from: contractOwner, gas: gasLimit});
	if (data.signature === '') {
		return res.status(200).json({message: `This person did not certify by 42 !`});
	} else {
		return res.status(200).json({message: 'Succes'})
	}
})

module.exports = router