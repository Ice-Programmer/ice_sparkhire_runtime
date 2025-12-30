-- 创建库
create database if not exists sparkhire;

-- 切换库
use sparkhire;

-- 用户表
create table if not exists `user`
(
    `id`          bigint                                   not null comment 'id' primary key,
    `username`    varchar(128)                             not null comment '用户昵称',
    `user_avatar` varchar(128)                             not null comment '用户头像',
    `email`       varchar(256)                             not null comment '邮箱',
    `gender`      tinyint     default 1                    not null comment '0-女 1-男',
    `profile`     text                                     null comment '自我评价',
    `user_role`   tinyint     default 0                    not null comment '用户角色（1-visitor 2-candidate 3-HR 4-admin）',
    `status`      tinyint     default 0                    not null comment '用户状态(0-正常 1-封禁)',
    `created_at`  datetime(3) default current_timestamp(3) not null comment '创建时间',
    `updated_at`  datetime(3) default current_timestamp(3) not null on update current_timestamp(3) comment '更新时间',
    `deleted_at`  datetime(3)                              null comment '删除时间',
    unique key uk_email (`email`),
    index idx_deleted_at (`deleted_at`)
) comment '用户' collate = utf8mb4_unicode_ci;

-- 专业表
create table if not exists `major`
(
    `id`         bigint                             not null comment 'id' primary key,
    `major_name` varchar(256)                       not null comment '专业名称',
    `post_num`   int      default 0                 not null comment '相关数量',
    `created_at` datetime default CURRENT_TIMESTAMP not null comment '创建时间'
) comment '专业' collate = utf8mb4_unicode_ci;

