CREATE TABLE t_user
(
    id         int NOT NULL AUTO_INCREMENT,
    username   varchar(255) DEFAULT NULL COMMENT '用户名',
    password   varchar(255) DEFAULT NULL COMMENT '密码',
    tel        varchar(20)  DEFAULT NULL COMMENT '手机号',
    created_at datetime     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at datetime     DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    deleted_at datetime     DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表';

CREATE TABLE t_balance
(
    id         int NOT NULL AUTO_INCREMENT,
    user_id    int      DEFAULT 0 COMMENT 'uid',
    balance    int      DEFAULT 0 COMMENT '余额',
    created_at datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at datetime DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    deleted_at datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '余额表';