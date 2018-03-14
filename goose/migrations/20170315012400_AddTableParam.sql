-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE param (
  id              SERIAL PRIMARY KEY,
  "type"          TEXT,
  "data"          TEXT,
  "value"         JSON
);


COMMENT ON COLUMN param."type" IS 'Config type';
COMMENT ON COLUMN param."data" IS 'Config data block';
COMMENT ON COLUMN param."value" IS 'Encoded data';


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE param CASCADE;
