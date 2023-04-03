import { describe, test, expect } from "vitest";
import fs from "fs";

import { JsonFiles } from "../src/shared/enums";
import { Account, Asset, Binary, Contract, Pool } from "../src/shared/schema";
import {
  duplicateCheck,
  getSchema,
  schemaErrorCheck,
} from "../src/validator/validator";

function loadJson(relativePath: string): object {
  const rawData = fs.readFileSync(__dirname + relativePath, {
    encoding: "utf-8",
  });
  return JSON.parse(rawData);
}

describe("given an account.json", () => {
  const fileType = JsonFiles.ACCOUNT;
  const schema = getSchema(fileType);

  describe("if it is invalid", () => {
    test("schemaErrorCheck should return true, indicating an error", () => {
      const data = loadJson("/data/validator/account_01.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("schemaErrorCheck should return false, indicating no error", () => {
      const data = loadJson("/data/validator/account_02.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeFalsy();
    });
  });

  describe("if there exists a duplicate id", () => {
    test("duplicateCheck returns true, indicating a duplicate", () => {
      const duplicateIdAccounts = loadJson(
        "/data/validator/account_03.json"
      ) as Account[];

      expect(duplicateCheck(fileType, duplicateIdAccounts)).toBeTruthy();
    });
  });
});

describe("given an asset.json", () => {
  const fileType = JsonFiles.ASSET;
  const schema = getSchema(fileType);

  describe("if it is invalid", () => {
    test("schemaErrorCheck should return true, indicating an error", () => {
      const data = loadJson("/data/validator/asset_01.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("schemaErrorCheck should return false, indicating no error", () => {
      const data = loadJson("/data/validator/asset_02.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeFalsy();
    });
  });

  describe("if there exists a duplicate id", () => {
    test("duplicateCheck returns true, indicating a duplicate", () => {
      const duplicateIdAssets = loadJson(
        "/data/validator/asset_03.json"
      ) as Asset[];

      expect(duplicateCheck(fileType, duplicateIdAssets)).toBeTruthy();
    });
  });

  describe("if an asset contains an emoji, expect no error", () => {
    test("schemaErrorCheck should return true, indicating an error", () => {
      const data = loadJson("/data/validator/asset_04.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeFalsy();
    });
  });
});

describe("given a binary.json", () => {
  const fileType = JsonFiles.BINARY;
  const schema = getSchema(fileType);

  describe("if it is invalid", () => {
    test("schemaErrorCheck should return true, indicating an error", () => {
      const data = loadJson("/data/validator/binary_01.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("schemaErrorCheck should return false, indicating no error", () => {
      const data = loadJson("/data/validator/binary_02.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeFalsy();
    });
  });

  describe("if there exists a duplicate id", () => {
    test("duplicateCheck returns true, indicating a duplicate", () => {
      const duplicateIdBinary = loadJson(
        "/data/validator/binary_03.json"
      ) as Binary[];

      expect(duplicateCheck(fileType, duplicateIdBinary)).toBeTruthy();
    });
  });
});

describe("given a contract.json", () => {
  const fileType = JsonFiles.CONTRACT;
  const schema = getSchema(fileType);

  describe("if it is invalid", () => {
    test("schemaErrorCheck should return true, indicating an error", () => {
      const data = loadJson("/data/validator/contract_01.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("schemaErrorCheck should return false, indicating no error", () => {
      const data = loadJson("/data/validator/contract_02.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeFalsy();
    });
  });

  describe("if there exists a duplicate id", () => {
    test("duplicateCheck returns true, indicating a duplicate", () => {
      const duplicateIdContract = loadJson(
        "/data/validator/contract_03.json"
      ) as Contract[];

      expect(duplicateCheck(fileType, duplicateIdContract)).toBeTruthy();
    });
  });
});

describe("given an entity.json", () => {
  const fileType = JsonFiles.ENTITY;
  const schema = getSchema(fileType);

  describe("if it is invalid", () => {
    test("schemaErrorCheck should return true, indicating an error", () => {
      const data = loadJson("/data/validator/entity_01.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("schemaErrorCheck should return false, indicating no error", () => {
      const data = loadJson("/data/validator/entity_02.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeFalsy();
    });
  });

  describe("if there exists a duplicate entity", () => {
    test("duplicateCheck returns true, indicating a duplicate", () => {
      const duplicateIdPool = loadJson(
        "/data/validator/pool_03.json"
      ) as Pool[];

      expect(duplicateCheck(fileType, duplicateIdPool)).toBeTruthy();
    });
  });
});

describe("given a pool.json", () => {
  const fileType = JsonFiles.POOL;
  const schema = getSchema(fileType);

  describe("if it is invalid", () => {
    test("schemaErrorCheck should return true, indicating an error", () => {
      const data = loadJson("/data/validator/pool_01.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeTruthy();
    });
  });

  describe("if it is valid", () => {
    test("schemaErrorCheck should return false, indicating no error", () => {
      const data = loadJson("/data/validator/pool_02.json");
      const hasError = schemaErrorCheck(schema, data);

      expect(hasError).toBeFalsy();
    });
  });

  describe("if there exists a duplicate id", () => {
    test("duplicateCheck returns true, indicating a duplicate", () => {
      const duplicateIdPool = loadJson(
        "/data/validator/pool_03.json"
      ) as Pool[];

      expect(duplicateCheck(fileType, duplicateIdPool)).toBeTruthy();
    });
  });
});
