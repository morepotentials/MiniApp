INSERT INTO users (user_type, first_name, last_name, email) VALUES ('client', 'Erik', 'Haight', 'erik(at)sojern.com');
INSERT INTO users (user_type, first_name, last_name, email) VALUES ('client', 'Paulo', 'Engelke', 'peengelke(at)gmail.com');
INSERT INTO users (user_type, first_name, last_name, email) VALUES ('broker', 'Paul', 'Brinkmann', 'brinkmann(at)gmail.com');

SELECT id FROM users WHERE email = 'peengelke(at)gmail.com';
SELECT id FROM users WHERE email = 'brinkmann(at)gmail.com';

INSERT INTO clients (client_id, broker_id) VALUES (2, 3);
INSERT INTO clients (client_id, broker_id) VALUES (1, 3);

SELECT client_id FROM clients WHERE broker_id = 3;

SELECT id, first_name, last_name, email FROM users U RIGHT JOIN clients C
ON U.id = C.client_id
WHERE C.broker_id = 3;

SELECT id,first_name,last_name,email FROM users