import { JTDSchemaType } from "ajv/dist/jtd";

export type Account = {
  id: string;
  entity: string;
  label: string;
};

export const AccountSchema: JTDSchemaType<Account[]> = {
  elements: {
    properties: {
      id: { type: "string" },
      entity: { type: "string" },
      label: { type: "string" },
    },
    additionalProperties: false,
  },
};

export type Asset = {
  id: string;
  entity?: string;
  name: string;
  symbol: string;
  decimals: string;
  circ_supply_api?: string | undefined;
  icon?: string | undefined;
  website?: string | undefined;
  telegram?: string | undefined;
  twitter?: string | undefined;
  discord?: string | undefined;
  coinmarketcap?: string | undefined;
  coingecko?: string | undefined;
};

export const AssetSchema: JTDSchemaType<Asset[]> = {
  elements: {
    properties: {
      id: { type: "string" },
      name: { type: "string" },
      symbol: { type: "string" },
      decimals: { type: "string" },
    },
    optionalProperties: {
      entity: { type: "string" },
      circ_supply_api: { type: "string" },
      icon: { type: "string" },
      website: { type: "string" },
      telegram: { type: "string" },
      twitter: { type: "string" },
      discord: { type: "string" },
      coinmarketcap: { type: "string" },
      coingecko: { type: "string" },
    },
    additionalProperties: false,
  },
};

export type Binary = {
  id: string;
  entity: string;
  label: string;
};

export const BinarySchema: JTDSchemaType<Binary[]> = {
  elements: {
    properties: {
      id: { type: "string" },
      entity: { type: "string" },
      label: { type: "string" },
    },
    additionalProperties: false,
  },
};

export type Contract = {
  id: string;
  entity: string;
  label: string;
};

export const ContractSchema: JTDSchemaType<Contract[]> = {
  elements: {
    properties: {
      id: { type: "string" },
      entity: { type: "string" },
      label: { type: "string" },
    },
    additionalProperties: false,
  },
};

export type Pool = {
  id: string;
  asset_ids: string[];
  dex: string;
  type: string;
};

export const PoolSchema: JTDSchemaType<Pool[]> = {
  elements: {
    properties: {
      id: { type: "string" },
      asset_ids: { elements: { type: "string" } },
      dex: { type: "string" },
      type: { type: "string" },
    },
    additionalProperties: false,
  },
};

export type LabelledType = Account | Binary | Contract;
