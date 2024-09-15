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

INSERT INTO posts (title, content, author_id)
values
("Post 1", "Content 1", 1),
("Post 2", "Content 2", 1),
("Post 3", "Content 3", 2),
("Post 4", "Content 4", 2),
("Post 5", "Content 5", 3),
("Post 6", "Content 6", 3),
("Post 7", "Content 7", 4),
("Post 8", "Content 8", 4),
("Post 9", "Content 9", 5),
("Post 10", "Content 10", 5);