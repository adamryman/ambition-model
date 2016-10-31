CREATE TABLE actions(id SERIAL PRIMARY KEY, action_name varchar(255), set_id integer, user_id integer, trello_id varchar(255))
CREATE TABLE occurrences(id SERIAL PRIMARY KEY, action_id integer, datetime varchar(255), data varchar(255))
