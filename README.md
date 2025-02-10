# investment-service

## Running Locally

## Running Tests

## Environment Variables

## Problem statement

- A Bank already offers ISAs and Pensions to Employees of Companies (Employers) who have an existing arrangement with
the Bank. A Bank would like to be able to offer ISA investments to retail (direct) customers who are not associated with an
employer. A Bank would like to keep the functionality for retail ISA customers separate from it’s Employer based offering
where practical.

- When customers invest into a Bank ISA they should be able to select a single fund from a list of available options. Currently
they will be restricted to selecting a single fund however in the future we would anticipate allowing selection of multiple
options.
Once the customer’s selection has been made, they should also be able to provide details of the amount they would like to
invest.

- Given the customer has both made their selection and provided the amount the system should record these values and allow
these details to be queried at a later date.
As a specific use case please consider a customer who wishes to deposit £25,000 into a Bank ISA all into the Bank
Equities Fund.

## Requirements
- A customer should be able to invest in a single fund from a predefined list.
- The system should store the selected fund and investment amount.
- Initially they can only invest in one fund (future support for multiple funds).
- Customers must specify how much they want to invest.
- When investing the amount must be greater than zero
- Customers must have a unique ID for tracking investments.
- Customers should be able to fetch their investment details: Fund name, Amount invested, Investment date
- There should be an API to list all available funds.
- Authentication is required using JWT tokens.

## Assumptions
- Customers already exist - we wont manage users creation or authentication
- Authentication with JWT - given this service is likely going to be consumed by a front end application, i have made the assumption we will be working with JWT tokens
- When investing the amount must be greater than zero and meet any minimum deposit limits (we dont have any info on this).

## Decision log
- **Use a microservice style architecture with domain driven elements** - although only one service, we are using microservice principles (only handles investments/single responsibility, independent database storage, API driven, and doesnt store session state etc)
- **Use GraphQL over REST** - Service is likely going to be used by a front end that might need flexible queries, we are working with relational data and this gives us the chance to build on the API is a more scalable fashion.
- **Use Postgres** - well suited for the data we are working with (relational), we could use any relational DB but i have opted for postgres.
- **Dont use a ORM** - lots of overhead and using SQL queries allows us to take a more performant, fine grained approach. This does come at a tradeoff for readability
- **Dont use internal package** - Given this is a self contained API (not a library), there is no risk of accidental imports by external projects, so to keep it simple i have decided not to use an internal package (although it is idiomatic to use one).
- **DDD approach** - use investment module to encapsulate business logic. We could extend this to have more domain related modules in the future
- **Seperate repository & db modules** - seperation of concerns, db package focused on db setup and migrations, whereas repository interacts with the database. helps if we decide to change db provider, and with mocking.
- **Seperate repository from investment (domain) module** - although with DDD the repository could be considered part of the investment domain, i have decided to seperate them. This leans into better seperation of concerns and makes our repository module reusable across domain modules.
- **keep the main.go file clean** - I have opted to use a service module to handle the service and dependency setup.
