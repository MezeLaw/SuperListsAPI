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