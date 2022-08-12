CREATE TABLE IF NOT EXISTS goods_sku
(
    id              bigserial PRIMARY KEY,
    goods_id        bigint                      NOT NULL,
    goods_attr_path jsonb                        NOT NULL,
    price           integer                     NOT NULL,
    stock           integer                     NOT NULL DEFAULT 0,
    created_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_goods
        FOREIGN KEY (goods_id)
            REFERENCES goods (id)
            ON DELETE CASCADE
);