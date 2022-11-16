CREATE TABLE IF NOT EXISTS phones
(
    id              SERIAL PRIMARY KEY,
    first_phone     VARCHAR(255),
    secondary_phone VARCHAR(255),
    created_at      TIMESTAMP default NOW(),
    updated_at      TIMESTAMP default NOW()
);

CREATE TABLE IF NOT EXISTS addresses
(
    id           SERIAL PRIMARY KEY,
    street       VARCHAR(255) NOT NULL,
    number       VARCHAR(255) NOT NULL,
    zip_code     VARCHAR(255),
    city         VARCHAR(255) NOT NULL,
    state        VARCHAR(255) NOT NULL,
    district     VARCHAR(255),
    reference    VARCHAR(255),
    address_name VARCHAR(255),
    created_at   TIMESTAMP default NOW(),
    updated_at   TIMESTAMP default NOW()
);

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    address_id INT,
    phone_id   INT,
    first_name VARCHAR(255)           NOT NULL,
    last_name  VARCHAR(255)           NOT NULL,
    email      VARCHAR(255)           NOT NULL,
    login      VARCHAR(255) unique    NOT NULL,
    password   VARCHAR(255)           NOT NULL,
    active     BOOLEAN   default True NOT NULL,
    created_at TIMESTAMP default NOW(),
    updated_at TIMESTAMP default NOW(),
    FOREIGN KEY (phone_id) REFERENCES phones (id) ON DELETE CASCADE,
    FOREIGN KEY (address_id) REFERENCES addresses (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS companies
(
    id           SERIAL PRIMARY KEY,
    user_id      INT          NOT NULL,
    address_id   INT,
    phone_id     INT,
    name         VARCHAR(255) NOT NULL,
    cnpj         VARCHAR(255) NOT NULL,
    fantasy_name VARCHAR(255) NOT NULL,
    ie           VARCHAR(255),
    logo         VARCHAR(255),
    created_at   TIMESTAMP default NOW(),
    updated_at   TIMESTAMP default NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (phone_id) REFERENCES phones (id) ON DELETE CASCADE,
    FOREIGN KEY (address_id) REFERENCES addresses (id) ON DELETE CASCADE
);

CREATE TYPE shippingType AS ENUM ('CIF', 'FOB');

CREATE TYPE status AS ENUM ('QUOTE', 'APPROVED', 'CANCELLED');

CREATE TABLE IF NOT EXISTS orders
(
    id                  SERIAL PRIMARY KEY,
    user_id             INT          NOT NULL,
    shipping_company_id INT,
    customer_id         INT          NOT NULL,
    status              status       NOT NULL,
    shipping_type       shippingType NOT NULL,
    purchaser           VARCHAR(255),
    observation         VARCHAR(255),
    created_at          TIMESTAMP default NOW(),
    updated_at          TIMESTAMP default NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES companies (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    user_id     INT                    NOT NULL,
    company_id  INT                    NOT NULL,
    name        VARCHAR(255)           NOT NULL,
    description VARCHAR(255),
    cod         VARCHAR(255),
    price       DECIMAL(10, 2)         NOT NULL,
    ipi         DECIMAL(10, 2),
    active      BOOLEAN   default True NOT NULL,
    created_at  TIMESTAMP default NOW(),
    updated_at  TIMESTAMP default NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INT            NOT NULL,
    product_id INT            NOT NULL,
    quantity   INT            NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    discount   DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);