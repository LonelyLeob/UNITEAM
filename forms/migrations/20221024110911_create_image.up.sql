CREATE TABLE IF NOT EXISTS form (
    uuid UUID not null primary key,
    quizName varchar(50) not null,
    quizDesc varchar(200) not null,
    anonym boolean not null,
    authorName varchar(200) not null
);

CREATE TABLE IF NOT EXISTS field (
    id serial not null,
    formUuid UUID,
    fieldName varchar(50) not null,
    primary key(id),
    constraint fk_forms
        foreign key(formUuid)
            references form(uuid)
);

CREATE TABLE IF NOT EXISTS answer (
    id serial not null,
    answer varchar(100) not null,
    fieldId int,
        primary key (id),
        constraint fk_field
            foreign key(fieldId)
                references field(id)
);

CREATE TABLE IF NOT EXISTS users (
    name text not null primary key,
    password varchar(200) not null,
    role text not null,
    email text not null unique,
    lastv bigint not null
);