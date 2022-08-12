CREATE TABLE IF NOT EXISTS brands (
                                          id bigserial PRIMARY KEY,
                                          brand_name text NOT NULL,
                                          created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                          updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);