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

## Repository structure
- a well-organized repo helps teams to navigate and maintain the codebase effectively.
- a unified structure also makes it easier for developers to move between systems and services.
```
myapi/
├── README.md          # Overall system overview, setup guide and architecture overview
├── design/            # Shared design elements across all services
│   ├── design.go      # Top-level design for unified approach using Goa's DSL for a unified approach
│   └── types/         # Shared type definitions defined with Meta("struct:pkg:path")
├── gen/               # Generated code (unified design approach)
│   ├── http/          # HTTP transport layer code
│   ├── grpc/          # gRPC transport layer code
│   └── types/         # Generated shared types
├── scripts/           # Automation scripts for common development and deployment tasks for both approaches
└── services/          # Individual service implementations
    ├── users/         # Example: User service; each follows the consistent structure
    │   ├── cmd/       # Service entry points and executables
    │   ├── design/    # Service-specific API design
    │   ├── gen/       # Service-specific Generated code (independent design approach)
    │   ├── users.go   # Business logic implementation
    │   └── README.md  # Service-specific documentation
    └── products/      # Example: Product service
        └── ...
```

The above structure supports both monolithic and microservice deployments, allowing
- clear separation of concerns
- shared types and design elements
- independent service evolution
- easy service discovery and navigation

<br>

## Service communication patterns
A common architecture pattern is to have 
- a few services(sometimes only one) -> exposes application platform's capabilities to external clients
- with multiple back services handling the actual business logic 

### Front services
- public-facing services
- use HTTP as transport for broad client compatibility
- Focus on orchestrating requests to back services
- handle authentication and authorization of external requests
- initiate observability contexts(traces, metrics)
- define broad APIs with shallow implementations

### Back services
- internal services
- often use gRPC for performance benefits
- implement core business logic
- may use private identity mechanisms (e.g. spiffe)
- contribute to existing observability contexts
- define focused APIs with deep implementations


## Scripts
### Unified design approach
- single command generates all service code
- centralized testing across services
- coordinated builds and deployments
- shared dependency management 

### Independent design approach
- per-service code generation
- isolated testing envs
- independent build processes
- service-specific deployments

<br>

## Service implementation
- each service runs as a separate executable -> promoting isolation and independent scaling

## Best practices
1. Choose appropriate transport: Use HTTP for external and gRPC for internal
2. Plan for evolution: version the services and plan for backward compatibility
3. Implement robust error handling: Define clear error types and handle cross service failure gracefully
4. Document service interactions: Maintain clear documentation of service APIs and dependencies