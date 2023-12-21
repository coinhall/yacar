<!-- omit from toc -->
# `yacar`

> Yet Another Cosmos Asset Registry

- [Files](#files)
  - [`account.json`](#accountjson)
  - [`contract.json`](#contractjson)
  - [`binary.json`](#binaryjson)
  - [`asset.json`](#assetjson)
    - [Example](#example)
  - [`entity.json`](#entityjson)
  - [`pool.json`](#pooljson)
- [FAQs](#faqs)
  - [How do I add a new pool?](#how-do-i-add-a-new-pool)
  - [How do I add a new asset?](#how-do-i-add-a-new-asset)
  - [How do I update my asset's circulating/total supply, icon, and other links?](#how-do-i-update-my-assets-circulatingtotal-supply-icon-and-other-links)
- [Contributing Guide](#contributing-guide)

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

Contains all verified and unverified **native / IBC / CW20 / CW721 assets**. This file will update automatically if all required fields of an asset can be inferred. The optional fields must be updated manually.

```ts
type Asset = {
  // The contract address of the cw20 tokens or denom of the ibc/native coins
  id: string;
  // The entity which created or controls `id`
  // Corresponds to the `name` field of an entity in `entity.json`
  entity?: string | undefined;
  // The canonical name of the asset (eg. "Axelar Wrapped Bitcoin")
  name: string;
  // The ticker of the asset (eg. "axlWBTC")
  symbol: string;
  // The number of decimals of the asset
  decimals: string;
  // The type of this asset: "native" | "ibc" | "cw20" | "cw721" | "tokenfactory"
  type: string;

  // The following fields are all optional
  // The transaction hash that contains "Coinhall verification" memo
  verification_tx?: string | undefined;

  // Static supply values; do not populate both static and dynamic values (see example below)
  // NOTE: these are decimal numbers (ie. NOT the on-chain integers)
  circ_supply?: string | undefined;
  total_supply?: string | undefined;

  // Dynamic supply values; do not populate both static and dynamic values (see example below)
  circ_supply_api?: string | undefined;
  total_supply_api?: string | undefined;

  // These fields are all URL links
  icon?: string | undefined;
  coinmarketcap?: string | undefined;
  coingecko?: string | undefined;
};
```

#### Example

Valid example with static total supply, and dynamic circulating supply:

```js
{
  "id": "ibc/example_asset_id",
  "name": "Example Asset",
  "symbol": "EA",
  "decimals": "6",
  "type": "ibc",
  "circ_supply_api": "https://exampleasset.com/circulating_supply",
  "total_supply": "1000000"
}
```

Invalid example with both static and dynamic circ supply (applies to total supply too):

```js
{
  "id": "ibc/example_asset_id",
  "name": "Example Asset",
  "symbol": "EA",
  "decimals": "6",
  "type": "ibc",
  "circ_supply": "1000000",
  "circ_supply_api": "https://exampleasset.com/circulating_supply",
}
```

### `entity.json`

Contains all **social information of an entity**. This file is updated manually.

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
  asset_ids: string[];
  dex: string;
  // The liquidity pool type: "xyk" | "stable" | "orderbook" | "balancerV1"
  type: string;
  // The contract address of the LP token (if it exists)
  lp_token_id?: string | undefined;
};
```

## FAQs

### How do I add a new pool?

You do *not* need to add pools manually. Our bot will periodically detect and add new pools to the `pool.json` file, provided that the pools come from a dex that we already support.

If you are a new dex that we do not yet support, [please reach out to us](https://t.me/coinhall_org).

### How do I add a new asset?

You do *not* need to add assets manually. Our bot will periodically detect and add new assets to the `asset.json` file, provided that the asset belong to at least one pool.

### How do I update my asset's circulating/total supply, icon, and other links?

**For circulating and total supply**: we accept either a static number, or an API endpoint (if the supply is dynamic). The values should be added to the [`asset.json` file](#assetjson).

**For token icon**: we only accept a valid hosted link to an image (SVG is preferred over PNG/JPEG/others). Please ensure that the link given shows an image only and nothing else (ie. it should NOT lead to an HTML webpage). The link should be added to the [`asset.json` file](#assetjson).

**For CoinMarketCap and Coingecko links**: these links should be added to the [`asset.json` file](#assetjson).

**For social links**: we accept website, Twitter, Telegram, and Discord links. Firstly, the links should be added to the [`entity.json` file](#entityjson). Then, update your asset in `asset.json` with the `entity` field, ensuring that the value is the same as the `name` field of the entity that you have created.

Then, follow the [contributing guide](#contributing) to open a pull request.

## Contributing Guide

1. [Fork this repo](https://github.com/coinhall/yacar/fork)
2. Read the [FAQs](#faqs) and make the necessary changes on your fork
3. Before committing any changes, run the `format.sh` script which automatically validates and formats your local files
4. If the validation passes, [create a pull request](https://github.com/coinhall/yacar/compare)
5. Enable GitHub Actions to run on your fork and ensure that your PR passes all checks
6. If necessary, [contact us](https://t.me/coinhall_org) to seek for a review
