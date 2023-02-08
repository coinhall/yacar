import { Static, StringOptions, Type } from "@sinclair/typebox";
import { TypeCompiler } from "@sinclair/typebox/compiler";

// Typebox options
const integerRegEx = /^\d+$/;
const stringOptions: StringOptions<string> = { minLength: 1 };

// Raw typebox schema
export const AccountType = Type.Object({
  id: Type.String(stringOptions),
  entity: Type.String(stringOptions),
  label: Type.String(),
});

const AssetType = Type.Object({
  id: Type.String(stringOptions),
  entity: Type.Optional(Type.String(stringOptions)),
  name: Type.String(stringOptions),
  symbol: Type.String(stringOptions),
  decimals: Type.RegEx(integerRegEx),
  circ_supply_api: Type.Optional(Type.String()),
  icon: Type.Optional(Type.String()),
  website: Type.Optional(Type.String()),
  telegram: Type.Optional(Type.String()),
  twitter: Type.Optional(Type.String()),
  discord: Type.Optional(Type.String()),
  coinmarketcap: Type.Optional(Type.String()),
  coingecko: Type.Optional(Type.String()),
});

const BinaryType = Type.Object({
  id: Type.RegEx(integerRegEx),
  entity: Type.String(stringOptions),
  label: Type.String(),
});

const ContractType = Type.Object({
  id: Type.String(stringOptions),
  entity: Type.String(stringOptions),
  label: Type.String(),
});

const EntityType = Type.Object({
  entity: Type.String(stringOptions),
  website: Type.Optional(Type.String()),
  telegram: Type.Optional(Type.String()),
  twitter: Type.Optional(Type.String()),
  discord: Type.Optional(Type.String()),
  coinmarketcap: Type.Optional(Type.String()),
  coingecko: Type.Optional(Type.String()),
});

const PoolType = Type.Object({
  id: Type.String(stringOptions),
  lp_token_id: Type.String(stringOptions),
  asset_ids: Type.Tuple([
    Type.String(stringOptions),
    Type.String(stringOptions),
  ]),
  dex: Type.String(stringOptions),
  type: Type.String(stringOptions),
});

// Inferred types
export type Account = Static<typeof AccountType>;
export type Asset = Static<typeof AssetType>;
export type Binary = Static<typeof BinaryType>;
export type Contract = Static<typeof ContractType>;
export type Entity = Static<typeof EntityType>;
export type Pool = Static<typeof PoolType>;

// Compiled schemas to compare against JSON
export const AccountSchema = TypeCompiler.Compile(Type.Array(AccountType));
export const AssetSchema = TypeCompiler.Compile(Type.Array(AssetType));
export const BinarySchema = TypeCompiler.Compile(Type.Array(BinaryType));
export const ContractSchema = TypeCompiler.Compile(Type.Array(ContractType));
export const EntitySchema = TypeCompiler.Compile(Type.Array(EntityType));
export const PoolSchema = TypeCompiler.Compile(Type.Array(PoolType));
