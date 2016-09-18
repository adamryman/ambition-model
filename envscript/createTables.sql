CREATE TABLE actions(id SERIAL PRIMARY KEY, action_name varchar(255), set_id integer, user_id integer)
CREATE TABLE occurrences(id SERIAL PRIMARY KEY, action_id integer, time timestamp)
