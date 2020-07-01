const FtDiploma = artifacts.require("FtDiplomaBase");

module.exports = function(deployer) {
  deployer.deploy(FtDiploma);
};