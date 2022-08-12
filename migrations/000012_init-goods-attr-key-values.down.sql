TRUNCATE goods_attr_value CASCADE;
TRUNCATE goods_attr_key CASCADE;
DROP FUNCTION IF EXISTS init_goods_attr_key_values(specs jsonb) CASCADE;