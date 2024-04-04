-- +goose Up
-- +goose StatementBegin
CREATE TABLE "files_metadata"
(
    "id"             UUID                           NOT NULL,
    "file_id"        UUID                           NOT NULL,
    "size"           BIGINT                         NOT NULL,
    "file_type"      VARCHAR(10)                    NOT NULL,
    "file_extention" VARCHAR(10) NOT NULL
);
ALTER TABLE
    "files_metadata"
    ADD PRIMARY KEY ("id");
CREATE TABLE "files"
(
    "id"          UUID                           NOT NULL,
    "file_id"     UUID                           NOT NULL,
    "owner_id" UUID NOT NULL
);
ALTER TABLE
    "files"
    ADD PRIMARY KEY ("id");
COMMENT
    ON COLUMN
    "files"."file_id" IS 'Id файла из s3 (minio) хранилища';
CREATE TABLE "tags"
(
    "id"      UUID         NOT NULL,
    "tag"     VARCHAR(255) NOT NULL,
    "file_id" UUID         NOT NULL
);
ALTER TABLE
    "tags"
    ADD PRIMARY KEY ("id");
ALTER TABLE
    "files_metadata"
    ADD CONSTRAINT "files_metadata_file_id_foreign" FOREIGN KEY ("file_id") REFERENCES "files" ("id");
ALTER TABLE
    "tags"
    ADD CONSTRAINT "tags_file_id_foreign" FOREIGN KEY ("file_id") REFERENCES "files" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop foreign key constraints
ALTER TABLE "tags"
    DROP CONSTRAINT IF EXISTS "tags_file_id_foreign";
ALTER TABLE "files_metadata"
    DROP CONSTRAINT IF EXISTS "files_metadata_file_id_foreign";

-- Drop tables
DROP TABLE IF EXISTS "tags";
DROP TABLE IF EXISTS "files_metadata";
DROP TABLE IF EXISTS "files";


-- +goose StatementEnd
