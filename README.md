# Project Introduction
    Device Manage is a device management service developed in Go. 
    It provides efficient functions for device status updates and metadata management. 
    The project is currently divided into two implementations, Legacy and Enhance. 
    Legacy is based on MySQL and Enhance is based on MongoDB. 
    Device status data is time series data, which is stored in flat tables in MySQL and time series collections in MongoDB.
    Each implementation provides Http service and processor service to provide REST API access to the outside world. 
    You can write your own code to call these interfaces.
# Prerequisites
## Prepare the environment
    Go Environment: Ensure Go (version 1.18 or later) is installed.
    Database Support: Install one of the following databases:
        MySQL: No specific requirements, Just Mysql instance with a connection url.
        MongoDB: Recommended version 5.0 or above(With ts collection supported) , and instance with a connection url.
    ETCD: These services depend on etcd service, so make sure etcd is installed and started 
## Clone the Project
    git clone https://github.com/finishy1995/device-manager.git
    cd device-manage
## Create Mysql table 
    You need to create all table structures in the mysql instance.
    Create table script as below:
```shell
        legacy/model/device_camera_data.sql:u
        legacy/model/device_metadata.sql
        legacy/model/device_sweeper_data.sql
```
    You can find the data model in file:
```shell
        legacy_data_models.md
```
## Create MongoDB table
    You need to create  Collections in the MongoDB instance or run generate/* APIs to generate collection Automatically.
    You can find the data model in file:
```shell
        enhanced_data_models.md
```
## Create Vector Index in MongoDB
    Create Vector Index in MongoDB So you can use vector search later in the enhanced demo.
```json
{
 "fields": [
   {
     "numDimensions": 600,
     "path": "vector_euclidean_distance",
     "similarity": "euclidean",
     "type": "vector"
   }
 ]
}
{
 "fields": [
   {
     "numDimensions": 600,
     "path": "vector_angle",
     "similarity": "cosine",
     "type": "vector"
   }
 ]
}
```
## Configuration File
1. Update Legacy processor service config file
    Location: legacy/processor/etc/processor.yaml
```yaml
Name: processor.rpc
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: processor.rpc
DataSource: "root:123456@tcp(localhost:3306)/device?charset=utf8mb4&parseTime=True&loc=Local"
```
2. Update Legacy http service config file
    Location: legacy/http/etc/http-api.yaml
```yaml
Name: http-api
Host: 0.0.0.0
Port: 8888
DataSource: "root:123456@tcp(localhost:3306)/device?charset=utf8mb4&parseTime=True&loc=Local"
Etcd: "127.0.0.1:2379"
Processor: "processor.rpc"
```
3. Update enhanced processor service config file
    Location: enhanced/processor/etc/processor.yaml
```yaml
Name: processor.rpc
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: processor.rpc
DataSource: "mongodb+srv://admin:admin@cluster0.1lpgl.mongodb.net/"
```
4. Update enhanced http service config file
    Location: enhanced/http/etc/http-api.yaml
```yaml
Name: http-api
Host: 0.0.0.0
Port: 8888
DataSource: "root:123456@tcp(localhost:3306)/device?charset=utf8mb4&parseTime=True&loc=Local"
Etcd: "127.0.0.1:2379"
Processor: "processor.rpc"
```
# How to run
## Legacy Deployment and Running the Service
1. Build the Legacy Processor
```shell
    go build -o legacy-processor legacy/processor/processor.go
```
2. Run the Processor Service
```shell
    ./legacy-processor
```
3. Build the http
    Open another shell tab and run:
```shell
    go build -o legacy-http legacy/http/http.go
```
4. Run the Processor Service
```shell
    ./legacy-http
```

## Enhanced Deployment and Running the Service
1. Build the Processor
```shell
    go build -o enhanced-processor enhanced/processor/processor.go
```
2. Run the Processor Service
```shell
    ./enhanced-processor
```
3. Build the http
    Open another shell tab and run:
```shell
    go build -o enhanced-http enhanced/http/http.go
```
4. Run the Processor Service
```shell
    ./enhanced-http
```
# How to use 
    Both Legacy and Enhanced provide services through http API. For service interfaces, please refer to
```shell
    legacy/http/http.api
    enhanced/http/http.api
```
    You can write some demo code to call these interfaces.
    It is best if you can call the /generate/* APIs first to generate some sample data and then call other APIs.
    The API of Enhanced as below:
```

service http-api {
	@handler GetMetadataHandler
	get /metadata (GetMetadataRequest) returns (GetMetadataResponse) // Get metadata by SN

	@handler UpdateMetadataHandler
	get /metadata/update (UpdateMetadataRequest) returns (UpdateMetadataResponse) // Update metadata by SN

	@handler GetMetricsHandler
	get /metrics (GetMetricsRequest) returns (GetMetricsResponse) // Get device metrics

	@handler GenerateDemoMetadataHandler
	get /generate/metadata (GenerateDemoMetadataReq) returns (GenerateDemoMetadataResp) // DEMO use

	@handler RandomUpdateMetadataHandler
	get /update/metadata (UpdateMetadataReq) returns (UpdateMetadataResp) // Benchmark use

	@handler GetRandomUpdateResultHandler
	get /update/result (GetUpdateResultReq) returns (GetUpdateResultResp) // Benchmark use

	@handler GenerateDemoDeviceDataHandler
	get /generate/devicedata (GenerateDemoDeviceDataReq) returns (GenerateDemoDeviceDataResp) // DEMO use

	@handler GenerateDeviceVectorDataHandler
	get /generate/vector (GenerateDeviceVectorDataReq) returns (GenerateDeviceVectorDataResp) // Using camera rotation to generate device vector data

	@handler SearchBadDeviceHandler
	get /search/vector (SearchBadDeviceReq) returns (SearchBadDeviceResp) // Search bad device by Atlas MongoDB Vector Search
}
```