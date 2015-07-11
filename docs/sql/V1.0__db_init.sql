
CREATE SCHEMA IF NOT EXISTS shrutidb;

SET SCHEMA 'shrutidb';


CREATE TABLE "settings" (
    "id"        SERIAL PRIMARY KEY,
    "key"       TEXT UNIQUE NOT NULL,
    "value"     JSONB NOT NULL,
    "provider"  INTEGER NOT NULL
);


CREATE TABLE "providers" (
    "id"            SERIAL PRIMARY KEY,
    "name"          TEXT UNIQUE NOT NULL,
    "display_name"  TEXT NOT NULL,
    "description"   TEXT,
    "web_url"       TEXT,
    "icon_url"      TEXT,
    "active"        BOOL NOT NULL DEFAULT 'true'
);

CREATE TABLE "notifications" (
    "id"            SERIAL PRIMARY KEY,
    "title"         TEXT NOT NULL,
    "url"           TEXT,
    "key"           TEXT UNIQUE NOT NULL,
    "provider"      INTEGER NOT NULL,
    "created_on"    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

ALTER TABLE "settings" ADD CONSTRAINT settings_fk0 FOREIGN KEY (provider) REFERENCES providers(id);
ALTER TABLE "notifications" ADD CONSTRAINT notifications_fk_provider FOREIGN KEY (provider) REFERENCES providers(id);

