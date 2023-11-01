CREATE TABLE "balance" (
    id SERIAL PRIMARY KEY,
    total_balance DECIMAL DEFAULT 0.0 NOT NULL,
    reserved_balance DECIMAL DEFAULT 0.0 NOT NULL,
    available_balance DECIMAL DEFAULT 0.0 NOT NULL
);

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    f_name VARCHAR(255) DEFAULT NULL,
    l_name VARCHAR(255) DEFAULT NULL,
    s_name VARCHAR(255) DEFAULT NULL,
    balance_id INT,
    FOREIGN KEY (balance_id) REFERENCES "balance" (id)
);

CREATE TABLE "ad" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT DEFAULT NULL,
    website_link VARCHAR(255) NOT NULL,
    budget DECIMAL DEFAULT 0.0 NOT NULL,
    target_id INT,
    image_link VARCHAR(255) NOT NULL,
    owner_id INT,
    FOREIGN KEY (owner_id) REFERENCES "user" (id),
    FOREIGN KEY (target_id) REFERENCES "target" (id)
);

CREATE TABLE "target" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    owner_id INT,
    gender VARCHAR(10) DEFAULT NULL,
    min_age INT DEFAULT 0, 
    max_age INT DEFAULT 127,
    tags TEXT DEFAULT NULL,
    keys TEXT DEFAULT NULL,
    regions TEXT DEFAULT NULL,
    interests TEXT DEFAULT NULL
);

CREATE TABLE "user_ads" (
    user_id INT,
    ad_id INT,
    FOREIGN KEY (user_id) REFERENCES "user" (id),
    FOREIGN KEY (ad_id) REFERENCES "ad" (id)
);
