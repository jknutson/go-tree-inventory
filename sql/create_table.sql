-- Table: public.tree_inventory_v1

-- DROP TABLE IF EXISTS public.tree_inventory_v1;

CREATE TABLE IF NOT EXISTS public.tree_inventory_v1
(
    id integer NOT NULL DEFAULT nextval('tree_inventory_v1_id_seq'::regclass),
    type character varying COLLATE pg_catalog."default",
    location character varying COLLATE pg_catalog."default",
    notes text COLLATE pg_catalog."default",
    diameter_breast_height_inches numeric,
    diameter_dripline_feet numeric,
    CONSTRAINT tree_inventory_v1_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;
