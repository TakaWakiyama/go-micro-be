-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE "tasks" (
    "id" varchar(255) NOT NULL,
    "user_id" varchar(255) NOT NULL,
    "parent_id" varchar(255),
    "title" varchar(255) NOT NULL,
    "is_template" boolean NOT NULL DEFAULT false,
    "detail" text NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMPTZ,
    PRIMARY KEY(id)
);
CREATE INDEX "task_parent_idx" ON "tasks" ("parent_id");

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
