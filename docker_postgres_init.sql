CREATE TABLE customers
(
    id        uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    firstname text COLLATE pg_catalog."default",
    lastname  text COLLATE pg_catalog."default",
    birthdate date,
    gender    text COLLATE pg_catalog."default",
    email     text COLLATE pg_catalog."default",
    address   text COLLATE pg_catalog."default"
);

INSERT INTO customers(id, firstname, lastname, birthdate, gender, email, address)
VALUES ('346f6f5b-52a8-48e7-bbb6-99d0e368ed72', 'testFirstname', 'testLastname', '2000-01-01',
        'MALE', 'email@gmail.com', 'testAddress'),
       (gen_random_uuid(), 'Tatjana', 'Putskova', '1993-05-21', 'FEMALE', 'email@gmail.com',
        '19 Sütiste tee, Tallinn, Estonia'),
       (gen_random_uuid(), 'Tatjana', 'Larina', '2001-05-21', 'FEMALE', 'email@gmail.com',
        'Pskovskaja gubernja, Russian Empire'),
       (gen_random_uuid(), 'Homer ', 'Simpson', '1956-03-12', 'MALE', 'email@gmail.com',
        '742 Evergreen Terrace, Springfield, United States'),
       (gen_random_uuid(), 'Bart  ', 'Simpson', '1981-02-23', 'MALE', 'email@gmail.com',
        '742 Evergreen Terrace, Springfield'),
       (gen_random_uuid(), 'Lisa  ', 'Simpson', '1981-05-09', 'FEMALE', 'email@gmail.com',
        '742 Evergreen Terrace, Springfield'),
       (gen_random_uuid(), 'Maggie  ', 'Simpson', '1987-04-19', 'FEMALE', 'email@gmail.com',
        '742 Evergreen Terrace, Springfield'),
       (gen_random_uuid(), 'Abraham  ', 'Simpson II', '2003-08-10', 'MALE', 'email@gmail.com',
        '742 Evergreen Terrace, Springfield'),
       (gen_random_uuid(), 'Ned  ', 'Flanders', '1962-05-21', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Charles', 'Montgomery Burns', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Krusty ', 'Clown', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Seymour', 'Skinner', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Milhouse ', 'Houten', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Edna  ', 'Krabappel', '1963-11-22', 'FEMALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Apu  ', 'Nahasapeemapetilon', '1993-05-21', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Eleanor   ', 'Abernathy', '1962-05-21', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Akira', 'Akira', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Ms. Albright ', 'Albright', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Woody ', 'Skinner', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Tattoo  ', 'Annie', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Aristotle  ', 'Amadopolis', '1963-11-22', 'FEMALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'State ', 'Comptroller Atkins', '1993-05-21', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Mary ', 'Bailey', '1962-05-21', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Jasper ', 'Beardly', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Benjamin  ', 'nerd', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Bill ', 'KBBL DJ', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Billy ', 'actor', '1963-11-22', 'MALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Birch ', 'Barlow', '1963-11-22', 'FEMALE', 'email@gmail.com', ''),
       (gen_random_uuid(), 'Black ', 'Weasel', '1993-05-21', 'MALE', 'email@gmail.com', '');