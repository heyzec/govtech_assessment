
-- +migrate Up
CREATE TABLE teachers (
    id    BIGSERIAL PRIMARY KEY,
    email VARCHAR
);

CREATE TABLE students (
    id    BIGSERIAL PRIMARY KEY,
    email VARCHAR
);

CREATE TABLE students_teachers (
    id    BIGSERIAL PRIMARY KEY,
    student_id BIGINT NOT NULL REFERENCES students,
    teacher_id BIGINT NOT NULL REFERENCES teachers
);
-- +migrate Down
DROP TABLE students;
DROP TABLE teachers;
DROP TABLE students_teachers;
