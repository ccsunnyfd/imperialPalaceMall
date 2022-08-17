-- 颜色、尺码之类的规格信息的值
CREATE TABLE IF NOT EXISTS goods_attr_value
(
    id          bigserial PRIMARY KEY,
    goods_id    bigint                      NOT NULL,
    attr_key_id bigint                      NOT NULL,
    attr_value  text                        NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_goods
        FOREIGN KEY (goods_id)
            REFERENCES goods (id)
            ON DELETE CASCADE,
    CONSTRAINT fk_goods_attr_key
        FOREIGN KEY (attr_key_id)
            REFERENCES goods_attr_key (id)
            ON DELETE CASCADE
);