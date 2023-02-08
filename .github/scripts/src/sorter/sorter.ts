import {
  Account,
  Asset,
  Binary,
  Contract,
  Entity,
  Pool,
} from "../shared/schema";

export function sortLabelledTypes(
  data: {
    id: string;
    entity: string;
    label: string;
  }[]
) {
  return data.sort((a, b) => {
    return (
      a.entity.localeCompare(b.entity, "en-US") ||
      a.label.localeCompare(b.label, "en-US") ||
      a.id.localeCompare(b.id, "en-US")
    );
  });
}

export function sortAsset(jsonData: Asset[]): Asset[] {
  const withEntities = jsonData
    .filter((v) => v.entity)
    .sort((a, b) => {
      return (
        a.entity!.localeCompare(b.entity!, "en-US") ||
        a.name.localeCompare(b.name, "en-US") ||
        a.id.localeCompare(b.id, "en-US")
      );
    });

  const withoutEntities = jsonData
    .filter((v) => !v.entity)
    .sort((a, b) => {
      return (
        a.name.localeCompare(b.name, "en-US") ||
        a.id.localeCompare(b.id, "en-US")
      );
    });

  return [...withEntities, ...withoutEntities];
}

export function sortBinary(jsonData: Binary[]): Binary[] {
  return jsonData.sort((a, b) => {
    return a.id.localeCompare(b.id, "en-US", { numeric: true });
  });
}

export function sortEntity(jsonData: Entity[]): Entity[] {
  return jsonData.sort((a, b) => {
    return a.entity.localeCompare(b.entity, "en-US");
  });
}

export function sortPool(jsonData: Pool[]): Pool[] {
  return jsonData.sort((a, b) => {
    return (
      a.dex.localeCompare(b.dex, "en-US") ||
      a.type.localeCompare(b.type, "en-US") ||
      a.id.localeCompare(b.id, "en-US")
    );
  });
}

export function getSortedPathJsonMap(
  pathJsonMap: Record<string, object>
): Record<string, object> {
  const sortedPathJsonMap: Record<string, object> = {};
  const errorPaths: string[] = [];

  for (const [path, jsonData] of Object.entries(pathJsonMap)) {
    if (path.endsWith("account.json")) {
      sortedPathJsonMap[path] = sortLabelledTypes(jsonData as Account[]);
    } else if (path.endsWith("asset.json")) {
      sortedPathJsonMap[path] = sortAsset(jsonData as Asset[]);
    } else if (path.endsWith("binary.json")) {
      sortedPathJsonMap[path] = sortBinary(jsonData as Binary[]);
    } else if (path.endsWith("contract.json")) {
      sortedPathJsonMap[path] = sortLabelledTypes(jsonData as Contract[]);
    } else if (path.endsWith("entity.json")) {
      sortedPathJsonMap[path] = sortEntity(jsonData as Entity[]);
    } else if (path.endsWith("pool.json")) {
      sortedPathJsonMap[path] = sortPool(jsonData as Pool[]);
    } else {
      errorPaths.push(path);
    }
  }

  if (errorPaths.length !== 0) {
    console.error(`Unable to sort:\n${errorPaths.join("\n")}`);
    process.exit(1);
  }

  return sortedPathJsonMap;
}
