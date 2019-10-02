const	request = require('sync-request');

async function getRandomUser(country, counter) {
	const url = `https://uinames.com/api/?amount=${counter}&region=${country}&gender=female&ext`
	req = request('GET', url).getBody();
	console.log(req);
};

module.exports = {
	getRandomUser: getRandomUser
}