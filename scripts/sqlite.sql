CREATE TABLE "todos"
(
    "id"                INTEGER       NOT NULL,
    "content"           VARCHAR(1024) NOT NULL,
    "priority"          TINYINT       NOT NULL DEFAULT '0',
    "is_remind_at_time" TINYINT       NOT NULL DEFAULT '0',
    "remind_at"         DATETIME      NOT NULL,
    "repeat_method"     VARCHAR(50)   NOT NULL DEFAULT '',
    "repeat_rule"       VARCHAR(256)  NOT NULL DEFAULT '',
    "repeat_end_at"     DATETIME      NOT NULL,
    "category"          VARCHAR(50)   NOT NULL DEFAULT '',
    "remark"            VARCHAR(1024) NOT NULL DEFAULT '',
    "complete"          TINYINT       NOT NULL,
    "created_at"        DATETIME      NOT NULL,
    "updated_at"        DATETIME      NOT NULL,
    PRIMARY KEY ("id")
);
