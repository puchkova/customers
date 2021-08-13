# Customers
The project is a simple **Golang** **Echo** framework & **PostgreSQL** CRUD REST API with **Gorm** library.

## Prerequisites
- The **Go** programming language https://golang.org/dl/
- **Docker Desktop** https://www.docker.com/products/docker-desktop
- **GoLand** or similar IDE https://www.jetbrains.com/go/promo/?source=google&medium=cpc&campaign=10156131500&gclid=CjwKCAjwsNiIBhBdEiwAJK4khrn3IDTmD-Xv1BFZ9HQeeSUwIeIFaG69dxoHLW1ACvjxdrZxD5Dn9RoCpXQQAvD_BwE
- **Postman** or similar API client https://www.postman.com/

## API
- GET http://localhost:1323/customers 
- GET http://localhost:1323/customers ?firstname=[*first name to search*]
- GET http://localhost:1323/customers ?lastname=*[*last name to search*]
- POST http://localhost:1323/customers ?[query param key]=[*new value*]&...
- PUT http://localhost:1323/customers ?[id]=[customer to update id(uuid)]&[field to update key]=[*new value*]...

## Installing&Running
- Clone thist repo https://github.com/puchkova/customers.git
- Start docker desktop app

### Terminal
- Open your terminal and run the `docker-compose up` to create and run docker containers
- Run `go run main.go` command to start the application

### Database
- Go to the browser to use Pgadmin 4 http://localhost:16543/browser/
- Sign in 
	 - login: **test@gmail.com**
	 - password: **test123!**
- Add new server 
- general > name > **test**
- connection > 
	 - host name/address > **your IP address** 
	 - *To check the IP address use ipconfig /all command. You need the IPv4 Address line(192.168.1....)(Preferred)*
	 - port: **5432** 
	 - maintenance database: **root**
	 - username: **root**
	 - password: **root**
- Press save
- *Now you can check the customer table servers>test>databases>root>schemas>tables*

### API
- Open the Postman
- Send the **GET** request using URL http://localhost:1323/customers
- You can see customers from database in the request body
- Send the **POST** request using Query Params *(key : firstname ; value : customer name)* 
or use the URL http://localhost:1323/customers?firstname=Hugo&lastname=Nospmis&birthdate=2000-05-25&gender=Male&email=hsimpson@gmail.com&address=742 Evergreen Terrace, Springfield 
- You can see *"The customer is added"* message in response body 
- You can check that customer with name Hugo was added to database using GET request http://localhost:1323/customers?firstname=hugo
- Send the **PUT** request using id as first query parameter, put fields to update as query parameters. Get the id from GET response

## Application Structure
- `config` and `storage` packages contains files with name `db.go`
- `controller`, `model` and `service` packages contains files with name `customer.go`
- Main directory contains `docker-compose.yml` and `docker_itit.sql` files
- API routes are in `main.go` file that is in main directory

## Cases for Manual API Testing
### GET
- GET http://localhost:1323/customers?firstname=tatjana *returns 2 rows*
- GET http://localhost:1323/customers?firstname=Tatjana *returns 2 rows*
- GET http://localhost:1323/customers?firstname=  TATJANA *returns 2 rows*
- GET http://localhost:1323/customers?firstname=ATjan *returns 2 rows*

- GET http://localhost:1323/customers?lastname=  SIMPSON  *returns 5 rows* 
- GET http://localhost:1323/customers?lastname=imPSO  *returns 5 rows*

### POST
- POST http://localhost:1323/customers?firstname=JoJo *request is not allowed and message in response body "Last Name is required field"*
- POST http://localhost:1323/customers?firstname=JoJo&lastname=Bouvier *request is not allowed and message in response body "Birthdate is required field"*
- POST http://localhost:1323/customers?firstname=JoJo&lastname=Bouvier&birthdate=2002-02-02 *request is not allowed and message in response body "Gender is required field"*
- POST http://localhost:1323/customers?firstname=JoJo&lastname=Bouvier&birthdate=2029-02-02&gender=Male *request is not allowed and message in response body "Email is required field"*
- POST http://localhost:1323/customers?firstname=JoJo&lastname=Bouvier&birthdate=2009-02-02&gender=Male&email=jojo.bouvier@gmail.com *request is not allowed and message in response body "Age should be in the range from 18 to 60 years"*
- POST http://localhost:1323/customers?firstname=JoJo&lastname=Bouvier&birthdate=2000-02-02&gender=Male&email=jojojojojo *request is not allowed and message in response body "Invalid email address format"*
- POST http://localhost:1323/customers?firstname=JoJo&lastname=Bouvier&birthdate=2000-02-02&gender=Unknown&email=jojo.bouvier@gmail.com *request is not allowed and message in response body "Gender should be Male or Female"*

### PUT
        {
            "ID": "346f6f5b-52a8-48e7-bbb6-99d0e368ed72",
            "firstname": "Maggie  ",
            "lastname": "Simpson",
            "birthdate": "1987-04-19T00:00:00Z",
            "gender": "Female",
            "email": "email@gmail.com",
            "address": "742 Evergreen Terrace, Springfield"
        },

- PUT http://localhost:1323/customers?id=346f6f5b-52a8-48e7-bbb6-99d0e368ed72&gender=Male *successful request and message "The customer is updated"*
- PUT http://localhost:1323/customers?id=346f6f5b-52a8-48e7-bbb6-99d0e368ed72&gender=Male&firstname=A *request is not allowed and message "Invalid First Name"*
- PUT http://localhost:1323/customers?id=346f6f5b-52a8-48e7-bbb6-99d0e368ed72&gender=Male&firstname=Sven&email=kkkk.lgggg *request is not allowed and message "Invalid email address format"*
- PUT http://localhost:1323/customers?id=346f6f5b-52a8-48e7-bbb6-99d0e368ed72&gender=Male&firstname=Sven&email=sven.simpson@gmail.com&gender= Male&adderess=755 Evergreen Terrace, Springfield&birthdate=1900-02-03 *request is not allowed "Age should be in the range from 18 to 60 years"*
- PUT http://localhost:1323/customers?id=346f6f5b-52a8-48e7-bbb6-99d0e368ed72&gender=Male&firstname=Sven&email=sven.simpson@gmail.com&gender= Male&adderess=755 Evergreen Terrace, Springfield&birthdate=1999-02-03 *successful request and message "The customer is updated"*


