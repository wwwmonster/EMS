CREATE TABLE public.users (
    id serial4 NOT NULL,
    email varchar(100) NOT NULL,
    "password" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    age numeric(3) NULL,
    isactive bool NULL,
    created timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    img varchar NULL,
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_pkey PRIMARY KEY (id)
);