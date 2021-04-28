CREATE DATABASE "devbook";

CREATE TABLE "usuarios" (
  "id" serial NOT NULL,
  PRIMARY KEY ("id"),
  "nome" character varying(50) NOT NULL,
  "nick" character varying(50) NOT NULL,
  "email" character varying(50) NOT NULL,
  "senha" character varying(20) NOT NULL,
  "criado_em" timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);