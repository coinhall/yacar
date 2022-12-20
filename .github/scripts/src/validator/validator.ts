import { TypeCheck } from "@sinclair/typebox/compiler";

import { JsonFiles } from "../shared/enums";
import {
  AccountSchema,
  AssetSchema,
  BinarySchema,
  ContractSchema,
  PoolSchema,
} from "../shared/schema";

function getSchema(key: string): TypeCheck<any> {
  switch (key) {
    case JsonFiles.ACCOUNT:
      return AccountSchema;
    case JsonFiles.ASSET:
      return AssetSchema;
    case JsonFiles.BINARY:
      return BinarySchema;
    case JsonFiles.CONTRACT:
      return ContractSchema;
    case JsonFiles.POOL:
      return PoolSchema;
    default:
      console.error(`Unable to get schema for "${key}"`);
      process.exit(1);
  }
}

export function validate(enumJsonMap: Record<string, object[]>): boolean {
  let hasDetectedError = false;

  for (const [key, jsons] of Object.entries(enumJsonMap)) {
    const schema = getSchema(key);
    for (const json of jsons) {
      const errors = [...schema.Errors(json)].map((v) => {
        const { path, value, message } = v;
        return { path, value, message };
      });
      if (errors.length !== 0) {
        console.error(errors);
        hasDetectedError = true;
      }
    }
  }
  return hasDetectedError;
}
