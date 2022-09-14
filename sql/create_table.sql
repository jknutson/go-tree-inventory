-- SEQUENCE: public.tree_inventory_v2_id_seq

-- DROP SEQUENCE IF EXISTS public.tree_inventory_v2_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.tree_inventory_v2_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY tree_inventory_v2.id;

--- --- ---

-- Table: public.tree_inventory_v2

-- DROP TABLE IF EXISTS public.tree_inventory_v2;

CREATE TABLE IF NOT EXISTS public.tree_inventory_v2
(
    id integer NOT NULL DEFAULT nextval('tree_inventory_v2_id_seq'::regclass),
    type character varying COLLATE pg_catalog."default",
    notes text COLLATE pg_catalog."default",
    diameter_breast_height_inches numeric,
    diameter_dripline_feet numeric,
    geom geometry,
    CONSTRAINT tree_inventory_v2_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;
