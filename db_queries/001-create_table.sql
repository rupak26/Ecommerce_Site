CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT CHECK (age >= 0),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    occupation VARCHAR(100)
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    productname VARCHAR(100),
	url TEXT,
	quantity INT
);


UPDATE users 
SET 
   name = :name, 
   occupation = :occupation 
WHERE id = :id RETURNING id, 
   name, age, email, 
   password, occupation;