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
		login: web3.utils.padRight(web3.utils.utf8ToHex(`${data.name.charAt(0)}${data.surname}`.toLowerCase()), 64),
		firstName: web3.utils.padRight(web3.utils.utf8ToHex(data.name), 64),
		lastName: web3.utils.padRight(web3.utils.utf8ToHex(data.surname), 64),
		intraLevel: web3.utils.padRight(web3.utils.utf8ToHex(randomInt(7, 21)), 64),
		promoYears: '2013',
		graduateYears: '2020',
	}
	return (ret_data);
};

function getRandomRegion() {
	const regions = ['Russia', 'China', 'Japan', 'France', 'Egypt', 'Greece', 'Korea', 'United States'];
	return (regions[randomInt(0, 8)]);
}

function getRandomUsers() {
	const url = `https://uinames.com/api/?amount=500&ext`
	const res = request('GET', url).getBody();
	// const res2 = request('GET', url).getBody();
	const data = JSON.parse(res);// + JSON.parse(res2);
	let ret_data = new Array();
	for (let i = 0; i < data.length; i++) {
		ret_data.push({
			login: web3.utils.padRight(web3.utils.utf8ToHex(`${data[i].name.charAt(0)}${data[i].surname.substring(0,8)}`.toLowerCase()), 64),
			firstName: web3.utils.padRight(web3.utils.utf8ToHex(data[i].name), 64),
			lastName: web3.utils.padRight(web3.utils.utf8ToHex(data[i].surname), 64),
			intraLevel: web3.utils.padRight(web3.utils.utf8ToHex(randomInt(7, 21)), 64),
			promoYears: '2013',
			graduateYears: '2020',
		});
	}
	return (ret_data);
}

module.exports = {
	randomInt: randomInt,
	getRandomUser: getRandomUser,
	getRandomRegion: getRandomRegion,
	getRandomUsers: getRandomUsers
}