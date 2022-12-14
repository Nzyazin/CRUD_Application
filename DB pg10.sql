PGDMP         ,                z            library    10.22    10.22     ?
           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            ?
           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            ?
           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                       false            ?
           1262    16393    library    DATABASE     ?   CREATE DATABASE library WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'Russian_Russia.1251' LC_CTYPE = 'Russian_Russia.1251';
    DROP DATABASE library;
             postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
             postgres    false            ?
           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                  postgres    false    3                        3079    12924    plpgsql 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
    DROP EXTENSION plpgsql;
                  false                        0    0    EXTENSION plpgsql    COMMENT     @   COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
                       false    1            ?            1259    16409    author    TABLE     z   CREATE TABLE public.author (
    author_id integer NOT NULL,
    author_name character varying(128),
    birthday date
);
    DROP TABLE public.author;
       public         postgres    false    3            ?            1259    16427    books    TABLE     ?   CREATE TABLE public.books (
    book_id integer NOT NULL,
    book_name character varying,
    fk_publishing_house_id integer,
    fk_author_id integer
);
    DROP TABLE public.books;
       public         postgres    false    3            ?            1259    16422    publishing_house    TABLE     ?   CREATE TABLE public.publishing_house (
    publishing_house_id integer NOT NULL,
    ph_name character varying(128),
    city character varying(128)
);
 $   DROP TABLE public.publishing_house;
       public         postgres    false    3            ?
          0    16409    author 
   TABLE DATA               B   COPY public.author (author_id, author_name, birthday) FROM stdin;
    public       postgres    false    196   ?       ?
          0    16427    books 
   TABLE DATA               Y   COPY public.books (book_id, book_name, fk_publishing_house_id, fk_author_id) FROM stdin;
    public       postgres    false    198   g       ?
          0    16422    publishing_house 
   TABLE DATA               N   COPY public.publishing_house (publishing_house_id, ph_name, city) FROM stdin;
    public       postgres    false    197          v
           2606    16413    author author_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.author
    ADD CONSTRAINT author_pkey PRIMARY KEY (author_id);
 <   ALTER TABLE ONLY public.author DROP CONSTRAINT author_pkey;
       public         postgres    false    196            z
           2606    16434    books books_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (book_id);
 :   ALTER TABLE ONLY public.books DROP CONSTRAINT books_pkey;
       public         postgres    false    198            x
           2606    16426 &   publishing_house publishing_house_pkey 
   CONSTRAINT     u   ALTER TABLE ONLY public.publishing_house
    ADD CONSTRAINT publishing_house_pkey PRIMARY KEY (publishing_house_id);
 P   ALTER TABLE ONLY public.publishing_house DROP CONSTRAINT publishing_house_pkey;
       public         postgres    false    197            {
           2606    16435    books book_publ_house    FK CONSTRAINT     ?   ALTER TABLE ONLY public.books
    ADD CONSTRAINT book_publ_house FOREIGN KEY (fk_publishing_house_id) REFERENCES public.publishing_house(publishing_house_id);
 ?   ALTER TABLE ONLY public.books DROP CONSTRAINT book_publ_house;
       public       postgres    false    198    197    2680            |
           2606    16440    books fk_author_book    FK CONSTRAINT     ?   ALTER TABLE ONLY public.books
    ADD CONSTRAINT fk_author_book FOREIGN KEY (fk_author_id) REFERENCES public.author(author_id);
 >   ALTER TABLE ONLY public.books DROP CONSTRAINT fk_author_book;
       public       postgres    false    196    2678    198            ?
   ?   x?-??
?0D?ٯ?ؐmҴ=?'{?Pz	%`	?R??߻????{̰??-???|p?x׊??bV?szĜ?5???5?ʂ??PX???2?j?s??\,?"?Ъ?t?????;2???N??{Jq?Pt
E??t?-Tj?%P???=j ??K*?      ?
   ?  x?eR[N?@???bO?ؤ?rӴ??@T | ??G?Z??p???{?
??(??x<3v)??:4z??Z?D/"ڈ%|???J/????$Ex0?#?X?Y??yԊ???V?s??	?HC??Q?x։1??h??0?ka???=+b?X?F?k??F?o??c??l?2#?,yl?????IO??j??%???<h?Y????t????B?B?<Rt?a?Xg?@G?Y?$???Ȧ1?5??2Z?+O˺C*/v???|r??eԦg????lc???G??,???5???*??Uv?O)e?????uT?n??l??Sv??،:?C??Ad???Bz?ќ?e???gq??S??M???d(M??m?N?Ny???F??????Sw7h3L?Jɽ?6??[?f????3r?ԭ??:?b?p?B?8??T      ?
   ?   x?3???".#N?̒??bN?????r.c΀???????Լ?TN??ĒԢ̼???"?=... ??     