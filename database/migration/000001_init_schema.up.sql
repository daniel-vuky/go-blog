-- 1 - Male
-- 2 - Female
-- 3 - Not Identify
CREATE TYPE "gender" AS ENUM (
    '1',
    '2',
    '3'
);

-- 1 - Category
-- 2 - Post
CREATE TYPE "url_rewrite_entity" AS ENUM (
    '1',
    '2'
);

CREATE TABLE "authorization_roles" (
    "role_id" serial PRIMARY KEY,
    "role_name" varchar(255) NOT NULL,
    "is_administrator" bool NOT NULL DEFAULT true,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

INSERT INTO authorization_roles (role_id, role_name, is_administrator) VALUES (1, 'Administrator', true);

CREATE TABLE "authorization_rules" (
    "rule_id" bigserial PRIMARY KEY,
    "role_id" bigint NOT NULL,
    "permission_code" varchar(128) NOT NULL,
    "is_allowed" bool NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "admin" (
    "admin_id" serial PRIMARY KEY,
    "role_id" bigint NOT NULL,
    "email" varchar(255) UNIQUE NOT NULL,
    "hashed_password" varchar NOT NULL,
    "firstname" varchar(32) NOT NULL,
    "lastname" varchar(32),
    "active" bool,
    "lock_expires" timestamptz,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00',
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "user" (
    "user_id" bigserial PRIMARY KEY,
    "email" varchar(255) UNIQUE NOT NULL,
    "firstname" varchar(32) NOT NULL,
    "lastname" varchar(32) NOT NULL,
    "subscribe" bool DEFAULT false,
    "gender" gender,
    "dob" timestamptz,
    "hashed_password" varchar NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00',
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "refresh_tokens" (
    "refresh_token_id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "refresh_token" varchar NOT NULL,
    "user_agent" varchar NOT NULL,
    "client_ip" varchar NOT NULL,
    "is_blocked" bool NOT NULL DEFAULT false,
    "expired_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "category" (
    "category_id" bigserial PRIMARY KEY,
    "parent_id" bigint NOT NULL,
    "name" varchar(255) NOT NULL,
    "url_key" text UNIQUE,
    "short_description" text,
    "description" text,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "post" (
    "post_id" bigserial PRIMARY KEY,
    "name" varchar(512) NOT NULL,
    "short_description" text,
    "description" text,
    "content" text,
    "url_key" text UNIQUE,
    "thumbnail" text,
    "author_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()',
    "updated_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "post_links" (
    "link_id" bigserial PRIMARY KEY,
    "category_id" bigint NOT NULL,
    "post_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "comment" (
    "comment_id" bigserial PRIMARY KEY,
    "post_id" bigint NOT NULL,
    "user_id" bigint NOT NULL,
    "parent_id" bigint,
    "comment" text NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()',
    "updated_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "url_rewrite" (
    "url_rewrite_id" bigserial PRIMARY KEY,
    "entity_type" url_rewrite_entity,
    "entity_id" bigint,
    "url_key" text,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE UNIQUE INDEX ON "authorization_rules" ("role_id", "permission_code");

ALTER TABLE "admin" ADD FOREIGN KEY ("role_id") REFERENCES "authorization_roles" ("role_id") ON DELETE SET NULL ON UPDATE NO ACTION;

ALTER TABLE "authorization_rules" ADD FOREIGN KEY ("role_id") REFERENCES "authorization_roles" ("role_id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "refresh_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "post_links" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("category_id") ON DELETE CASCADE;

ALTER TABLE "post_links" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id") ON DELETE CASCADE;

ALTER TABLE "comment" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "comment" ADD FOREIGN KEY ("parent_id") REFERENCES "comment" ("comment_id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "comment" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id") ON DELETE CASCADE;

CREATE OR REPLACE FUNCTION update_password_changed_at()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.hashed_password IS DISTINCT FROM OLD.hashed_password THEN NEW.password_changed_at = NOW();
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER admin_password_changed_trigger
    BEFORE UPDATE ON "admin"
    FOR EACH ROW
    EXECUTE FUNCTION update_password_changed_at();

CREATE TRIGGER user_password_changed_trigger
    BEFORE UPDATE ON "user"
    FOR EACH ROW
    EXECUTE FUNCTION update_password_changed_at();