create or replace function batch_insert_goods(gg text[]) returns
    boolean AS
$$
declare
    brand_id bigint;
    spu_no text;
    idx integer;
    goods_name text;
    start_price numeric(9, 2);
    category_id bigint;
begin
    for i in 1..60 LOOP
        SELECT id into brand_id FROM brands order by random() limit 1;
        spu_no = format('%s%s%s', brand_id, extract ( year from CURRENT_TIMESTAMP ), gen_random_uuid ());
        idx = floor(array_length(gg, 1) * random()) + 1;
        goods_name = format('%s%s', gg[idx], floor(random()*1000));
        start_price = 1000 * random();
        SELECT id into category_id FROM categories order by random() limit 1;

        INSERT INTO goods (spu_no, goods_name, start_price, category_id, brand_id) VALUES (spu_no, goods_name, start_price, category_id, brand_id);
    end loop;
    return true;
end;
$$
    LANGUAGE plpgsql;

select * from batch_insert_goods(array['夏凉被', '冬暖被', '秋躺椅', '春旗袍']) as tab;