CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "email" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expire_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "order_type_id" bigint NOT NULL,
  "discount" float8 NOT NULL DEFAULT 1,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "order_types" (
  "id" bigserial PRIMARY KEY,
  "days" bigint NOT NULL,
  "price" bigint NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

COMMENT ON COLUMN "order_types"."days" IS 'must be positive';

ALTER TABLE "orders" ADD FOREIGN KEY ("user_name") REFERENCES "users" ("username");

ALTER TABLE "orders" ADD FOREIGN KEY ("order_type_id") REFERENCES "order_types" ("id");
