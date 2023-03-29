CREATE DATABASE h8_project_book;

CREATE TABLE books(
  "id" SERIAL PRIMARY KEY,
  "title" varchar(255) UNIQUE NOT NULL ,
  "author" varchar(255) NOT NULL ,
  "desc" text NOT NULL
);