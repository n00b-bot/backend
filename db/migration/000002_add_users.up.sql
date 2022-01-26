CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_passwd" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT 'now()',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");