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
- Customers should be able to fetch their investment details: Fund name, Amount invested, Investment date
- There should be an API to list all available funds.
- There should be a seperation of business rules for retail and employee investments

## Assumptions
- Customers already exist - we wont manage users creation or authentication
- Authentication with JWT - although not implemented in this solution, i assume we are using JWTs and need to consider implementing this in the future in this service.
- Fund data already exists in the database. I have not implemented an endpoint to create fund data as i assume this will be handled in another microservice. I have however inserted some default data for funds in a migrations file
- When investing the amount must be greater than zero and meet any minimum deposit limits (we dont have any info on this).
- We dont need to worry about cancelling investments (although this could be a feature in the future)

## Decision log
- **Use a microservice style architecture with domain driven elements** - although only one service, we are using microservice principles (only handles investments/single responsibility, independent database storage, API driven, and doesnt store session state etc)
- **Use GraphQL over REST** - Service is likely going to be used by a front end that might need flexible queries, we are working with relational data and this gives us the chance to build on the API is a more scalable fashion.
- **Use Postgres** - well suited for the data we are working with (relational), we could use any relational DB but i have opted for postgres.
- **Dont use a ORM** - lots of overhead and using SQL queries allows us to take a more performant, fine grained approach. This does come at a tradeoff for readability
- **Dont use internal package** - Given this is a self contained API (not a library), there is no risk of accidental imports by external projects, so to keep it simple i have decided not to use an internal package (although it is idiomatic to use one).
- **DDD approach** - use investment module to encapsulate business logic. We could extend this to have more domain related modules in the future
- **Seperate database from investment (domain) module** - although with DDD the database could be considered part of the investment domain, i have decided to seperate them. This leans into better seperation of concerns and makes our database module reusable across domain modules.
- **keep the main.go file clean** - I have opted to use a service module to handle the service and dependency setup.
- **Make a clear distinction between employee and retail customers in investment module** - seperate the creating of investments for retail and employee customers, so we can easily add more business rules for each case in the future.

## Improvement
- implement auth middleware to check JWT in request and extract customerID (for example) form the claims in the context in the gql handlers.
- implement better errors returning from graphQL requests (making sure we dont leak implementation details)

TODO - write up how we would implement investing in multilple funds and what could be any issues with this (different fund rules?)

## Local dev testing

- getFunds
![get funds via postman](./test_screenshots/get-funds.png)
