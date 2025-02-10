CREATE TYPE fund_category AS ENUM ('EMPLOYEE_ISA', 'RETAIL_ISA', 'EMPLOYEE_PENSION');

CREATE TABLE fund (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    category fund_category NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
