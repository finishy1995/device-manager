CREATE TABLE device_camera_data (
    id bigint(20) NOT NULL AUTO_INCREMENT,
    device_sn char(15) NOT NULL,
    timestamp BIGINT NOT NULL,
    is_fixed TINYINT(1) NOT NULL DEFAULT 0,
    battery_level SMALLINT UNSIGNED NOT NULL,
    rotation_x FLOAT NOT NULL,
    rotation_y FLOAT NOT NULL,
    rotation_z FLOAT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_device_time` (`device_sn`,`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='device camera data table';