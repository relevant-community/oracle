node ./scripts/setparams.js

echo "Collecting genesis txs..."
oracled collect-gentxs

echo "Validating genesis file..."
oracled validate-genesis