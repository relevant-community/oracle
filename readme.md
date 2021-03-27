# Cosmos Oracle

The Oracle module allows validators to run arbitrary off-chain worker processes and report the results for inclusion in the on-chain state.

Unlike onchain smart contracts, offchain workers are able to run non-deterministic code, like fetching exchange price data via an api call, and long-running computations, i.e. AI alogrithms or graph analysis. All of the validators are expected to run the workers and come to a consensus on the results. Exactly how consensus is reached can be decided by the app developer.

The module supports running an arbitrary amount of different oracles simultaniously.

The module is inspired by the [Terra Oracle Module](https://docs.terra.money/dev/spec-oracle.html#concepts) as well as a more recent iteration of it by [Sommelier](https://github.com/PeggyJV/sommelier/tree/main/x/oracle).

**oracle** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Getting Started

- Tutorial: https://relevant-community.github.io/oracle/tutorial/
- Documentation: https://relevant-community.github.io/oracle/

Run the demo app that fetches ATOM/USD price:

Clone the repo into your local folder and run:

```
$ starport serve
```

In a new browser window, star the oracle worker process:

```
$ oracled tx oracle start-worker --from alice -o text -y
```

After some time, check the on-chain Atom-USD price:

```
$ oracled query atom atomUsd
```

You should see the latest Atom/USD price and the blockHeight at which it was captured.

---

## Starport Docs:

```
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Launch

To launch your blockchain live on mutliple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)
