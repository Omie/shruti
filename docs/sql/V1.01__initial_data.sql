
-- SET SCHEMA 'shrutidb';


INSERT INTO providers (name, display_name, description, web_url, icon_url, active) VALUES ('SYSTEM', 'SYSTEM', 'System settings', '', '', true);

INSERT INTO settings (key, value, provider) VALUES ('system_settings',  '{ "use_ivona": true }', (SELECT id FROM providers WHERE name='SYSTEM'));

