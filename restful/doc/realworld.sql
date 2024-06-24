create table user (
                      id BIGINT(20) not null auto_increment comment 'id',
                      user_name varchar(128) not null default '' comment '用户名',
                      email varchar(128) not null default '' comment '邮箱',
                      bio varchar(512) not null default '' comment 'bio',
                      del_state TINYINT(4) not null default '0' comment '删除状态 0：未删除 1：已删除',
                      create_time DATETIME not null default CURRENT_TIMESTAMP COMMENT '创建时间',
                      update_time DATETIME not null default CURRENT_TIMESTAMP comment '更新时间',
                      delete_time DATETIME not null default CURRENT_TIMESTAMP comment '删除时间',
                      primary key (id)
)

create table article (
                         id bigint(20) not null auto_increment comment 'id',
                         auther_id BIGINT(20) not null default '0' comment '文章作者',
                         title varchar(128) not null default '' comment '文章标题',
                         description varchar(256) not null default '' comment '文章描述',
                         body LONGTEXT comment '文章内容',
                         del_state TINYINT(4) not null default '0' comment '删除状态 0：未删除 1：已删除',
                         create_time DATETIME not null default CURRENT_TIMESTAMP COMMENT '创建时间',
                         update_time DATETIME not null default CURRENT_TIMESTAMP comment '更新时间',
                         delete_time DATETIME not null default CURRENT_TIMESTAMP comment '删除时间',
                         primary key (id)
)

create table article_favorited (
                                   id bigint(20) not null auto_increment comment 'id',
                                   article_id bigint(20) not null default '0' comment '文章id',
                                   favorite_user_id bigint(20) not null default '0' comment '喜爱用户id',
                                   del_state TINYINT(4) not null default '0' comment '删除状态 0：未删除 1：已删除',
                                   create_time DATETIME not null default CURRENT_TIMESTAMP COMMENT '创建时间',
                                   update_time DATETIME not null default CURRENT_TIMESTAMP comment '更新时间',
                                   delete_time DATETIME not null default CURRENT_TIMESTAMP comment '删除时间',
                                   primary key (id)
)

create table follow (
                        id bigint(20) not null auto_increment comment 'id',
                        user_id bigint(20) not null default '0' comment '用户id',
                        follow_user_id bigint(20) not null default '0' comment '关注用户id',
                        del_state TINYINT(4) not null default '0' comment '删除状态 0：未删除 1：已删除',
                        create_time DATETIME not null default CURRENT_TIMESTAMP COMMENT '创建时间',
                        update_time DATETIME not null default CURRENT_TIMESTAMP comment '更新时间',
                        delete_time DATETIME not null default CURRENT_TIMESTAMP comment '删除时间',
                        primary key (id)
)

create table tag (
                     id bigint(20) not null auto_increment comment 'id',
                     tag_name varchar(256) not null default '' comment '标签名',
                     del_state TINYINT(4) not null default '0' comment '删除状态 0：未删除 1：已删除',
                     create_time DATETIME not null default CURRENT_TIMESTAMP COMMENT '创建时间',
                     update_time DATETIME not null default CURRENT_TIMESTAMP comment '更新时间',
                     delete_time DATETIME not null default CURRENT_TIMESTAMP comment '删除时间',
                     primary key (id)
)

create table article_tag (
                             id bigint(20) not null auto_increment comment 'id',
                             article_id bigint(20) not null default '0' comment '文章id',
                             tag_id bigint(20) not null default '0' comment '标签id',
                             del_state TINYINT(4) not null default '0' comment '删除状态 0：未删除 1：已删除',
                             create_time DATETIME not null default CURRENT_TIMESTAMP COMMENT '创建时间',
                             update_time DATETIME not null default CURRENT_TIMESTAMP comment '更新时间',
                             delete_time DATETIME not null default CURRENT_TIMESTAMP comment '删除时间',
                             primary key (id)
)


create table article_comment (
                                 id bigint(20) not null auto_increment comment 'id',
                                 auther_id BIGINT(20) not null default '0' comment '评论作者',
                                 article_id BIGINT(20) not null default '0' comment '文章id',
                                 body LONGTEXT comment '评论内容',
                                 del_state TINYINT(4) not null default '0' comment '删除状态 0：未删除 1：已删除',
                                 create_time DATETIME not null default CURRENT_TIMESTAMP COMMENT '创建时间',
                                 update_time DATETIME not null default CURRENT_TIMESTAMP comment '更新时间',
                                 delete_time DATETIME not null default CURRENT_TIMESTAMP comment '删除时间',
                                 primary key (id)
)
