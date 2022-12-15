import Ajv from "ajv/dist/jtd";
import betterAjvErrors from "better-ajv-errors";

import { JsonFiles } from "../shared/enums";
import {
  AccountSchema,
  AssetSchema,
  BinarySchema,
  ContractSchema,
  PoolSchema,
} from "../shared/types";

function getSchema(key: string) {
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
  const ajv = new Ajv({ allErrors: true });
  let hasDetectedError = false;

  for (const [key, jsons] of Object.entries(enumJsonMap)) {
    const schema = getSchema(key);
    const validate = ajv.compile(schema);
    for (const json of jsons) {
      const valid = ajv.validate(schema, json);
      if (!valid) {
        hasDetectedError = true;
        const output = betterAjvErrors(schema, json, validate.errors!, {
          indent: 2,
        });
        console.error(output);
      }
    }
  }
  return hasDetectedError;
}
