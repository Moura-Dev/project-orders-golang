CREATE TABLE IF NOT EXISTS contacts (
    id SERIAL PRIMARY KEY,
    name         TEXT,
    website      TEXT,
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

CREATE TABLE IF NOT EXISTS companies
(
    id           SERIAL PRIMARY KEY,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users
(
    id              SERIAL PRIMARY KEY,
    name            TEXT,
    email           TEXT,
    login           TEXT,
    password        TEXT,
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW()

);

CREATE TABLE IF NOT EXISTS user_companies
(
    company_id INT,
    user_id INT,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE

);

CREATE TABLE IF NOT EXISTS factories
(
    id         SERIAL PRIMARY KEY,
    company_id INT,
    contact_id INT,
    name       TEXT,
    telephone  TEXT,
    email      TEXT,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES contacts (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS costumers
(
    id         SERIAL PRIMARY KEY,
    company_id INT,
    contact_id INT,
    name       TEXT,
    telephone  TEXT,
    email      TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES contacts (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS catalogs
(
    id         SERIAL PRIMARY KEY,
    company_id INT,
    factory_id INT,
    name       TEXT,
    year       TIMESTAMP,
    month      TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id) REFERENCES factories (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS items
(
    id          SERIAL PRIMARY KEY,
    company_id INT,
    catalog_id  INT,
    name        TEXT,
    code        TEXT,
    reference   TEXT,
    description TEXT,
    image_url   TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (catalog_id) REFERENCES catalogs  (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders
(
    id         SERIAL PRIMARY KEY,
    company_id INT,
    costumer_id INT,
    factory_id INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id)  REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (costumer_id) REFERENCES costumers (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id)  REFERENCES factories (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_items
(
    order_id INT,
    item_id  INT,
    company_id INT,
    quantity INT,
    FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS items_price
(
    item_id  INT,
    company_id INT,
    price DECIMAL(10, 2),
    ipi INT,
    FOREIGN KEY (item_id)    REFERENCES items (id)     ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE
);

select *
from companies;
INSERT INTO companies (created_at, updated_at) VALUES (created_at, updated_at)
RETURNING *;

select *
from user_companies;
INSERT INTO user_companies (company_id)
VALUES (1);

select *
from users;
INSERT INTO users (name, login, password, email)
VALUES ('gabriel', 'gabrieler', '123', 'gabriel@gmail')
RETURNING *;


select *
from factories;
INSERT INTO factories (company_id)
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
from orders;
INSERT INTO orders (factory_id)
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
         join companies c2 on c2.id = f.company_id
WHERE c.id = 2;


-- GetOrderInfo
SELECT o.id, c2.name AS factory_name, c.name catalog_name, i.name AS item_name, oi.quantity
FROM orders AS o
         JOIN factory f ON f.id = o.factory_id
         join companies c2 on c2.id = f.company_id
         JOIN order_item oi ON o.id = oi.order_id
         JOIN item i ON oi.item_id = i.id
         JOIN catalog c ON c.id = i.catalog_id
WHERE o.id = 1;