-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE model1 (
  id              SERIAL PRIMARY KEY,
  created_at      TIMESTAMPTZ NOT NULL DEFAULT (now())
);


COMMENT ON COLUMN model1.created_at IS 'Created';


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE model1 CASCADE;
