import { describe, expect, test } from "vitest";
import fs from "fs";

import { Account, Asset, Binary, Contract, Pool } from "../src/shared/schema";
import { sortAsset, sortLabelledJson, sortPool } from "../src/sorter/sorter";
import {
  orderAssetKeys,
  orderLabelledTypeKeys,
  orderPoolKeys,
} from "../src/sorter/order";

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

      expect(orderLabelledTypeKeys(unorderedKeyAccounts)).toStrictEqual(
        expectedAccounts
      );
    });
  });

  describe("if the objects are out of order", () => {
    test("objects should be sorted (id, entity, label)", () => {
      const shuffledAccounts = loadJson(
        "/data/sorter/raw/account_02.json"
      ) as Account[];

      const expectedAccounts = loadJson(
        "/data/sorter/expected/account_02.json"
      ) as Account[];

      expect(sortLabelledJson(shuffledAccounts)).toStrictEqual(
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
        "/data/sorter/raw/asset_01.json"
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

      expect(orderLabelledTypeKeys(unorderedKeyBinary)).toStrictEqual(
        expectedBinary
      );
    });
  });

  describe("if the objects are out of order", () => {
    test("objects should be sorted (id, entity, label)", () => {
      const shuffledBinary = loadJson(
        "/data/sorter/raw/binary_02.json"
      ) as Binary[];

      const expectedBinary = loadJson(
        "/data/sorter/expected/binary_02.json"
      ) as Binary[];

      expect(sortLabelledJson(shuffledBinary)).toStrictEqual(expectedBinary);
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

      expect(orderLabelledTypeKeys(unorderedKeyContracts)).toStrictEqual(
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

      expect(sortLabelledJson(shuffledContracts)).toStrictEqual(
        expectedContracts
      );
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
