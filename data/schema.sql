CREATE TABLE edomain (
    id integer NOT NULL,
    domain_id integer DEFAULT 0,
    url character varying(200),
    title character varying(50),
    status integer DEFAULT 1,
    pre_text text,
    post_text text,
    analytics character varying(15),
    header_text text
);

CREATE TABLE forum (
    id integer NOT NULL,
    domain_id integer DEFAULT 0,
    title character varying(50) DEFAULT '',
    body text,
    status smallint DEFAULT 1,
    sorder smallint DEFAULT 1,
    created timestamp NOT NULL,
    oid integer DEFAULT 0,
    updated timestamp,
    topics integer DEFAULT 0,
    group_name character varying(30)
);

CREATE TABLE post (
    id integer NOT NULL,
    pid integer DEFAULT 0,
    forum_id integer DEFAULT 0,
    title character varying(150) DEFAULT '',
    tripcode character varying(10) DEFAULT '',
    body text,
    status smallint DEFAULT 0,
    created timestamp,
    updated timestamp ,
    oid integer DEFAULT 0, 
    topics integer DEFAULT 0
);
