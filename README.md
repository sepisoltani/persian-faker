![persian-faker](https://socialify.git.ci/sepisoltani/persian-faker/image?description=1&font=Inter&forks=1&issues=1&language=1&name=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Light)
# Persian Faker
[![codecov](https://codecov.io/gh/sepisoltani/persian-faker/graph/badge.svg?token=H0C0NFE9PM)](https://codecov.io/gh/sepisoltani/persian-faker)

The `persian-faker` is a specialized Go package designed to generate realistic Persian fake data, ideal for populating test data or simulating user data in applications that require Persian locale-specific data.

## Installation

To install this package, run the following command in your terminal:

```bash
go get github.com/sepisoltani/persian-faker
```

## Quick Start
Here's how you can start using the persian-faker to generate various types of fake data:

### Importing the Package
```go
import persianfaker "github.com/sepisoltani/persian-faker"
```

### Creating a Data Generator

```go
var generator = persianfaker.NewDataGenerator()
```

### Name Provider
* Generate a random Persian first name:
```go
var firstName = generator.Name.GenerateFirstName()
```
* Generate a random Persian last name:
```go
var firstName = generator.Name.GenerateLastName()
```
* Generate a random Persian full name:
```go
var firstName = generator.Name.GenerateFullName()
```

### Phone Number Provider
* Generate a random Persian mobile number:
```go
var mobileNumber = generator.PhoneNumber.GeneratePersianMobileNumber()
```

### Location Provider
* Generate a random province:
```go
var province = generator.Location.GenerateProvince()
```

* Generate a random city:
```go
var province = generator.Location.GenerateCity()
```

* Generate a random country:
```go
var province = generator.Location.GenerateCountry()
```

* Generate a random address:
```go
var address = generator.Location.GeneratePersianAddress()
```

### Bank Provider
* Generate a random Persian bank name:
```go
var bankName = generator.Bank.GenerateBankName()
```

* Generate a random Persian bank IBAN:
```go
var iban = generator.Bank.GenerateBankIBAN()
```

* Generate a random bank card number:
```go
var cardNumber = generator.Bank.GenerateBankCardNumber()
```

### Digit Provider
* Generate a Persian digit:
```go
var digit = generator.Digit.GeneratePersianDigit()
```

### Bill Provider
* Generate a random bill type:
```go
var billType = generator.Bill.GenerateBillType()
```

## Contributing
We welcome contributions to improve the package. If you have suggestions or improvements, please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for more information.