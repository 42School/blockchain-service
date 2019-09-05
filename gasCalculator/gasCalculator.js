const	fs = require('fs');
const	web3 = require('web3')
const	wildcard = require('wildcard');
const	folderGas = './gasCalculator';

let average = 0;
let total = 0;
let etherAverageCost = 0;
let etherCost = 0;
let y = 0;
let nbStudent = 0;

fs.readdirSync(folderGas).forEach(file => {
	if (wildcard('gasUsed_*.json', file)) {
		let gasUsed = fs.readFileSync(`${folderGas}/${file}`);
		gasUsed = JSON.parse(gasUsed);
		nbStudent = nbStudent + 500;
		for (var i in gasUsed) {
			total = total + gasUsed[i];
			y++;
		}
	}
});

average = Math.round(total / y);
etherAverageCost = web3.utils.toWei(`${average}`, 'Gwei');
etherCost = web3.utils.toWei(`${total}`, 'Gwei');
etherAverageCost = web3.utils.fromWei(`${etherAverageCost}`, 'ether');
etherCost = web3.utils.fromWei(`${etherCost}`, 'ether');
console.log(`Total of ${nbStudent} graduates: ${total}`);
console.log(`Average cost of one graduate: ${average}`);
console.log(`Transalte average cost in Ether: ${etherAverageCost}`);
console.log(`Transalte ${nbStudent} graduates cost in Ether: ${etherCost}`);