DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS acl;

CREATE TABLE user_roles (
    id string PRIMARY KEY,
    role_name VARCHAR(255),
    created_at INT,
    created_by string,
    updated_at INT,
    updated_by string,
    deleted_at INT
);

CREATE TABLE users (
    id string PRIMARY KEY ,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    role_id string,
    name VARCHAR(255),
    address VARCHAR(255),
    phone VARCHAR(255),
    profile_picture VARCHAR(255),
    created_at INT,
    created_by string,
    updated_at INT,
    updated_by string,
    deleted_at INT,
    FOREIGN KEY (role_id) REFERENCES user_roles(id)
);

CREATE TABLE acl (
    user_id string PRIMARY KEY,
    role_id string,
    name VARCHAR(255),
    created_at INT,
    created_by string,
    updated_at INT,
    updated_by string,
    deleted_at INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES user_roles(id)
);