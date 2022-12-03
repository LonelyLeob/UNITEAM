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
    id UUID not null primary key,
    name text not null unique,
    email text not null unique,
    password varchar(200) not null
);

CREATE TABLE IF NOT EXISTS meta (
    id serial not null primary key,
    uuid UUID not null,
    lv bigint not null,
    browser varchar(50),
    os varchar(50),
    refresh varchar(200)
);

CREATE TABLE IF NOT EXISTS courses (
    id serial not null primary key,
    title text not null,
    cdesc text not null,
    author varchar(200) not null
);

CREATE TABLE IF NOT EXISTS sections (
    id serial not null primary key,
    course_id int not null,
    content text not null,
        constraint fk_course
            foreign key(course_id)
                references courses(id)
);