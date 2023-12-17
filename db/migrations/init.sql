CREATE TABLE "balance" (
    id SERIAL PRIMARY KEY,
    total_balance DECIMAL DEFAULT 0.0 NOT NULL,
    reserved_balance DECIMAL DEFAULT 0.0 NOT NULL,
    available_balance DECIMAL DEFAULT 0.0 NOT NULL
);

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    f_name TEXT DEFAULT NULL,
    l_name TEXT DEFAULT NULL,
    s_name TEXT DEFAULT 'company',
    avatar TEXT DEFAULT 'default.jpg',
    balance_id INT,
    FOREIGN KEY (balance_id) REFERENCES "balance" (id)
);

CREATE TABLE "target" (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    owner_id INT,
    gender TEXT DEFAULT NULL,
    min_age INT DEFAULT 0, 
    max_age INT DEFAULT 127,
    tags TEXT DEFAULT NULL,
    keys TEXT DEFAULT NULL,
    regions TEXT DEFAULT NULL,
    interests TEXT DEFAULT NULL,
    FOREIGN KEY (owner_id) REFERENCES "user" (id)
);

CREATE TABLE "ad" (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT DEFAULT NULL,
    website_link TEXT NOT NULL,
    budget DECIMAL DEFAULT 0.0 NOT NULL,
    click_cost DECIMAL DEFAULT 0.0 NOT NULL,
    target_id INT,
    image_link TEXT NOT NULL,
    owner_id INT,
    FOREIGN KEY (owner_id) REFERENCES "user" (id),
    FOREIGN KEY (target_id) REFERENCES "target" (id)
);


CREATE TABLE "pad" (
    id SERIAL PRIMARY KEY,
    clicks INT DEFAULT 0,
    balance DECIMAL DEFAULT 0.0,
    name TEXT NOT NULL,
    owner_id INT,
    target_id INT,
    website_link TEXT NOT NULL,
    price DECIMAL DEFAULT 0.0 NOT NULL,
    description TEXT DEFAULT NULL,
    FOREIGN KEY (owner_id) REFERENCES "user" (id),
    FOREIGN KEY (target_id) REFERENCES "target" (id)
);
