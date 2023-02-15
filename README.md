<!-- omit from toc -->
# `yacar`

> Yet Another Cosmos Asset Registry

- [Files](#files)
  - [`account.json`](#accountjson)
  - [`contract.json`](#contractjson)
  - [`binary.json`](#binaryjson)
  - [`asset.json`](#assetjson)
  - [`entity.json`](#entityjson)
  - [`pool.json`](#pooljson)
- [Contributing](#contributing)

## Files

### `account.json`

Contains **notable user addresses** including native multisig and CW3 multisig accounts. This file is updated manually.

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

Contains **notable cosmwasm smart contracts** excluding CW3 multisig accounts. This file is updated manually.

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

Contains **notable cosmwasm binaries**. This file is updated manually.

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

Contains all verified and unverified **native / IBC / CW20 assets**. This file will update automatically if all required fields of an asset can be inferred. The optional fields must be updated manually.

```ts
type Asset = {
  // The contract address of the cw20 tokens
  // or denom of the ibc/native coins
  id: string;
  // The entity which created or controls `id`
  // A nullish value means that the asset is "unverified"
  entity?: string | undefined;
  // The canonical name of the asset (eg. "Axelar Wrapped Bitcoin")
  name: string;
  // The ticker of the asset (eg. "axlWBTC")
  symbol: string;
  // The number of decimals of the asset
  decimals: string;
  // Following optional fields are all URL links
  circ_supply_api?: string | undefined;
  icon?: string | undefined;
  coingecko?: string | undefined;
  coinmarketcap?: string | undefined;
};
```

### `entity.json`

Contains all social information of a project. This file is updated manually.

```ts
type Entity = {
  name: string;
  website?: string | undefined;
  telegram?: string | undefined;
  twitter?: string | undefined;
  discord?: string | undefined;
}
```

### `pool.json`

Contains **dexes' liquidity pools**. This file will update automatically if all required fields of a pool can be inferred (specifically, `dex` and `type`). Otherwise, the missing fields must be updated manually.

```ts
type Pool = {
  // The contract address of the liquidity pool
  id: string;
  // The contract address of the LP token
  lp_token_id: string;
  asset_ids: string[];
  dex: string;
  // The liquidity pool type; typically "xyk" or "stable"
  type: string;
};
```

## Contributing

1. [Fork this repo](https://github.com/coinhall/yacar/fork)
2. Push changes to your fork
3. The files will be validated and formatted automatically
4. If validation passes, [create a pull request](https://github.com/coinhall/yacar/compare)
5. If necessary, seek for a review via [Telegram](https://t.me/coinhall_org)
