// Strips unused transitive dependencies from @chainlink/contracts.
// We only consume its Solidity sources; its declared deps (arbitrum, optimism,
// scroll, zksync, extra OpenZeppelin copies, ...) are never imported.
module.exports = {
  hooks: {
    readPackage(pkg) {
      if (pkg.name === '@chainlink/contracts') {
        pkg.dependencies = {};
      }
      return pkg;
    },
  },
};
