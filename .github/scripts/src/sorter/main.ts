import { getOrderededPathJsonMap } from "./order";
import { getSortedPathJsonMap } from "./sorter";

import { getRootDir } from "../shared/config";
import { ChainDirectories, JsonFiles } from "../shared/enums";
import {
  absoluteJsonGlob,
  filePathsFromGlob,
  getPathJsonMap,
} from "../shared/reader";
import { writeToJson } from "../shared/writer";

function main(): void {
  // Get file paths to all matching chain-data JSON
  const projRoot = getRootDir();
  const fileNamesGlob = `{${Object.values(JsonFiles).join(",")}}.json`;
  const chainDirsGlob = `{${Object.values(ChainDirectories).join(",")}}`;
  const fileGlobPattern = absoluteJsonGlob(
    projRoot,
    fileNamesGlob,
    chainDirsGlob
  );
  const filePaths = filePathsFromGlob(fileGlobPattern);
  const pathJsonMap = getPathJsonMap(filePaths);

  // Sort JSON keys according to their types
  const sortedPathJsonMap = getSortedPathJsonMap(pathJsonMap);

  // Ensure keys are in order
  const orderedPathJsonMap = getOrderededPathJsonMap(sortedPathJsonMap);

  // Write changes
  for (const [filePath, json] of Object.entries(orderedPathJsonMap)) {
    writeToJson(filePath, json);
  }
  console.log(`Sorted the following files:\n  ${filePaths.join("\n  ")}`);
}

main();
