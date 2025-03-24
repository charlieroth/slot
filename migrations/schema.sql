
--
-- Postgres SQL Schema dump automatic generated by geni
--


-- EXTENSIONS 

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

-- TABLES 

CREATE TABLE maker_schedules (
 created_at timestamp without time zone  NOT NULL,
 day_of_week integer  NOT NULL,
 end_time time without time zone  NOT NULL,
 id uuid  NOT NULL,
 is_active boolean  NOT NULL,
 maker_id uuid  NOT NULL,
 start_time time without time zone  NOT NULL,
 updated_at timestamp without time zone  NOT NULL
);

CREATE TABLE schema_migrations (
 id character varying (255) NOT NULL
);

CREATE TABLE users (
 created_at timestamp without time zone  NOT NULL,
 email character varying (255) NOT NULL,
 id uuid  NOT NULL,
 name character varying (255) NOT NULL,
 phone character varying (255),
 time_zone character varying (50),
 updated_at timestamp without time zone  NOT NULL,
 user_type character varying (255) NOT NULL
);

-- CONSTRAINTS 

ALTER TABLE maker_schedules ADD CONSTRAINT maker_schedules_maker_id_fkey FOREIGN KEY (maker_id) REFERENCES users(id);

ALTER TABLE maker_schedules ADD CONSTRAINT maker_schedules_pkey PRIMARY KEY (id);

ALTER TABLE schema_migrations ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (id);

ALTER TABLE users ADD CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE users ADD CONSTRAINT users_pkey PRIMARY KEY (id);

-- INDEXES 

CREATE INDEX idx_maker_schedules_maker_id ON public.maker_schedules USING btree (maker_id)

CREATE INDEX idx_users_email ON public.users USING btree (email)

CREATE UNIQUE INDEX maker_schedules_pkey ON public.maker_schedules USING btree (id)

CREATE UNIQUE INDEX schema_migrations_pkey ON public.schema_migrations USING btree (id)

CREATE UNIQUE INDEX users_email_key ON public.users USING btree (email)

CREATE UNIQUE INDEX users_pkey ON public.users USING btree (id)

