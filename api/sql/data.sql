INSERT INTO users (name, nick,  email, password) 
VALUES
("USER 1","user_1", "user_1@gmail.com", "$2a$10$RVJmIdLQko/VhYDcQaangOBIcMpkgMktUngIGXUzUXzSzd5ZBqSqK" ),
("USER 2","user_2", "user_2@gmail.com", "$2a$10$RVJmIdLQko/VhYDcQaangOBIcMpkgMktUngIGXUzUXzSzd5ZBqSqK" ),
("USER 3","user_3", "user_3@gmail.com", "$2a$10$RVJmIdLQko/VhYDcQaangOBIcMpkgMktUngIGXUzUXzSzd5ZBqSqK" ),
("USER 4","user_4", "user_4@gmail.com", "$2a$10$RVJmIdLQko/VhYDcQaangOBIcMpkgMktUngIGXUzUXzSzd5ZBqSqK" ),
("USER 5","user_5", "user_5@gmail.com", "$2a$10$RVJmIdLQko/VhYDcQaangOBIcMpkgMktUngIGXUzUXzSzd5ZBqSqK" );
INSERT INTO followers (user_id, follower_id)
VALUES
(1,2),
(3,1),
(1,3);