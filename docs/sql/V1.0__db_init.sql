
-- CREATE SCHEMA IF NOT EXISTS shrutidb;

-- SET SCHEMA 'shrutidb';

CREATE TABLE "providers" (
    "id"                SERIAL PRIMARY KEY,
    "name"              TEXT UNIQUE NOT NULL,
    "display_name"      TEXT NOT NULL,
    "description"       TEXT,
    "web_url"           TEXT,
    "icon_url"          TEXT,
    "active"            BOOL NOT NULL DEFAULT 'true',
    "voice"             TEXT NOT NULL DEFAULT 'Nicole'
);


CREATE TABLE "notifications" (
    "id"            SERIAL PRIMARY KEY,
    "title"         TEXT NOT NULL,
    "url"           TEXT,
    "key"           TEXT UNIQUE NOT NULL,
    "heard"         INTEGER DEFAULT 10,
    "provider"      INTEGER NOT NULL,
    "priority"      INTEGER DEFAULT 20,
    "action"        INTEGER DEFAULT 10,
    "created_on"    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);
ALTER TABLE "notifications" ADD CONSTRAINT notifications_fk_provider FOREIGN KEY (provider) REFERENCES providers(id);
-- heard, priority and action are supposed to be enums
-- since flyway has not accepted my PR to enable non-transactional migrations, I am going to manage these on Golang side
-- fyi, altering types in postgres require to be done without transaction

-- we will be looking up at all those notifications based on timestamp
CREATE INDEX notifications_created_on on notifications(created_on);

CREATE TABLE "settings" (
    "id"        SERIAL PRIMARY KEY,
    "key"       TEXT UNIQUE NOT NULL,
    "value"     JSONB NOT NULL,
    "provider"  INTEGER NOT NULL
);
ALTER TABLE "settings" ADD CONSTRAINT settings_fk0 FOREIGN KEY (provider) REFERENCES providers(id);


