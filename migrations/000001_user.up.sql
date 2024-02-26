

create type user_role_enum as enum ('admin', 'teacher', 'student');

CREATE TABLE "user" (
    "id" uuid PRIMARY KEY,
    "firstname" VARCHAR(255) NOT NULL,
    "lastname" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "role" user_role_enum,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "verified" BOOLEAN NOT NULL DEFAULT FALSE
);