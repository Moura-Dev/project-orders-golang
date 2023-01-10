package db

const Schema = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	login VARCHAR(255) unique NOT NULL,
	password VARCHAR(255) NOT NULL,
	active BOOLEAN default True NOT NULL
);

CREATE TABLE IF NOT EXISTS sellers (
    	id SERIAL PRIMARY KEY,
    	user_id INT NOT NULL,
    	name VARCHAR(255) NOT NULL,
    	tax_id VARCHAR(255) NOT NULL unique,
    	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	name VARCHAR(255) NOT NULL,
	cnpj VARCHAR(255) NOT NULL,
	fantasy_name VARCHAR(255) NOT NULL,
	ie VARCHAR(255),
	fone VARCHAR(255),
	fone2 VARCHAR(255),
	logo VARCHAR(255),
	company_type VARCHAR(255) NOT NULL,
	address_id INT,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS addresses (
	id SERIAL PRIMARY KEY,
	street VARCHAR(255) NOT NULL,
	number VARCHAR(255) NOT NULL,
	zip_code VARCHAR(255) ,
	city VARCHAR(255) NOT NULL,
	state VARCHAR(255) NOT NULL,
	district VARCHAR(255) ,
	reference VARCHAR(255) ,
	addressname VARCHAR(255),
	user_id INT NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);



CREATE TABLE IF NOT EXISTS orders (
	id SERIAL PRIMARY KEY,
	status VARCHAR(255) NOT NULL,
	shipping VARCHAR(10) NOT NULL,
	user_id INT NOT NULL,
	portage_id INT,
	customer_id INT NOT NULL,
	purchaser VARCHAR(255),
	observation VARCHAR(255),
	created_at TIMESTAMP default NOW(),
	updated_at TIMESTAMP default NOW(),
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
	FOREIGN KEY (portage_id) REFERENCES companies (id) ON DELETE CASCADE,
	FOREIGN KEY (customer_id) REFERENCES companies (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(255),
	cod VARCHAR(255),
	price DECIMAL(10,2) NOT NULL,
	ipi DECIMAL(10,2),
	active BOOLEAN default True NOT NULL,
	company_id INT NOT NULL,
    user_id INT NOT NULL,
	FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
	order_id INT NOT NULL,
	product_id INT NOT NULL,
	quantity INT NOT NULL,
	price DECIMAL(10,2) NOT NULL,
	discount DECIMAL(10,2) NOT NULL,
	FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
	FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);
CREATE TABLE "sessions" (
                            "id" uuid PRIMARY KEY,
                            "email" varchar NOT NULL,
                            "refresh_token" varchar NOT NULL,
                            "user_agent" varchar NOT NULL,
                            "client_ip" varchar NOT NULL,
                            "is_blocked" boolean NOT NULL DEFAULT false,
                            "expires_at" timestamptz NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            FOREIGN KEY (email) REFERENCES "users" (email) ON DELETE CASCADE
);`
