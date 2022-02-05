CREATE TABLE public.users (
                              id serial PRIMARY KEY,
                              "name" text not null,
                              email text NOT NULL,
                              role text NOT NULL,
                              created_at timestamp without time zone NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone NULL,
                              deleted_at timestamp without time zone NULL
);

CREATE TABLE public.lists (
                              id serial PRIMARY KEY,
                              "name" varchar(150) NULL,
                              description varchar(150) NULL,
                              invite_code text NULL,
                              created_at timestamp without time zone null DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone null DEFAULT NULL,
                              deleted_at timestamp without time zone null DEFAULT NULL
);


CREATE TABLE public.user_lists (
                                   id serial PRIMARY KEY,
                                   list_id serial NOT NULL,
                                   user_id bigint NOT NULL,
                                   created_at timestamp without time zone NULL DEFAULT CURRENT_TIMESTAMP,
                                   updated_at timestamp without time zone NULL,
                                   deleted_at timestamp without time zone NULL
);


ALTER TABLE public.user_lists ADD CONSTRAINT user_lists_list_id_fk FOREIGN KEY (list_id) REFERENCES public.lists(id) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE public.user_lists ADD CONSTRAINT user_lists_user_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE RESTRICT ON UPDATE RESTRICT;



CREATE TABLE public.tasks (
                              id serial PRIMARY KEY,
                              creator_user_id bigint not null,
                              description varchar(350) NULL,
                              task_done boolean NULL DEFAULT false,
                              list_id bigint not null,
                              created_at timestamp without time zone NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone NULL,
                              deleted_at timestamp without time zone NULL
);

ALTER TABLE public.tasks ADD CONSTRAINT tasks_creator_user_id_fk FOREIGN KEY (creator_user_id) REFERENCES public.users(id) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE public.tasks ADD CONSTRAINT tasks_list_id_fk FOREIGN KEY (list_id) REFERENCES public.lists(id) ON DELETE RESTRICT ON UPDATE RESTRICT;