-- 初始化sku表
create or replace function init_goods_sku()
    returns
        boolean
AS
$$
declare
    goods_attr_path      bigint[];
    goods_id             bigint;
    goods_attr_value_id1 bigint;
    goods_attr_value_id2 bigint;
    price                integer;
    stock                integer;
    cur_goods_attr_values CURSOR FOR SELECT a.goods_id, a.id, b.id
                                     FROM goods_attr_value as a
                                              left join goods_attr_value as b on a.attr_key_id != b.attr_key_id
                                         and a.goods_id = b.goods_id
                                         and a.attr_key_id < b.attr_key_id;

begin
    OPEN cur_goods_attr_values;

    LOOP
        FETCH cur_goods_attr_values INTO goods_id, goods_attr_value_id1, goods_attr_value_id2;
        EXIT WHEN NOT FOUND;

        if goods_attr_value_id2 IS NOT NULL then
            goods_attr_path = ARRAY [goods_attr_value_id1, goods_attr_value_id2];
            price = floor(random() * 10000 + 100);
            stock := floor(random() * 100 + 1);

            insert into goods_sku(goods_id, goods_attr_path, price, stock) values (goods_id, goods_attr_path, price, stock);
        end if;

    END LOOP;

    CLOSE cur_goods_attr_values;

    RETURN true;
end;
$$
    LANGUAGE plpgsql;

BEGIN TRANSACTION;

select *
from init_goods_sku() as tab;

-- ROLLBACK;
COMMIT;