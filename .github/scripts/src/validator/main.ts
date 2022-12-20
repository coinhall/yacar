import { validate } from "./validator";
import { getRootDir } from "../shared/config";
import { ChainDirectories, JsonFiles } from "../shared/enums";
import {
  absoluteJsonGlob,
  filePathsFromGlob,
  getEnumContentMap,
} from "../shared/reader";

function main(): void {
  const projRoot = getRootDir();

  // Get jsons to validate and load into memory
  const chainDirsGlob = `{${Object.values(ChainDirectories).join(",")}}`;
  const fileNamesGlob = `{${Object.values(JsonFiles).join(",")}}`;
  const fileGlobPattern = `${absoluteJsonGlob(
    projRoot,
    fileNamesGlob,
    chainDirsGlob
  )}.json`;
  const filePaths = filePathsFromGlob(fileGlobPattern);
  const enumJsonMap = getEnumContentMap(filePaths);

  // Validate matching chain data JSONS to their respective schema
  const hasError = validate(enumJsonMap);

  // Exit with an error if any schema is invalid
  if (hasError) {
    console.error("Invalid files were detacted!");
    process.exit(1);
  }

  console.log(`Validated the following files:\n  ${filePaths.join("\n  ")}`);
}

main();
