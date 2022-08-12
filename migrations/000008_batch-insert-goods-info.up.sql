create or replace function batch_insert_goods_info(images text[])
    returns
        boolean
AS
$$
declare
    img                   text;
    str_goods_name        text;
    goods_id              bigint;
    content               text;
    COVER_IMAGE_KIND      smallint = 0;
    CONTENT_DESCRIBE_KIND smallint = 1;
    cur_goods CURSOR FOR SELECT id, goods_name
                         FROM goods;
begin
    OPEN cur_goods;

    LOOP
        FETCH cur_goods INTO goods_id, str_goods_name;
        EXIT WHEN NOT FOUND;

        foreach img in array images
            LOOP
                content = img;
                INSERT INTO goods_info(goods_id, kind, content) VALUES (goods_id, COVER_IMAGE_KIND, content);
            END LOOP;

        content = format('%s%s', str_goods_name, '商品描述文本');
        INSERT INTO goods_info(goods_id, kind, content) VALUES (goods_id, CONTENT_DESCRIBE_KIND, content);
    END LOOP;

    CLOSE cur_goods;

    RETURN true;
end;
$$
    LANGUAGE plpgsql;

select *
from batch_insert_goods_info(array ['https://cloud-1252822131.cos.ap-beijing.myqcloud.com/img/20200820221945.png', 'https://cloud-1252822131.cos.ap-beijing.myqcloud.com/img/20200820222013.png', 'https://cloud-1252822131.cos.ap-beijing.myqcloud.com/img/20200820222033.png']) as tab;