CREATE TABLE "user" (
    id           SERIAL PRIMARY KEY,
    login        VARCHAR(255) UNIQUE,
    password     VARCHAR(255)
);
