CREATE TABLE
    "users" (
        "id" SERIAL PRIMARY KEY,
        "username" varchar NOT NULL,
        "email" varchar NOT NULL,
        "password" varchar NOT NULL
    )