CREATE TABLE officer
(
    id serial NOT NULL CONSTRAINT officer_pk PRIMARY KEY,
    name varchar(255) NOT NULL
);

CREATE UNIQUE index officer_id_uindex ON officer(id);

CREATE TABLE "case"
(
    id serial NOT NULL CONSTRAINT case_pk PRIMARY KEY,
    owner varchar(255),
    color varchar(50),
    brand varchar(100),
    resolved boolean DEFAULT FALSE,
    moment timestamp default now(),
    officer_id integer constraint case_officer_id_fk references officer
);

CREATE UNIQUE index case_id_uindex ON "case"(id);