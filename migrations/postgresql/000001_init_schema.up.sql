CREATE TABLE "wallets" (
  "id" bigserial PRIMARY KEY,
  "uuid" text NOT NULL UNIQUE,
  "name" text NOT NULL,
  "personal_id" text NOT NULL UNIQUE,
  "email" text NOT NULL UNIQUE,
  "password" text NOT NULL,
  "is_a_shopkeeper" boolean NOT NULL,
  "balance" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "entries" text [],
  "transfers" text []
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "uuid" text NOT NULL UNIQUE,
  "wallet_uuid" text NOT NULL,
  "amount" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_wallet FOREIGN KEY("wallet_uuid") REFERENCES wallets("uuid") ON DELETE CASCADE
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_wallet_uuid" text NOT NULL,
  "to_wallet_uuid" text NOT NULL,
  "amount" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_from_wallet FOREIGN KEY("from_wallet_uuid") REFERENCES wallets("uuid") ON DELETE CASCADE,
  CONSTRAINT fk_to_wallet FOREIGN KEY("to_wallet_uuid") REFERENCES wallets("uuid") ON DELETE CASCADE
);