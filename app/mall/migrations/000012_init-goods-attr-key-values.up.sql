create or replace function init_goods_attr_key_values(specs jsonb)
    returns
        boolean
AS
$$
declare
    spec     record;
    goods_id bigint;
    goods_attr_key_id bigint;
    spec_val text;
    cur_goods CURSOR FOR SELECT id
                         FROM goods;
begin
    OPEN cur_goods;

    LOOP
        FETCH cur_goods INTO goods_id;
        EXIT WHEN NOT FOUND;

        FOR spec IN select * from jsonb_each(specs)
            LOOP
                INSERT INTO goods_attr_key(goods_id, attr_key)
                VALUES (goods_id, spec.key)
                RETURNING id INTO goods_attr_key_id;

                FOR spec_val IN select * from jsonb_array_elements_text(spec.value)
                    LOOP
                        INSERT INTO goods_attr_value(goods_id, attr_key_id, attr_value)
                        VALUES (goods_id, goods_attr_key_id, spec_val);
                    END LOOP;

            END LOOP;

    END LOOP;

    CLOSE cur_goods;

    RETURN true;
end;
$$
    LANGUAGE plpgsql;

BEGIN TRANSACTION;

select *
from init_goods_attr_key_values('{
  "颜色": [
    "蓝色",
    "粉色"
  ],
  "尺码": [
    "120x150*240",
    "120x180*300",
    "120x200*360"
  ]
}'::jsonb) as tab;

-- ROLLBACK;
COMMIT;