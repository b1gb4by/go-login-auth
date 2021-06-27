---- DROP ----
drop schema if exists go_auth;

create schema go_auth;

use go_auth;

drop table if exists `users`;

---- CREATE ----
create table if not exists `users` (
    `id` int(20) auto_increment,
    `first_name` varchar(255),
    `last_name` varchar(255),
    `email` varchar(255),
    `password` varchar(255),
    `created_at` datetime default current_timestamp,
    `updated_at` default current_timestamp on update current_timestamp,
    primary key(`id`)
) default charset = utf8 collate = utf8_bin;

---- INSERT ----
insert into
    users (first_name, last_name, email, password)
values
    (
        "Alan",
        "Bob",
        "alan@example.com",
        "test_password"
    );

insert into
    users (first_name, last_name, email, password)
values
    (
        "Callen",
        "Den",
        "callen@example.com",
        "test_password1"
    );
