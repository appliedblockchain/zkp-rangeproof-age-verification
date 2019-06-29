const { exec } = require('child_process')
const Web3 = require('web3')
const contracts = require('./contracts')

const web3 = new Web3(provider)
const provider = 'http://localhost:8545'

const { RangeProofValidator } = contracts
const { abi, address, bytecode } = RangeProofValidator

const contract = new web3.eth.Contract(abi, address, { from: '0xfFb383e52D4006690a5215a4a978220863147aD2', data: bytecode })

function execAsync(cmd, options = { delimiter: '|' }) {
  return new Promise((resolve, reject) => {
    exec(cmd, (error, stdout, stderr) => {
      if (error) {
        reject(error)
        return
      }

      const trimmed = stdout.trim()

      const result = options.delimiter
        ? trimmed.split(options.delimiter).map(item => item.trim())
        : [ trimmed ]

      resolve(result)
    })
  })
}

;(async () => {
  const lowerBound = 18
  const upperBound = 65
  const [ commitment, proof ] = await execAsync('java -jar ../java/zkrp/classes/artifacts/zkrp_jar/zkrp.jar "generate proof"')
  console.log('COMMITMENT: \n', commitment)
  console.log('PROOF: \n', proof)

  const validated = await contract.methods.validate(lowerBound, upperBound, commitment, proof).call();
  console.log('VALIDATED: ', validated)
})()