
 INSERT INTO users (id, name, email, password) VALUES
    ("1", "user1", "user1@example.com", "$2a$10$CwTycUXWue0Thq9StjUM0uJ5ISzSZYw1UChgr6FwNiMduX.O/wrxC"), -- password: password1
    ("2", "user2", "user2@example.com", "$2a$10$uBtZ5h3uZ/1c8XseQf/VCeNWDzKfjd6iKgK6tvQrFXy8EwOMCNOlO"), -- password: password2
    ("3", "user3", "user3@example.com", "$2a$10$TrF.0Cjv8HqQ3sGd.C/PUOejmL4DpWJ10.DABKu1.4oJ3d5v/0B2y"), -- password: password3
    ("4", "user4", "user4@example.com", "$2a$10$ZnQ3e5XZo2w1xE6u6cZgXuR/g3hx8E6s5xJnB7Ih7Wz8De9gNhJxq"), -- password: password4
    ("5", "user5", "user5@example.com", "$2a$10$zQ1kQW/9lHsMsJLUbOFYTeXKo0/7yJ8WhNdF27u4i7AukxOS2jDh6"), -- password: password5
    ("6", "user6", "user6@example.com", "$2a$10$ljzRA8EbR3NfRkN5UBJHse2QuPda7RvkgAsptXG5THMRt6gMKgj62"), -- password: password6
    ("7", "user7", "user7@example.com", "$2a$10$J3Z9Xr.wO0iX68xL3BB9peHhY4a3qx4FbdAc8rOz8A.4S5U3mnU8O"), -- password: password7
    ("8", "user8", "user8@example.com", "$2a$10$FxWyM6.c5lF1gMlK1QiDxe/bWgN5wA8Whd8YAjMiT6wyo8QqgHrmG"), -- password: password8
    ("9", "user9", "user9@example.com", "$2a$10$HbXCuXK3P/xv5ZnAAEPEh.mX9MDVp5xgox2Q5PPf8U/8C9SDbDwHe"), -- password: password9
    ("10", "user10", "user10@example.com", "$2a$10$/0FLyFQUCCXybFSq9A06Je2FlUAnPPoJzC6Vbk4rIMg4Yze/EX3aW"); -- password: password10

INSERT INTO posts (id, title, user_id, content, created_at)
VALUES
    ('post1', 'Post Title 1', 'user1', 'Post content 1', DATETIME('now')),
    ('post2', 'Post Title 2', 'user2', 'Post content 2', DATETIME('now')),
    ('post3', 'Post Title 3', 'user3', 'Post content 3', DATETIME('now')),
    ('post4', 'Post Title 4', 'user4', 'Post content 4', DATETIME('now')),
    ('post5', 'Post Title 5', 'user5', 'Post content 5', DATETIME('now')),
    ('post6', 'Post Title 6', 'user6', 'Post content 6', DATETIME('now')),
    ('post7', 'Post Title 7', 'user7', 'Post content 7', DATETIME('now')),
    ('post8', 'Post Title 8', 'user8', 'Post content 8', DATETIME('now')),
    ('post9', 'Post Title 9', 'user9', 'Post content 9', DATETIME('now')),
    ('post10', 'Post Title 10', 'user10', 'Post content 10', DATETIME('now'));

-- Likes_Dislikes tablosuna örnek veri ekleme
INSERT INTO likes_dislikes (id, post_id, user_id, is_like)
VALUES
    ('like1', 'post1', 'user1', 1),
    ('like2', 'post2', 'user2', 0),
    ('like3', 'post3', 'user3', 1),
    ('like4', 'post4', 'user4', 0),
    ('like5', 'post5', 'user5', 1),
    ('like6', 'post6', 'user6', 0),
    ('like7', 'post7', 'user7', 1),
    ('like8', 'post8', 'user8', 0),
    ('like9', 'post9', 'user9', 1),
    ('like10', 'post10', 'user10', 0);

-- Commits tablosuna örnek veri ekleme
INSERT INTO commits (id, user_id, post_id, content, created_at)
VALUES
    ('commit1', 'user1', 'post1', 'Commit content 1', DATETIME('now')),
    ('commit2', 'user2', 'post2', 'Commit content 2', DATETIME('now')),
    ('commit3', 'user3', 'post3', 'Commit content 3', DATETIME('now')),
    ('commit4', 'user4', 'post4', 'Commit content 4', DATETIME('now')),
    ('commit5', 'user5', 'post5', 'Commit content 5', DATETIME('now')),
    ('commit6', 'user6', 'post6', 'Commit content 6', DATETIME('now')),
    ('commit7', 'user7', 'post7', 'Commit content 7', DATETIME('now')),
    ('commit8', 'user8', 'post8', 'Commit content 8', DATETIME('now')),
    ('commit9', 'user9', 'post9', 'Commit content 9', DATETIME('now')),
    ('commit10', 'user10', 'post10', 'Commit content 10', DATETIME('now'));

-- Categories tablosuna örnek veri ekleme
INSERT INTO categories (id, name, post_id, go, php, python, c, csharp, cplus, rust, java, javascript, html, css, chat)
VALUES
    ('category1', 'Category 1', 'post1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'),
    ('category2', 'Category 2', 'post2', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'),
    ('category3', 'Category 3', 'post3', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'),
    ('category4', 'Category 4', 'post4', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'),
    ('category5', 'Category 5', 'post5', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'),
    ('category6', 'Category 6', 'post6', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'),
    ('category7', 'Category 7', 'post7', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'),
    ('category8', 'Category 8', 'post8', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'),
    ('category9', 'Category 9', 'post9', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'),
    ('category10', 'Category 10', 'post10', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1');

