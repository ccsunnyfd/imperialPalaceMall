create or replace function batch_insert_brands(cc text[]) returns
    boolean AS
$$
declare
    c text;
begin
    foreach c in array cc LOOP
            INSERT INTO brands (brand_name) VALUES (c);
        end loop;
    return true;
end;
$$
    LANGUAGE plpgsql;

select * from batch_insert_brands(array['故宫文化', '朝阳国贸CBD', '西单购物广场']) as tab;