-- Write your down sql migration here

DROP INDEX idx_availability_exceptions_date;
DROP INDEX idx_availability_exceptions_maker_id;
DROP TABLE availability_exceptions;