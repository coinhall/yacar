import { describe, test, expect } from "vitest";
import fs from "fs";

import { JsonFiles } from "../src/shared/enums";
import { validate } from "../src/validator/validator";

function loadJson(relativePath: string): object {
  const rawData = fs.readFileSync(__dirname + relativePath, {
    encoding: "utf-8",
  });
  return JSON.parse(rawData);
}

function toEnumJsonMap(
  fileType: JsonFiles,
  matchingJsons: object[]
): Record<JsonFiles, object[]> {
  const enumJsonMap: Record<string, object[]> = {};
  enumJsonMap[fileType] = matchingJsons;
  return enumJsonMap;
}

describe("given an account.json", () => {
  const fileType = JsonFiles.ACCOUNT;

  describe("if it is invalid", () => {
    test("validate should return true, indicating an error", () => {
      const invalidAccounts = loadJson("/data/validator/account_01.json");
      const enumJsonMap = toEnumJsonMap(fileType, [invalidAccounts]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("validate should return false, indicating an error", () => {
      const validAccounts = loadJson("/data/validator/account_02.json");
      const enumJsonMap = toEnumJsonMap(fileType, [validAccounts]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeFalsy();
    });
  });
});

describe("given an asset.json", () => {
  const fileType = JsonFiles.ASSET;

  describe("if it is invalid", () => {
    test("validate should return true, indicating an error", () => {
      const invalidAssets = loadJson("/data/validator/asset_01.json");
      const enumJsonMap = toEnumJsonMap(fileType, [invalidAssets]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("validate should return false, indicating an error", () => {
      const validAssets = loadJson("/data/validator/asset_02.json");
      const enumJsonMap = toEnumJsonMap(fileType, [validAssets]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeFalsy();
    });
  });
});

describe("given a binary.json", () => {
  const fileType = JsonFiles.BINARY;

  describe("if it is invalid", () => {
    test("validate should return true, indicating an error", () => {
      const invalidBinary = loadJson("/data/validator/binary_01.json");
      const enumJsonMap = toEnumJsonMap(fileType, [invalidBinary]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("validate should return false, indicating an error", () => {
      const validBinary = loadJson("/data/validator/binary_02.json");
      const enumJsonMap = toEnumJsonMap(fileType, [validBinary]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeFalsy();
    });
  });
});

describe("given a contract.json", () => {
  const fileType = JsonFiles.CONTRACT;

  describe("if it is invalid", () => {
    test("validate should return true, indicating an error", () => {
      const invalidContracts = loadJson("/data/validator/contract_01.json");
      const enumJsonMap = toEnumJsonMap(fileType, [invalidContracts]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("validate should return false, indicating an error", () => {
      const validContracts = loadJson("/data/validator/contract_02.json");
      const enumJsonMap = toEnumJsonMap(fileType, [validContracts]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeFalsy();
    });
  });
});

describe("given a pool.json", () => {
  const fileType = JsonFiles.POOL;

  describe("if it is invalid", () => {
    test("validate should return true, indicating an error", () => {
      const invalidPools = loadJson("/data/validator/pool_01.json");
      const enumJsonMap = toEnumJsonMap(fileType, [invalidPools]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("validate should return false, indicating an error", () => {
      const validPools = loadJson("/data/validator/pool_02.json");
      const enumJsonMap = toEnumJsonMap(fileType, [validPools]);

      const hasError = validate(enumJsonMap);
      expect(hasError).toBeFalsy();
    });
  });
});
