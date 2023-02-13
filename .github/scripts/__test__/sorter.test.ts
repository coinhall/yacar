import { describe, expect, test } from "vitest";
import fs from "fs";

import {
  Account,
  Asset,
  Binary,
  Contract,
  Entity,
  Pool,
} from "../src/shared/schema";
import {
  orderAssetKeys,
  orderEntityKeys,
  orderLabelledKeys,
  orderPoolKeys,
} from "../src/sorter/order";
import {
  sortAsset,
  sortBinary,
  sortEntity,
  sortLabelledTypes,
  sortPool,
} from "../src/sorter/sorter";

function loadJson(relativePath: string): object {
  const rawData = fs.readFileSync(__dirname + relativePath, {
    encoding: "utf-8",
  });
  return JSON.parse(rawData);
}

describe("given an account.json", () => {
  describe("if the object's keys are out of order", () => {
    test("sort the object's keys", () => {
      const unorderedKeyAccounts = loadJson(
        "/data/sorter/raw/account_01.json"
      ) as Account[];

      const expectedAccounts = loadJson(
        "/data/sorter/expected/account_01.json"
      ) as Account[];

      expect(orderLabelledKeys(unorderedKeyAccounts)).toStrictEqual(
        expectedAccounts
      );
    });
  });

  describe("if the objects are out of order", () => {
    test("objects should be sorted (entity, label, id)", () => {
      const shuffledAccounts = loadJson(
        "/data/sorter/raw/account_02.json"
      ) as Account[];

      const expectedAccounts = loadJson(
        "/data/sorter/expected/account_02.json"
      ) as Account[];

      expect(sortLabelledTypes(shuffledAccounts)).toStrictEqual(
        expectedAccounts
      );
    });
  });
});

describe("given an asset.json", () => {
  describe("if the object's keys are out of order", () => {
    test("sort the object's keys", () => {
      const unorderedKeyAssets = loadJson(
        "/data/sorter/raw/asset_01.json"
      ) as Asset[];

      const expectedAssets = loadJson(
        "/data/sorter/expected/asset_01.json"
      ) as Asset[];

      expect(orderAssetKeys(unorderedKeyAssets)).toStrictEqual(expectedAssets);
    });
  });

  describe("if the objects are out of order", () => {
    test("objects should be sorted (entity, name, id)", () => {
      const shuffledAssets = loadJson(
        "/data/sorter/raw/asset_02.json"
      ) as Asset[];

      const expectedAssets = loadJson(
        "/data/sorter/expected/asset_02.json"
      ) as Asset[];

      expect(sortAsset(shuffledAssets)).toStrictEqual(expectedAssets);
    });
  });
});

describe("given an binary.json", () => {
  describe("if the object's keys are out of order", () => {
    test("sort the object's keys (id, entity, label)", () => {
      const unorderedKeyBinary = loadJson(
        "/data/sorter/raw/binary_01.json"
      ) as Binary[];

      const expectedBinary = loadJson(
        "/data/sorter/expected/binary_01.json"
      ) as Binary[];

      expect(orderLabelledKeys(unorderedKeyBinary)).toStrictEqual(
        expectedBinary
      );
    });
  });

  describe("if the objects are out of order", () => {
    test("objects should be sorted (id)", () => {
      const shuffledBinary = loadJson(
        "/data/sorter/raw/binary_02.json"
      ) as Binary[];

      const expectedBinary = loadJson(
        "/data/sorter/expected/binary_02.json"
      ) as Binary[];

      expect(sortBinary(shuffledBinary)).toStrictEqual(expectedBinary);
    });
  });
});

describe("given a contract.json", () => {
  describe("if the object's keys are out of order", () => {
    test("sort the object's keys (id, entity, label)", () => {
      const unorderedKeyContracts = loadJson(
        "/data/sorter/raw/contract_01.json"
      ) as Contract[];

      const expectedContracts = loadJson(
        "/data/sorter/expected/contract_01.json"
      ) as Contract[];

      expect(orderLabelledKeys(unorderedKeyContracts)).toStrictEqual(
        expectedContracts
      );
    });
  });

  describe("if the objects are out of order", () => {
    test("objects should be sorted (id, entity, label)", () => {
      const shuffledContracts = loadJson(
        "/data/sorter/raw/contract_02.json"
      ) as Contract[];

      const expectedContracts = loadJson(
        "/data/sorter/expected/contract_02.json"
      ) as Contract[];

      expect(sortLabelledTypes(shuffledContracts)).toStrictEqual(
        expectedContracts
      );
    });
  });
});

describe("given an entity.json", () => {
  describe("if the object's keys are out of order ", () => {
    test("sort the object's keys", () => {
      const unorderedKeyEntity = loadJson(
        "/data/sorter/raw/entity_01.json"
      ) as Entity[];

      const expectedEntity = loadJson(
        "/data/sorter/expected/entity_01.json"
      ) as Entity[];

      expect(orderEntityKeys(unorderedKeyEntity)).toStrictEqual(expectedEntity);
    });
  });

  describe("if the objects are out of order", () => {
    test("objects should be sorted", () => {
      const shuffledEntities = loadJson(
        "/data/sorter/raw/entity_02.json"
      ) as Entity[];

      const expectedEntities = loadJson(
        "/data/sorter/expected/entity_02.json"
      ) as Entity[];

      expect(sortEntity(shuffledEntities)).toStrictEqual(expectedEntities);
    });
  });
});

describe("given a pool.json", () => {
  describe("if the object's keys are out of order ", () => {
    test("sort the object's keys", () => {
      const unorderedKeyPool = loadJson(
        "/data/sorter/raw/pool_01.json"
      ) as Pool[];

      const expectedPool = loadJson(
        "/data/sorter/expected/pool_01.json"
      ) as Pool[];

      expect(orderPoolKeys(unorderedKeyPool)).toStrictEqual(expectedPool);
    });
  });

  describe("if the objects are out of order", () => {
    test("objects should be sorted", () => {
      const shuffledPools = loadJson("/data/sorter/raw/pool_02.json") as Pool[];

      const expectedPools = loadJson(
        "/data/sorter/expected/pool_02.json"
      ) as Pool[];

      expect(sortPool(shuffledPools)).toStrictEqual(expectedPools);
    });
  });
});
