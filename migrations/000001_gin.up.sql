CREATE TABLE IF NOT EXISTS book (
    id VARCHAR(200) PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    author VARCHAR(50) NOT NULL,
    published_date VARCHAR(50),
    isbn INT
);
