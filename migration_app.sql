

CREATE SEQUENCE users_id_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

CREATE TYPE usertype AS ENUM ('broker', 'client');

CREATE TABLE public.users (
  id integer NOT NULL DEFAULT nextval('users_id_seq') UNIQUE,
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone DEFAULT now() NOT NULL,
  user_type usertype,
  first_name text,
  last_name text,
  email text UNIQUE
);

ALTER SEQUENCE users_id_seq OWNED BY public.users.id;

CREATE TABLE public.broker (
  broker_id integer REFERENCES public.users(id),
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE public.clients (
  client_id integer REFERENCES public.users(id),
  broker_id integer REFERENCES public.users(id),
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone DEFAULT now() NOT NULL
);