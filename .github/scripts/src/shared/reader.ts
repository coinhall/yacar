import fs from "fs";
import { glob } from "glob";
import { JsonFiles } from "./enums";

export const absoluteJsonGlob = (
  projRoot: string,
  files: string,
  dirs?: string
): string => {
  return dirs ? `${projRoot}/${dirs}/${files}` : `${projRoot}/${files}`;
};

export function filePathsFromGlob(globPattern: string): string[] {
  return glob.sync(globPattern);
}

function fileEnumFromPath(path: string): JsonFiles {
  if (path.endsWith("account.json")) {
    return JsonFiles.ACCOUNT;
  } else if (path.endsWith("asset.json")) {
    return JsonFiles.ASSET;
  } else if (path.endsWith("binary.json")) {
    return JsonFiles.BINARY;
  } else if (path.endsWith("contract.json")) {
    return JsonFiles.CONTRACT;
  } else if (path.endsWith("pool.json")) {
    return JsonFiles.POOL;
  } else {
    console.error("Unkown file detected, unable to provide file enum");
    process.exit(1);
  }
}

export function getEnumContentMap(
  filePaths: string[]
): Record<JsonFiles, object[]> {
  const enumJsonMap: Record<string, object[]> = {};
  for (const path of filePaths) {
    const mapKey = fileEnumFromPath(path);
    const rawData = fs.readFileSync(path, { encoding: "utf-8" });
    const jsonData = JSON.parse(rawData);
    if (!enumJsonMap[mapKey]) {
      enumJsonMap[mapKey] = [jsonData];
    } else {
      enumJsonMap[mapKey].push(jsonData);
    }
  }
  return enumJsonMap;
}

export function getPathJsonMap(filePaths: string[]): Record<string, object> {
  const pathJsonMap: Record<string, object> = {};
  for (const path of filePaths) {
    const rawData = fs.readFileSync(path, { encoding: "utf-8" });
    const jsonData = JSON.parse(rawData);
    pathJsonMap[path] = jsonData;
  }
  return pathJsonMap;
}
