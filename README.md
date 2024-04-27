![persian-faker](https://socialify.git.ci/sepisoltani/persian-faker/image?description=1&font=Inter&forks=1&issues=1&language=1&name=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Light)
# Persian Faker

The `persian-faker` is a specialized Go package designed to generate realistic Persian fake data, ideal for populating test data or simulating user data in applications that require Persian locale-specific data.

## Installation

To install this package, run the following command in your terminal:

```bash
go get github.com/sepisoltani/persian-faker
```

## Usage
Below are examples of how to use the package to generate Persian first names, last names, and full names. These examples assume you have successfully installed the package and have Go set up on your system.

## Generating a Random Persian First Name

```go
package main

import (
    "fmt"
    "github.com/sepisoltani/persian-faker/providers/name"
)

func main() {
    loader := names.FileDataLoader{}
    ng, err := names.NewNameGenerator(loader, 12345) // Example seed for reproducibility
    if err != nil {
        fmt.Println("Error initializing name generator:", err)
        return
    }

    fmt.Println("Random First Name:", ng.RandomFirstName())
}
```



## Generating a Random Persian Last Name

```go
package main

import (
    "fmt"
    "github.com/sepisoltani/persian-faker/providers/name"
)

func main() {
    loader := names.FileDataLoader{}
    ng, err := names.NewNameGenerator(loader, 12345) // Example seed for reproducibility
    if err != nil {
        fmt.Println("Error initializing name generator:", err)
        return
    }

    fmt.Println("Random Last Name:", ng.RandomLastName())
}
```



## Generating a Random Persian Full Name

```go
package main

import (
    "fmt"
    "github.com/sepisoltani/persian-faker/providers/name"
)

func main() {
    loader := names.FileDataLoader{}
    ng, err := names.NewNameGenerator(loader, 12345) // Example seed for reproducibility
    if err != nil {
        fmt.Println("Error initializing name generator:", err)
        return
    }

    fmt.Println("Random Last Name:", ng.RandomFullName())
}
```

## Contributing
We welcome contributions to improve the package. If you have suggestions or improvements, please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for more information.