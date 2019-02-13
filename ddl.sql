CREATE TABLE groups
(
	id serial NOT NULL,
	name text NOT NULL,
	description text NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	PRIMARY KEY (id)
) WITHOUT OIDS;


CREATE TABLE group_members
(
	id bigserial NOT NULL,
	user_id int NOT NULL,
	group_id int NOT NULL,
	role text NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	PRIMARY KEY (id),
	CONSTRAINT UQ_group_members_user_group UNIQUE (user_id, group_id)
) WITHOUT OIDS;

CREATE TABLE users
(
	id serial NOT NULL,
	email text UNIQUE,
	password_digest text,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER TABLE group_members
	ADD CONSTRAINT FK_group_members_groups FOREIGN KEY (group_id)
	REFERENCES groups (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE group_members
	ADD CONSTRAINT FK_group_members_users FOREIGN KEY (user_id)
	REFERENCES users (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;
