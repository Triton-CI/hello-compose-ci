package main

/*
Application Flow Diagram:

```mermaid
graph TD
    A[Client Request] --> B{Route Handler}
    B -->|/| C[rootHandler]
    B -->|/text| D[textHandler]
    B -->|/html| E[htmlHandler]
    B -->|/json| F[jsonHandler]

    C --> G[Return HTML with Available Routes]

    D --> H[Create Human struct]
    H --> I[Format as Plain Text]
    I --> J[Return text/plain response]

    E --> K[Create Human struct]
    K --> L[Format as HTML page]
    L --> M[Return text/html response]

    F --> N[Create Human struct]
    N --> O[Encode as JSON]
    O --> P[Return application/json response]

    style C fill:#e1f5ff
    style D fill:#e1f5ff
    style E fill:#e1f5ff
    style F fill:#e1f5ff
    style H fill:#ffe1e1
    style K fill:#ffe1e1
    style N fill:#ffe1e1
```

```mermaid
classDiagram
    class Human {
        +string Name
        +int Age
        +string City
    }

    class HTTPServer {
        +rootHandler(w, r)
        +textHandler(w, r)
        +htmlHandler(w, r)
        +jsonHandler(w, r)
        +main()
    }

    HTTPServer ..> Human : uses

    note for Human "Data structure used by all handlers"
    note for HTTPServer "HTTP server listening on port 8080"
```
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	human := Human{
		Name: "John Doe",
		Age:  30,
		City: "Paris",
	}

	fmt.Fprintf(w, "Name: %s\nAge: %d\nCity: %s", human.Name, human.Age, human.City)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	human := Human{
		Name: "John Doe",
		Age:  30,
		City: "Paris",
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Human Info</title>
</head>
<body>
    <h1>Human Information</h1>
    <p><strong>Name:</strong> %s</p>
    <p><strong>Age:</strong> %d</p>
    <p><strong>City:</strong> %s</p>
</body>
</html>
`, human.Name, human.Age, human.City)

	fmt.Fprint(w, html)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	human := Human{
		Name: "John Doe",
		Age:  30,
		City: "Paris",
	}

	json.NewEncoder(w).Encode(human)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>API Routes</title>
</head>
<body>
    <h1>Available Routes</h1>
    <ul>
        <li><a href="/text">/text</a> - Text response</li>
        <li><a href="/html">/html</a> - HTML response</li>
        <li><a href="/json">/json</a> - JSON response</li>
    </ul>
</body>
</html>
`
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/text", textHandler)
	http.HandleFunc("/html", htmlHandler)
	http.HandleFunc("/json", jsonHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
