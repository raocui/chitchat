drop table posts;
drop table threads;
drop table sessions;
drop table users;

use test;
set names utf8;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` varchar(64) NOT NULL UNIQUE,
  `name` varchar(255) NOT NULL UNIQUE COMMENT '用户名',
  `email` varchar(255) NOT NULL UNIQUE COMMENT '邮箱',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `created_at` int(10) NOT NULL COMMENT '创建时间',
  `updated_at` int(10) NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';



CREATE TABLE `sessions` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` varchar(64) NOT NULL UNIQUE,
  `email` varchar(255) NOT NULL UNIQUE COMMENT '邮箱',
  `user_id` int NOT NULL COMMENT '用户id',
  `created_at` int(10) NOT NULL COMMENT '创建时间',
  `updated_at` int(10) NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会话表';


CREATE TABLE `threads` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` varchar(64) NOT NULL UNIQUE,
  `topic` text,
  `user_id` int NOT NULL,
  `created_at` int(10) NOT NULL COMMENT '创建时间',
  `updated_at` int(10) NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `posts` (
  `id` int  NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` varchar(64) NOT NULL UNIQUE,
  `body` text,
  `user_id` int NOT NULL,
  `thread_id` int NOT NULL,
  `created_at` int(10) NOT NULL COMMENT '创建时间',
  `updated_at` int(10) NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `comments` (
  `id` int  NOT NULL AUTO_INCREMENT COMMENT 'id',
  `content` TEXT,
  `user_id` int NOT NULL,
  `post_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
alter table comments add constraint FK_POST_ID foreign key(post_id) REFERENCES posts(id);