--
-- PostgreSQL database dump
--

-- Dumped from database version 10.21 (Ubuntu 10.21-1.pgdg22.04+1)
-- Dumped by pg_dump version 10.21 (Ubuntu 10.21-1.pgdg22.04+1)

-- Started on 2023-05-03 12:09:30 WIB

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

DROP DATABASE upn_belajar_go;
--
-- TOC entry 3030 (class 1262 OID 17021)
-- Name: upn_belajar_go; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE upn_belajar_go WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';


ALTER DATABASE upn_belajar_go OWNER TO postgres;

\connect upn_belajar_go

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
-- TOC entry 1 (class 3079 OID 13104)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 3033 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 17022)
-- Name: jenis_mitra; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.jenis_mitra (
    id character varying(36) NOT NULL,
    nama_jenis_mitra character varying(200),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.jenis_mitra OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 17044)
-- Name: kelas_siswa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.kelas_siswa (
    id character varying(36) NOT NULL,
    id_kelas character varying(36),
    tahun_ajaran text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.kelas_siswa OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 17053)
-- Name: kelas_siswa_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.kelas_siswa_detail (
    id character varying(36) NOT NULL,
    id_kelas_siswa character varying(36),
    id_siswa character varying(36),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.kelas_siswa_detail OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 17034)
-- Name: m_kelas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_kelas (
    id character varying(36) NOT NULL,
    kode character varying(100),
    nama character varying(100),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.m_kelas OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 17028)
-- Name: m_siswa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_siswa (
    id character varying(36) NOT NULL,
    nama character varying(100),
    kelas character varying(100),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false,
    berkas character varying,
    id_kelas character varying
);


ALTER TABLE public.m_siswa OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 27817)
-- Name: role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role (
    id character varying(36) NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.role OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 27808)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    id_role character varying(36),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3018 (class 0 OID 17022)
