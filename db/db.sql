create table if not exists product (
   product_id serial primary key,
   name varchar(50) not null,
   price numeric(9,2) not null,
   image varchar(255),
   is_in_stock boolean not null,

   created_at date not null,
   updated_at date,
   deleted_at date,

   category_id int not null
);

create table if not exists category (
    category_id serial primary key,
    name varchar(50) not null unique,

    created_at date not null,
    updated_at date,
    deleted_at date
);

alter table product add constraint fk_category foreign key (category_id) references category(category_id);

insert into category(name, created_at, updated_at, deleted_at) values ('Food', now(), null, null);
insert into category(name, created_at, updated_at, deleted_at) values ('Other', now(), null, null);

insert into product(name, price, image, is_in_stock, created_at, updated_at, deleted_at, category_id) values ('Hamburger', 5.99, 'link/to/image/hamburger', true, now(), null, null, 1);
insert into product(name, price, image, is_in_stock, created_at, updated_at, deleted_at, category_id) values ('Corn', 2, 'link/to/image/corn', true, now(), null, null, 1);
insert into product(name, price, image, is_in_stock, created_at, updated_at, deleted_at, category_id) values ('Bean', 1.99, 'link/to/image/bean', false, now(), null, null, 1);
insert into product(name, price, image, is_in_stock, created_at, updated_at, deleted_at, category_id) values ('HDMI cable', 15.79, 'link/to/image/hdmi-cable', true, now(), null, null, 2);
insert into product(name, price, image, is_in_stock, created_at, updated_at, deleted_at, category_id) values ('Sun screen', 3.29, 'link/to/image/sun-screen', true, now(), null, null, 2);
