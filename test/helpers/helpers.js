const	request = require('sync-request');

function randomInt(min, max) {
	min = Math.ceil(min);
	max = Math.floor(max);
	return Math.floor(Math.random() * (max - min)) + min;
}

async function getRandomUser(country) {
	const url = `https://uinames.com/api/?amount=1&region=${country}&gender=female&ext`
	const res = request('GET', url).getBody();
	const data = JSON.parse(res)
	const ret_data = {
		intraLevel: web3.utils.padRight(web3.utils.utf8ToHex(randomInt(7, 21)), 64),
		promoYears: '2013',
		graduateYears: '2020',
		flags: '0'
	}
	return ({data: ret_data, name: `${data.name} ${data.surname}`});
};

function getRandomRegion() {
	const regions = ['Russia', 'China', 'Japan', 'France', 'Egypt', 'Greece', 'Korea', 'United States'];
	return (regions[randomInt(0, 8)]);
}

function getRandomUsers() {
	const url = `https://uinames.com/api/?amount=500&ext`
	const res = request('GET', url).getBody();
	const data = JSON.parse(res);
	let ret_data = new Array();
	for (let i = 0; i < data.length; i++) {
		ret_data.push({data: {
			intraLevel: web3.utils.padRight(web3.utils.utf8ToHex(randomInt(7, 21)), 64),
			promoYears: '2013',
			graduateYears: '2020',
			flags: '0'
		}, name: `${data[i].name} ${data[i].surname}`});
	}
	return (ret_data);
}

function multiSig(data, accounts) {
	let		signatures = [];
	for (let i = 0; i < accounts.length; i++) {
		const sign = web3.eth.accounts.sign(data, accounts[i].privateKey);
		signatures[i] = sign.signature;
	}
	const hash = web3.utils.sha3(signatures);
	return (hash);
}

module.exports = {
	randomInt: randomInt,
	getRandomUser: getRandomUser,
	getRandomRegion: getRandomRegion,
	getRandomUsers: getRandomUsers,
	multiSig: multiSig
}