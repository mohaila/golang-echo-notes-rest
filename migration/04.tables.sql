CREATE TABLE go.notes
(
  id serial NOT NULL,
  title text NOT NULL,
  description text NOT NULL,
  CONSTRAINT pk_notes PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE go.notes
  OWNER TO go;