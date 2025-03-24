-- Write your up sql migration here

CREATE TABLE maker_schedules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    maker_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    day_of_week INT NOT NULL CHECK (day_of_week BETWEEN 0 AND 6), -- 0-6 for Monday-Sunday
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT check_valid_time_range CHECK (start_time < end_time)
);

CREATE INDEX idx_maker_schedules_maker_id ON maker_schedules(maker_id);