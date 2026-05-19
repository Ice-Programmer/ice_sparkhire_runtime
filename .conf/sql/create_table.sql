-- 创建库
create database if not exists sparkhire;

-- 切换库
use sparkhire;

-- 用户表
create table if not exists `user`
(
    `id`          bigint                                   not null comment 'id' primary key,
    `username`    varchar(128)                             not null comment '用户昵称',
    `user_avatar` varchar(256)                             not null comment '用户头像',
    `email`       varchar(256)                             not null comment '邮箱',
    `gender`      tinyint     default 1                    not null comment '0-女 1-男',
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

alter table `school`
    add column `school_icon` varchar(256) default '' not null comment '学校 icon' after `post_num`;


-- 求职者表
create table if not exists `candidate`
(
    `id`                  bigint                                 not null comment 'id' primary key,
    `user_id`             bigint                                 not null comment '用户id',
    `age`                 int          default 20                not null comment '年龄',
    `profile`             text                                   null comment '简介',
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
    `career_id`           bigint                                 not null comment '职业id',
    `description`         text                                   not null comment '职位详情',
    `requirement`         text                                   not null comment '职位要求',
    `education_type`      tinyint                                null comment '最低学历要求',
    `job_type`            tinyint                                not null comment '职业类型（实习、兼职、春招等）',
    `apply_count`         int          default 0                 not null comment '投递次数',
    `favorite_count`      int          default 0                 not null comment '收藏次数',
    `first_geo_level_id`  bigint                                 not null comment '一级地理位置 id',
    `second_geo_level_id` bigint                                 not null comment '二级地理位置 id',
    `third_geo_level_id`  bigint                                 not null comment '三级地理位置 id',
    `forth_geo_level_id`  bigint                                 not null comment '四级地理位置 id',
    `address`             varchar(512) default ''                not null comment '具体地址',
    `latitude`            decimal(10, 7)                         null comment '纬度',
    `longitude`           decimal(10, 7)                         null comment '经度',
    `salary_upper`        int          default 0                 not null comment '薪水上限',
    `salary_lower`        int          default 0                 not null comment '薪水下限',
    `currency_type`       int          default 1                 not null comment '薪水货币类型',
    `frequency_type`      tinyint      default 1                 not null comment '类型',
    `status`              tinyint      default 0                 not null comment '招聘状态（0 - 招聘中 1 - 结束招聘）',
    `created_at`          datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`          datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`          datetime                               null comment '删除时间',
    index idx_userId (`user_id`, `deleted_at`),
    index idx_companyId (`company_id`, `deleted_at`)
) comment '招聘信息' collate = utf8mb4_unicode_ci;

-- 职位申请与进度主表
create table if not exists `recruitment_application`
(
    `id`             bigint auto_increment comment 'id' primary key,
    `user_id`        bigint                             not null comment '用户 id',
    `recruitment_id` bigint                             not null comment '招聘岗位 id',
    `company_id`     bigint                             not null comment '公司 id (对应 company 表 id)',
    `status`         tinyint  default 1                 not null comment '当前推进状态：1-已投递/新投递, 2-简历筛选中, 3-约面/面试中, 4-待发Offer, 5-已发Offer, 6-已录用/入职, 7-不合适/淘汰, 8-求职者放弃',
    `remark`         varchar(512)                       null comment '最新进度备注（如淘汰原因、面试评语摘要）',
    `created_at`     datetime default CURRENT_TIMESTAMP not null comment '投递时间/创建时间',
    `updated_at`     datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`     datetime                           null comment '删除时间',
    unique key `uk_candidate_recruitment` (`user_id`, `recruitment_id`, `deleted_at`),
    key `idx_recruitment_status` (`recruitment_id`, `status`),
    key `idx_company_id` (`company_id`),
    key `idx_candidate_id` (`user_id`)
) comment '职位申请与进度主表' collate = utf8mb4_unicode_ci;

-- 职位进度流转历史表
create table if not exists `recruitment_application_log`
(
    `id`             bigint auto_increment comment 'id' primary key,
    `application_id` bigint                             not null comment '职位申请 id (对应 job_application 表 id)',
    `from_status`    tinyint  default 0                 not null comment '变更前状态(0表示初始投递)',
    `to_status`      tinyint                            not null comment '变更后状态',
    `operator_id`    bigint                             not null comment '操作人 id (对应 user 表 id，可以是 HR 或系统或求职者自己)',
    `remark`         varchar(1024)                      null comment '阶段操作备注/评语/原因',
    `created_at`     datetime default CURRENT_TIMESTAMP not null comment '操作时间',
    key `idx_application_id` (`application_id`)
) comment '职位进度流转历史表' collate = utf8mb4_unicode_ci;

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

create table if not exists `user_favorite`
(
    `id`          bigint auto_increment comment '主键 id' primary key,
    `user_id`     bigint                                                         not null comment '用户 id',
    `target_type` tinyint                                                        not null comment '收藏目标类型: 1-公司, 2-职位, 3-文章等',
    `target_id`   bigint                                                         not null comment '目标 id (公司id/职位id等)',
    `created_at`  datetime default current_timestamp                             not null comment '创建时间',
    `updated_at`  datetime default current_timestamp on update current_timestamp not null comment '更新时间',
    `deleted_at`  datetime                                                       null comment '删除时间',
    unique key `uk_user_target` (`user_id`, `target_type`, `target_id`, `deleted_at`),
    index `idx_target` (`target_type`, `target_id`)
) comment '用户通用收藏记录表' collate = utf8mb4_unicode_ci;

-- 应聘者期望岗位
create table if not exists `candidate_wish_career`
(
    `id`             bigint auto_increment comment 'id' primary key,
    `user_id`        bigint                             not null comment '用户id',
    `career_id`      bigint   default 0                 not null comment '职业id',
    `salary_upper`   int      default 0                 not null comment '薪水上限',
    `salary_lower`   int      default 0                 not null comment '薪水下限',
    `currency_type`  int      default 1                 not null comment '薪水货币类型',
    `frequency_type` tinyint  default 1                 not null comment '类型',
    `created_at`     datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`     datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`     datetime                           null comment '删除时间',
    unique key uk_user_industry_career (`user_id`, `deleted_at`),
    index idx_user_id (`user_id`, `deleted_at`)
) comment '应聘者期望岗位' collate = utf8mb4_unicode_ci;

-- 职业表
create table if not exists `career`
(
    `id`          bigint auto_increment comment 'id' primary key,
    `career_name` varchar(256)                       not null comment '职业名称',
    `description` varchar(1024)                      null comment '职业介绍',
    `career_icon` varchar(256)                       null comment 'icon',
    `career_type` int                                not null comment '职业类型',
    `created_at`  datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`  datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`  datetime                           null comment '删除时间'
) comment '职业' collate = utf8mb4_unicode_ci;

-- 职业类型表
create table if not exists `career_type`
(
    `id`               bigint auto_increment comment 'id' primary key,
    `career_type_name` varchar(256)                       not null comment '职业类型名称',
    `created_at`       datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`       datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`       datetime                           null comment '删除时间'
) comment '职业类型表' collate = utf8mb4_unicode_ci;

-- 公司评价表
create table if not exists `company_comment`
(
    `id`             bigint auto_increment comment 'id' primary key,
    `company_id`     bigint                             not null comment '公司 id',
    `user_id`        bigint                             not null comment '用户 id',
    `content`        text                               null comment '评论内容',
    `root_id`        bigint   default 0                 not null comment '所属根评论 id，0 表示自身为根',
    `parent_id`      bigint   default 0                 not null comment '被回复帖子，0 为根节点',
    `reply_user_id`  bigint   default 0                 not null comment '被回复用户 id',
    `reply_count`    int      default 0                 not null comment '若是根评论，记录其子评论总数',
    `favorite_count` int      default 0                 not null comment '点赞数量',
    `created_at`     datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`     datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`     datetime                           null comment '删除时间',
    index idx_company_root (company_id, root_id, created_at, deleted_at),
    index idx_root_id (root_id, deleted_at)
) comment '公司评价表' collate = utf8mb4_unicode_ci;

-- 公司福利分类
create table if not exists `company_benefit_category`
(
    `id`         bigint auto_increment comment 'id' primary key,
    `company_id` bigint       not null comment '公司 id',
    `title`      varchar(128) not null comment '福利分组标题',
    `subtitle`   varchar(255)          default null comment '分组说明',
    `sort`       int          not null default 0 comment '排序，越小越靠前',
    `status`     tinyint      not null default 1 comment '状态：1正常，-1禁用',
    `created_at` datetime              default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at` datetime              default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at` datetime     null comment '删除时间',
    key `idx_company_sort` (`company_id`, `sort`),
    key `idx_company_status` (`company_id`, `status`)
) comment '公司福利分类表' collate = utf8mb4_unicode_ci;

-- 公司福利细则
create table if not exists `company_benefit_item`
(
    `id`          bigint auto_increment comment 'id' primary key,
    `category_id` bigint       not null comment '福利分组id',
    `title`       varchar(128) not null comment '福利条目标题',
    `content`     varchar(512)          default null comment '福利条目描述',
    `sort`        int          not null default 0 comment '排序，越小越靠前',
    `status`      tinyint      not null default 1 comment '状态：1正常，-1禁用',
    `created_at`  datetime              default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`  datetime              default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`  datetime     null comment '删除时间',
    key `idx_category_sort` (`category_id`, `sort`),
    key `idx_category_status` (`category_id`, `status`)
) comment '公司福利细则表' collate = utf8mb4_unicode_ci;

-- 面试计划表
create table if not exists `interview_schedule`
(
    `id`                 bigint auto_increment comment 'id' primary key,
    `candidate_id`       bigint                             not null comment '用户 id',
    `creator_id`         bigint                             not null comment '创建 id',
    `recruitment_id`     bigint                             not null comment '招聘 id',
    `company_id`         bigint                             not null comment '公司 id',
    `interview_ts`       bigint                             not null comment '面试开始时间',
    `interview_date`     varchar(64)                        not null comment '面试时间 yyyy-MM-dd',
    `interview_duration` int                                not null comment '面试时间（分钟）',
    `interview_type`     tinyint  default 1                 not null comment '面试类型：1-视频面试, 2-线下面试, 3-电话面试',
    `interview_link`     varchar(512)                       null comment '面试会议链接',
    `status`             tinyint  default 1                 not null comment '状态：1-未开始, 2-进行中, 3-已结束/历史, 4-已取消',
    `created_at`         datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`         datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`         datetime                           null comment '删除时间',
    key `idx_candidate_id_interview_date` (`candidate_id`, `interview_date`),
    key `idx_recruitment_id` (`recruitment_id`)
) comment '面试计划表' collate = utf8mb4_unicode_ci;

-- 社区帖子表
create table if not exists `forum_post`
(
    `id`             bigint auto_increment comment 'id' primary key,
    `user_id`        bigint                             not null comment '创建用户 id',
    `title`          varchar(256)                       not null comment '帖子标题',
    `content`        text                               null comment '内容',
    `favorite_count` int      default 0                 not null comment '收藏次数',
    `view_count`     bigint   default 0                 not null comment '浏览次数',
    `status`         tinyint  default 1                 not null comment '状态：1-正常 2-审核中 3-屏蔽',
    `type`           tinyint  default 1                 not null comment '帖子类型：1-普通 2-置顶 3-精华',
    `created_at`     datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`     datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`     datetime                           null comment '删除时间',
    key `idx_user_id` (`user_id`),
    KEY `idx_status_created` (`status`, `created_at`)
) comment '社区帖子表' collate = utf8mb4_unicode_ci;

-- 用户搜索记录表
create table if not exists `user_search_history`
(
    `id`             bigint auto_increment comment 'id' primary key,
    `user_id`        bigint                             not null comment '创建用户 id',
    `search_content` varchar(256)                       not null comment '搜索内容',
    `type`           tinyint  default 1                 not null comment '搜索类型 1-recruitment 2-post',
    `created_at`     datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `deleted_at`     datetime                           null comment '删除时间'
) comment '用户搜索记录表' collate = utf8mb4_unicode_ci;

-- Chat 会话表
create table if not exists `chat_session`
(
    `id`                bigint primary key auto_increment comment '会话id',
    `session_no`        varchar(64) not null unique comment '会话唯一编号',
    `recruitment_id`    bigint      not null comment '岗位id',
    `candidate_user_id` bigint      not null comment '求职者id',
    `hr_user_id`        bigint      not null comment 'hr用户id',
    `last_message_id`   bigint       default null comment '最后一条消息id',
    `last_message`      varchar(500) default null comment '最后消息内容',
    `last_message_type` tinyint      default 1 comment '最后消息类型',
    `last_message_time` datetime     default null comment '最后消息时间',
    `candidate_unread`  int          default 0 comment '求职者未读数',
    `hr_unread`         int          default 0 comment 'hr未读数',
    `status`            tinyint      default 1 comment '状态 1正常 2已结束 3已屏蔽',
    `create_time`       datetime     default current_timestamp,
    `update_time`       datetime     default current_timestamp on update current_timestamp,
    `deleted_at`        datetime    null comment '删除时间',
    index idx_candidate (`candidate_user_id`),
    index idx_hr (`hr_user_id`),
    index idx_recruitment (`recruitment_id`),
    index idx_last_time (`last_message_time`)
) comment 'Chat 会话表' collate = utf8mb4_unicode_ci;


-- Chat 消息表
create table if not exists `chat_message`
(
    `id`            bigint primary key auto_increment comment '消息id',
    `session_id`    bigint   not null comment '会话id',
    `sender_id`     bigint   not null comment '发送者id',
    `receiver_id`   bigint   not null comment '接收者id',
    `sender_type`   tinyint  not null comment '发送者类型 1候选人 2hr',
    `message_type`  tinyint  not null default 1 comment '1文本 2图片 3文件 4简历 5岗位卡片 6面试邀请 7系统消息',
    `content`       text     null comment '消息内容',
    `is_read`       tinyint  not null default 0 comment '是否已读',
    `revoke_status` tinyint  not null default 0 comment '撤回状态',
    `send_status`   tinyint  not null default 1 comment '1发送中 2发送成功 3发送失败',
    `create_time`   datetime not null default current_timestamp,
    `deleted_at`    datetime null comment '删除时间',
    index idx_session (`session_id`),
    index idx_sender (`sender_id`),
    index idx_receiver (`receiver_id`),
    index idx_create_time (`create_time`)
) comment 'chat 消息表' collate = utf8mb4_unicode_ci;


-- 消息已读记录表
create table `chat_message_read`
(
    `id`         bigint primary key auto_increment,
    `message_id` bigint not null comment '消息id',
    `user_id`    bigint not null comment '用户id',
    `read_time`  datetime default current_timestamp,
    unique key uk_message_user (`message_id`, `user_id`),
    index idx_user (`user_id`)
) comment '消息已读记录表' collate = utf8mb4_unicode_ci;