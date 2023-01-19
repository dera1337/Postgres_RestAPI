CREATE TABLE IF NOT EXISTS "class" (
	class_id serial NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT class_pk PRIMARY KEY (class_id)
);

CREATE TABLE IF NOT EXISTS school (
	school_id serial NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT school_pk PRIMARY KEY (school_id)
);

CREATE TABLE IF NOT EXISTS principal (
	principal_id serial NOT NULL,
	"name" varchar NOT NULL,
	school_id int NOT NULL,
	CONSTRAINT principal_pk PRIMARY KEY (principal_id),
	CONSTRAINT principal_fk FOREIGN KEY (school_id) REFERENCES school(school_id)
);

CREATE TABLE IF NOT EXISTS student (
	student_id serial NOT NULL,
	"name" varchar NOT NULL,
	school_id int NOT NULL,
	CONSTRAINT student_pk PRIMARY KEY (student_id),
	CONSTRAINT student_fk FOREIGN KEY (school_id) REFERENCES school(school_id)
);

CREATE TABLE IF NOT EXISTS student_class_xref (
	student_id int NOT NULL,
	class_id int NOT NULL,
	CONSTRAINT student_class_xref_pk PRIMARY KEY (student_id,class_id),
	CONSTRAINT student_class_xref_fk FOREIGN KEY (student_id) REFERENCES student(student_id),
	CONSTRAINT student_class_xref_fk_1 FOREIGN KEY (class_id) REFERENCES "class"(class_id)
);
