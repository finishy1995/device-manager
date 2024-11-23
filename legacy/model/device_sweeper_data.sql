CREATE TABLE device_sweeper_data (
    id bigint(20) NOT NULL AUTO_INCREMENT,
    device_sn char(15) NOT NULL,
    timestamp BIGINT NOT NULL,
    is_charging TINYINT(1) NOT NULL DEFAULT 0,
    battery_level SMALLINT UNSIGNED NOT NULL,
    position_x FLOAT NOT NULL,
    position_y FLOAT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_device_time` (`device_sn`,`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='device sweeper data table';