/* Database */
DROP DATABASE IF EXISTS kudos;
CREATE DATABASE IF NOT EXISTS kudos;
USE kudos;

/* Tables */
CREATE TABLE IF NOT EXISTS person (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    points INT NOT NULL DEFAULT 0,
    reais INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS kudos (
    id SERIAL PRIMARY KEY,
    personIdFK INT NOT NULL,
    name VARCHAR(200) NOT NULL,
    CONSTRAINT kudos_person_id_fk 
        FOREIGN KEY (personIdFK) REFERENCES person (id)
);

/* Inserts */
INSERT INTO person (id, name, points, reais) VALUES 
    (DEFAULT, 'Felipe', DEFAULT, DEFAULT),
    (DEFAULT, 'Jonas', DEFAULT, DEFAULT),
    (DEFAULT, 'Valdinei', DEFAULT, DEFAULT),
    (DEFAULT, 'João', DEFAULT, DEFAULT);

-- INSERT INTO kudos (id, personIdFK, name) VALUES 
--     (DEFAULT, 1, 'SUPER'),
--     (DEFAULT, 1, 'GREAT'),
--     (DEFAULT, 1, 'GOOD'),
--     (DEFAULT, 1, 'NICE'),
--     (DEFAULT, 1, 'OK'),
--     (DEFAULT, 2, 'SUPER'),
--     (DEFAULT, 2, 'GREAT'),
--     (DEFAULT, 2, 'GOOD'),
--     (DEFAULT, 2, 'NICE'),
--     (DEFAULT, 3, 'OK');

/* Selects */
SELECT * FROM person;
SELECT * FROM kudos;
SELECT person.id, person.name, COUNT(kudos.id) AS qtd_kudos 
    FROM person INNER JOIN kudos ON (person.id = kudos.personIdFK) 
    GROUP BY person.id;
SELECT person.id, person.name, kudos.name FROM person 
    INNER JOIN kudos ON (person.id = kudos.personIdFK)
    WHERE person.id = 1;

/* Updates */
UPDATE person SET points = 0 WHERE id = 0;
UPDATE person SET reais = 0 WHERE id = 0;

-- /* Truncates */
-- SET FOREIGN_KEY_CHECKS = 0;
-- TRUNCATE TABLE kudos;
-- TRUNCATE TABLE person;
-- SET FOREIGN_KEY_CHECKS = 1;

/* DELETES */
-- DELETE FROM kudos WHERE id != 0;
-- DELETE FROM person WHERE id != 0;
