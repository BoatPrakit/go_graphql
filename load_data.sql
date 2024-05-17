INSERT INTO contact (contact_id, name, first_name, last_name, gender_id, dob, email, phone, address, photo_path, created_at, created_by)
VALUES
(1, 'John Doe', 'John', 'Doe', 1, '1990-01-01', 'john@example.com', '123-456-7890', '123 Main St', 'path/to/photo.jpg', '2024-04-16 12:00:00', 'Admin'),
(2, 'Jane Smith', 'Jane', 'Smith', 2, '1992-05-15', 'jane@example.com', '987-654-3210', '456 Oak St', 'path/to/photo2.jpg', '2024-04-16 12:30:00', 'Admin');
-- INSERT INTO contact (name,first_name,last_name,gender_id) VALUES ("example_name","a","d",1);
-- INSERT INTO contact (name,first_name,last_name,gender_id) VALUES ("example_name","b","e",2);
-- INSERT INTO contact (name,first_name,last_name,gender_id) VALUES ("example_name","c","f",3);

-- Generate 100 random contacts
-- INSERT INTO contact (name, first_name, last_name, gender_id, dob, email, phone, address, photo_path)
-- SELECT
--   SUBSTR(RANDOMBLOB(10), 3, 10) || ' ' || SUBSTR(RANDOMBLOB(10), 3, 10),
--   SUBSTR(RANDOMBLOB(10), 3, 10),
--   SUBSTR(RANDOMBLOB(10), 3, 10),
--   RANDOM() % 3 + 1,
--   STRFTIME('%Y-%m-%d', 'now', '-' || RANDOM() % 30 || ' days'),
--   SUBSTR(RANDOMBLOB(20), 3, 20) || '@' || SUBSTR(RANDOMBLOB(10), 3, 10) || '.com',
--   '+' || CAST(RANDOM() * 1000000000 AS TEXT) || '-' || CAST(RANDOM() * 10000000 AS TEXT),
--   SUBSTR(RANDOMBLOB(30), 3, 30) || ', ' || SUBSTR(RANDOMBLOB(10), 3, 10) || ', ' || SUBSTR(RANDOMBLOB(10), 3, 10),
--   'https://placeimg.com/640/480/people'
-- FROM generate_series(1, 100);

-- -- Update created_at timestamp for all rows
-- UPDATE contact SET created_at = STRFTIME('%Y-%m-%d %H:%M:%S', 'now');

-- Insert data into the `users` table
INSERT INTO users (user_id, email, password)
VALUES
('user1', 'user1@example.com', 'password1'),
('user2', 'user2@example.com', 'password2'),
('user3', 'user3@example.com', 'password3');
