# Working with Multiple Services in Goa

## Service
- a service in Goa represents a logical grouping of related endpoints that provive specific functionality
- The approach of splitting functionality across multiple services in larger applications, enables
    - better organization of API endpoints
    - clearer separation of concerns
    - easier maintenance and testing
    - independent deployment capabilities
    - granular security controls

<br>

## Service architecture
- front service(s)
    - exposed to the outside world 
- back service(s)
    - used by front service(s)

<br>

## Service organization
Goa provide two main approaches to organize the services' designs and generated code
1. Independent design
    - treats eac service as a standalone unit
    - `goa gen` is invoked for each service separately and each contain their own `gen/` directory
    - helps to move services easily to separate repos and to version them independently 
2. Unified design
    - defines a top-level design file that imports all the services design package and defines a common API
    - centralizes code generation and type sharing
    - generated code live in the root `gen/` directory
    - a single `goa gen` command generates all service code
    - shared types defined using the `Meta` `struct:pkg:path` key are automatically available across services
    - The system maintains a unified OpenAPI specification
    - well-suited for grouping related services
    - ensures that all the related services are versioned together and makes it easier to manage dependencies and updates 

<br>

## Transport Considerations
choice of transport protocol significantly impacts how services interact

### HTTP services
##### pros
- for external-facing services
- offers 
    - universal compatibility
    - a rich ecosystem of tools and middleware
    - familiar REST patterns
- easy to debug and test
- natural fit for web applications

##### cons
- servers potentially need to deal with different types of encoding
- need to choose a specific style for the the APIs
- poor encoding performance as compared to a binary protocol like protobuf 

### gRPC services
##### pros
- well-suited for internal service communication due to 
    - its high performance
    - low latency
    - built-in streaming support
- provides built-in service discovery when reflection is enabled
- enables efficient multiplexing of requests and responses over a single connection leads to significant performance gains when communicating between services

##### cons
- not well-suited for external facing services since it requires a client library to encode and decode the binary messages
- not widely adopted as HTTP; so may nto be as suitable for services that expect a wide range of clients

<br>


