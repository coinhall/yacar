<!-- omit from toc -->
# `yacar`

> Yet Another Cosmos Asset Registry

- [JSON Files](#json-files)
  - [`account.json`](#accountjson)
  - [`contract.json`](#contractjson)
  - [`binary.json`](#binaryjson)
  - [`asset.json`](#assetjson)
  - [`pool.json`](#pooljson)
- [Contributing](#contributing)

## JSON Files

### `account.json`

Contains **notable user accounts** including native multisig and CW3 multisig accounts.

```ts
type Account = {
  // The address of the wallet or smart contract
  id: string;
  // The entity which created or controls `id`
  entity: string;
  // A short descriptive label of `id`
  label: string;
};
```

### `contract.json`

Contains **notable cosmwasm smart contracts** excluding CW3 multisig accounts.

```ts
type Contract = {
  // The address of the smart contract
  id: string;
  // The entity which created or controls `id`
  entity: string;
  // A short descriptive label of `id`
  label: string;
};
```

### `binary.json`

Contains **notable cosmwasm binaries**.

```ts
type Binary = {
  // The code_id of the cosmwasm binary
  id: string;
  // The entity which created or controls `id`
  entity: string;
  // A short descriptive label of `id`
  label: string;
};
```

### `asset.json`

Contains **verified native / IBC / CW20 assets**.

```ts
type Asset = {
  // The contract address of the cw20 tokens
  // or denom of the ibc/native coins
  id: string;
  name: string;
  symbol: string;
  decimals: string;
  // The optional entity which created or controls `id`
  entity?: string | undefined;
  // Following optional fields are all URL links
  circ_supply_api?: string | undefined;
  icon?: string | undefined;
  website?: string | undefined;
  telegram?: string | undefined;
  twitter?: string | undefined;
  discord?: string | undefined;
  coinmarketcap?: string | undefined;
  coingecko?: string | undefined;
};
```

### `pool.json`

Contains **dexes' liquidity pools**.

```ts
type Pool = {
  // The contract address of the liquidity pool
  id: string;
  asset_ids: string[];
  dex: string;
  type: string; // typically "xyk" or "stable"
};
```

## Contributing

TODO
