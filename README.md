# creditcard

How to run the project

- Clone the repository onto your local machine

- Ensure that Docker is installed and configured properly

- All packages should resolve when initialising the project for the first time

- Run the below command, making sure to expose port :8080

<code>creditcard % docker run --rm -it -p 8080:8080 -v "`pwd`/../src:/src" -v "`pwd`/../data:/data" -w /src docker.io/creditcard/image</code>

- The test URL, once the project is running, is at http://localhost:8080/api/validate/{creditCardNumber}, to which you can pass credit card numbers as path params

