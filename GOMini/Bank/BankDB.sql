CREATE TYPE "Currency" AS ENUM (
  'USD',
  'GBP',
  'EURO',
  'INR'
);

CREATE TABLE "Account" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" Currency NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "acc_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" bigserial PRIMARY KEY,
  "from_account" bigInt,
  "to_account" bigInt,
  "amount" bigInt NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Account" ("owner");

CREATE INDEX ON "entries" ("acc_id");

CREATE INDEX ON "transfer" ("from_account");

CREATE INDEX ON "transfer" ("to_account");

CREATE INDEX ON "transfer" ("from_account", "to_account");

CREATE INDEX ON "transfer" ("created_at");

COMMENT ON COLUMN "entries"."amount" IS 'can be positive or negative';

ALTER TABLE "entries" ADD FOREIGN KEY ("acc_id") REFERENCES "Account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_account") REFERENCES "Account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account") REFERENCES "Account" ("id");
