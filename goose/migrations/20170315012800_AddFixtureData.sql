-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO param ("type", "data", "value") VALUES
  (
    'Develop.mr_robot',
    'Database.processing',
    '{"host":"localhost","port":"5432","database":"devdb","user":"mr_robot","password":"secret","schema":"public"}'
  ),(
    'Test.vpn',
    'Rabbit.log',
    '{"host":"10.0.5.42","port":"5671","virtualhost":"/","user":"guest","password":"guest"}'
  );

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM param WHERE "type" = 'Develop.mr_robot' AND "data" = 'Database.processing';
DELETE FROM param WHERE "type" = 'Test.vpn' AND "data" = 'Rabbit.log';