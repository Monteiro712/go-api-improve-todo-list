CREATE TABLE tasks (
    id serial PRIMARY KEY,
    title varchar,
    about varchar,
    completed boolean,
    userid serial
);

INSERT INTO tasks (title, about, completed, userid) VALUES
    ('Tarefa 1', 'Descrição da Tarefa 1', true, DEFAULT),
    ('Tarefa 2', 'Descrição da Tarefa 2', false, DEFAULT);
