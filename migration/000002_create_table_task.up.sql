CREATE TABLE IF NOT EXISTS Tasks(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    title TEXT NOT NULL ,
    description TEXT ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    completed_at TIMESTAMP,
    priority INT NOT NULL CHECK (priority BETWEEN 1 AND 5),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    

);