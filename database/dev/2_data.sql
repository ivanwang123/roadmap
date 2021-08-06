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
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, email, password, created_at, updated_at) FROM stdin;
1	test	test@test.com	$2a$14$QrCCnI4Kftg0WvdrbFrzOeW9Vy0q4K2TtD4mcWYTDA27Vc0X0cIWG	2021-07-15 20:55:25.877196-04	2021-07-15 20:55:25.877196-04
\.


--
-- Data for Name: roadmaps; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roadmaps (id, title, description, creator_id, created_at, updated_at) FROM stdin;
1	map	description	1	2021-07-15 20:55:30.702377-04	2021-07-15 20:55:30.702377-04
\.


--
-- Data for Name: checkpoints; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.checkpoints (id, title, instructions, links, roadmap_id, created_at, updated_at) FROM stdin;
1	Checkpoint one	Instructions	{https://tailwindcss.com/docs/justify-items#center}	1	2021-07-15 20:55:53.467122-04	2021-07-15 20:55:53.467122-04
2	Checkpoint two	Instructions	{https://tailwindcss.com/docs/justify-items#center}	1	2021-07-15 20:59:25.618622-04	2021-07-15 20:59:25.618622-04
3	Checkpoint three	Instructions	{https://tailwindcss.com/docs/justify-items#center}	1	2021-07-15 21:15:40.528612-04	2021-07-15 21:15:40.528612-04
4	Checkpoint four	Instructions	{https://tailwindcss.com/docs/justify-items#center}	1	2021-07-15 21:16:40.380588-04	2021-07-15 21:16:40.380588-04
5	Checkpoint four	Instructions	{https://tailwindcss.com/docs/justify-items#center}	1	2021-07-15 21:16:50.875662-04	2021-07-15 21:16:50.875662-04
6	Checkpoint four	Instructions	{https://tailwindcss.com/docs/justify-items#center}	1	2021-07-15 21:17:08.326898-04	2021-07-15 21:17:08.326898-04
\.


--
-- Data for Name: checkpoint_status; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.checkpoint_status (user_id, checkpoint_id, roadmap_id, status) FROM stdin;
1	1	1	COMPLETE
1	2	1	SKIP
\.


--
-- Data for Name: roadmap_followers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roadmap_followers (user_id, roadmap_id) FROM stdin;
1	1
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
3	f
\.


--
-- Name: checkpoints_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.checkpoints_id_seq', 6, true);


--
-- Name: roadmaps_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roadmaps_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- PostgreSQL database dump complete
--

