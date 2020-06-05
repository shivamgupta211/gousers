
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name text,
    last_name text,
    interests text,
    skills text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
