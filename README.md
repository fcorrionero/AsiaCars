# Asia Cars

### Install and compile project
```
mkdir -p build && cp .env.dist ./build/.env
go build -o build/asia_cars
```
### Run the application
```
cd build/
./asia_cars
```
When these commands are executed, you can test the application by going to http://localhost:8888/hello in your browser.
The default port is `8888` and could be configured in .env file.

### Run all tests
```
go test ./test/...
```
Inside `test` folder you will find unit tests for application and domain files. Integration tests for vehicle repository
and functional tests fo http controllers. Also, there is an example of benchmark for vehicle repository Save method.

To run infrastructure/ui/http tests you must execute first `cp .env.dist .env` from project root.

### Make requests
As in PROBLEM.md file is shown there are two types of requests, **operations** and **telemetries**.
Below are examples to make these requests.

#### Operations
<ul>
<li>
InFleet Vehicle

```
curl -H "Content-Type: application/json" --request POST -d '{"chassis_number" : "12345678901234567", "license_plate" : "ASd5646", "category" : "MBMR" }' http://localhost:8888/infleet
```
</li>
<li>
Install Vehicle. The chassis_number must exists in our persistence system, so you must first execute inFleet operation.
Otherwise a 404 response will be returned.

```
curl -H "Content-Type: application/json" --request POST -d '{"chassis_number" : "12345678901234567","serial_number"  : "G-34567"}' http://localhost:8888/install
```
</li>
</ul>

#### Telemetries
In order to obtain a correct response from the server, the vehicle have to be installed.
<ul>
<li>
Battery

```
curl -H "Content-Type: application/json" --request POST -d '{"serial_number": "G-34567","battery": 75}' http://localhost:8888/battery
```
</li>
<li>
Fuel

```
curl -H "Content-Type: application/json" --request POST -d '{"serial_number": "G-34567","fuel": 50,"type": "increment"}' http://localhost:8888/fuel
```
</li>
<li>
Milleage

```
curl -H "Content-Type: application/json" --request POST -d '{"serial_number": "G-34567","mileage": 1000,"unit": "km"}' http://localhost:8888/mileage
```
</li>

<li>
Get all telemetries

```
curl http://localhost:8888/telemetries?serial_number=G-34567
```
</li>
</ul>

### Notes
The application has not a real persistence system, all data it is stored in memory, so all data will be deleted when you
stop the application.

### Roadmap
In order to improve the application we could develop a queue system to handle large number of requests, in a manner that
when a request arrives to our API it will be enqueued to later processing.
Other improvement could be developed using CQRS to separate readings and writings, using separated persistence systems (raead
and write models).
Other services like redis, Amazon ElasticCache, etc are available solutions to use a shared database between multiples
instances of our app.