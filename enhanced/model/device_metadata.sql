CREATE TABLE `device_metadata` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `device_sn` char(15) NOT NULL COMMENT 'device sn',
  `param_type` smallint(5) NOT NULL COMMENT 'param type',
  `param_value` varchar(32) DEFAULT NULL COMMENT 'param value',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_device_param` (`device_sn`,`param_type`),
  KEY `idx_device_sn` (`device_sn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='device metadata table';