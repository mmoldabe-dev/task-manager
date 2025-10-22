
CREATE EXTENSION IF NOT EXISTS citext;


CREATE TABLE IF NOT EXISTS Tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name_tags CITEXT UNIQUE CHECK (char_length(name_tags) BETWEEN 1 AND 15)
);
