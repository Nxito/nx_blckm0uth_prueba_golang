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


CREATE DATABASE mytestdb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Spanish_Spain.1252';

ALTER DATABASE mytestdb OWNER TO admin;
alter system set idle_in_transaction_session_timeout='1min';

\connect mytestdb

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


/*Funciones*/

/* Ante una operacion update o insert, se añadirán datos */
CREATE FUNCTION public.dates_update_trigger() 
RETURNS trigger LANGUAGE plpgsql AS $$
DECLARE
	nid bigint = 0;
BEGIN

        IF (TG_OP = 'UPDATE') THEN
            NEW.update_date:= NOW();
        ELSIF (TG_OP = 'INSERT') THEN			
			NEW.creation_date:= NOW(); 
        END IF;		

        RETURN NEW;
    END;
$$;
ALTER FUNCTION public.dates_update_trigger() OWNER TO admin;

/*
    Ante una operacion se actualiza la propiedad status basada en maxplayers de la queue indicada en la session
*/

CREATE OR REPLACE FUNCTION public.update_session_status_trigger() 
RETURNS trigger  LANGUAGE plpgsql AS $$
DECLARE
    players_limit INTEGER;
BEGIN
  -- get limit players from specified queue 
  SELECT max_players INTO players_limit FROM public.queues WHERE id =  NEW.id_queue;

 IF players_limit IS NULL THEN
    RAISE EXCEPTION 'The players limit or the queue not founded';
  END IF;

  IF (array_length(NEW.players, 1) >= players_limit) THEN
    NEW.status := 'closed';
  END IF;

  RETURN NEW;
END;
$$ ;

ALTER FUNCTION public.update_session_status_trigger() OWNER TO admin;



SET default_tablespace = '';

SET default_table_access_method = heap;

-- Creo la tabla player 
CREATE TABLE public.players (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(50) NOT NULL,
    creation_date timestamp with time zone DEFAULT now() NOT NULL,
    update_date timestamp with time zone
);


ALTER TABLE public.players OWNER TO admin ;
ALTER TABLE public.players ADD CONSTRAINT unique_player_name UNIQUE (name);

-- Creo la tabla de colas 
CREATE TABLE public.queues (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(50) NOT NULL,
    max_players int DEFAULT 2 NOT NULL,
    creation_date timestamp with time zone DEFAULT now() NOT NULL,
    update_date timestamp with time zone
);


ALTER TABLE public.queues OWNER TO admin;
ALTER TABLE public.queues ADD CONSTRAINT unique_queue_name UNIQUE (name);

-- Creo la tabla de sesiones 
CREATE TABLE public.sessions (
    id uuid DEFAULT gen_random_uuid() NOT NULL, 
    id_queue uuid NOT NULL, 
    players UUID[] NOT NULL, 
    status character varying(50) DEFAULT 'opened' NOT NULL,
    creation_date timestamp with time zone DEFAULT now() NOT NULL,
    update_date timestamp with time zone
);


ALTER TABLE public.sessions OWNER TO admin;

-- -- Creo la tabla de parametros y configuración 
-- CREATE TABLE public.parameters (
--     name VARCHAR(255) PRIMARY KEY,
--     value  VARCHAR(255)
-- );

-- INSERT INTO public.parameters (name, value) VALUES
--     ('max_queues', '10');

-- ALTER TABLE public.parameters ADD CONSTRAINT unique_parameter_name UNIQUE (name);


--triggers 
--actualizo la fecha de update en las tablas al insertar o actualizar usando la funcion dates_update_trigger creada mas arriba
CREATE TRIGGER trigger_players
BEFORE INSERT OR UPDATE ON public.players
FOR EACH ROW EXECUTE FUNCTION public.dates_update_trigger();

CREATE TRIGGER trigger_queues
BEFORE INSERT OR UPDATE ON public.queues
FOR EACH ROW EXECUTE FUNCTION public.dates_update_trigger();

CREATE TRIGGER trigger_sessions
BEFORE INSERT OR UPDATE ON public.sessions
FOR EACH ROW EXECUTE FUNCTION public.dates_update_trigger();

-- reviso la longitud de la propiedad players para  cerrar la  cola
CREATE TRIGGER trigger_update_session_status
BEFORE UPDATE
ON public.sessions
FOR EACH ROW
EXECUTE FUNCTION public.update_session_status_trigger();


