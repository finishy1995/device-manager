# Enhanced MongoDB Data Models

## Data Model 1: DeviceMetadata

The `DeviceMetadata` model in MongoDB leverages the **Embedded Model** to store metadata efficiently. Instead of creating multiple rows for each `device_sn` and its associated parameters, all parameters are stored as a **map** within a single document. This structure significantly improves storage efficiency and enhances read and update performance.

### MongoDB Data Structure

```go
type Param struct {
	PV string    `bson:"pv"` // Parameter value
	CT time.Time `bson:"ct"` // Creation time
	UT time.Time `bson:"ut"` // Update time
}

type DeviceMetadata struct {
	DeviceSn string           `bson:"device_sn"` // Device serial number
	Params   map[string]Param `bson:"params"`    // Map of parameter types to their values
}
```

### Example Document
```json
{
  "device_sn": "SN-100000000000",
  "params": {
    "0": { "pv": "49462", "ct": "2024-11-24T09:19:08Z", "ut": "2024-11-26T07:02:27Z" },
    "1": { "pv": "10928", "ct": "2024-11-24T09:19:08Z", "ut": "2024-11-26T07:02:27Z" },
    "2": { "pv": "14525", "ct": "2024-11-24T09:19:08Z", "ut": "2024-11-26T07:02:27Z" },
    "3": { "pv": "15911", "ct": "2024-11-24T09:19:08Z", "ut": "2024-11-26T07:02:27Z" },
    "4": { "pv": "96200", "ct": "2024-11-24T09:19:08Z", "ut": "2024-11-26T07:02:27Z" }
  }
}
```

### Key Advantages
1. **Embedded Model**:  
   By storing parameters as an embedded map (params), MongoDB avoids the need for multiple rows in a relational database, improving storage efficiency.
2. **Dynamic Schema**:  
   New parameter types can be added directly to the params map without altering the schema, making it highly flexible.
3. **Improved Performance**:  
   Updates and reads for a specific device’s metadata are more efficient since all data is contained within a single document.

---

## Data Model 2: DeviceCameraData

The `DeviceCameraData` model stores timestamped data for devices with camera-specific metrics. Each record corresponds to a single data entry for a device at a specific time.

### MongoDB Data Structure

```go
type DeviceCameraData struct {
	Id           int64     `bson:"id"`            // Unique identifier
	DeviceSn     string    `bson:"device_sn"`     // Device serial number
	Timestamp    time.Time `bson:"timestamp"`     // Timestamp of the data entry
	IsFixed      int64     `bson:"is_fixed"`      // Indicates if the device is fixed (1 = true, 0 = false)
	BatteryLevel uint64    `bson:"battery_level"` // Battery level of the device
	RotationX    float64   `bson:"rotation_x"`    // Rotation along the X-axis
	RotationY    float64   `bson:"rotation_y"`    // Rotation along the Y-axis
	RotationZ    float64   `bson:"rotation_z"`    // Rotation along the Z-axis
}
```

---

## Data Model 3: DeviceSweeperData

The `DeviceSweeperData` model stores timestamped data for sweeper devices. Each record corresponds to a single data entry for a device at a specific time, with metrics relevant to its operation.

### MongoDB Data Structure

```go
type DeviceSweeperData struct {
	Id           int64     `bson:"id"`            // Unique identifier
	DeviceSn     string    `bson:"device_sn"`     // Device serial number
	Timestamp    time.Time `bson:"timestamp"`     // Timestamp of the data entry
	IsCharging   int64     `bson:"is_charging"`   // Indicates if the device is charging (1 = true, 0 = false)
	BatteryLevel uint64    `bson:"battery_level"` // Battery level of the device
	PositionX    float64   `bson:"position_x"`    // X-coordinate of the device's position
	PositionY    float64   `bson:"position_y"`    // Y-coordinate of the device's position
}
```

---

## Key Insight: Combining DeviceCameraData and DeviceSweeperData in MongoDB

Although `DeviceCameraData` and `DeviceSweeperData` have different data structures, they can be stored in the **same MongoDB collection**. This is possible because MongoDB's flexible schema allows documents in a collection to have different fields.

### Unified Collection Example

```json
[
  {
    "id": 1,
    "device_sn": "SN-100000000000",
    "timestamp": "2024-11-28T12:00:00Z",
    "is_fixed": 1,
    "battery_level": 95,
    "rotation_x": 0.123,
    "rotation_y": -0.456,
    "rotation_z": 0.789
  },
  {
    "id": 2,
    "device_sn": "SN-100000000001",
    "timestamp": "2024-11-28T12:05:00Z",
    "is_charging": 0,
    "battery_level": 87,
    "position_x": 12.34,
    "position_y": 56.78
  }
]
```

### Key Advantages

1. **Shared Collection**:  
   Both DeviceCameraData and DeviceSweeperData can exist in the same collection, reducing the management overhead of maintaining separate collections for every device type.
2. **Time Series Optimization**:  
   Since these records are essentially **time series data**, MongoDB's **Time Series Collections** can be used. This provides:
      - **Data Compression**: Time series collections are highly compressed, reducing storage costs.
      - **Improved Query Performance**: Optimized for queries based on time ranges, which are common in telemetry data.
3. **Extensibility**:  
   Adding support for new types of devices requires only adding new fields to the documents, rather than creating entirely new tables or collections.
4. **Efficiency**:  
   By combining similar data structures in a single collection, MongoDB reduces duplication and improves query performance, especially when analyzing data across multiple device types.