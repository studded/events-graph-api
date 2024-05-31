-- EVENTS --
CREATE TABLE events (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text,
    start_date date,
    end_date date,
    location text,
    description text
);

-- ACTIVITIES --
CREATE TABLE activities (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    event_id integer,
    name text,
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    description text
);
CREATE INDEX activities_event_id_index ON activities(event_id int4_ops);

-- USERS --
CREATE TABLE users (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text,
    email text,
    phone text
);
CREATE UNIQUE INDEX users_email_index ON users(email text_ops);

-- ROLES --
CREATE TYPE user_role AS ENUM ('admin', 'contributor', 'attendee');
CREATE TABLE roles (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer,
    event_id integer,
    type user_role
);
CREATE UNIQUE INDEX roles_user_id_event_id_index ON roles(user_id int4_ops,event_id int4_ops);
CREATE INDEX roles_user_id_index ON roles(user_id int4_ops);
CREATE INDEX roles_event_id_index ON roles(event_id int4_ops);

-- EXPENSES --
CREATE TABLE expenses (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    event_id integer,
    name text,
    cost double precision,
    description text,
    category text
);
CREATE INDEX expenses_event_id_index ON expenses(event_id int4_ops);