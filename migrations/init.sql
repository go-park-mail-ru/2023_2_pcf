CREATE TABLE "user" (
    id           SERIAL PRIMARY KEY,
    login        VARCHAR(255) UNIQUE,
    password     VARCHAR(255),
    f_name       VARCHAR(255),
    l_name       VARCHAR(255)
);

CREATE TABLE "ad" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    sector VARCHAR(255),
    owner_id INT,
    FOREIGN KEY (owner_id) REFERENCES "user" (id)
);