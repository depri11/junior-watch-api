-- Active: 1704123704989@@127.0.0.1@5432@user_db_test
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id TEXT PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    role INTEGER,
    name VARCHAR(255),
    address VARCHAR(255),
    phone VARCHAR(255),
    profile_picture VARCHAR(255),
    created_at NUMBER,
    created_by VARCHAR(255),
    updated_at NUMBER,
    updated_by VARCHAR(255),
    deleted_at NUMBER
);