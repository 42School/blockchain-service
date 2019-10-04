const	web3 = require('web3');

function str2bytes32(data) {
	const	new_data = {
		login: web3.utils.padRight(web3.utils.utf8ToHex(data.login), 64),
		firstName: web3.utils.padRight(web3.utils.utf8ToHex(data.firstName), 64),
		lastName: web3.utils.padRight(web3.utils.utf8ToHex(data.lastName), 64),
		intraLevel: web3.utils.padRight(web3.utils.utf8ToHex(data.intraLevel), 64),
		birthDate: web3.utils.padRight(web3.utils.utf8ToHex(data.birthDate), 64),
		birthCity: web3.utils.padRight(web3.utils.utf8ToHex(data.birthCity), 64),
		birthCountry: web3.utils.padRight(web3.utils.utf8ToHex(data.birthCountry), 64),
		promoYears: data.promoYears,
		graduateYears: data.graduateYears,
	}
	return (new_data);
}

module.exports = {
	str2bytes32: str2bytes32
}