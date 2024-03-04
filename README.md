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
  - [How do I update my asset's market cap or FDV?](#how-do-i-update-my-assets-market-cap-or-fdv)
  - [Why is my GitHub Actions CI failing?](#why-is-my-github-actions-ci-failing)
  - [How can I expedite merging of PRs?](#how-can-i-expedite-merging-of-prs)
  - [How can I verify my asset?](#how-can-i-verify-my-asset)
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

You do *not* need to add assets manually. Our bot will periodically detect and add new assets to the `asset.json` file, provided that the asset belong to at least one pool, and that the correct metadata of the asset exists. Follow these [troubleshooting steps](https://coinhall.gitbook.io/coinhall-wiki/developer-docs/token-listing-update-info-verification#troubleshooting-steps) to ensure that your asset has the correct on-chain metadata.

### How do I update my asset's circulating/total supply, icon, and other links?

**For circulating and total supply**: we accept either a static number, or an API endpoint (if the supply is dynamic). The values should be added to the [`asset.json` file](#assetjson).

**For asset icon**: we only accept a valid hosted link to an image (SVG is preferred over PNG/JPEG/others). Please ensure that the link given shows an image only and nothing else (ie. it should NOT lead to an HTML webpage). The link should be added to the [`asset.json` file](#assetjson).

**For CoinMarketCap and Coingecko links**: these links should be added to the [`asset.json` file](#assetjson).

**For social links**: we accept website, Twitter, Telegram, and Discord links. Firstly, the links should be added to the [`entity.json` file](#entityjson). Then, update your asset in `asset.json` with the `entity` field, ensuring that the value is the same as the `name` field of the entity that you have created.

Then, follow the [contributing guide](#contributing) to open a pull request.

### How do I update my asset's market cap or FDV?

Market cap is derived using circulating supply and FDV is derived using total supply. If an asset's market cap or FDV is wrong/missing, it usually means that the circulating or total supply values are wrong/missing. Updating the circulating or total supply values should fix the issue.

### Why is my GitHub Actions CI failing?

Please ensure that you run the formatting scripts, `format.sh` (for MacOS/Linux users) or `format.bat` (for Windows users), before committing and pushing any changes. The GitHub Actions CI will run the same checks as the script. If running the script fails on your local repository, then the CI for your remote repository will fail as well.

### How can I expedite merging of PRs?

We do a routine check every two business days to manually merge open PRs. If you have an urgent need to list your pool or update your asset info, please refer to the [Coinhall docs](https://coinhall.gitbook.io/coinhall-wiki/developer-docs/token-listing-update-info-verification).

### How can I verify my asset?

Please refer to the [Coinhall docs](https://coinhall.gitbook.io/coinhall-wiki/developer-docs/token-listing-update-info-verification#token-verification).

## Contributing Guide

1. Read the [FAQs](#faqs) and the [Coinhall docs](https://coinhall.gitbook.io/coinhall-wiki/developer-docs/token-listing-update-info-verification)
2. [Fork this repo](https://github.com/coinhall/yacar/fork) and clone the forked repo locally
3. Make the necessary changes on your local repo
4. Run the `format.sh` (for MacOS/Linux users) or `format.bat` (for Windows users) script which automatically validates and formats your local files
5. If the validation and formatting passes, commit your changes and [create a pull request](https://github.com/coinhall/yacar/compare)
6. Enable GitHub Actions to run on your fork and ensure that your PR passes all CI checks
7. Wait for your PRs to be merged (up to two business days)
