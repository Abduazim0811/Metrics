CREATE TABLE IF NOT EXISTS book (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) NOT NULL,
    author VARCHAR(50) NOT NULL,
    published_date VARCHAR(50),
    isbn INT
);
