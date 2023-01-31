CREATE TABLE "persons" (
  "id" bigserial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "name" varchar NOT NULL,
  "personal_id" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "is_a_shopkeeper" boolean NOT NULL
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "owner" varchar NOT NULL,
  "balance" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "account_uuid" varchar NOT NULL,
  "amount" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" varchar NOT NULL,
  "to_account_id" varchar NOT NULL,
  "amount" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "wallets" (
  "id" bigserial PRIMARY KEY,
  "person_uuid" string NOT NULL,
  "account_uuid" string NOT NULL,
  "entries" string[],
  "transfers" bigserial[]
);

CREATE INDEX ON "persons" ("uuid");

CREATE INDEX ON "persons" ("personal_id");

CREATE INDEX ON "persons" ("email");

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_uuid");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "persons" ("uuid") ON DELETE CASCADE;

ALTER TABLE "entries" ADD FOREIGN KEY ("account_uuid") REFERENCES "accounts" ("uuid") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("uuid") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("uuid") ON DELETE CASCADE;

ALTER TABLE "wallets" ADD FOREIGN KEY ("person_uuid") REFERENCES "persons" ("uuid") ON DELETE CASCADE;

ALTER TABLE "wallets" ADD FOREIGN KEY ("account_uuid") REFERENCES "accounts" ("uuid") ON DELETE CASCADE;

ALTER TABLE "entries" ADD FOREIGN KEY ("uuid") REFERENCES "wallets" ("entries") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("id") REFERENCES "wallets" ("transfers") ON DELETE CASCADE;
