CREATE TABLE IF NOT EXISTS addresses (
                                     id bigserial PRIMARY KEY,
                                     user_id bigint NOT NULL,
                                     user_name text NOT NULL,
                                     tel text NOT NULL,
                                     region text[] NOT NULL,
                                     detail_info text NOT NULL,
                                     post_code text NOT NULL,
                                     deleted_at timestamp(0) with time zone,
                                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                     updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
