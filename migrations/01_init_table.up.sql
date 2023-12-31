CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS user_roles CASCADE;
DROP TABLE IF EXISTS acl CASCADE;

CREATE TABLE users (
    user_id UUID PRIMARY KEY  DEFAULT uuid_generate_v4(),
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    role_id UUID,
    name VARCHAR(255),
    address VARCHAR(255),
    phone VARCHAR(255),
    profile_picture VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMPTZ,
    updated_by UUID
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (role_id) REFERENCES UserRoles(role_id)
);

CREATE TABLE user_roles (
    role_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_name VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    updated_by UUID
    deleted_at TIMESTAMPTZ
);

CREATE TABLE acl (
    acl_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id UUID,
    name VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    updated_by UUID,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (role_id) REFERENCES Roles(role_id)
);