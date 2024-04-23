--
-- PostgreSQL database dump
--

-- Dumped from database version 14.0 (Debian 14.0-1.pgdg110+1)
-- Dumped by pg_dump version 14.0 (Debian 14.0-1.pgdg110+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: test_models; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.test_models (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text
);


ALTER TABLE public.test_models OWNER TO postgres;

--
-- Name: test_models_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.test_models_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.test_models_id_seq OWNER TO postgres;

--
-- Name: test_models_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.test_models_id_seq OWNED BY public.test_models.id;


--
-- Name: test_models id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_models ALTER COLUMN id SET DEFAULT nextval('public.test_models_id_seq'::regclass);


--
-- Data for Name: test_models; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.test_models (id, created_at, updated_at, deleted_at, name) FROM stdin;
1	\N	\N	\N	omid
2	\N	\N	\N	ali
\.


--
-- Name: test_models_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.test_models_id_seq', 1, false);


--
-- Name: test_models test_models_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_models
    ADD CONSTRAINT test_models_pkey PRIMARY KEY (id);


--
-- Name: idx_test_models_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_test_models_deleted_at ON public.test_models USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

