<br />
<div align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/3771/3771518.png" alt="Logo" width="80" height="80">

  <h3 align="center">Contact List</h3>

  <p align="center">
    An Golang API for manage a simple contact list!
  </p>

  <br />
</div>

## About The Project

It's here a way to manage a contact list using essential methods of HTTP requests and responses!

4 methods, 4 ways to use CRUD:
* It's possible create a new contact
* Want read all contacts? You can
* Update a contact is easy
* Warning! You can delete a contact
<br />

## Built With

This API was developed using Go(Golang) and some features/packages.

* [![Go]][Go-url]
<br />

## Getting Started

This is an example of how you can implements and use this API from Postman.

### Prerequisites

* Postman
* Golang

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/brunodeev/contact-list.git
   ```
2. Initialize go mod and run go mod tidy for update packages
   ```sh
   go mod init example-module
   ```
   ```sh
   go mod tidy
   ```
<br />

## Get all contacts (GET)
   ```sh
   http://ip:port/contacts
   ```
## Create contact (POST)
   ```sh
   http://ip:port/contacts
   ```
#### In Postman
   ```json
   {
       "name": "Example",
       "phones": [
           {
               "code": "+00",
               "ddd": "000",
               "number": "000000000"
           }
        ]
    }
   ```

## Update contact (PUT)
   ```sh
   http://ip:port/contacts?id=example
   ```
#### In Postman
   ```json
   {
       "name": "NewExample",
       "phones": [
           {
               "code": "+11",
               "ddd": "111",
               "number": "111111111"
           }
        ]
    }
   ```
## Delete contact (DELETE)
   ```sh
   http://ip:port/contacts?id=example
   ```

<br />

## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. Thanks!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/contact-feature`)
3. Commit your Changes (`git commit -m 'Add some feature contact'`)
4. Push to the Branch (`git push origin feature/contact-feature`)
5. Open a Pull Request

## Contact

Bruno César - bcgmeireles@gmail.com

Project Link: [https://github.com/brunodeev/contact-list](https://github.com/brunodeev/contact-list)


[contributors-url]: https://github.com/brunodev/contact-list/graphs/contributors
[forks-url]: https://github.com/brunodev/contact-list/network/members
[stars-url]: https://github.com/brunodev/contact-list/stargazers
[issues-url]: https://github.com/brunodev/contact-list/issues
[license-url]: https://github.com/brunodev/contact-list/blob/master/LICENSE.txt
[linkedin-url]: https://linkedin.com/in/brunodeev
[product-screenshot]: images/screenshot.png
[Go]: https://img.shields.io/badge/go-00add8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev
