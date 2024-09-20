CREATE TYPE country AS ENUM ('Azerbaijan', 'Kazakhstan', 'Kyrgyzstan', 'Turkey', 'Uzbekistan');
CREATE TYPE content AS ENUM ('text', 'post', 'photo', 'video');

CREATE TABLE IF NOT EXISTS hashtag
(
    id          UUID UNIQUE DEFAULT gen_random_uuid(),
    name        VARCHAR PRIMARY KEY,
    description VARCHAR
);

CREATE TABLE IF NOT EXISTS countries
(
    id          UUID DEFAULT gen_random_uuid(),
    city_name   VARCHAR UNIQUE,
    country     country,
    nationality VARCHAR PRIMARY KEY,
    flag        VARCHAR
);


CREATE TABLE IF NOT EXISTS posts
(
    id          UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    user_id     UUID    NOT NULL,
    nationality VARCHAR REFERENCES countries (nationality),
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
    user_id    UUID NOT NULL,
    post_id    UUID REFERENCES posts (id) ON DELETE CASCADE,
    content    TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at BIGINT                   DEFAULT 0
);

CREATE TABLE IF NOT EXISTS likes
(
    user_id    UUID NOT NULL,
    post_id    UUID REFERENCES posts (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    UNIQUE (user_id, post_id)
);


-- ------------------- table for chat ------------------------------

CREATE TABLE IF NOT EXISTS chat
(
    id         UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    user1_id   UUID NOT NULL,
    user2_id   UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at BIGINT                   DEFAULT 0
);

CREATE TABLE messages
(
    id              UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    chat_id uuid REFERENCES chat (id) ON DELETE CASCADE,
    sender_id       UUID,
    content_type    content NOT NULL,
    content         TEXT NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_aut     TIMESTAMP WITH TIME ZONE,
    is_read         BOOLEAN                  DEFAULT FALSE,
    deleted_at BIGINT                   DEFAULT 0
);
