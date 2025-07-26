CREATE TABLE IF NOT EXISTS todo_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
)
