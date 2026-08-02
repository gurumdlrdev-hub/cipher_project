
CREATE TABLE cipher_code (
    serialNumber INT PRIMARY KEY,
    codeNumber INT,
    codeCharacter TEXT
);

SELECT * FROM cipher_code;

INSERT INTO cipher_code (serialNumber, codeNumber, codeCharacter)
VALUES (1, 1, '@'),
(2, 2, '#'),
(3, 3, '$'),
(4, 4, '%'),
(5, 5, '*'),
(6, 6, '^'),
(7, 7, '&'),
(8, 8, '\'),
(9, 9, '?'),
(10, 0, '!');