-- Seed data for development
-- Run: psql -h localhost -U turf_user -d turf_booking -f scripts/seed.sql

-- Insert sample users (passwords are bcrypt hash of "password123")
-- You'll generate real hashes from the app; these are placeholders
INSERT INTO users (name, email, phone, password_hash, role) VALUES
  ('Rahul Sharma', 'rahul@example.com', '9876543210', '$placeholder_hash', 'player'),
  ('Priya Patel', 'priya@example.com', '9876543211', '$placeholder_hash', 'owner'),
  ('Admin User', 'admin@example.com', '9876543212', '$placeholder_hash', 'admin')
ON CONFLICT (email) DO NOTHING;

-- Insert sample turfs
INSERT INTO turfs (name, location, city, owner_id, sport_type, price_per_hour, amenities)
SELECT
  'Cricket Zone Arena',
  '123 MG Road, Near Metro Station',
  'Bangalore',
  u.id,
  'cricket',
  1500.00,
  ARRAY['floodlights', 'changing_room', 'parking', 'drinking_water']
FROM users u WHERE u.email = 'priya@example.com'
ON CONFLICT DO NOTHING;

INSERT INTO turfs (name, location, city, owner_id, sport_type, price_per_hour, amenities)
SELECT
  'Smash Box Cricket',
  '45 Koramangala 5th Block',
  'Bangalore',
  u.id,
  'cricket',
  1200.00,
  ARRAY['floodlights', 'parking']
FROM users u WHERE u.email = 'priya@example.com'
ON CONFLICT DO NOTHING;
