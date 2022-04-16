-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" varchar(255) NOT NULL,
    "name" varchar(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMPTZ,
    PRIMARY KEY(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users"
-- +goose StatementEnd
