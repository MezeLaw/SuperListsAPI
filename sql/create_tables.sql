CREATE TABLE IF NOT EXISTS product (
                                       product_id INT NOT NULL,
                                       name varchar(250) NOT NULL,
    PRIMARY KEY (product_id)
    );


CREATE TABLE IF NOT EXISTS users (
                              id serial PRIMARY KEY,
                              "name" text not null,
                              email text NOT NULL UNIQUE,
                              password text NOT NULL,
                              role text NOT NULL,
                              created_at timestamp without time zone NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone NULL,
                              deleted_at timestamp without time zone NULL
);


CREATE TABLE IF NOT EXISTS lists (
                              id serial PRIMARY KEY,
                              "name" varchar(150) NULL,
                              description varchar(150) NULL,
                              invite_code text NULL,
                              user_creator_id bigint NOT NULL,
                              created_at timestamp without time zone null DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone null DEFAULT NULL,
                              deleted_at timestamp without time zone null DEFAULT NULL
);


CREATE TABLE IF NOT EXISTS user_lists (
                                   id serial PRIMARY KEY,
                                   list_id serial NOT NULL,
                                   user_id bigint NOT NULL,
                                   created_at timestamp without time zone NULL DEFAULT CURRENT_TIMESTAMP,
                                   updated_at timestamp without time zone NULL,
                                   deleted_at timestamp without time zone NULL
);

ALTER TABLE user_lists ADD CONSTRAINT user_lists_list_id_fk FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE user_lists ADD CONSTRAINT user_lists_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE RESTRICT;

CREATE TABLE list_items (
                              id serial PRIMARY KEY,
                              list_id bigint NOT NULL,
                              user_id bigint not null,
                              title varchar(150) NULL,
                              description TEXT NULL,
                              is_done boolean NULL DEFAULT false,
                              created_at timestamp without time zone NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone NULL,
                              deleted_at timestamp without time zone NULL
);

ALTER TABLE list_item ADD CONSTRAINT item_creator_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE list_item ADD CONSTRAINT list_id_fk FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE RESTRICT ON UPDATE RESTRICT;