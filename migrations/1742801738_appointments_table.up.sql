-- Write your up sql migration here

-- Appointments
CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    maker_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT, -- Prevent deleting a maker with appointments
    customer_id UUID REFERENCES users(id) ON DELETE SET NULL, -- Allow deleting a customer with appointments
    service_id UUID NOT NULL REFERENCES services(id) ON DELETE RESTRICT, -- Prevent deleting a service with appointments
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'confirmed'
        CHECK (status IN ('pending', 'confirmed', 'cancelled', 'completed', 'no_show')),
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT check_valid_time_range CHECK (start_time < end_time)
);

CREATE INDEX idx_appointments_maker_id ON appointments(maker_id);
CREATE INDEX idx_appointments_customer_id ON appointments(customer_id);
CREATE INDEX idx_appointments_time_range ON appointments(start_time, end_time);

-- Add basic constraints to ensure appointments dont' overlap
CREATE OR REPLACE FUNCTION check_appointment_overlap()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM appointments
        WHERE maker_id = NEW.maker_id
          AND id != NEW.id -- Skip the current appointment being updated
          AND status IN ('pending', 'confirmed')
          AND (
            (NEW.start_time < end_time AND NEW.end_time > start_time)
          )
    ) THEN
        RAISE EXCEPTION 'Appointment overlaps with an existing appointment';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_appointment_overlap_trigger
BEFORE INSERT OR UPDATE ON appointments
FOR EACH ROW
EXECUTE FUNCTION check_appointment_overlap();
