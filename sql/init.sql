create table if not exists users(
    id bigint primary key auto_increment,
    name varchar(200) not null,
    dob varchar(10) not null,
    description varchar(300),
    address varchar(300),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)
engine = InnoDB
default charset = utf8;