CREATE TABLE "user" (
    "id" uuid PRIMARY KEY,
    "user_name" varchar NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "phone" varchar NOT NULL,
    "created_at" timestamp NOT NULL,
    "modified_at" timestamp NOT NULL
)