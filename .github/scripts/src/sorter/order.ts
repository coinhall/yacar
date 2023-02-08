import {
  Account,
  Asset,
  Binary,
  Contract,
  Entity,
  Pool,
} from "../shared/schema";

export function orderLabelledKeys(
  data: {
    id: string;
    entity: string;
    label: string;
  }[]
) {
  return data.map((v) => {
    const { id, entity, label } = v;
    return { id, entity, label };
  });
}

export function orderAssetKeys(assetData: Asset[]): Asset[] {
  return assetData.map((v) => {
    const {
      id,
      entity,
      name,
      symbol,
      decimals,
      circ_supply_api,
      icon,
      website,
      telegram,
      twitter,
      discord,
      coinmarketcap,
      coingecko,
    } = v;

    const sortedAsset: Asset = {
      id,
      entity,
      name,
      symbol,
      decimals,
      circ_supply_api,
      icon,
      website,
      telegram,
      twitter,
      discord,
      coinmarketcap,
      coingecko,
    };

    const filteredAsset = Object.fromEntries(
      Object.entries(sortedAsset).filter(([_key, value]) => value != null)
    );

    return filteredAsset as Asset;
  });
}

export function orderEntityKeys(entityData: Entity[]): Entity[] {
  return entityData.map((v) => {
    const {
      entity,
      website,
      telegram,
      twitter,
      discord,
      coinmarketcap,
      coingecko,
    } = v;

    const sortedEntity = {
      entity,
      website,
      telegram,
      twitter,
      discord,
      coinmarketcap,
      coingecko,
    };

    const filteredEntity = Object.fromEntries(
      Object.entries(sortedEntity).filter(([_key, value]) => value != null)
    );

    return filteredEntity as Entity;
  });
}

export function orderPoolKeys(poolData: Pool[]): Pool[] {
  return poolData.map((v) => {
    const { id, lp_token_id, asset_ids, dex, type } = v;
    return { id, lp_token_id, asset_ids, dex, type };
  });
}

export function getOrderededPathJsonMap(
  sortedJsonMap: Record<string, object>
): Record<string, object> {
  const orderedPathJsonMap: Record<string, object> = {};
  const errorPaths: string[] = [];

  for (const [path, jsonData] of Object.entries(sortedJsonMap)) {
    if (path.endsWith("account.json")) {
      orderedPathJsonMap[path] = orderLabelledKeys(jsonData as Account[]);
    } else if (path.endsWith("asset.json")) {
      orderedPathJsonMap[path] = orderAssetKeys(jsonData as Asset[]);
    } else if (path.endsWith("binary.json")) {
      orderedPathJsonMap[path] = orderLabelledKeys(jsonData as Binary[]);
    } else if (path.endsWith("contract.json")) {
      orderedPathJsonMap[path] = orderLabelledKeys(jsonData as Contract[]);
    } else if (path.endsWith("entity.json")) {
      orderedPathJsonMap[path] = orderEntityKeys(jsonData as Entity[]);
    } else if (path.endsWith("pool.json")) {
      orderedPathJsonMap[path] = orderPoolKeys(jsonData as Pool[]);
    } else {
      errorPaths.push(path);
    }
  }

  if (errorPaths.length !== 0) {
    console.error(`Unable to order:\n${errorPaths.join("\n")}`);
    process.exit(1);
  }

  return orderedPathJsonMap;
}
