CREATE TABLE IF NOT EXISTS "todo_list"(
    "id" bigserial primary key,
    "title" text,
    "description" text,
    "reminder" timestamp default current_timestamp,
    PRIMARY KEY ("id")
);