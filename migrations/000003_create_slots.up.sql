CREATE TABLE IF NOT EXISTS slots (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    turf_id     UUID NOT NULL REFERENCES turfs(id) ON DELETE CASCADE,
    date        DATE NOT NULL,
    start_time  TIME NOT NULL,
    end_time    TIME NOT NULL,
    status      VARCHAR(20) DEFAULT 'available',
    created_at  TIMESTAMP DEFAULT NOW(),
    UNIQUE(turf_id, date, start_time)
);

CREATE INDEX idx_slots_turf_date ON slots(turf_id, date);
CREATE INDEX idx_slots_status ON slots(status);
