create or replace function batch_insert_categories(cc text[]) returns
    boolean AS
$$
declare
    c text;
begin
    foreach c in array cc LOOP
            INSERT INTO categories (category_name) VALUES (c);
    end loop;
    return true;
end;
$$
    LANGUAGE plpgsql;

select * from batch_insert_categories(array['生活方式', '文创好物', '创意玩具', '微瑕特卖', '宫廷服饰', '甄选珠宝', '宫廷美玉']) as tab;