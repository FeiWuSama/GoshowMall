create database if not exists goshow_mall;

use goshow_mall;

CREATE TABLE `permission`
(
    `id`        bigint                                                        NOT NULL AUTO_INCREMENT,
    `code`      varchar(255) CHARACTER SET Utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限编码',
    `type`      tinyint                                                       NOT NULL COMMENT '1:菜单 2:操作',
    `name`      varchar(255) CHARACTER SET Utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限名称',
    `page_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单路径',
    `parent_id` bigint                                                        NOT NULL DEFAULT '-1' COMMENT '父级权限ID',
    `status`    tinyint                                                       NOT NULL COMMENT '1:正常-1:禁用',
    `sort`      int                                                           NOT NULL DEFAULT 1,
    `desc`      varchar(255) CHARACTER SET Utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限描述',
    `create_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime                                                      NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    `create_by` bigint                                                        NOT NULL DEFAULT '0',
    `update_by` bigint                                                        NOT NULL DEFAULT '0',
    PRIMARY KEY (id) USING BTREE,
    UNIQUE KEY idx_code (`code`, `status`) USING BTREE,
    KEY `idx_name` (`name`) USING BTREE,
    KEY `idx_parent_id` (`parent_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC COMMENT ='权限消单表';

CREATE TABLE `admin`
(
    `id`           bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '主键-管理员ID表',
    `name`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名字',
    `nick_name`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
    `mobile`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
    `lark_open_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '飞书OpenID',
    `password`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
    `status`       tinyint                                                       NOT NULL DEFAULT '1' COMMENT '1:正常-1:禁用',
    `create_at`    datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at`    datetime                                                      NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    `create_by`    bigint                                                        NOT NULL DEFAULT '0',
    `update_by`    bigint                                                        NOT NULL DEFAULT '0',
    `sex`          tinyint                                                       NOT NULL DEFAULT 3 COMMENT '3:其他1:男2:女',
    `is_delete`    tinyint                                                       NOT NULL DEFAULT 0,
    PRIMARY KEY (id) USING BTREE,
    UNIQUE KEY idx_mobile (`mobile`) USING BTREE,
    KEY idx_name (`name`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC COMMENT ='管理员表';

CREATE TABLE `role`
(
    `id`          bigint                                                        NOT NULL AUTO_INCREMENT,
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
    `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色描述',
    `status`      tinyint                                                       NOT NULL DEFAULT 1 COMMENT '1:正常 -1:禁用',
    `create_at`   datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at`   datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `create_by`   bigint                                                        NOT NULL DEFAULT 0,
    `update_by`   bigint                                                        NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_name` (`name`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = DYNAMIC COMMENT = '角色表';

CREATE TABLE `role_permission`
(
    `id`            BIGINT   NOT NULL AUTO_INCREMENT,
    `role_id`       BIGINT   NOT NULL COMMENT '角色id',
    `permission_id` BIGINT   NOT NULL COMMENT '权限id',
    `create_at`     DATETIME NOT NULL,
    `update_at`     DATETIME NOT NULL,
    `create_by`     BIGINT   NOT NULL DEFAULT 0,
    `update_by`     BIGINT   NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    KEY `idx_role_id` (`role_id`)

) ENGINE = INNODB
  AUTO_INCREMENT = 1
  CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='角色权限表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `user`
(
    `id`        bigint                                                        NOT NULL,
    `nick_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '',
    `sex`       tinyint                                                       NOT NULL DEFAULT 0 COMMENT '0：男，1：女',
    `password`  varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '',
    `status`    tinyint                                                       NOT NULL DEFAULT 1 COMMENT '1：正常，-1：禁用',
    `avatar`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像地址（云存储key）',
    `create_at` datetime(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `update_at` datetime(0)                                                   NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    `create_by` bigint                                                        NOT NULL DEFAULT 0,
    `update_by` bigint                                                        NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_nick_name` (`nick_name`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `wechat_user`
(
    `id`             bigint                                                        NOT NULL AUTO_INCREMENT,
    `user_id`        bigint                                                        NOT NULL DEFAULT 0 COMMENT '用户id',
    `wechat_user_id` bigint                                                        NOT NULL DEFAULT 0 COMMENT '微信用户uuid',
    `avatar`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
    `nick_name`      varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
    `create_at`      datetime(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `update_at`      datetime(0)                                                   NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_user_id` (`user_id`) USING BTREE,
    UNIQUE INDEX `idx_wechat_user_id` (`wechat_user_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '微信用户表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `mobile_user`
(
    `id`            bigint                                                        NOT NULL AUTO_INCREMENT,
    `user_id`       bigint                                                        NOT NULL DEFAULT 0 COMMENT '用户id',
    `mobile_aes`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'aes加密的手机号',
    `mobile_sha256` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'sha256加密后的手机号',
    `create_at`     datetime(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `update_at`     datetime(0)                                                   NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_mobile` (`mobile_sha256`) USING BTREE,
    UNIQUE INDEX `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '手机用户表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `app_user`
(
    `id`        bigint                                                        NOT NULL AUTO_INCREMENT,
    `user_id`   bigint                                                        NOT NULL DEFAULT 0 COMMENT '用户id',
    `app_code`  int                                                           NOT NULL COMMENT '1000：公众号，1001：小程序',
    `open_id`   varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '微信应用openid',
    `status`    tinyint                                                       NOT NULL DEFAULT 1 COMMENT 'create_at\r\nupdate_at\r\n',
    `create_at` datetime(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `update_at` datetime(0)                                                   NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_user_id_app_code` (`user_id`, `app_code`) USING BTREE,
    UNIQUE INDEX `idx_open_id` (`open_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '应用用户表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `user_goods`
(
    `id`          bigint  NOT NULL AUTO_INCREMENT,
    `user_id`     bigint  NOT NULL DEFAULT 0 COMMENT '用户id',
    `order_id`    bigint  NOT NULL COMMENT '订单id',
    `goods_id`    bigint  NOT NULL COMMENT '商品id',
    `goods_type`  tinyint NOT NULL COMMENT '商品类型',
    `buy_time`    bigint  NOT NULL COMMENT '购买时间',
    `expire_time` bigint  NOT NULL COMMENT '到期时间',
    PRIMARY KEY (`id`),
    INDEX `idx_user_id` (`user_id`) USING BTREE,
    INDEX `idx_goods_id` (`goods_id`) USING BTREE,
    INDEX `idx_expire_time` (`expire_time`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户购买商品表'
  ROW_FORMAT = REDUNDANT;

CREATE TABLE `goods`
(
    `id`             bigint                                                        NOT NULL AUTO_INCREMENT,
    `name`           varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '',
    `cover`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品封面',
    `introduction`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品介绍（云存储地址）',
    `price`          bigint                                                        NOT NULL DEFAULT 0 COMMENT '商品价格',
    `effective_time` tinyint                                                       NOT NULL COMMENT '有效时长 1：1个月，2：半年，3：一年，4：永久',
    `status`         tinyint                                                       NOT NULL DEFAULT 1 COMMENT '1：正常，-1：下架',
    `sale_type`      tinyint                                                       NOT NULL COMMENT '1：免费，2：收费',
    `create_at`      datetime(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `update_at`      datetime(0)                                                   NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    `create_by`      bigint                                                        NOT NULL DEFAULT 0,
    `update_by`      bigint                                                        NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_name` (`name`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '商品表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `goods_catalog`
(
    `id`        bigint                                                       NOT NULL AUTO_INCREMENT,
    `parent_id` bigint                                                       NOT NULL DEFAULT -1,
    `level`     int                                                          NOT NULL DEFAULT 1,
    `name`      varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
    `goods_id`  bigint                                                       NOT NULL,
    `sort`      bigint                                                       NOT NULL DEFAULT 0,
    `update_at` datetime(0)                                                  NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    `update_by` bigint                                                       NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_goods_id` (`goods_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '商品目录表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `goods_catalog_detail`
(
    `id`         bigint                                                        NOT NULL AUTO_INCREMENT,
    `goods_id`   bigint                                                        NOT NULL,
    `catalog_id` bigint                                                        NOT NULL,
    `name`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
    `is_trial`   tinyint                                                       NOT NULL COMMENT '1：试看，-1：付费',
    `status`     tinyint                                                       NOT NULL DEFAULT 1 COMMENT '1：正常，-1：下架',
    `cover`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '封面',
    `detail`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         NOT NULL COMMENT '详情',
    `homework`   text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         NOT NULL COMMENT '作业',
    `sort`       bigint                                                        NOT NULL DEFAULT 0,
    `update_at`  datetime(0)                                                   NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    `update_by`  bigint                                                        NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_goods_id` (`goods_id`) USING BTREE,
    INDEX `idx_catalog_id` (`catalog_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '商品目录详情表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `order`
(
    `id`                   bigint                                                        NOT NULL AUTO_INCREMENT,
    `order_number`         char(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci     NOT NULL COMMENT '订单号',
    `user_id`              bigint                                                        NOT NULL DEFAULT 0 COMMENT '用户id',
    `status`               tinyint                                                       NOT NULL DEFAULT 1 COMMENT '1：正常，-1：下架',
    `order_source`         tinyint                                                       NOT NULL COMMENT '1：用户下单，2：管理后台，3：活动赠送',
    `order_amount`         bigint                                                        NOT NULL COMMENT '订单金额',
    `order_orginal_amount` bigint                                                        NOT NULL COMMENT '订单原始金额',
    `payment_amount`       bigint                                                        NOT NULL COMMENT '订单支付金额',
    `trade_number`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '交易号',
    `inner_trade_number`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内部交易号',
    `order_desc`           mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci   NOT NULL COMMENT '订单描述',
    `payment_at`           bigint                                                        NOT NULL DEFAULT 0 COMMENT '订单交易时间',
    `user_remark`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户备注',
    `confirm_at`           bigint                                                        NULL     DEFAULT NULL COMMENT '收货时间',
    `confirm_type`         tinyint                                                       NULL     DEFAULT NULL COMMENT '收货类型，1：用户确认，2：自动发货',
    `refund_amount`        bigint                                                        NOT NULL COMMENT '订单退款金额',
    `refund_at`            bigint                                                        NULL     DEFAULT NULL COMMENT '退款时间',
    `cancel_at`            bigint                                                        NULL     DEFAULT NULL COMMENT '取消时间',
    `cancel_by`            bigint                                                        NULL COMMENT '退款用户id',
    `cancel_type`          tinyint                                                       NULL COMMENT '1：用户取消，2：超时取消，3：客服取消',
    `cancel_reason`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '取消原因',
    `create_at`            datetime(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `create_by`            bigint                                                        NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_order_number` (`order_number`) USING BTREE,
    INDEX `idx_create_at` (`create_at`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_trade_number` (`trade_number`) USING BTREE,
    INDEX `idx_payment_at` (`payment_at`),
    INDEX `idx_order_source` (`order_source`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '订单表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `order_item`
(
    `id`         bigint                                                NOT NULL AUTO_INCREMENT,
    `order_id`   bigint                                                NOT NULL COMMENT '订单id',
    `user_id`    bigint                                                NOT NULL DEFAULT 0 COMMENT '用户id',
    `goods_id`   bigint                                                NOT NULL,
    `goods_type` tinyint                                               NOT NULL DEFAULT 1 COMMENT '1：课程，2：文本',
    `quantity`   int                                                   NOT NULL DEFAULT 1 COMMENT '商品数量',
    `goods_snap` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商品快照',
    PRIMARY KEY (`id`),
    INDEX `idx_order_id` (`order_id`),
    INDEX `idx_goods_id` (`goods_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '订单对象表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `sms_template`
(
    `id`              bigint                                                        NOT NULL AUTO_INCREMENT,
    `scene_code`      varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '场景编码',
    `sign_name`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '短信签名',
    `platform_tmp_id` int                                                           NOT NULL COMMENT '短信平台的模板id',
    `tmp_content`     varchar(255)                                                  NOT NULL COMMENT '短信模板内容',
    `status`          tinyint                                                       NOT NULL DEFAULT 1 COMMENT '1：正常，-1：下架',
    `create_at`       datetime(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `update_at`       datetime(0)                                                   NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    `user_id`         bigint                                                        NOT NULL DEFAULT 0,
    `platform`        char CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_scene_code` (`scene_code`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '短信模板表'
  ROW_FORMAT = DYNAMIC;

CREATE TABLE `upload_file`
(
    `id`        bigint                                                        NOT NULL AUTO_INCREMENT,
    `scene`     char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci     NOT NULL COMMENT '业务场景',
    `file_key`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '云存储key',
    `user_id`   bigint                                                        NOT NULL DEFAULT 0,
    `user_type` tinyint                                                       NOT NULL COMMENT '1：普通用户，2：管理员',
    `file_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci      NOT NULL COMMENT '文件类型',
    `file_size` bigint                                                        NOT NULL COMMENT '文件字节数',
    `file_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件名',
    `upload_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '上传ip',
    `create_at` datetime(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_file_key` (`file_key`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '云对象存储上传文件表';