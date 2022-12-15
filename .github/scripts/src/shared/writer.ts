import fs from "fs";

export function writeToJson(filePath: string, json: object): void {
  fs.writeFileSync(filePath, JSON.stringify(json, null, 2).concat("\n"));
}
