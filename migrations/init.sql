CREATE TABLE "Regions" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE "Interests" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE "Tags" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE "User" (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    f_name VARCHAR(255),
    l_name VARCHAR(255),
    s_name VARCHAR(255),
    balance_id INT,
    FOREIGN KEY (balance_id) REFERENCES "Balance" (id)
);

CREATE TABLE "Balance" (
    id SERIAL PRIMARY KEY,
    total_balance DECIMAL,
    reserved_balance DECIMAL,
    available_balance DECIMAL
);

CREATE TABLE "Ad" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    website_link VARCHAR(255),
    budget DECIMAL,
    audience_id INT,
    image_link VARCHAR(255),
    owner_id INT,
    FOREIGN KEY (owner_id) REFERENCES "User" (id)
    FOREIGN KEY (audience_id) REFERENCES "Audience" (id)
);

CREATE TABLE "Audience" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    owner_id INT,
    gender VARCHAR(10),
    min_age INT,
    max_age INT 
);

CREATE TABLE "AudienceRegions" (
    audience_id INT,
    region_id INT,
    FOREIGN KEY (audience_id) REFERENCES "Audience" (id),
    FOREIGN KEY (region_id) REFERENCES "Regions" (id)
);

CREATE TABLE "AudienceInterests" (
    audience_id INT,
    interest_id INT,
    FOREIGN KEY (audience_id) REFERENCES "Audience" (id),
    FOREIGN KEY (interest_id) REFERENCES "Interests" (id)
);

CREATE TABLE "AudienceTags" (
    audience_id INT,
    tag_id INT,
    FOREIGN KEY (audience_id) REFERENCES "Audience" (id),
    FOREIGN KEY (tag_id) REFERENCES "Tags" (id)
);

CREATE TABLE "UserAds" (
    user_id INT,
    ad_id INT,
    FOREIGN KEY (user_id) REFERENCES "User" (id),
    FOREIGN KEY (ad_id) REFERENCES "Ad" (id)
);
