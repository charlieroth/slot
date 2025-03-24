-- Write your down sql migration here

DROP TRIGGER IF EXISTS check_appointment_overlap_trigger ON appointments;
DROP FUNCTION IF EXISTS check_appointment_overlap();
DROP INDEX IF EXISTS idx_appointments_time_range;
DROP INDEX IF EXISTS idx_appointments_customer_id;
DROP INDEX IF EXISTS idx_appointments_maker_id;
DROP TABLE IF EXISTS appointments;