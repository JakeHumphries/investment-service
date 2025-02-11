CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO fund (id, name, category, customer_type)
SELECT gen_random_uuid(), 'Cushon Equities Fund', 'ISA', 'RETAIL'
WHERE NOT EXISTS (SELECT 1 FROM fund WHERE name = 'Cushon Equities Fund');

INSERT INTO fund (id, name, category, customer_type)
SELECT gen_random_uuid(), 'Cushon Pension Growth Fund', 'PENSION', 'EMPLOYEE'
WHERE NOT EXISTS (SELECT 1 FROM fund WHERE name = 'Cushon Pension Growth Fund');

INSERT INTO fund (id, name, category, customer_type)
SELECT gen_random_uuid(), 'Cushon Managed ISA Fund', 'ISA', 'EMPLOYEE'
WHERE NOT EXISTS (SELECT 1 FROM fund WHERE name = 'Cushon Managed ISA Fund');
