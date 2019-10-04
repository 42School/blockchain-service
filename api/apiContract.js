var		express = require('express');
var		bodyParser = require('body-parser');
var		request = require('request');
var		router = express.Router();
var		helpers = require('./helpers');
var		Web3 = require('web3');
var		web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:7545'));

web3.eth.defaultAccount = web3.eth.accounts[0];
const	GraduateMarvinCore = require("../build/contracts/GraduateMarvinCore.json");
const	contractInstance = new web3.eth.Contract(GraduateMarvinCore.abi, '0x4727f3efAE5EfF5166F0ECDc383772BdbEd827a4');

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
			birthDate: '21/21/1998',
			birthCity: 'Lyon',
			birthCountry: 'France',
			promoYears: apiData.pool_year,
			graduateYears: `${new Date().getFullYear()}`,
		};
		data = helpers.str2bytes32(data);
		dataBytes = web3.utils.toHex(data);
		signature = web3.eth.accounts.sign(dataBytes, account42.privateKey);
		contractInstance.methods.createGraduate(data, signature.signature)
		.send({from: '0xf8f9CBFd9DCc907A141a51384e92FB499b5D889a', gas:6721975})
		.then((event) => {
			const	graduateId = event.events.CreateGraduate.returnValues.graduateId;
			const	gasUsed = event.gasUsed;
			if (event) {
				console.log(graduateId, gasUsed);
				return res.status(201).json({message: `Graduate creation success, the cost is ${gasUsed}`});
			}
		})
		.catch((err) => {
			return res.status(500).json({message: `Graduate creation failed`});
		});
	});
})

module.exports = router