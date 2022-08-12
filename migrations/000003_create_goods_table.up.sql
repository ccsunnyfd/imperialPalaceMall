CREATE TABLE IF NOT EXISTS goods
(
    id          bigserial PRIMARY KEY,
    spu_no      text                        NOT NULL,
    goods_name        text                        NOT NULL,
    start_price numeric(9, 2),
    category_id bigint,
    brand_id    bigint,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_category
        FOREIGN KEY (category_id)
            REFERENCES categories (id)
            ON DELETE CASCADE,
    CONSTRAINT fk_brand
        FOREIGN KEY (brand_id)
            REFERENCES brands (id)
            ON DELETE CASCADE
);