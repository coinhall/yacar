import { Asset, LabelledType, Pool } from "../shared/schema";

export function orderLabelledTypeKeys(
  labelledTypeData: LabelledType[]
): LabelledType[] {
  return labelledTypeData.map((v) => {
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

export function orderPoolKeys(poolData: Pool[]): Pool[] {
  return poolData.map((v) => {
    const { id, asset_ids, dex, type } = v;
    return { id, asset_ids, dex, type };
  });
}

export function getOrderededPathJsonMap(
  sortedJsonMap: Record<string, object>
): Record<string, object> {
  const orderedPathJsonMap: Record<string, object> = {};
  const errorPaths: string[] = [];

  for (const [path, jsonData] of Object.entries(sortedJsonMap)) {
    if (
      path.endsWith("account.json") ||
      path.endsWith("binary.json") ||
      path.endsWith("contract.json")
    ) {
      orderedPathJsonMap[path] = orderLabelledTypeKeys(
        jsonData as LabelledType[]
      );
    } else if (path.endsWith("asset.json")) {
      orderedPathJsonMap[path] = orderAssetKeys(jsonData as Asset[]);
    } else if (path.endsWith("pool.json")) {
      orderedPathJsonMap[path] = orderPoolKeys(jsonData as Pool[]);
    } else {
      errorPaths.push(path);
    }
  }
  if (errorPaths.length !== 0) {
    console.warn(`Unable to sort:\n${errorPaths.join("\n")}`);
  }
  return orderedPathJsonMap;
}
