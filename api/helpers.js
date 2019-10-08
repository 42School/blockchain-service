const	web3 = require('web3');

function str2bytes32(data) {
	const	new_data = {
		intraLevel: web3.utils.padRight(web3.utils.utf8ToHex(data.intraLevel), 64),
		promoYears: data.promoYears,
		graduateYears: data.graduateYears,
		flags: data.flags
	}
	return (new_data);
}

function multiSig(data, accounts) {
	let		signatures = [];
	for (let i = 0; i < accounts.length; i++) {
		const sign = accounts[i].sign(data);
		signatures[i] = sign.signature;
	}
	const hash = web3.utils.sha3(signatures);
	return (hash);
}

module.exports = {
	str2bytes32: str2bytes32,
	multiSig: multiSig
}