-- 行业表
create table if not exists `industry`
(
    `id`            bigint                             not null comment 'id' primary key,
    `industry_name` varchar(256)                       not null comment '行业名称',
    `industry_type` bigint                             not null comment '行业类型',
    `post_num`      bigint   default 0                 not null comment '相关数量',
    `created_at`    datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`    datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '行业' collate = utf8mb4_unicode_ci;

-- 行业类型表
create table if not exists `industry_type`
(
    `id`         bigint                             not null comment 'id' primary key,
    `name`       varchar(256)                       not null comment '行业类型名称',
    `post_num`   bigint   default 0                 not null comment '相关数量',
    `created_at` datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at` datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '行业类型' collate = utf8mb4_unicode_ci;

-- 学校
create table if not exists `school`
(
    `id`          bigint                             not null comment 'id' primary key,
    `school_name` varchar(256)                       not null comment '学校名称',
    `post_num`    bigint   default 0                 not null comment '相关数量',
    `created_at`  datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`  datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '学校' collate = utf8mb4_unicode_ci;

-- 求职者表
create table if not exists `candidate`
(
    `id`                  bigint                                 not null comment 'id' primary key,
    `user_id`             bigint                                 not null comment '用户id',
    `age`                 int          default 20                not null comment '年龄',
    `education`           int          default 1                 not null comment '最高学历(1-本科 2-研究生 3-博士生 4-大专 5-高中 6-高中以下)',
    `phone`               varchar(128) default ''                not null comment '联系方式',
    `graduation_year`     int                                    not null comment '毕业年份',
    `job_status`          tinyint                                not null comment '求职状态',
    `first_geo_level_id`  bigint                                 not null comment '一级地理位置 id',
    `second_geo_level_id` bigint                                 not null comment '二级地理位置 id',
    `third_geo_level_id`  bigint                                 not null comment '三级地理位置 id',
    `forth_geo_level_id`  bigint                                 not null comment '四级地理位置 id',
    `address`             varchar(512) default ''                not null comment '具体地址',
    `latitude`            decimal(10, 7)                         null comment '纬度',
    `longitude`           decimal(10, 7)                         null comment '经度',
    `created_at`          datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`          datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`          datetime                               null comment '删除时间',
    unique key `uk_user_id` (`user_id`, `deleted_at`)
) comment '求职者' collate = utf8mb4_unicode_ci;

-- 一级地理位置
create table if not exists `geo_first_level`
(
    `id`         bigint auto_increment                  not null comment 'id' primary key,
    `geo_name`   varchar(128)                           not null comment '地理名称',
    `code`       varchar(128) default ''                not null comment 'code',
    `created_at` datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at` datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '一级地理位置' collate = utf8mb4_unicode_ci;


-- 二级地理位置
create table if not exists `geo_second_level`
(
    `id`                 bigint auto_increment                  not null comment 'id' primary key,
    `geo_name`           varchar(128)                           not null comment '地理名称',
    `code`               varchar(128) default ''                not null comment 'code',
    `first_geo_level_id` bigint                                 not null comment '一级地理位置 id',
    `created_at`         datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`         datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '二级地理位置' collate = utf8mb4_unicode_ci;

-- 三级地理位置
create table if not exists `geo_third_level`
(
    `id`                  bigint auto_increment                  not null comment 'id' primary key,
    `geo_name`            varchar(128)                           not null comment '地理名称',
    `code`                varchar(128) default ''                not null comment 'code',
    `second_geo_level_id` bigint                                 not null comment '二级地理位置 id',
    `created_at`          datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`          datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '三级地理位置' collate = utf8mb4_unicode_ci;

-- 四级地理位置
create table if not exists `geo_forth_level`
(
    `id`                 bigint auto_increment                  not null comment 'id' primary key,
    `geo_name`           varchar(128)                           not null comment '地理名称',
    `code`               varchar(128) default ''                not null comment 'code',
    `third_geo_level_id` bigint                                 not null comment '三级地理位置 id',
    `created_at`         datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`         datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '四级地理位置' collate = utf8mb4_unicode_ci;

-- tag 表
create table if not exists `tag`
(
    `id`             bigint auto_increment comment 'id' primary key,
    `tag_name`       varchar(256)                       not null comment '标签名称',
    `create_user_id` bigint                             not null comment '创建用户 id',
    `created_at`     datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`     datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    unique key `uk_tag_name_del` (`tag_name`)
) comment 'tag 表' collate = utf8mb4_unicode_ci;

-- tag 关系表
create table if not exists `tag_obj_rel`
(
    `id`         bigint auto_increment comment 'id' primary key,
    `tag_id`     bigint                             not null comment 'tag id',
    `obj_id`     bigint                             not null comment 'obj_id',
    `obj_type`   int                                not null comment 'obj type(1-candidate/2-recruitment)',
    `created_at` datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    unique key `uk_tag_obj` (`tag_id`, `obj_id`, `obj_type`),
    index `uk_obj_id_type` (`obj_id`, `obj_type`)
) comment 'tag_obj_rel' collate = utf8mb4_unicode_ci;

-- 教育经历表
create table if not exists `education_experience`
(
    `id`               bigint auto_increment comment 'id' primary key,
    `user_id`          bigint                             not null comment '用户id',
    `school_id`        bigint                             not null comment '学校id',
    `education_status` tinyint                            not null comment '学历类型',
    `begin_year`       int                                not null comment '开始年份',
    `end_year`         int                                not null comment '结束年份',
    `major_id`         bigint                             not null comment '专业id',
    `activity`         text                               null comment '在校经历',
    `created_at`       datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`       datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`       datetime                           null comment '删除时间',
    unique key uk_user_education_tye (`user_id`, `education_status`, `deleted_at`),
    index idx_user_id (`user_id`)
) comment '教育经历' collate = utf8mb4_unicode_ci;

-- 应聘者经历表
create table if not exists `career_experience`
(
    `id`              bigint auto_increment comment 'id' primary key,
    `user_id`         bigint                             not null comment '用户id',
    `experience_name` varchar(256)                       not null comment '经历名称',
    `begin_ts`        bigint   default 0                 not null comment '开始时间',
    `end_ts`          bigint   default 0                 not null comment '结束时间',
    `job_role`        varchar(256)                       not null comment '担任职务',
    `description`     text                               not null comment '经历描述',
    `created_at`      datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`      datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`      datetime                           null comment '删除时间',
    index idx_user_id (`user_id`, `deleted_at`)
) comment '应聘者经历' collate = utf8mb4_unicode_ci;

-- 资格证书表
create table if not exists `qualification`
(
    `id`                 bigint auto_increment comment 'id' primary key,
    `qualification_name` varchar(256)                       not null comment '资格证书名称',
    `qualification_type` int                                not null comment '资格证书类型',
    `created_at`         datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`         datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '资格证书' collate = utf8mb4_unicode_ci;

-- 招聘信息表
create table if not exists `recruitment`
(
    `id`                  bigint auto_increment comment 'id' primary key,
    `name`                varchar(512)                           not null comment '岗位招聘标题',
    `user_id`             bigint                                 not null comment '岗位发布者id',
    `company_id`          bigint                                 not null comment '公司id',
    `position_id`         bigint                                 not null comment '职业id',
    `industry_id`         bigint                                 not null comment '行业id',
    `description`         text                                   not null comment '职位详情',
    `requirement`         text                                   not null comment '职位要求',
    `education_type`      tinyint                                null comment '最低学历要求',
    `job_type`            int                                    not null comment '职业类型（实习、兼职、春招等）',
    `apply_count`         int          default 0                 not null comment '投递次数',
    `favorite_count`      int          default 0                 not null comment '收藏次数',
    `first_geo_level_id`  bigint                                 not null comment '一级地理位置 id',
    `second_geo_level_id` bigint                                 not null comment '二级地理位置 id',
    `third_geo_level_id`  bigint                                 not null comment '三级地理位置 id',
    `forth_geo_level_id`  bigint                                 not null comment '四级地理位置 id',
    `address`             varchar(512) default ''                not null comment '具体地址',
    `latitude`            decimal(10, 7)                         null comment '纬度',
    `longitude`           decimal(10, 7)                         null comment '经度',
    `salary_upper`        int                                    null comment '薪水上限',
    `salary_lower`        int                                    null comment '薪水下限',
    `salary_type`         bigint                                 not null comment '薪水货币类型',
    `status`              tinyint      default 0                 not null comment '招聘状态（0 - 招聘中 1 - 结束招聘）',
    `created_at`          datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`          datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`          datetime                               null comment '删除时间',
    index idx_userId (`user_id`, `deleted_at`),
    index idx_companyId (`company_id`, `deleted_at`)
) comment '招聘信息' collate = utf8mb4_unicode_ci;

-- 公司信息表
create table if not exists `company`
(
    `id`                  bigint auto_increment comment 'id' primary key,
    `company_name`        varchar(256)                            not null comment '公司名称',
    `create_user_id`      bigint                                  not null comment '创建用户 id',
    `description`         text                                    not null comment '公司介绍',
    `favorite_count`      int           default 0                 not null comment '收藏次数',
    `logo`                varchar(256)                            not null comment '公司 Logo',
    `background_img`      varchar(256)                            null comment '公司背景图片',
    `company_img_list`    varchar(1024) default '[]'              null comment '公司图片',
    `industry_id`         bigint                                  not null comment '公司产业',
    `first_geo_level_id`  bigint                                  not null comment '一级地理位置 id',
    `second_geo_level_id` bigint                                  not null comment '二级地理位置 id',
    `third_geo_level_id`  bigint                                  not null comment '三级地理位置 id',
    `forth_geo_level_id`  bigint                                  not null comment '四级地理位置 id',
    `address`             varchar(512)  default ''                not null comment '具体地址',
    `latitude`            decimal(10, 7)                          null comment '纬度',
    `longitude`           decimal(10, 7)                          null comment '经度',
    `created_at`          datetime      default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`          datetime      default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`          datetime                                null comment '删除时间'
) comment '公司信息' collate = utf8mb4_unicode_ci;

