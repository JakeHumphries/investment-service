CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO fund (id, name, category)
SELECT gen_random_uuid(), 'Cushon Equities Fund', 'RETAIL_ISA'
WHERE NOT EXISTS (SELECT 1 FROM fund WHERE name = 'Cushon Equities Fund');

INSERT INTO fund (id, name, category)
SELECT gen_random_uuid(), 'Cushon Pension Growth Fund', 'EMPLOYEE_PENSION'
WHERE NOT EXISTS (SELECT 1 FROM fund WHERE name = 'Cushon Pension Growth Fund');

INSERT INTO fund (id, name, category)
SELECT gen_random_uuid(), 'Cushon Managed ISA Fund', 'EMPLOYEE_ISA'
WHERE NOT EXISTS (SELECT 1 FROM fund WHERE name = 'Cushon Managed ISA Fund');
