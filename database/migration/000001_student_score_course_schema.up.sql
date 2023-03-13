CREATE TABLE "student" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" bigint NOT NULL,
  "student_number" varchar NOT NULL,
  "updated_time" timestamptz NOT NULL DEFAULT (now()),
  "created_time" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "score" (
  "id" bigserial PRIMARY KEY,
  "score" bigint NOT NULL,
  "student_id" bigint NOT NULL,
  "course_id" bigint NOT NULL,
  "updated_time" timestamptz NOT NULL DEFAULT (now()),
  "created_time" timestamptz NOT NULL DEFAULT (now())
    -- 另一種方式新增
    -- PRIMARY KEY (student_id, course_id)
    -- CONSTRAINT fk_student FOREIGN KEY(student_id) REFERENCES student(id)
    -- CONSTRAINT fk_course FOREIGN KEY(course_id) REFERENCES course(id)
);

CREATE TABLE "course" (
  "id" bigserial PRIMARY KEY,
  "subject" varchar NOT NULL,
  "subject_id" varchar NOT NULL,
  "updated_time" timestamptz NOT NULL DEFAULT (now()),
  "created_time" timestamptz NOT NULL DEFAULT (now())
);
-- 另一種方式新增
ALTER TABLE "score" ADD FOREIGN KEY ("student_id") REFERENCES "student" ("id");
ALTER TABLE "score" ADD FOREIGN KEY ("course_id") REFERENCES "course" ("id");