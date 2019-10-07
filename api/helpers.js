const	web3 = require('web3');

function str2bytes32(data) {
	const	new_data = {
		login: web3.utils.sha3(data.login),
		firstName: web3.utils.sha3(data.firstName),
		lastName: web3.utils.sha3(data.lastName),
		intraLevel: web3.utils.padRight(web3.utils.utf8ToHex(data.intraLevel), 64),
		promoYears: data.promoYears,
		graduateYears: data.graduateYears,
	}
	return (new_data);
}

module.exports = {
	str2bytes32: str2bytes32
}