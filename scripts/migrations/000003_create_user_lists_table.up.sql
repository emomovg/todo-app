CREATE TABLE IF NOT EXISTS users_lists
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
)
