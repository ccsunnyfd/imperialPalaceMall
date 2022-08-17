-- 颜色、尺码之类的规格信息的名称
CREATE TABLE IF NOT EXISTS goods_attr_key
(
    id         bigserial PRIMARY KEY,
    goods_id   bigint                      NOT NULL,
    attr_key   text                        NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_goods
        FOREIGN KEY (goods_id)
            REFERENCES goods (id)
            ON DELETE CASCADE
);