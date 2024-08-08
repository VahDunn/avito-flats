-- Создание последовательности для flats
CREATE SEQUENCE flats_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Создание таблицы flats
CREATE TABLE public.flats (
    id integer NOT NULL DEFAULT nextval('flats_id_seq'::regclass),
    house_id integer NOT NULL,
    flat_number integer NOT NULL,
    rooms integer NOT NULL,
    price numeric(10,2) NOT NULL,
    user_id integer NOT NULL,
    moderation_status smallint NOT NULL,
    CONSTRAINT flats_pkey PRIMARY KEY (id),
    CONSTRAINT flats_moderation_status_check CHECK (moderation_status = ANY (ARRAY[0, 1])),
    CONSTRAINT flats_house_id_fkey FOREIGN KEY (house_id)
        REFERENCES public.houses (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT flats_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

-- Создание последовательности для houses
CREATE SEQUENCE houses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Создание таблицы houses
CREATE TABLE public.houses (
    id integer NOT NULL DEFAULT nextval('houses_id_seq'::regclass),
    address character varying(255) NOT NULL,
    build_year integer NOT NULL,
    developer character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT houses_pkey PRIMARY KEY (id)
);

-- Триггер для обновления поля updated_at в таблице houses
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_houses_updated_at
BEFORE UPDATE ON public.houses
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Создание последовательности для subscribers
CREATE SEQUENCE subscribers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Создание таблицы subscribers
CREATE TABLE public.subscribers (
    id integer NOT NULL DEFAULT nextval('subscribers_id_seq'::regclass),
    house_id integer NOT NULL,
    email character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT subscribers_pkey PRIMARY KEY (id),
    CONSTRAINT subscribers_house_id_email_key UNIQUE (house_id, email),
    CONSTRAINT fk_subscribers_house FOREIGN KEY (house_id)
        REFERENCES public.houses (id) MATCH SIMPLE
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);

-- Триггер для обновления поля updated_at в таблице subscribers
CREATE TRIGGER update_subscribers_updated_at
BEFORE UPDATE ON public.subscribers
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Создание последовательности для users
CREATE SEQUENCE users_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Создание таблицы users
CREATE TABLE public.users (
    user_id integer NOT NULL DEFAULT nextval('users_user_id_seq'::regclass),
    usertype smallint NOT NULL,
    email character varying(255) NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_id),
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_role_check CHECK (usertype = ANY (ARRAY[0, 1]))
);