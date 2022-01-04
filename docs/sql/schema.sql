 resource | CREATE TABLE `resource` (
  `id` char(64) CHARACTER SET latin1 NOT NULL,
  `vendor` tinyint(1) NOT NULL,
  `region` varchar(64) CHARACTER SET latin1 NOT NULL,
  `zone` varchar(64) CHARACTER SET latin1 NOT NULL,
  `create_at` bigint(13) NOT NULL,
  `expire_at` bigint(13) DEFAULT NULL,
  `category` varchar(64) CHARACTER SET latin1 NOT NULL,
  `type` varchar(120) CHARACTER SET latin1 NOT NULL,
  `instance_id` varchar(120) CHARACTER SET latin1 NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` varchar(255) CHARACTER SET latin1 NOT NULL,
  `update_at` bigint(13) DEFAULT NULL,
  `sync_at` bigint(13) DEFAULT NULL,
  `sync_accout` varchar(255) CHARACTER SET latin1 DEFAULT NULL,
  `public_ip` varchar(64) CHARACTER SET latin1 DEFAULT NULL,
  `private_ip` varchar(64) CHARACTER SET latin1 DEFAULT NULL,
  `pay_type` varchar(255) CHARACTER SET latin1 DEFAULT NULL,
  `describe_hash` varchar(255) NOT NULL,
  `resource_hash` varchar(255) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `instance_id` (`vendor`,`instance_id`) USING BTREE,
  KEY `name` (`name`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `private_ip` (`public_ip`) USING BTREE,
  KEY `public_ip` (`public_ip`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

CREATE TABLE `host` (
  `resource_id` varchar(64) NOT NULL,
  `cpu` tinyint(4) NOT NULL,
  `memory` int(13) NOT NULL,
  `gpu_amount` tinyint(4) DEFAULT NULL,
  `gpu_spec` varchar(255) DEFAULT NULL,
  `os_type` varchar(255) DEFAULT NULL,
  `os_name` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL,
  `serial_number` varchar(120) DEFAULT NULL,
  `image_id` char(64) DEFAULT NULL,
  `internet_max_bandwidth_out` int(10) DEFAULT NULL,
  `internet_max_bandwidth_in` int(10) DEFAULT NULL,
  `key_pair_name` varchar(255) DEFAULT NULL,
  `security_groups` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`resource_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `rds` (
  `engine_type` varchar(255) NOT NULL,
  `engine_version` varchar(255) NOT NULL,
  `instance_class` varchar(255) NOT NULL,
  `class_type` varchar(255) NOT NULL,
  `export_type` varchar(255) NOT NULL,
  `network_type` varchar(255) NOT NULL,
  `type` varchar(255) NOT NULL,
  `cpu` int(11) NOT NULL,
  `memory` int(11) NOT NULL,
  `db_max_quantity` int(11) NOT NULL,
  `account_max_quantity` int(11) NOT NULL,
  `max_connections` int(11) NOT NULL,
  `max_iops` int(11) NOT NULL,
  `collation` varchar(255) NOT NULL,
  `time_zone` varchar(64) NOT NULL,
  `storage_capacity` int(11) NOT NULL,
  `storage_type` varchar(255) NOT NULL,
  `security_ip_mode` varchar(255) NOT NULL,
  `security_ip_list` text NOT NULL,
  `connection_mode` varchar(255) NOT NULL,
  `ip_type` varchar(255) NOT NULL,
  `lock_mode` varchar(255) NOT NULL,
  `lock_reason` varchar(255) NOT NULL,
  `deploy_mode` varchar(255) NOT NULL,
  `port` int(11) NOT NULL,
  `extra` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1

 CREATE TABLE `secret` (
  `id` varchar(64) NOT NULL,
  `create_at` bigint(13) NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
  `vendor` tinyint(1) NOT NULL,
  `address` varchar(255) NOT NULL,
  `allow_regions` text NOT NULL,
  `crendential_type` tinyint(1) NOT NULL,
  `api_key` varchar(255) NOT NULL,
  `api_secret` text NOT NULL,
  `request_rate` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_key` (`api_key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `task` (
  `id` varchar(64) NOT NULL,
  `region` varchar(64) NOT NULL,
  `resource_type` tinyint(1) NOT NULL,
  `secret_id` varchar(64) NOT NULL,
  `secret_desc` text CHARACTER SET utf8mb4 NOT NULL,
  `timeout` int(11) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `message` text NOT NULL,
  `start_at` bigint(20) NOT NULL,
  `end_at` bigint(20) NOT NULL,
  `total_succeed` int(11) NOT NULL,
  `total_failed` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `task_record` (
  `instance_id` varchar(64) NOT NULL COMMENT '实例Id',
  `instance_name` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT '实例名称',
  `is_success` tinyint(2) NOT NULL COMMENT '是否同步成功',
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '通过异常时的信息',
  `task_id` varchar(64) NOT NULL COMMENT 'task id',
  `create_at` bigint(20) NOT NULL COMMENT '记录创建时间',
  KEY `idx_task` (`task_id`) USING HASH,
  KEY `idx_instance` (`instance_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=latin1