CREATE TABLE IF NOT EXISTS task_tags(
    id_task UUID NOT NULL  REFERENCES Tasks(id),
    id_tags UUID NOT NULL  REFERENCES Tags(id),
    PRIMARY KEY (id_task, id_tags)
)