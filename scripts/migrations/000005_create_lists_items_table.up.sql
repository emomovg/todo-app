CREATE TABLE IF NOT EXISTS lists_items
(
    id serial not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
    )
