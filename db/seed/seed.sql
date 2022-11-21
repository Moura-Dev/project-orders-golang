INSERT INTO companies (created_at, updated_at) VALUES (NOW(), NOW()) RETURNING *; -- 1 Representante
INSERT INTO companies (created_at, updated_at) VALUES (NOW(), NOW()) RETURNING *; -- 2 Representante
INSERT INTO companies (created_at, updated_at) VALUES (NOW(), NOW()) RETURNING *; -- 3 Representante
INSERT INTO companies (created_at, updated_at) VALUES (NOW(), NOW()) RETURNING *; -- 4 Representante

INSERT INTO factories (company_id, name, telephone, email) VALUES (1, 'Moureili', '30758452', 'moureli@gmail.com') RETURNING *;
INSERT INTO factories (company_id, name, telephone, email) VALUES (2, 'Nike', '9998888', 'nike@company.com') RETURNING *;
INSERT INTO factories (company_id, name, telephone, email) VALUES (3, 'Razer', '6663333', 'razer@mtcaro.com') RETURNING *;
INSERT INTO factories (company_id, name, telephone, email) VALUES (4, 'Bulb', '6663333', 'bulb@mtcaro.com') RETURNING *;
----------------------------------------------------------------------
----------------------------------------------------------------------
-- Factories criados por um representante e que são disponíveis para um outro representante.
INSERT INTO available_factories_by_company (factory_id, company_id) VALUES (1, 3) RETURNING *;
INSERT INTO available_factories_by_company (factory_id, company_id) VALUES (2, 3) RETURNING *;
INSERT INTO available_factories_by_company (factory_id, company_id) VALUES (3, 4) RETURNING *;

DELETE FROM available_factories_by_company WHERE company_id = 3 AND factory_id = 1;
DELETE FROM available_factories_by_company WHERE company_id = 3 AND factory_id = 2;
DELETE FROM available_factories_by_company WHERE company_id = 3 AND factory_id = 4;

select * from available_factories_by_company;
SELECT * from factories;

SELECT f.id as factory_id, f.name as factory_name  FROM factories f
WHERE f.company_id IN (
    SELECT factory_id FROM available_factories_by_company afbc
    WHERE afbc.company_id = 3
) OR f.company_id = 3;

INSERT INTO catalogs (factory_id, company_id, name, year, month) VALUES (1, 1, 'tubos', NOW(), NOW()) RETURNING *;
INSERT INTO catalogs (factory_id, company_id, name, year, month) VALUES (2, 2, 'tênis', NOW(), NOW()) RETURNING *;
INSERT INTO catalogs (factory_id, company_id, name, year, month) VALUES (3, 3, 'teclados', NOW(), NOW()) RETURNING *;
INSERT INTO catalogs (factory_id, company_id, name, year, month) VALUES (4, 4, 'lampadas', NOW(), NOW()) RETURNING *;
----------------------------------------------------------------------
----------------------------------------------------------------------
-- Catálogos criados por um representante e que são disponíveis para um outro representante.
INSERT INTO available_catalogs_by_company (catalog_id, company_id) VALUES (1, 3) RETURNING *;
INSERT INTO available_catalogs_by_company (catalog_id, company_id) VALUES (2, 3) RETURNING *;
INSERT INTO available_catalogs_by_company (catalog_id, company_id) VALUES (4, 3) RETURNING *;

DELETE FROM available_catalogs_by_company WHERE company_id = 3 AND catalog_id = 1;
DELETE FROM available_catalogs_by_company WHERE company_id = 3 AND catalog_id = 2;
DELETE FROM available_catalogs_by_company WHERE company_id = 3 AND catalog_id = 4;

select * from available_catalogs_by_company;
select * from catalogs;

SELECT f.id             AS catalog_id,
       c.name           AS catalog_name,
       f.name           AS factory_name,
       f.company_id     AS f_company_id,
       c.company_id     AS c_company_id
FROM catalogs c
    JOIN factories f ON c.factory_id = f.id
WHERE c.company_id IN (
        SELECT a.catalog_id FROM available_catalogs_by_company a WHERE a.company_id = 3
    )
OR c.company_id = 3;

