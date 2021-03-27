---
title: Cosmos Oracle Module
description: The documentation of the Cosmos Oracle Module.
footer:
  newsletter: false
aside: true
---

# Cosmos Oracle Module

The Oracle module allows validators to run arbitrary off-chain worker processes and report the results for inclusion in the on-chain state.

Unlike onchain smart contracts, offchain workers are able to run non-deterministic code, like fetching exchange price data via an api call, and long-running computations, i.e. AI alogrithms or graph analysis. All of the validators are expected to run the workers and come to a consensus on the results. Exactly how consensus is reached can be decided by the app developer.

The module is inspired by the [Terra Oracle Module](https://docs.terra.money/dev/spec-oracle.html#concepts) as well as a more recent iteration of it by [Sommelier](https://github.com/PeggyJV/sommelier/tree/main/x/oracle).

## Getting Started

- **[The Tutoral](./tutorial)** goes through the steps required to incorporate the Oracle module into your codebase. As an example, we fetch the Atom/USDC price data from Binance and write it on-chain.

- **[Module Documentation](./modules/oracle)**: Describes the high-level architecture and module apis.

## Learn More About Oracles

- This [Terra Blog Post](https://medium.com/stakewithus/terra-oracle-voter-by-stakewith-us-d54a1321beb9) has a nice general overview of core concepts.
- [Terra Oracle Module Docs](https://docs.terra.money/dev/spec-oracle.html#concepts) are also very insightful
- A useful [talk about oracles by Chainlink's Sergey Nazarov](https://youtu.be/UAP6--JTAlU)
- [Sommelier Oracle Module](https://github.com/PeggyJV/sommelier/tree/b2f81e9007db479ac5c88bf4d6edbc17a27120fc/x/oracle)

## Other Resources

- **[Cosmos SDK API Reference](https://godoc.org/github.com/cosmos/cosmos-sdk)**: Godocs of the Cosmos SDK.
