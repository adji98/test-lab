--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5 (Ubuntu 11.5-1.pgdg18.04+1)
-- Dumped by pg_dump version 11.5 (Ubuntu 11.5-1.pgdg18.04+1)

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

SET default_with_oids = false;

--
-- Name: genres; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.genres (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.genres OWNER TO streamingweb;

--
-- Name: qualities; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.qualities (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.qualities OWNER TO streamingweb;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO streamingweb;

--
-- Name: season_details; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.season_details (
    id uuid NOT NULL,
    season_id uuid NOT NULL,
    video_id uuid NOT NULL,
    episode_no integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.season_details OWNER TO streamingweb;

--
-- Name: season_tags; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.season_tags (
    id uuid NOT NULL,
    season_id uuid NOT NULL,
    tag_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.season_tags OWNER TO streamingweb;

--
-- Name: seasons; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.seasons (
    id uuid NOT NULL,
    series_id uuid NOT NULL,
    season_no integer NOT NULL,
    plot character varying(255),
    completed boolean NOT NULL,
    episode_count integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.seasons OWNER TO streamingweb;

--
-- Name: series; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.series (
    id uuid NOT NULL,
    title character varying(255) NOT NULL,
    plot character varying(255) NOT NULL,
    completed boolean NOT NULL,
    season_count integer NOT NULL,
    episode_count integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.series OWNER TO streamingweb;

--
-- Name: series_genres; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.series_genres (
    id uuid NOT NULL,
    video_id uuid NOT NULL,
    series_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.series_genres OWNER TO streamingweb;

--
-- Name: tags; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.tags (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tags OWNER TO streamingweb;

--
-- Name: users; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    fullname character varying(255),
    shortname character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO streamingweb;

--
-- Name: video_genres; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.video_genres (
    id uuid NOT NULL,
    video_id uuid NOT NULL,
    genre_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.video_genres OWNER TO streamingweb;

--
-- Name: video_qualities; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.video_qualities (
    id uuid NOT NULL,
    video_id uuid NOT NULL,
    quality_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.video_qualities OWNER TO streamingweb;

--
-- Name: video_tags; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.video_tags (
    id uuid NOT NULL,
    video_id uuid NOT NULL,
    tag_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.video_tags OWNER TO streamingweb;

--
-- Name: video_views; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.video_views (
    id uuid NOT NULL,
    video_id uuid NOT NULL,
    quality_id uuid NOT NULL,
    user_id uuid NOT NULL,
    completed boolean NOT NULL,
    paused_on integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.video_views OWNER TO streamingweb;

--
-- Name: videoes; Type: TABLE; Schema: public; Owner: streamingweb
--

CREATE TABLE public.videoes (
    id uuid NOT NULL,
    title character varying(255) NOT NULL,
    plot character varying(255) NOT NULL,
    synopsis character varying(255) NOT NULL,
    length integer NOT NULL,
    view_count integer NOT NULL,
    like_count integer NOT NULL,
    dislike_count integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.videoes OWNER TO streamingweb;

--
-- Name: genres genres_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (id);


--
-- Name: qualities qualities_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.qualities
    ADD CONSTRAINT qualities_pkey PRIMARY KEY (id);


--
-- Name: season_details season_details_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.season_details
    ADD CONSTRAINT season_details_pkey PRIMARY KEY (id);


--
-- Name: season_tags season_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.season_tags
    ADD CONSTRAINT season_tags_pkey PRIMARY KEY (id);


--
-- Name: seasons seasons_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.seasons
    ADD CONSTRAINT seasons_pkey PRIMARY KEY (id);


--
-- Name: series_genres series_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.series_genres
    ADD CONSTRAINT series_genres_pkey PRIMARY KEY (id);


--
-- Name: series series_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: video_genres video_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.video_genres
    ADD CONSTRAINT video_genres_pkey PRIMARY KEY (id);


--
-- Name: video_qualities video_qualities_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.video_qualities
    ADD CONSTRAINT video_qualities_pkey PRIMARY KEY (id);


--
-- Name: video_tags video_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.video_tags
    ADD CONSTRAINT video_tags_pkey PRIMARY KEY (id);


--
-- Name: video_views video_views_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.video_views
    ADD CONSTRAINT video_views_pkey PRIMARY KEY (id);


--
-- Name: videoes videoes_pkey; Type: CONSTRAINT; Schema: public; Owner: streamingweb
--

ALTER TABLE ONLY public.videoes
    ADD CONSTRAINT videoes_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: streamingweb
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

CREATE OR REPLACE VIEW public.vw_continue_watching
AS SELECT a.user_id,
    a.video_id,
    a.quality_id,
    a.paused_on,
    a.updated_at AS last_view_at,
    c.title,
        CASE
            WHEN ss.* IS NULL THEN 0
            ELSE ss.season_no
        END AS season_no,
    COALESCE(se.title, ''::character varying) AS series_title
   FROM video_views a
     JOIN video_qualities b ON a.quality_id = b.id
     JOIN videoes c ON b.video_id = c.id
     LEFT JOIN season_details sd ON c.id = sd.video_id
     LEFT JOIN seasons ss ON sd.season_id = ss.id
     LEFT JOIN series se ON ss.series_id = se.id
  WHERE a.completed = false;


