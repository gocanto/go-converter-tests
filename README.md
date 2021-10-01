## Converter Playground

This repository contains examples that make use of the given currency conversion [library](https://github.com/voyago/converter).

- [x] Fetch data using the currency layer service.
- [x] Fetching data using example (mock) data to demonstrate how different drivers would work.

## How do I try them?

First, make sure you have `GO` installed in your computer. Please, follow this [guide](https://golang.org/dl/)
in the case you do not have `GO` installed in your computer yet.

Second, you will have to clone this project to you local machine and cd into the cloned folder. Like so:

```bash
git@github.com:voyago/converter-tests.git

cd converter-tests
```

Third, you would have to create an `.env` file in order for you to run these examples since the given library
is environment aware. Thus, it needs this information to resolve desired drivers. 

You can do this by typing the following.

```bash
cp .prod.env.example .env
```

Lastly, you would be able to run the program by typing the following command from your command line

```bash
go run main.go
```

### Example output

```bash
 gus@gocanto  ~/Sites/converter-test   main  go run main.go

+ --------------------------------------------------- +
|          CURRENCY LAYER API DATA                    |
+ --------------------------------------------------- +

--- Environment
App name: converter-test
App env: production [live?: true]

--- Currencies exchange rate
Available currencies rate: 157
USD rate: 0.736022

--- Conversions
Result: [SGD 1.00], Amount: [1.00]

+ --------------------------------------------------- +
|               TESTING DATA                          |
+ --------------------------------------------------- +

--- Environment
App name: converter-test
App env: local [live?: false]

--- Currencies exchange rate
Available currencies rate: 157
SGD rate: 1.34258

--- Conversions
Result: [SGD 0.74], Amount: [0.74]

```

## Contributing

Please, feel free to fork this repository to contribute to it by submitting a functionalities/bugs-fixes pull request to enhance it.

## License

Please see [License File](https://github.com/voyago/converter-tests/blob/main/LICENSE) for more information.

## How can I thank you?

Why not star this GitHub repository and share its link on your social network?

> Don't forget to [follow me on twitter](https://twitter.com/gocanto)!

