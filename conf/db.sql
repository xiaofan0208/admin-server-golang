

CREATE TABLE backend_user (
	`id`           int(11) NOT NULL AUTO_INCREMENT,
	`username`     varchar(32) NOT NULL      COMMENT '管理员名',
	`password`     varchar(255) NOT NULL     COMMENT '密码',
	`remember_me`  TINYINT(1)                COMMENT '记住我',
	`created`      int(11)                   COMMENT '创建时间',
	`status`       TINYINT(1)  DEFAULT '0'   COMMENT '状态 1:正常 2:删除',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;