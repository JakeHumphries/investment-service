CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE category AS ENUM ('ISA', 'PENSION');

CREATE TYPE customer_type AS ENUM ('RETAIL', 'EMPLOYEE');

CREATE TABLE fund (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    category category NOT NULL,
    customer_type customer_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
