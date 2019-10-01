const	GraduateMarvin = artifacts.require("GraduateMarvinCore");

contract("GraduateMarvinCore", async (accounts) => {

	it("Testing function createGraduate", async () => {
		let instance = await GraduateMarvin.deployed();
		let new_graduate = await instance.createGraduate(
				web3.utils.utf8ToHex('Login'),
				web3.utils.utf8ToHex('FirstName'),
				web3.utils.utf8ToHex('LastName'),
				web3.utils.utf8ToHex('IntraLevel'),
				web3.utils.utf8ToHex('BirthDate'),
				web3.utils.utf8ToHex('City'),
				web3.utils.utf8ToHex('Country'),
				2017,
				2020
			);
		let graduate = await instance.getGraduate(web3.utils.utf8ToHex('Login'));
	})

})