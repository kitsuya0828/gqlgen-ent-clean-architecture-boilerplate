# gqlgen-ent-clean-architecture-boilerplate

A Clean Architecture boilerplate for Golang applications built with gqlgen and ent.

This boilerplate is inspired by [Clean Architecture with ent and gqlgen \| by Manato Kuroda \| Better Programming](https://betterprogramming.pub/clean-architecture-with-ent-and-gqlgen-a789933a3665), but with some original customizations based on the [official docs of ent.](https://entgo.io/ja/docs/tutorial-todo-gql).

<img src="./assets/logo_compressed.png" />

[The Go gopher](https://go.dev/blog/gopher) was designed by [Renée French](https://www.instagram.com/reneefrench/).

## Usage
### Run Docker

```bash
cd docker
docker comopse up
```

### Set up database

```bash
make setup_db
make migrate_schema
```

### Codegen

```bash
make generate
```

### Start server

```bash
make start
```
You can see the **GraphQL Playground** at [http://localhost:8080/playground](http://localhost:8080/playground).

### Testing

```bash
make setup_test_db
make test_repository
```

### E2E

```bash
make setup_e2e_db
make e2e
```

## FAQ
### How to add a schema?
> Now Writing.

## References
* [gqlgen](https://gqlgen.com/)
    * [Getting Started](https://gqlgen.com/getting-started/)
* [ent.](https://entgo.io/)
    * [First Steps](https://entgo.io/docs/tutorial-setup)
    * [GraphQL Basics](https://entgo.io/docs/tutorial-todo-gql)