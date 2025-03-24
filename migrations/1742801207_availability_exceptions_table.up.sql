-- Write your up sql migration here

-- Availability exceptions (holidays, vacations, special hours)
CREATE TABLE availability_exceptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    maker_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    exception_date DATE NOT NULL,
    start_time TIME, -- NULL means unavailable all day
    end_time TIME, -- NULL means unavailable all day
    is_available BOOLEAN DEFAULT false, -- true=extra availability, false=unavailable
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT check_valid_time_range CHECK (
        (start_time IS NULL AND end_time IS NULL) OR
        (start_time IS NOT NULL AND end_time IS NOT NULL AND start_time < end_time)
    )
);

CREATE INDEX idx_availability_exceptions_maker_id ON availability_exceptions(maker_id);
CREATE INDEX idx_availability_exceptions_date ON availability_exceptions(exception_date);