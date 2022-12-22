import { TypeCheck } from "@sinclair/typebox/compiler";

import { JsonFiles } from "../shared/enums";
import {
  AccountSchema,
  AssetSchema,
  BinarySchema,
  ContractSchema,
  PoolSchema,
} from "../shared/schema";

export function getSchema(key: string): TypeCheck<any> {
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

/** Returns true if duplicate exists */
export function duplicateCheck(type: string, data: { id: string }[]): boolean {
  const uniqueIdSet = new Set<string>();
  const duplicateIdSet = new Set<string>();

  for (const { id } of data) {
    if (uniqueIdSet.has(id)) {
      duplicateIdSet.add(id);
    } else {
      uniqueIdSet.add(id);
    }
  }

  if (duplicateIdSet.size !== 0) {
    console.error("Duplicate(s) detected in:", type);
    console.error([...duplicateIdSet], "\n");
    return true;
  }

  return false;
}

/** Returns true if error exists */
export function schemaErrorCheck(schema: TypeCheck<any>, data: object) {
  const schemaErrors = [...schema.Errors(data)].map((v) => {
    const { path, value, message } = v;
    return { path, value, message };
  });

  if (schemaErrors.length !== 0) {
    console.error(schemaErrors);
    return true;
  }

  return false;
}

export function validate(enumJsonMap: Record<string, object[]>) {
  let hasDetectedError = false;

  for (const [key, jsons] of Object.entries(enumJsonMap)) {
    const schema = getSchema(key);
    for (const json of jsons) {
      const hasDuplicate = duplicateCheck(key, json as { id: string }[]);
      if (hasDuplicate) {
        hasDetectedError = true;
      }

      const hasSchemaERror = schemaErrorCheck(schema, json);
      if (hasSchemaERror) {
        hasDetectedError = true;
      }
    }
  }

  if (hasDetectedError) {
    process.exit(1);
  }
}
