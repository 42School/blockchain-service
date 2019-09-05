var		express = require('express');
var		bodyParser = require('body-parser');
var		router = express.Router();
var		Web3 = require('web3');
var		web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:7545'));

router.use(bodyParser.json());

router.get('/sophie/:login', (req, res) => {
	const	login = req.params.login;
	const	account = {
		address: process.env.AD1,
		privateKey: process.env.PK1
	};
	const	signature = web3.eth.accounts.sign(login, account.privateKey);
	const	signData = {
		signature: signature.signature,
		address: account.address
	}
	return res.status(200).json({message: 'Success', signData: signData})
})

router.get('/benny/:login', (req, res) => {
	const	login = req.params.login;
	const	account = {
		address: process.env.AD2,
		privateKey: process.env.PK2
	};
	const	signature = web3.eth.accounts.sign(login, account.privateKey);
	const	signData = {
		signature: signature.signature,
		address: account.address
	}
	return res.status(200).json({message: 'Success', signData: signData})
})

router.get('/niel/:login', (req, res) => {
	const	login = req.params.login;
	const	account = {
		address: process.env.AD3,
		privateKey: process.env.PK3
	};
	const	signature = web3.eth.accounts.sign(login, account.privateKey);
	const	signData = {
		signature: signature.signature,
		address: account.address
	}
	return res.status(200).json({message: 'Success', signData: signData})
})

router.post('/get-hash/:login', (req, res) => {
	const	data = req.body.data;
	const	login = req.params.login;
	let		sign = [];
	let		y = 0;
	for (i in data) {
		const recoverAdd = web3.eth.accounts.recover(login, data[i].signature);
		sign[y] = data[i].signature;
		y++;
	}
	if (sign.length === 3) {
		const hash = web3.utils.sha3(`${sign[0]},${sign[1]},${sign[2]}`);
		return res.status(200).json({message: 'Success', hash: hash});
	} else {
		return res.status(500).json({message: 'Failure: Data not valide'});
	}
})

module.exports = router