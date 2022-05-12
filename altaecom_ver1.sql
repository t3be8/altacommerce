CREATE DATABASE altacommerce;

use altacommerce;

DROP TABLE IF EXISTS users; 
DROP TABLE IF EXISTS address; 
DROP TABLE IF EXISTS product_categories; 
DROP TABLE IF EXISTS products; 


CREATE TABLE `users` (
	id int AUTO_INCREMENT PRIMARY KEY,
	name varchar(255),
	dob date,
	gender enum('L','P'),
	email varchar(50),
	phone varchar(13),
	password varchar(255),
	created_at timestamp default now(),
	updated_at timestamp
);

CREATE TABLE `address` (
	id int PRIMARY KEY AUTO_INCREMENT,
	address varchar(255),
	kodepos int(5),
	user_id int,
	CONSTRAINT FK_ADDRESS_USER foreign key(user_id) REFERENCES users(user_id) on update cascade on delete cascade,
	created_at timestamp default now(),
	updated_at timestamp
)

CREATE TABLE `product_categories` (
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(255),
	created_at timestamp default now(),
	updated_at timestamp
)

CREATE TABLE `products` (
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(255),
	description text,
	price numeric(18,2),
	stok int(5),
	images varchar(255),
	status 
	user_id int,
	product_category_id int,
	CONSTRAINT FK_PRODUCT_USER foreign key(user_id) REFERENCES users(user_id) on update cascade on delete cascade,
	constraint FK_PRODUCT_CATPRODUCT foreign key(product_category_id) references product_categories(product_category_id) on update cascade on delete cascade,
	created_at timestamp default current_timestamp,
	updated_at timestamp
)