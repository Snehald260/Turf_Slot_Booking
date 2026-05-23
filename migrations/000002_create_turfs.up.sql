CREATE TABLE IF NOT EXISTS turfs (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(200) NOT NULL,
    location        TEXT NOT NULL,
    city            VARCHAR(100) NOT NULL,
    owner_id        UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    sport_type      VARCHAR(50) DEFAULT 'cricket',
    price_per_hour  DECIMAL(10,2) NOT NULL,
    amenities       TEXT[],
    is_active       BOOLEAN DEFAULT true,
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_turfs_city ON turfs(city);
CREATE INDEX idx_turfs_owner ON turfs(owner_id);
CREATE INDEX idx_turfs_sport ON turfs(sport_type);
CREATE INDEX idx_turfs_active ON turfs(is_active);
