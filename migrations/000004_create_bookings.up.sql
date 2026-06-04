CREATE TABLE IF NOT EXISTS bookings (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slot_id         UUID NOT NULL REFERENCES slots(id) ON DELETE CASCADE,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status          VARCHAR(20) DEFAULT 'pending',
    payment_status  VARCHAR(20) DEFAULT 'unpaid',
    amount          DECIMAL(10,2) NOT NULL,
    booked_at       TIMESTAMP DEFAULT NOW(),
    cancelled_at    TIMESTAMP,
    UNIQUE(slot_id)
);

CREATE INDEX idx_bookings_user ON bookings(user_id);
CREATE INDEX idx_bookings_status ON bookings(status);
CREATE INDEX idx_bookings_slot ON bookings(slot_id);
