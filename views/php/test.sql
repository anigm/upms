create database zatweb;

use zatweb;

create table ZatUser (
    id int key auto_increment,
	f_user nvarchar(255),
    f_password nvarchar(255)
);

insert into ZatUser(`f_user`,`f_password`) values('guest','guest');