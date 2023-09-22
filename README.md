# Photogram

## Functional requirements:

System supports uploading a picture, deleting an existing picture, update a previously uploaded picture, fetch a picture file, fetch a list of uploaded pictures.

## Non-functional requirements:

The system must store the data for re-use in a later point of time. Images upload should be persisted in a storage so that it can be retrieved at any point of time.

## Installation:

To run this project, you require Docker and Docker-Compose to be pre-installed in your system.

`Makefile` contains all commands you require to run the server or run tests.

To build & run the API server along with the database, use command:

```bash
make whole
```

The above command builds the API layer and the database, runs the application server on port 8000. You can open http://localhost:8000 to check if your server is running and returns a 200 response code.

Other commands for frequent usage can found inside the Makefile. Use the following command to list all possible ones with their descriptions.

```bash
make help
```

## Packages:

This project is written in Golang. Packages used with it are:

- [Gin](https://github.com/gin-gonic/gin) - HTTP framework layer
- [Gorm](https://github.com/go-gorm/gorm) - ORM library to connect to Postgres (database used)
- [Image](https://github.com/golang/image) - Libraries to help decode various kinds of image and their metadata
- [Swagger](https://github.com/swaggo/gin-swagger) - API documentation and manual testing
- [Viper](https://github.com/spf13/viper) - Configuration management via a config file
- [Testify](https://github.com/stretchr/testify) - Toolkit for common test assertions and mocks


## Endpoints:

- GET [/healthcheck](http://localhost:8000/healthcheck) - Get server uptime metrics
- GET [/swagger/index.html](http://localhost:8000/swagger/index.html) - Access Swagger API docs & playground
- GET [/](http://localhost:8000/) - List of all upload images with their metadata. (This supports pagination starting from page 1 as default). Page size is restricted to 10 from the backend.
- POST [/](http://localhost:8000/) - Upload an image. Request type should be form-data with key `image` storing the file.
- PUT [/picture/<id>](http://localhost:8000/picture/1) - Updates a picture for a row id. Request syntax is same as POST.
- GET [/picture/<id>](http://localhost:8000/picture/1) - Returns an upload picture with its pre-compute metadata as JSON
- GET [/picture/<id>/image](http://localhost:8000/picture/1/image) - Returns an upload picture with its pre-compute metadata
- DELETE [/picture/<id>](http://localhost:8000/picture/1) - Deletes an entry from the database

## DB Architecture:

I have used [PostgresSQL](https://www.postgresql.org/) as my RDMS data store. This is running inside Docker Compose environment itself. Following is a description of its columns.

| Name | Type | Description |
| --- | --- | --- |
| id | integer, Primary Key | Auto-generated and auto-incremented row id |
| created_on | integer | Gets auto-created by the db whenever a new row is inserted |
| updated_on | integer | Gets auto-updated by the db whenever a row is inserted or updated |
| deleted | boolean (default false) | Gets to true whenever a row is deleted (soft-delete) |
| name | string | Name of the image file |
| destination | string | Destination of the image file in the host system |
| height | integer | Height of the image in pixels |
| width | integer | Width of the image in pixels |
| size | integer | Size of the image in bytes |
| content_type | string | Content Type of the image file |

We can find the above implementation in `db/models.go` in `Picture` struct.

Some decisions made during the modelling:

1. Usage of incremental integer rather than UUID:
I have not used UUID for this small project as it would be very difficult to search for a row using UUID. In a production use-case, all the APIs would default to using UUIDs which are unique, random and not iterable. Right now, it would be simpler for the end-user to search 1, 2, 3 and so on. Given a UI layer, UUID would be my way forward.
2. I don’t hard-delete the row whenever the records needed to be deleted. This is so that all data is retained. This does provide an overhead to put a filter on all APIs for listing, fetching etc.
3. created_on & updated_on fields are introduced to know when actions are performed. List API gives the result sorted based on latest updated on.


## API architecture:

I have used the Service-Repository pattern to write the API handlers, configurations, services, data-access, file storage. These are all implemented using interfaces and concept of Dependency Inversion (from D in SOLID principles).

Usage of interfaces allows us to swap any struct in Go which implicitly implements the interface to simply be used at runtime. With this pattern, it would be simple to implement some of the following use-cases too:

- A command line application to upload and list pictures, given the handlers use the existing `service.PicturesService` defined in `service/pictures.go`. The application would write their own handler to return an os.Stdout rather than JSON response.
- Swap the underlying database to MongoDB or ElasticSearch for which we would have to implement `db.PicturesRepository` defined in `repository/pictures.go`. This interface defines the functions needed by the service.
- Swap the storage from local system storage to a cloud storage like S3 for which we would have to implement `storage.ImageStorage` defined in `storage/storage.go`.

We have also defined DTO (Data Transfer Objects) which are basic struct architecture containing data to be sent or received from one entity to another. These can be configuration or API request/response in `dto/api.go`.

## Image Manipulation:

The project also tries to validate the request data, especially the file. Users won’t be allowed to upload any arbitrary file. Valid content types are:

- image/jpeg
- image/png
- image/gif
- image/tiff
- image/webp
- image/bmp

The system also tries to pre-process the data and store some information about it in its database - its name, content type, height, width, size in bytes.

## Testing:

To run the tests, use command:

```bash
make test
```

Tests defined implement fake storage and fake repository. These implement the interfaces defined above `storage.ImageStorage` & `db.PicturesRepository` respectively.

- Fake storage is responsible for not actually creating a file, but keeping the components in a variable.
- Fake repository doesn’t make a db connection, instead stores the data temporarily.

The project contains basic tests to validate the working of our service, as this is where the core part of the application lies. Service connects to the data layer and storage layer. Tests are written using random strings and random length of data to ensure we have not hardcoded any use-case.

## Future work:

1. Dominant colors:
We can pre-process the image to find dominant colors which can be used by the UI to show as a placeholder till the actual image loads. Same can be done using article linked below:

[Extract dominant colors of an image using Python - GeeksforGeeks](https://www.geeksforgeeks.org/extract-dominant-colors-of-an-image-using-python/amp/)

2. Image resizing:
We can also implement multiple sizes for images and store the compressed ones depending on various viewports. This kind of functionality is provided in services like Cloudinary.

For the above improvements, we would go for Python as it is very supportive of image handling. We can use a Message System like RabbitMQ. The tasks will be pushed by our API server in Go whenever an image gets created or updated. A Python task server can pick up these tasks to compute various meta-data. These data can then be stored in our database for future references.
