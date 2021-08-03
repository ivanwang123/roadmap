--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: fuzzystrmatch; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS fuzzystrmatch WITH SCHEMA public;


--
-- Name: EXTENSION fuzzystrmatch; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION fuzzystrmatch IS 'determine similarities and distance between strings';


--
-- Name: status_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.status_type AS ENUM (
    'COMPLETE',
    'INCOMPLETE',
    'SKIP'
);


ALTER TYPE public.status_type OWNER TO postgres;

--
-- Name: revise_updated_at(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.revise_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW; 
END;
$$;


ALTER FUNCTION public.revise_updated_at() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: checkpoint_status; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.checkpoint_status (
    user_id integer NOT NULL,
    checkpoint_id integer NOT NULL,
    roadmap_id integer NOT NULL,
    status public.status_type DEFAULT 'INCOMPLETE'::public.status_type NOT NULL
);


ALTER TABLE public.checkpoint_status OWNER TO postgres;

--
-- Name: checkpoints; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.checkpoints (
    id integer NOT NULL,
    title character varying(256) NOT NULL,
    instructions text NOT NULL,
    links text[] DEFAULT '{}'::text[] NOT NULL,
    roadmap_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.checkpoints OWNER TO postgres;

--
-- Name: checkpoints_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.checkpoints_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.checkpoints_id_seq OWNER TO postgres;

--
-- Name: checkpoints_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checkpoints_id_seq OWNED BY public.checkpoints.id;


--
-- Name: roadmap_followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roadmap_followers (
    user_id integer NOT NULL,
    roadmap_id integer NOT NULL
);


ALTER TABLE public.roadmap_followers OWNER TO postgres;

--
-- Name: roadmaps; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roadmaps (
    id integer NOT NULL,
    title character varying(256) NOT NULL,
    description text NOT NULL,
    creator_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.roadmaps OWNER TO postgres;

--
-- Name: roadmaps_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roadmaps_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roadmaps_id_seq OWNER TO postgres;

--
-- Name: roadmaps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roadmaps_id_seq OWNED BY public.roadmaps.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(64) NOT NULL,
    email character varying(256) NOT NULL,
    password character varying(64) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: checkpoints id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkpoints ALTER COLUMN id SET DEFAULT nextval('public.checkpoints_id_seq'::regclass);


--
-- Name: roadmaps id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roadmaps ALTER COLUMN id SET DEFAULT nextval('public.roadmaps_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: checkpoint_status checkpoint_status_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkpoint_status
    ADD CONSTRAINT checkpoint_status_pkey PRIMARY KEY (user_id, checkpoint_id);


--
-- Name: checkpoints checkpoints_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkpoints
    ADD CONSTRAINT checkpoints_pkey PRIMARY KEY (id);


--
-- Name: roadmap_followers roadmap_followers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roadmap_followers
    ADD CONSTRAINT roadmap_followers_pkey PRIMARY KEY (user_id, roadmap_id);


--
-- Name: roadmaps roadmaps_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roadmaps
    ADD CONSTRAINT roadmaps_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: checkpoints revise_checkpoints_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER revise_checkpoints_updated_at BEFORE UPDATE ON public.checkpoints FOR EACH ROW EXECUTE FUNCTION public.revise_updated_at();


--
-- Name: roadmaps revise_roadmap_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER revise_roadmap_updated_at BEFORE UPDATE ON public.roadmaps FOR EACH ROW EXECUTE FUNCTION public.revise_updated_at();


--
-- Name: users revise_user_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER revise_user_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.revise_updated_at();


--
-- Name: checkpoint_status checkpoint_status_checkpoint_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkpoint_status
    ADD CONSTRAINT checkpoint_status_checkpoint_id_fkey FOREIGN KEY (checkpoint_id) REFERENCES public.checkpoints(id);


--
-- Name: checkpoint_status checkpoint_status_roadmap_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkpoint_status
    ADD CONSTRAINT checkpoint_status_roadmap_id_fkey FOREIGN KEY (roadmap_id) REFERENCES public.roadmaps(id);


--
-- Name: checkpoint_status checkpoint_status_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkpoint_status
    ADD CONSTRAINT checkpoint_status_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: checkpoints checkpoints_roadmap_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkpoints
    ADD CONSTRAINT checkpoints_roadmap_id_fkey FOREIGN KEY (roadmap_id) REFERENCES public.roadmaps(id) ON DELETE CASCADE;


--
-- Name: roadmap_followers roadmap_followers_roadmap_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roadmap_followers
    ADD CONSTRAINT roadmap_followers_roadmap_id_fkey FOREIGN KEY (roadmap_id) REFERENCES public.roadmaps(id);


--
-- Name: roadmap_followers roadmap_followers_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roadmap_followers
    ADD CONSTRAINT roadmap_followers_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: roadmaps roadmaps_creator_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roadmaps
    ADD CONSTRAINT roadmaps_creator_id_fkey FOREIGN KEY (creator_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

