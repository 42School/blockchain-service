const	fs = require('fs');
const	pdfkit = require('pdfkit');
var		Web3 = require('web3');
var		web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:7545'));

function str2bytes32(data) {
	const	new_data = {
		intraLevel: web3.utils.padRight(web3.utils.utf8ToHex(data.intraLevel), 64),
		promoYears: data.promoYears,
		graduateYears: data.graduateYears,
		flags: data.flags
	}
	return (new_data);
}

function multiSig(data, accs) {
	let		signatures = [];
	for (let i = 0; i < accs.length; i++) {
		const sign = web3.eth.accounts.sign(data, accs[i].privateKey);
		signatures[i] = sign.signature;
	}
	const hash = web3.utils.sha3(`${signatures[0]},${signatures[1]},${signatures[2]}`);
	return (hash);
}

function createPdf(graduateData, name, login) {
	const	promoYears = graduateData[0].promoYears;
	const	levelIntra = web3.utils.toUtf8(graduateData[0].intraLevel);
	const	doc = new pdfkit({layout: 'landscape', size: 'A4', margins: {top: 0, bottom: 0, left: 0, right: 0}});
	doc.pipe(fs.createWriteStream(`./graduate/${login}.pdf`));
	doc.image('./assets/graduate_marvin_test.jpg', 0, 0, {
		align: 'center',
		valign: 'center'
	 });
	doc.registerFont('Futura', './assets/futura.ttf');
	doc.font('Futura').fontSize(42).text(name, 0, 224, {align: 'center'});
	doc.font('Futura').fontSize(21).text(promoYears, 367, 405);
	doc.font('Futura').fontSize(21).text(levelIntra.split('.')[0], 486, 405);
	doc.font('Futura').fontSize(21).text(levelIntra.split('.')[1], 515, 405);
	doc.font('Futura').fontSize(21).text('21', 115, 528);
	doc.font('Futura').fontSize(21).text('12', 146, 528);
	doc.font('Futura').fontSize(21).text('2020', 181, 528);
	doc.end()
}

module.exports = {
	str2bytes32: str2bytes32,
	multiSig: multiSig,
	createPdf: createPdf
}