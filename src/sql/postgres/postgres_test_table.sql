CREATE TABLE people(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  surname VARCHAR(50)
);
INSERT INTO public.people(name, surname) VALUES ( 'Ala', 'Kowalska');
INSERT INTO public.people(name, surname) VALUES ( 'Jane', 'Doe');