CREATE TABLE IF NOT EXISTS goods_info
(
    id          bigserial PRIMARY KEY,
    kind        smallint                    NOT NULL,
    content     text                        NOT NULL DEFAULT '',
    goods_id    bigint,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_goods
        FOREIGN KEY (goods_id)
            REFERENCES goods (id)
            ON DELETE CASCADE
);