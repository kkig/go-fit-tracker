## Design

```
.
└── internal
    └── app
        ├── adapter
        │   ├── controller.go    # 4. Dependency Injection
        │   └── repository
        │       └── parameter.go # 3. Implementation
        ├── application
        │   └── usecase
        │       └── parameter.go # 2. Interface Function Call
        └── domain
            ├── parameter.go
            └── repository
                └── parameter.go # 1. Interface
```
