const fs = require('fs');
const web3 = require('web3')

let gasUsed0 = fs.readFileSync('./gasUsedFor500Graduates.json');
let gasUsed1 = fs.readFileSync('./UsersCreationGasUsed.json');
gasUsed0 = JSON.parse(gasUsed0);
gasUsed1 = JSON.parse(gasUsed1);

let average = 0;
let total = 0;
let etherAverageCost = 0;
let etherCost = 0;
let y = 0;

for (var i in gasUsed0) {
	total = total + gasUsed0[i];
	y++;
}
for (var i in gasUsed1) {
	total = total + gasUsed1[i];
	y++;
}
average = Math.round(total / y);
etherAverageCost = web3.utils.toWei(`${average}`, 'Gwei');
etherCost = web3.utils.toWei(`${total}`, 'Gwei');
etherAverageCost = web3.utils.fromWei(`${etherAverageCost}`, 'ether');
etherCost = web3.utils.fromWei(`${etherCost}`, 'ether');
console.log(`Total of 1000 graduates: ${total}`);
console.log(`Average cost of one graduate: ${average}`);
console.log(`Transalte average cost in Ether: ${etherAverageCost}`);
console.log(`Transalte 1000 graduates cost in Ether: ${etherCost}`);