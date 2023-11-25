CREATE TABLE "survey" (
    id SERIAL PRIMARY KEY,
    type INT DEFAULT 0,
    question TEXT DEFAULT NULL
);

CREATE TABLE "rate" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    rate INT NOT NULL,
    survey_id INT NOT NULL,
    FOREIGN KEY (survey_id) REFERENCES "survey" (id)
);