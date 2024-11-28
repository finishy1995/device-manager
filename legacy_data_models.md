# Legacy (Mysql) Data Models

## Data Model 1: device_metadata

The device_metadata table is designed to store metadata information for devices. Each record corresponds to a specific metadata parameter of a device, identified by a unique combination of the device serial number (device_sn) and parameter type (param_type).

You can see table schema in [./legacy/model/device_metadata.sql](./legacy/model/device_metadata.sql)

### Column Descriptions

| Column Name   | Data Type        | Description                                                                 |
|---------------|------------------|-----------------------------------------------------------------------------|
| `id`          | `bigint(20)`     | Primary key. Auto-incremented unique identifier for each row.              |
| `device_sn`   | `char(15)`       | Device serial number. Acts as a unique device identifier.                  |
| `param_type`  | `smallint(5)`    | Type of parameter associated with the device.                              |
| `param_value` | `varchar(32)`    | Value of the parameter. Can be `NULL`.                                     |
| `create_time` | `timestamp`      | Timestamp when the record was created. Defaults to the current timestamp.  |
| `update_time` | `timestamp`      | Timestamp when the record was last updated. Automatically updated on change. |

### Keys and Indexes

1. **Primary Key**:  
   The `id` column serves as the primary key, ensuring each row is uniquely identified.

2. **Unique Key (`uk_device_param`)**:  
   A combination of `device_sn` and `param_type` is enforced to be unique, ensuring that each device can only have one entry per parameter type.

3. **Index (`idx_device_sn`)**:  
   An index on the `device_sn` column improves query performance when filtering by device serial number.

### Table Characteristics

- **Storage Engine**: InnoDB is used for transaction support and referential integrity.
- **Character Set**: `utf8mb4` ensures support for extended Unicode characters.

### Sample Data

Here are the first 5 rows of data in the `device_metadata` table:

| `id` | `device_sn`      | `param_type` | `param_value` | `create_time`         | `update_time`         |
|------|------------------|--------------|---------------|-----------------------|-----------------------|
| 1   | SN-100000000000  | 0            | 49462         | 2024-11-24 09:19:08   | 2024-11-26 07:02:27   |
| 2   | SN-100000000000  | 1            | 10928         | 2024-11-24 09:19:08   | 2024-11-26 07:02:27   |
| 3   | SN-100000000000  | 2            | 14525         | 2024-11-24 09:19:08   | 2024-11-26 07:02:27   |
| 4   | SN-100000000000  | 3            | 15911         | 2024-11-24 09:19:08   | 2024-11-26 07:02:27   |
| 5   | SN-100000000000  | 4            | 96200         | 2024-11-24 09:19:08   | 2024-11-26 07:02:27   |


---

## Data Model 2: device_camera_data

The `device_camera_data` table is designed to store camera-related data for devices. Each record corresponds to a timestamped data entry for a specific device.

You can see the table schema in [./legacy/model/device_camera_data.sql](./legacy/model/device_camera_data.sql).

### Column Descriptions

| Column Name      | Data Type          | Description                                                                 |
|-------------------|--------------------|-----------------------------------------------------------------------------|
| `id`             | `bigint(20)`       | Primary key. Auto-incremented unique identifier for each row.              |
| `device_sn`      | `char(15)`         | Device serial number. Acts as a unique device identifier.                  |
| `timestamp`      | `BIGINT`           | Epoch timestamp (in milliseconds) of the camera data entry.                |
| `is_fixed`       | `TINYINT(1)`       | Indicates if the device is fixed (1 for true, 0 for false).                |
| `battery_level`  | `SMALLINT UNSIGNED`| Current battery level of the device.                                       |
| `rotation_x`     | `FLOAT`            | Rotation of the device along the X-axis.                                   |
| `rotation_y`     | `FLOAT`            | Rotation of the device along the Y-axis.                                   |
| `rotation_z`     | `FLOAT`            | Rotation of the device along the Z-axis.                                   |

### Keys and Indexes

1. **Primary Key**:  
   The `id` column serves as the primary key, ensuring each row is uniquely identified.

2. **Unique Key (`uk_device_time`)**:  
   A combination of `device_sn` and `timestamp` is enforced to be unique, ensuring that each device can only have one entry per timestamp.

### Table Characteristics

- **Storage Engine**: InnoDB is used for transaction support and referential integrity.
- **Character Set**: `utf8mb4` ensures support for extended Unicode characters.

### Sample Data

Here are some example rows of data in the `device_camera_data` table:

| `id` | `device_sn`      | `timestamp`   | `is_fixed` | `battery_level` | `rotation_x` | `rotation_y` | `rotation_z` |
|------|------------------|---------------|------------|-----------------|--------------|--------------|--------------|
| 1    | SN-100000000000  | 1732438800 | 0          | 2291              | 304.91        | 28.45      | 149.37        |
| 2    | SN-100000000000  | 1732438801 | 0          | 2288              | 304.76       | 27.98        | 149.44       |
| 3    | SN-100000000000  | 1732438802 | 0          | 2279              | 304.35        | 28.39       | 149.15        |
| 4    | SN-100000000000  | 1732438803 | 0          | 2271              | 304.49       |28.21        | 149.31       |
| 5    | SN-100000000000  | 1732438804 | 0          | 2269              | 304.33        | 28.34       | 149.29        |


---

## Data Model 3: device_sweeper_data

The `device_sweeper_data` table is designed to store sweeper-related data for devices. Each record corresponds to a timestamped data entry for a specific device.

You can see the table schema in [./legacy/model/device_sweeper_data.sql](./legacy/model/device_sweeper_data.sql).

### Column Descriptions

| Column Name      | Data Type          | Description                                                                 |
|-------------------|--------------------|-----------------------------------------------------------------------------|
| `id`             | `bigint(20)`       | Primary key. Auto-incremented unique identifier for each row.              |
| `device_sn`      | `char(15)`         | Device serial number. Acts as a unique device identifier.                  |
| `timestamp`      | `BIGINT`           | Epoch timestamp (in milliseconds) of the sweeper data entry.               |
| `is_charging`    | `TINYINT(1)`       | Indicates if the device is charging (1 for true, 0 for false).             |
| `battery_level`  | `SMALLINT UNSIGNED`| Current battery level of the device.                                       |
| `position_x`     | `FLOAT`            | X-coordinate of the device's position.                                     |
| `position_y`     | `FLOAT`            | Y-coordinate of the device's position.                                     |

### Keys and Indexes

1. **Primary Key**:  
   The `id` column serves as the primary key, ensuring each row is uniquely identified.

2. **Unique Key (`uk_device_time`)**:  
   A combination of `device_sn` and `timestamp` is enforced to be unique, ensuring that each device can only have one entry per timestamp.

### Table Characteristics

- **Storage Engine**: InnoDB is used for transaction support and referential integrity.
- **Character Set**: `utf8mb4` ensures support for extended Unicode characters.

---

### Key Insight

The structures of the `device_camera_data` and `device_sweeper_data` tables are largely similar. Both tables store timestamped data about devices, along with some device-specific metrics (e.g., rotation for cameras, position for sweepers). However, due to differences in the types of sensors and the data they collect, it becomes necessary to store the data for each device type in separate tables.  

This approach means that for every new type of device, a new table must be created to accommodate the unique dimensions of data collected by that device's sensors.