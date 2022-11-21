CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS contacts
(
    id           SERIAL PRIMARY KEY NOT NULL,
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

CREATE OR REPLACE TRIGGER set_timestamp
    BEFORE UPDATE ON contacts
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS companies
(
    id         SERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY NOT NULL,
    name       TEXT,
    email      TEXT,
    login      TEXT,
    password   TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
    );

CREATE OR REPLACE TRIGGER set_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS user_companies
(
    company_id INT NOT NULL,
    user_id    INT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (company_id, user_id)
);

CREATE TABLE IF NOT EXISTS factories
(
    id         SERIAL PRIMARY KEY NOT NULL,
    company_id INT NOT NULL,
    contact_id INT,
    name       TEXT,
    telephone  TEXT,
    email      TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES contacts (id) ON DELETE CASCADE
);

CREATE OR REPLACE TRIGGER set_timestamp
    BEFORE UPDATE ON factories
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS available_factories_by_company
(
    company_id INT NOT NULL,
    factory_id INT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id) REFERENCES factories (id) ON DELETE CASCADE,
    PRIMARY KEY (company_id, factory_id)

);

CREATE TABLE IF NOT EXISTS costumers
(
    id         SERIAL PRIMARY KEY NOT NULL,
    company_id INT NOT NULL,
    contact_id INT NOT NULL,
    name       TEXT,
    telephone  TEXT,
    email      TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES contacts (id) ON DELETE CASCADE
);

CREATE OR REPLACE TRIGGER set_timestamp
    BEFORE UPDATE ON costumers
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS catalogs
(
    id         SERIAL PRIMARY KEY NOT NULL,
    company_id INT NOT NULL,
    factory_id INT NOT NULL,
    name       TEXT,
    year       TIMESTAMP,
    month      TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id) REFERENCES factories (id) ON DELETE CASCADE
);

CREATE OR REPLACE TRIGGER set_timestamp
    BEFORE UPDATE ON catalogs
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS available_catalogs_by_company
(
    company_id INT NOT NULL,
    catalog_id INT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (catalog_id) REFERENCES catalogs (id) ON DELETE CASCADE,
    PRIMARY KEY (company_id, catalog_id)
);

CREATE TABLE IF NOT EXISTS items
(
    id          SERIAL PRIMARY KEY NOT NULL,
    company_id  INT NOT NULL,
    catalog_id  INT NOT NULL,
    factory_id  INT NOT NULL,
    name        TEXT,
    code        TEXT,
    reference   TEXT,
    description TEXT,
    image_url   TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (catalog_id) REFERENCES catalogs  (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id) REFERENCES factories (id) ON DELETE CASCADE
);

CREATE OR REPLACE TRIGGER set_timestamp
    BEFORE UPDATE ON items
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS orders
(
    id          SERIAL PRIMARY KEY NOT NULL,
    company_id  INT NOT NULL,
    costumer_id INT NOT NULL,
    factory_id  INT NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (costumer_id) REFERENCES costumers (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id) REFERENCES factories (id) ON DELETE CASCADE
);

CREATE OR REPLACE TRIGGER set_timestamp
    BEFORE UPDATE ON orders
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS order_items
(
    company_id INT NOT NULL,
    order_id   INT NOT NULL,
    item_id    INT NOT NULL,
    quantity   INT,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    PRIMARY KEY (company_id, order_id, item_id)
);

CREATE TABLE IF NOT EXISTS item_prices
(
    company_id INT NOT NULL,
    item_id    INT NOT NULL,
    price      DECIMAL(10, 2) CONSTRAINT price_non_negative_value CHECK (price >= 0),
    ipi        INT CONSTRAINT ipi_non_negative_value CHECK (ipi >= 0),
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    PRIMARY KEY (company_id, item_id)
);

CREATE OR REPLACE TRIGGER set_timestamp
    BEFORE UPDATE ON item_prices
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
