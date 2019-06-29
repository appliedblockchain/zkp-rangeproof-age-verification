const fs = require('fs')
const { join } = require('path')
const Web3 = require('web3')
const PROVIDER = 'http://0.0.0.0:8545'

;(async () => {
  try {
    const web3 = new Web3(new Web3.providers.HttpProvider(PROVIDER))
    const coinbase = await web3.eth.getCoinbase()
    const from = coinbase

    const filePath = join(__dirname, '../../../solidity/build/contracts/RangeProofValidator.json')
    const fileContent = fs.readFileSync(filePath)
    const contract = JSON.parse(fileContent)
    const sendParams = { from, gas: 50000000 }

    const deployContract = async contract => {
      const { abi, bytecode } = contract
      const contractInstance = new web3.eth.Contract(abi, { from, data: bytecode })
      return contractInstance.deploy({ arguments: [] }).send(sendParams)
    }

    const RangeProofValidator = await deployContract(contract, 'RangeProofValidator')

    const contractsJSON = `module.exports = ${
      JSON.stringify({
        RangeProofValidator: { address: RangeProofValidator.options.address, abi: RangeProofValidator._jsonInterface, bytecode: RangeProofValidator.bytecode }
      }, null, 2).replace(/"/g, '\'')
    }\n`

    const path = join(__dirname, '../contracts.js')
    fs.writeFileSync(path, contractsJSON)

    console.log(`Contract information saved at ${path}`)
  } catch (err) {
    err.message === 'Invalid JSON RPC response: ""'
      ? console.error('Error: Unable to connect to network, is parity running?')
      : console.error(err)
  }
})()
