CREATE TABLE IF NOT EXISTS character (
    id SMALLSERIAL NOT NULL
        CONSTRAINT character_pk
            PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

INSERT INTO character (id, name, created_at) VALUES (1, 'Elesis', '2023-01-09 00:00:00+07:00');
INSERT INTO character (id, name, created_at) VALUES (2, 'Lire', '2023-01-09 00:00:00+07:00');
INSERT INTO character (id, name, created_at) VALUES (3, 'Arme', '2023-01-09 00:00:00+07:00');
INSERT INTO character (id, name, created_at) VALUES (4, 'Lass', '2023-01-09 00:00:00+07:00');
INSERT INTO character (id, name, created_at) VALUES (5, 'Ryan', '2023-01-09 00:00:00+07:00');
INSERT INTO character (id, name, created_at) VALUES (6, 'Ronan', '2023-01-09 00:00:00+07:00');
INSERT INTO character (id, name, created_at) VALUES (7, 'Amy', '2023-01-09 00:00:00+07:00');
INSERT INTO character (id, name, created_at) VALUES (8, 'Jin', '2023-01-09 00:00:00+07:00');
INSERT INTO character (id, name, created_at) VALUES (9, 'Sieghart', '2023-01-09 00:00:00+07:00');
INSERT INTO character (id, name, created_at) VALUES (10, 'Mari', '2023-01-09 00:00:00+07:00');

-- reset your sequence here because we've
-- inserted data into the tables manually,
-- so the sequence is not incremented

SELECT SETVAL(
    ( SELECT PG_GET_SERIAL_SEQUENCE('character', 'id') ),
    ( SELECT MAX(id) FROM character )
);
