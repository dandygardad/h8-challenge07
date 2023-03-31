CREATE DATABASE h8_project_book;

CREATE TABLE books(
  "id" SERIAL PRIMARY KEY,
  "name_book" varchar(255) NOT NULL,
  "author" varchar(255) NOT NULL,
  "created_at" timestamp NULL DEFAULT now(),
  "updated_at" timestamp NULL DEFAULT now()
);