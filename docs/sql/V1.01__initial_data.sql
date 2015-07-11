
SET SCHEMA 'shrutidb';


INSERT INTO providers (name, display_name, description) VALUES ('SYSTEM', 'SYSTEM', 'System settings');

INSERT INTO settings (key, value, provider) VALUES ('system_settings',  '{ "use_ivona": true }', (SELECT id FROM providers WHERE name='SYSTEM'));

