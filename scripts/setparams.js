const fs = require('fs')
const path = '/Users/slavab/.oracle/config/genesis.json'
const genesis = require(path)
genesis.app_state.oracle.params.claim_params.AtomClaim = {
  claim_type: 'AtomClaim',
  prevote: true,
  vote_period: 3,
  vote_threshold: '0.5',
}
fs.writeFile(path, JSON.stringify(genesis), (err) => {
  if (err) {
    console.error(err)
    return
  }
})
