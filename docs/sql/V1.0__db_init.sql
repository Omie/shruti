
CREATE SCHEMA IF NOT EXISTS contentdb;

SET SCHEMA 'shrutidb';


CREATE TABLE "settings" (
    "id"    SERIAL PRIMARY KEY,
    "key"   TEXT UNIQUE NOT NULL,
    "value" jsonb NOT NULL
);


CREATE TABLE "plugins" (
    "id"            SERIAL PRIMARY KEY,
    "name"          TEXT UNIQUE NOT NULL,
    "display_name"  TEXT NOT NULL,
    "description"   TEXT,
    "web_url"       TEXT NOT NULL,
    "icon_url"      TEXT,
    "active"        BOOL NOT NULL DEFAULT 'true'
);

CREATE TABLE "notifications" (
    "id"        SERIAL PRIMARY KEY,
    "title"     TEXT NOT NULL,
    "url"       TEXT,
    "key"       TEXT UNIQUE NOT NULL,
    "provider"  INTEGER NOT NULL,
    "created_on" timestamp without time zone NOT NULL DEFAULT now()
);

ALTER TABLE "notifications" ADD CONSTRAINT notifications_fk_provider FOREIGN KEY (provider) REFERENCES plugins(id);

