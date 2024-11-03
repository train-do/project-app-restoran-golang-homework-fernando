create database "Restoran";

create table if not exists "User" (
    "id" serial primary key,
    "username" varchar(100) unique not null,
    "password" varchar(100) not null,
    "created_at" timestamp not null
);

create table if not exists "Admin" (
    "id" serial primary key,
    "user_id" int references "User"("id") not null,
    "name" varchar(50) not null,
    "created_at" timestamp not null
);

create table if not exists "Chef" (
    "id" serial primary key,
    "user_id" int references "User"("id") not null,
    "name" varchar(50) not null,
    "created_at" timestamp not null
);

create table if not exists "Customer" (
    "id" serial primary key,
    "user_id" int references "User"("id") not null,
    "name" varchar(50) not null,
    "created_at" timestamp not null
);

create table if not exists "Item" (
    "id" serial primary key,
    "name" varchar(100) not null,
    "price" int not null,
    "created_at" timestamp not null
);

create table if not exists "Order" (
    "id" serial primary key,
    "customer_id" int references "Customer"("id") not null,
    "total_price" int,
    "discount" int,
    "rating" int,
    "created_at" timestamp not null
);

create table if not exists "OrderItem" (
    "id" serial primary key,
    "order_id" int references "Order"("id") not null,
    "item_id" int references "Item"("id") not null,
    "quantity" int not null,
    "status" varchar(30),
    "created_at" timestamp not null
);
