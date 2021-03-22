<!--
order: 0
title: "Overview"
parent:
  title: "Oracle Module"
-->

# Oracle & Offchain Workers

## Abstract

This document specifies the Oracle module for the Cosmos SDK.

The Oracle module allows validators to run arbitrary off-chain worker processes and report the results for inclusion in the on-chain state.

Unlike onchain smart contracts, offchain workers are able to run non-deterministic code, like fetching exchange price data via an api call, and long-running computations, i.e. AI alogrithms or graph analysis. All of the validators are expected to run the workers and come to a consensus on the results. Exactly how consensus is reached can be decided by the app developer.

The module is inspired by the [Terra Oracle Module](https://docs.terra.money/dev/spec-oracle.html#concepts) as well as a more recent iteration of it by [Sommelier](https://github.com/PeggyJV/sommelier/tree/main/x/oracle).

### Hi-Level Overview

At its core, the Oracle module accepts and stores arbitrary `Claims`. When submitting a claim, a validator casts a vote for the particular claim. The oracle module tracks all Claims by `ClaimType` and voting `Round`. The module includes some helper methods to tally the votes for each round, however it is the responsibility of external modules to implement this logic as needed.

## Contents

1. **[Concepts](01_concepts.md)**
   - [Worker](01_concepts.md#Worker)
   - [Claim](01_concepts.md#Claim)
   - [Vote](01_concepts.md#Vote)
   - [Prevote](01_concepts.md#Prevote)
   - [Round](01_concepts.md#Round)
   - [PendingRounds](01_concepts.md#PendingRounds)
   - [Tallying Votes](01_concepts.md#Tallying-Votes)
   - [HouseKeeping](01_concepts.md#HouseKeeping)
   - [Rewards and Punishments](01_concepts.md#Rewards-and-Punishments)
   - [Validator Delegation](01_concepts.md#Validator-Delegation)

2) **[State](02_state.md)**
   - [Round](02_state.md#Round)
   - [PendingRounds](02_state.md#PendingRounds)
   - [Prevote](02_state.md#Prevote)
   - [FeedDelegateKey](02_state.md#FeedDelegateKey)
3) **[Messages](03_messages.md)**
   - [MsgDelegate](03_messages.md#MsgDelegate)
   - [MsgPrevote](03_messages.md#MsgPrevote)
   - [MsgVote](03_messages.md#MsgVote)
4) **[Events](04_events.md)**
   - [MsgDelegate](04_events.md#MsgDelegate)
   - [MsgPrevote](04_events.md#MsgPrevote)
   - [MsgVote](04_events.md#MsgVote)
5) **[Params](05_params.md)**
