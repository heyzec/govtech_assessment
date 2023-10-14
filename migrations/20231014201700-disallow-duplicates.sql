
-- +migrate Up
ALTER TABLE students ADD CONSTRAINT unique_student_email UNIQUE (email);
ALTER TABLE teachers ADD CONSTRAINT unique_teacher_email UNIQUE (email);
ALTER TABLE students_teachers ADD CONSTRAINT unique_student_id_teacher_id UNIQUE (student_id, teacher_id);

-- +migrate Down
ALTER TABLE students_teachers DROP CONSTRAINT unique_student_id_teacher_id;
ALTER TABLE teachers DROP CONSTRAINT unique_teacher_email;
ALTER TABLE students DROP CONSTRAINT unique_student_email;
