CREATE TABLE customers2
(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    firstname text COLLATE pg_catalog."default",
    lastname text  COLLATE pg_catalog."default",
    birthdate date ,
    gender text COLLATE pg_catalog."default",
    email text COLLATE pg_catalog."default",
    address text COLLATE pg_catalog."default"
);

INSERT INTO customers2(id, firstname, lastname, birthdate, gender, email, address) VALUES
    (gen_random_uuid(), 'Tatjana', 'Putskova', '1993-05-21', 'Female', 'email@gmail.com', 'Sütiste tee 19, Tallinn' );