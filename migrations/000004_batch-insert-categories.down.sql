DELETE FROM categories WHERE category_name IN ('生活方式', '文创好物', '创意玩具', '微瑕特卖', '宫廷服饰', '甄选珠宝', '宫廷美玉');
DROP FUNCTION IF EXISTS batch_insert_categories(text[]) CASCADE;