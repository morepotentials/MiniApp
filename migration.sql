CREATE SEQUENCE clients_id_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

CREATE TABLE public.clients (
  id integer NOT NULL DEFAULT nextval('clients_id_seq'),
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone DEFAULT now() NOT NULL,
  broker_id integer NOT NULL
);

ALTER SEQUENCE clients_id_seq OWNED BY public.clients.id;