INSERT INTO items (company_id, factory_id, catalog_id, name, code, reference, description) VALUES (4, 3, 3, 'Naga', 'N1RGB', 'NRGBR', 'Mouse Gamer') RETURNING *;
INSERT INTO items (company_id, factory_id, catalog_id, name, code, reference, description) VALUES (4, 3, 3, 'Deather', 'D1RGB', 'DABR', 'Mouse Gamer') RETURNING *;
INSERT INTO items (company_id, factory_id, catalog_id, name, code, reference, description) VALUES (4, 3, 3, 'Monitor 24" 240hz', 'M24240M', 'M24BR', 'Monitor Gamer') RETURNING *;
----------------------------------------------------------------------
----------------------------------------------------------------------
-- Existem 3 items criados pelo representante 3 e que estão visíveis para o representante 4
-- pois esse tem o acesso ao factory 3 criadora dos items
INSERT INTO available_catalogs_by_company (catalog_id, company_id) VALUES (4, 3) RETURNING *;
select * from items where company_id = 4;
select * from available_factories_by_company;

SELECT i.id,
       f.id             AS catalog_id,
       i.name           AS catalog_name,
       f.name           AS factory_name,
       f.company_id     AS f_company_id,
       i.company_id     AS c_company_id
FROM items i
         JOIN factories f ON i.factory_id = f.id
WHERE i.company_id IN (
    SELECT a.factory_id FROM available_factories_by_company a WHERE a.company_id = 3
)
   OR i.company_id = 4;
----------------------------------------------------------------------
----------------------------------------------------------------------
-- Rascunho para testar outros usuários
INSERT INTO items (company_id, factory_id, catalog_id, name, code, reference, description) VALUES (3, 3, 3, 'Naga', 'N1RGB', 'NRGBR', 'Mouse Gamer') RETURNING *;
INSERT INTO items (company_id, factory_id, catalog_id, name, code, reference, description) VALUES (3, 3, 3, 'Deather', 'D1RGB', 'DABR', 'Mouse Gamer') RETURNING *;
INSERT INTO items (company_id, factory_id, catalog_id, name, code, reference, description) VALUES (3, 3, 3, 'Monitor 24" 240hz', 'M24240M', 'M24BR', 'Monitor Gamer') RETURNING *;
INSERT INTO item_prices (company_id, item_id, price, ipi) VALUES (3, 10, 1, 5) RETURNING *;
INSERT INTO item_prices (company_id, item_id, price, ipi) VALUES (3, 11, 233.99, 5) RETURNING *;
INSERT INTO item_prices (company_id, item_id, price, ipi) VALUES (3, 12, 9653, 5) RETURNING *;
INSERT INTO costumers (company_id, name, telephone, email) VALUES (3, 'Hiperer LTDA', '333333', 'hiperer@hip.com') RETURNING *;
INSERT INTO orders (company_id, costumer_id, factory_id) VALUES (3, 2, 3) RETURNING *;
INSERT INTO order_items (company_id, order_id, item_id, quantity) VALUES (3, 2, 10, 10) RETURNING *;
INSERT INTO order_items (company_id, order_id, item_id, quantity) VALUES (3, 2, 11, 5) RETURNING *;
INSERT INTO order_items (company_id, order_id, item_id, quantity) VALUES (3, 2, 12, 3) RETURNING *;
----------------------------------------------------------------------
----------------------------------------------------------------------
-- O cliente Ajex foi criado pelo representante 4 que possuí visibilidade dos items da factory 3 'Razer'
INSERT INTO costumers (company_id, name, telephone, email) VALUES (4, 'Ajex LTDA', '12344312', 'ajex@gg.com') RETURNING *;
select * from costumers;

INSERT INTO orders (company_id, costumer_id, factory_id) VALUES (4, 1, 3) RETURNING *;
select * from orders;

INSERT INTO order_items (company_id, order_id, item_id, quantity) VALUES (4, 1, 7, 10) RETURNING *;
INSERT INTO order_items (company_id, order_id, item_id, quantity) VALUES (4, 1, 8, 5) RETURNING *;
INSERT INTO order_items (company_id, order_id, item_id, quantity) VALUES (4, 1, 9, 3) RETURNING *;

select * from order_items;

SELECT DISTINCT
       i.id,
       c2.id  AS company,
       i.name AS product,
       c.name AS catalog,
       f.name AS factory,
       ip.price,
       oi.quantity,
       ip.ipi
FROM order_items oi
    JOIN item_prices ip ON  ip.item_id    = oi.item_id
    JOIN items i        ON  i.id          = ip.item_id
    JOIN catalogs c     ON  c.id          = i.catalog_id
    JOIN factories f    ON  f.id          = c.factory_id
    JOIN companies c2   ON  oi.company_id = c2.id
WHERE oi.company_id = 4 AND oi.order_id = 1;
