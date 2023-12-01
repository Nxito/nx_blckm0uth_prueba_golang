SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = '1min';
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;


CREATE DATABASE mydb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Spanish_Spain.1252';

ALTER DATABASE mydb OWNER TO admin;
alter system set idle_in_transaction_session_timeout='1min';

\connect mydb

SET statement_timeout = 0;
SET lock_timeout = 0;
--SET idle_in_transaction_session_timeout = '1min';
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;



CREATE FUNCTION public.entity_insert_update_trigger() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
	nid bigint = 0;
BEGIN
        /*
        Ante una operacion, se añaden datos
        */
        IF (TG_OP = 'UPDATE') THEN
            NEW.update_date:= NOW();
        ELSIF (TG_OP = 'INSERT') THEN			
			NEW.creation_date:= NOW(); 
        END IF;		
				        
        RETURN NEW;
    END;
$$;


ALTER FUNCTION public.entity_insert_update_trigger() OWNER TO admin;

SET default_tablespace = '';

SET default_table_access_method = heap;

-- Creo la tabla principal 
CREATE TABLE public.instance (
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    id_process_definition bigint NOT NULL,
    description character varying(250),
    creation_date timestamp with time zone DEFAULT now() NOT NULL,
    update_date timestamp with time zone,
    property jsonb
);


ALTER TABLE public.instance OWNER TO admin;
