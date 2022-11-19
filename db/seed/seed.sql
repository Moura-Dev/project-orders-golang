CREATE TABLE IF NOT EXISTS company
(
    id           SERIAL PRIMARY KEY,
    name         TEXT,
    site         TEXT,
    cnpj_cpf     TEXT,
    address      TEXT,
    telephone    TEXT,
    logo_url     TEXT,
    fantasy_name TEXT,
    ie           TEXT,
    email        TEXT,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_company
(
    id         SERIAL PRIMARY KEY,
    bill       TEXT,
    company_id INT,
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "user"
(
    id              SERIAL PRIMARY KEY,
    name            TEXT,
    login           TEXT,
    password        TEXT,
    email           TEXT,
    user_company_id INT,
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_company_id) REFERENCES user_company (id) ON DELETE CASCADE

);

CREATE TABLE IF NOT EXISTS factory
(
    id         SERIAL PRIMARY KEY,
    company_id INT,
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS catalog
(
    id         SERIAL PRIMARY KEY,
    factory_id INT,
    name       TEXT,
    year       TIMESTAMP,
    month      TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (factory_id) REFERENCES factory (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS item
(
    id          SERIAL PRIMARY KEY,
    catalog_id  INT,
    name        TEXT,
    code        TEXT,
    reference   TEXT,
    description TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (catalog_id) REFERENCES catalog (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS costumer
(
    id         SERIAL PRIMARY KEY,
    company_id INT,
    name       TEXT,
    telephone  TEXT,
    email      TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "order"
(
    id         SERIAL PRIMARY KEY,
    factory_id INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (factory_id) REFERENCES factory (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_item
(
    order_id INT,
    item_id  INT,
    quantity INT,
    FOREIGN KEY (item_id) REFERENCES item (id) ON DELETE CASCADE,
    FOREIGN KEY (order_id) REFERENCES "order" (id) ON DELETE CASCADE
);

select *
from company;
INSERT INTO company (name, site, cnpj_cpf, address, telephone, logo_url, fantasy_name, ie, email)
VALUES ('Pride', 'pride.com', '123', 'rua 1', '31234-1233', 'pride.png', 'pride company', '123',
        'pride@gmail')
RETURNING *;

select *
from user_company;
INSERT INTO user_company (bill, company_id)
VALUES ('R$1000', 1);

select *
from "user";
INSERT INTO "user" (name, login, password, email, user_company_id)
VALUES ('gabriel', 'gabrieler', '123', 'gabriel@gmail', 1)
RETURNING *;


select *
from factory;
INSERT INTO factory (company_id)
VALUES (1);

select *
from catalog;
INSERT INTO catalog (factory_id, name, year, month)
VALUES (1, 'tubos', NOW(), NOW())
RETURNING *;
INSERT INTO catalog (factory_id, name, year, month)
VALUES (1, 'GOD', NOW(), NOW())
RETURNING *;
INSERT INTO catalog (factory_id, name, year, month)
VALUES (1, 'ATE', NOW(), NOW())
RETURNING *;

select *
from item;
INSERT INTO item (catalog_id, name, code, reference, description)
VALUES (2, 'X', 'L1', 'LX1', 'É L')
RETURNING *;

INSERT INTO item (catalog_id, name, code, reference, description)
VALUES (3, 'Y', 'B1', 'BX1', 'O B')
RETURNING *;

INSERT INTO item (catalog_id, name, code, reference, description)
VALUES (2, 'w', 'C1', 'CX1', 'P L')
RETURNING *;

select *
from "order";
INSERT INTO "order" (factory_id)
VALUES (1)
RETURNING *;

select *
from order_item;
INSERT INTO order_item (order_id, item_id, quantity)
VALUES (1, 1, 10)
RETURNING *;
INSERT INTO order_item (order_id, item_id, quantity)
VALUES (1, 2, 10)
RETURNING *;
INSERT INTO order_item (order_id, item_id, quantity)
VALUES (1, 3, 10)
RETURNING *;

-- GetFactoryItems
select i.id, i.name as item_name, c.name as catalog_name, c2.name company_name
from item i
         join catalog c on c.id = i.catalog_id
         join factory f on f.id = c.factory_id
         join company c2 on c2.id = f.company_id
WHERE c.id = 2;


-- GetOrderInfo
SELECT o.id, c2.name AS factory_name, c.name catalog_name, i.name AS item_name, oi.quantity
FROM "order" AS o
         JOIN factory f ON f.id = o.factory_id
         join company c2 on c2.id = f.company_id
         JOIN order_item oi ON o.id = oi.order_id
         JOIN item i ON oi.item_id = i.id
         JOIN catalog c ON c.id = i.catalog_id
WHERE o.id = 1;