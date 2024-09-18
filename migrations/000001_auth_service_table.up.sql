CREATE TABLE IF NOT EXISTS hashtag
(
    id          UUID UNIQUE DEFAULT gen_random_uuid(),
    name        VARCHAR PRIMARY KEY,
    description VARCHAR
);

CREATE TABLE IF NOT EXISTS countries
(
    id          UUID DEFAULT gen_random_uuid(),
    city_name   VARCHAR,
    country     VARCHAR,
    nationality VARCHAR PRIMARY KEY
);


CREATE TABLE IF NOT EXISTS posts
(
    id          UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    username    UUID    NOT NULL,
    nationality VARCHAR REFERENCES countries(nationality),
    location    VARCHAR NOT NULL,
    title       VARCHAR NOT NULL,
    hashtag     VARCHAR REFERENCES hashtag (name),
    content     TEXT,
    image_url   VARCHAR                  DEFAULT 'no image',
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at  BIGINT                   DEFAULT 0
);

CREATE TABLE IF NOT EXISTS comments
(
    id         UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    username   UUID NOT NULL,
    post_id    UUID REFERENCES posts (id) ON DELETE CASCADE,
    content    TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at BIGINT                   DEFAULT 0
);

CREATE TABLE IF NOT EXISTS likes
(
    user_id    UUID NOT NULL ,
    post_id    UUID REFERENCES posts (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    UNIQUE (user_id, post_id)
);
