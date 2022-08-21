CREATE TABLE IF NOT EXISTS users (
                                          id bigserial PRIMARY KEY,
                                          open_id text NOT NULL,
                                          nick_name text NOT NULL DEFAULT '',
                                          gender bigint NOT NULL DEFAULT 0,
                                          city text NOT NULL DEFAULT '',
                                          province text NOT NULL DEFAULT '',
                                          country text NOT NULL DEFAULT '',
                                          avatar_url text NOT NULL DEFAULT '',
                                          union_id text NOT NULL DEFAULT '',
                                          deleted_at timestamp(0) with time zone,
                                          created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                          updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);