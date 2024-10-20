CREATE TABLE
    IF NOT EXISTS followers (
        user_id bigint NOT NULL,
        follower_id bigint NOT NULL,
        created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (user_id, follower_id),
        FOREIGN KEY (user_id) REFERENCES users (id),
        FOREIGN KEY (follower_id) REFERENCES users (id)
    )