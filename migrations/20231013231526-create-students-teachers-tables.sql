
-- +migrate Up
CREATE DOMAIN created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;
CREATE DOMAIN updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;
CREATE DOMAIN deleted_at TIMESTAMPTZ DEFAULT NULL;

CREATE TABLE teachers (
    id    BIGSERIAL PRIMARY KEY,
    email VARCHAR,

    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
);

CREATE TABLE students (
    id    BIGSERIAL PRIMARY KEY,
    email VARCHAR,

    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
);

CREATE TABLE students_teachers (
    id    BIGSERIAL PRIMARY KEY,
    student_id BIGINT NOT NULL REFERENCES students,
    teacher_id BIGINT NOT NULL REFERENCES teachers,

    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
);

-- +migrate Down
DROP TABLE students_teachers;
DROP TABLE students;
DROP TABLE teachers;

DROP DOMAIN deleted_at;
DROP DOMAIN updated_at;
DROP DOMAIN created_at;

