CREATE TABLE cipher_projects (
serialNumber INT PRIMARY KEY,
decodeNumber INT,
code TEXT
);

SELECT * FROM cipher_project;

CREATE TABLE IF NOT EXISTS cipher_project (
    serialnumber SERIAL PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
    decodenumber VARCHAR(10) NOT NULL
);



DROP TABLE IF EXISTS cipher_project CASCADE;

CREATE TABLE cipher_project (
    serialnumber SERIAL PRIMARY KEY,
    decodenumber VARCHAR(10) NOT NULL,
    code VARCHAR(10) NOT NULL
);

SELECT * FROM cipher_project;

\c cipher_project;

DROP TABLE IF EXISTS cipher_project CASCADE;

CREATE TABLE cipher_project (
    serialnumber INT PRIMARY KEY,
    decodenumber VARCHAR(10) NOT NULL,
    code VARCHAR(10) NOT NULL
);

INSERT INTO cipher_project (serialnumber, decodenumber, code) 
VALUES (1405, '0000', '~')
ON CONFLICT DO NOTHING;