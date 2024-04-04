-- +goose Up
-- +goose StatementBegin
CREATE TABLE "file_album"
(
    "id"       UUID NOT NULL,
    "file_id"  UUID NOT NULL,
    "album_id" UUID NOT NULL
);
ALTER TABLE
    "file_album"
    ADD PRIMARY KEY ("id");
CREATE TABLE "members"
(
    "user_id"  UUID        NOT NULL,
    "album_id" UUID        NOT NULL,
    "access"   VARCHAR(10) NOT NULL
);
CREATE TABLE "album"
(
    "id"         UUID                           NOT NULL,
    "title"      VARCHAR(50)                    NOT NULL,
    "owner_id"   UUID                           NOT NULL,
    "preview_id" UUID
);
ALTER TABLE
    "album"
    ADD PRIMARY KEY ("id");
COMMENT
    ON COLUMN
    "album"."owner_id" IS 'Id создателя альбома';
COMMENT
    ON COLUMN
    "album"."preview_id" IS 'Id фото или видео, которое будет использовано в пред просмотре альбома';
ALTER TABLE
    "file_album"
    ADD CONSTRAINT "file_album_album_id_foreign" FOREIGN KEY ("album_id") REFERENCES "album" ("id");
ALTER TABLE
    "members"
    ADD CONSTRAINT "members_album_id_foreign" FOREIGN KEY ("album_id") REFERENCES "album" ("id");
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
-- Удаление внешних ключей
ALTER TABLE file_album
    DROP CONSTRAINT file_album_album_id_foreign;
ALTER TABLE members
    DROP CONSTRAINT members_album_id_foreign;

-- Удаление таблицы album
DROP TABLE IF EXISTS album;

-- Удаление таблицы members
DROP TABLE IF EXISTS members;

-- Удаление таблицы file_album
DROP TABLE IF EXISTS file_album;

-- +goose StatementEnd
