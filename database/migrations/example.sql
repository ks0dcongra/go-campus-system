--
-- PostgreSQL database dump
--

-- Dumped from database version 10.18
-- Dumped by pg_dump version 10.18

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
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: example_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.example_item (
    id integer NOT NULL,
    data character varying(30) NOT NULL,
    updated_time timestamp with time zone,
    created_time timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.example_item OWNER TO postgres;

--
-- Name: example_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.example_item_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.example_item_id_seq OWNER TO postgres;

--
-- Name: example_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.example_item_id_seq OWNED BY public.example_item.id;


--
-- Name: example_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.example_item ALTER COLUMN id SET DEFAULT nextval('public.example_item_id_seq'::regclass);


--
-- Data for Name: example_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.example_item (id, data, updated_time, created_time) FROM stdin;
1	example1	0001-01-01 08:06:00+08:06	2022-06-28 15:35:13.759176+08
2	example1	0001-01-01 08:06:00+08:06	2022-06-28 15:35:16.095133+08
3	example1	2022-06-28 15:35:16.095+08	2022-06-28 15:35:16.095+08
4	example1	2022-06-28 15:36:53.978128+08	2022-06-28 15:36:53.978128+08
5	example1	2022-06-28 15:35:16.095+08	2022-06-28 15:35:16.095+08
6	example1	2022-06-28 15:35:16.095+08	2022-06-28 15:35:16.095+08
7	example1	2022-06-28 15:35:16.095+08	2022-06-28 15:35:16.095+08
8	example1	2022-06-28 15:35:16.095+08	2022-06-28 15:35:16.095+08
9	example1	2022-06-28 15:35:16.095+08	2022-06-28 15:35:16.095+08
10	example1	2022-06-28 15:35:16.095+08	2022-06-28 15:35:16.095+08
11	example1	2022-06-28 15:37:19.333808+08	2022-06-28 15:37:19.333808+08
12	example1	2022-06-28 15:37:29.496937+08	2022-06-28 15:37:29.496937+08
13	example1	2022-06-28 15:38:44.049342+08	2022-06-28 15:38:44.049342+08
14	example1	2022-06-28 15:38:47.59192+08	2022-06-28 15:38:47.59192+08
\.


--
-- Name: example_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.example_item_id_seq', 14, true);


--
-- PostgreSQL database dump complete
--