-- Dependencies: 196
-- Data for Name: jenis_mitra; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.jenis_mitra (id, nama_jenis_mitra, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('67153215-b6cd-44c1-a20e-12321f82e572', 'RISET', '2023-03-07 06:56:41.942254', NULL, '667ff26a-d8a0-4ae1-9a5a-277305f404d1', NULL, false);


--
-- TOC entry 3021 (class 0 OID 17044)
-- Dependencies: 199
-- Data for Name: kelas_siswa; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.kelas_siswa (id, id_kelas, tahun_ajaran, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('7a863728-1ae1-47d6-b3a5-a60c5b84b956', '279db6e8-6f00-4f20-8c22-a5ce914c878b', '2022/20231', '2023-04-11 11:08:43.931981', '2023-05-03 10:10:59.27072', '', '', false);
INSERT INTO public.kelas_siswa (id, id_kelas, tahun_ajaran, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('61d33b33-11d5-4bf9-a5ee-d134f7f4da08', '279db6e8-6f00-4f20-8c22-a5ce914c878b', '2022/2023', '2023-04-18 09:03:50.367867', '2023-05-03 10:49:23.231727', '', '', true);


--
-- TOC entry 3022 (class 0 OID 17053)
-- Dependencies: 200
-- Data for Name: kelas_siswa_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.kelas_siswa_detail (id, id_kelas_siswa, id_siswa, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('8398cad2-4584-46a6-ae4a-f7176cad29b9', '7a863728-1ae1-47d6-b3a5-a60c5b84b956', '91a2e856-2128-47ee-898d-f0dce054f7d1', '2023-04-18 09:01:05.648573', NULL, '', NULL, true);
INSERT INTO public.kelas_siswa_detail (id, id_kelas_siswa, id_siswa, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('5a5d91d0-1895-4e96-bede-ee3d17c62d2b', '7a863728-1ae1-47d6-b3a5-a60c5b84b956', '97f86370-5dde-476f-92cf-8a012cd23039', '2023-05-03 10:08:20.976283', NULL, '', NULL, false);


--
-- TOC entry 3020 (class 0 OID 17034)
-- Dependencies: 198
-- Data for Name: m_kelas; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.m_kelas (id, kode, nama, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('54c9cd72-9541-4dc4-8864-16dc8236cfcc', 'X-IPA-B', 'X IPA B', '2023-02-28 08:26:53.492433', '2023-02-28 08:26:53.492798', '', NULL, false);
INSERT INTO public.m_kelas (id, kode, nama, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('279db6e8-6f00-4f20-8c22-a5ce914c878b', 'X-IPA-A', 'X IPA A', '2023-02-28 08:25:41.270441', '2023-02-28 08:30:36.543048', '', '', false);


--
-- TOC entry 3019 (class 0 OID 17028)
-- Dependencies: 197
-- Data for Name: m_siswa; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.m_siswa (id, nama, kelas, created_at, updated_at, created_by, updated_by, is_deleted, berkas, id_kelas) VALUES ('5428ec3a-4a8e-444b-aa16-dd8f9d2cb0ee', 'Bambang Tri', 'A', '2023-02-28 11:08:22.69835', '2023-02-28 11:19:48.458771', '', '', true, NULL, NULL);
INSERT INTO public.m_siswa (id, nama, kelas, created_at, updated_at, created_by, updated_by, is_deleted, berkas, id_kelas) VALUES ('eb6a8af6-a218-4f41-a96e-6b982096bc89', 'Ilham', 'B', '2023-03-07 10:43:23.418977', '2023-03-28 12:06:20.016878', '', '', true, NULL, NULL);
INSERT INTO public.m_siswa (id, nama, kelas, created_at, updated_at, created_by, updated_by, is_deleted, berkas, id_kelas) VALUES ('97f86370-5dde-476f-92cf-8a012cd23039', 'Bambang', 'C', '2023-03-28 11:51:41.34781', '2023-04-04 11:54:30.85422', '', '', false, '/files/berkas_siswa/berkas_siswa_18e38cab-5fb2-466f-ad32-35b909055d26.jpg', '54c9cd72-9541-4dc4-8864-16dc8236cfcc');
INSERT INTO public.m_siswa (id, nama, kelas, created_at, updated_at, created_by, updated_by, is_deleted, berkas, id_kelas) VALUES ('f17f7e58-00f8-4c26-b43f-beb1aff4e781', 'Ali A', 'V', '2023-03-14 11:17:44.479128', '2023-04-04 11:54:35.341454', '', '', false, '/files/berkas_siswa/berkas_siswa_cdbc5831-5978-46c8-81c2-e0eb26fac507.jpg', '279db6e8-6f00-4f20-8c22-a5ce914c878b');
INSERT INTO public.m_siswa (id, nama, kelas, created_at, updated_at, created_by, updated_by, is_deleted, berkas, id_kelas) VALUES ('91a2e856-2128-47ee-898d-f0dce054f7d1', 'Dimas', '10', '2023-04-04 11:52:39.758261', '2023-04-11 08:40:13.557833', '', '', false, '/files/berkas_siswa/berkas_siswa_6bde05a7-abe7-4412-912a-efbca7a40da1.jpg', '279db6e8-6f00-4f20-8c22-a5ce914c878b');


--
-- TOC entry 3024 (class 0 OID 27817)
-- Dependencies: 202
-- Data for Name: role; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.role (id, name) VALUES ('HA01', 'SUPERADMIN');


--
-- TOC entry 3023 (class 0 OID 27808)
-- Dependencies: 201
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, name, username, email, password, id_role, created_at, updated_at, is_deleted) VALUES ('0274b088-dd88-11ed-9c44-8a7f1a7d5dd4', 'Muhammad Alkautsar', 'superadmin', 'superadmin@gmail.com', '$2a$10$0Eo6pnewf.GTxCz9HYp9m.Tv1S8UrR87pNBd/EKv64XqBM9Vr1R2u', 'HA01', '2023-04-18 08:26:42', '2023-04-18 08:26:42', false);


--
-- TOC entry 2882 (class 2606 OID 17027)
-- Name: jenis_mitra jenis_mitra_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jenis_mitra
    ADD CONSTRAINT jenis_mitra_pkey PRIMARY KEY (id);


--
-- TOC entry 2890 (class 2606 OID 17058)
-- Name: kelas_siswa_detail kelas_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kelas_siswa_detail
    ADD CONSTRAINT kelas_detail_pkey PRIMARY KEY (id);


--
-- TOC entry 2888 (class 2606 OID 17052)
-- Name: kelas_siswa kelas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kelas_siswa
    ADD CONSTRAINT kelas_pkey PRIMARY KEY (id);


--
-- TOC entry 2886 (class 2606 OID 17039)
-- Name: m_kelas m_kelas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_kelas
    ADD CONSTRAINT m_kelas_pkey PRIMARY KEY (id);


--
-- TOC entry 2884 (class 2606 OID 17033)
-- Name: m_siswa m_siswa_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_siswa
    ADD CONSTRAINT m_siswa_pkey PRIMARY KEY (id);


--
-- TOC entry 2894 (class 2606 OID 27823)
-- Name: role role_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_name_key UNIQUE (name);


--
-- TOC entry 2896 (class 2606 OID 27821)
-- Name: role role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);


--
-- TOC entry 2892 (class 2606 OID 27816)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3032 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2023-05-03 12:09:30 WIB

--
-- PostgreSQL database dump complete
--

