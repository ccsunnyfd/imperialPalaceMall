CREATE TABLE IF NOT EXISTS categories (
    id bigserial PRIMARY KEY,
    category_name text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);