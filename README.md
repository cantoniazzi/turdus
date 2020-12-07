# turdus 

> Turdus is a genus of birds in the Turdidae family, comprising 66 species of blackbirds, thrushes, caraxués, and medium and large-sized sabiás, characterized by their rounded heads and long pointed wings, and generally having a melodious song. [[Wikipedia](https://pt.wikipedia.org/wiki/Turdus)]

![turdus_rufiventris](docs/turdus_rufiventris.png)



Turdus is an interface to consume Twitter data based on specific searches, consolidating this data in file format for the user.

# architecture

in soon.

# how to install

```bash
make install
```

or via docker:

`make docker/build`

# how to run

#### locally

You should export envs `APP_PORT`, `TWITTER_API_URL` , and `TWITTER_TOKEN_BEARER` with respective values.

`make local/run/api`

#### via docker:

You should create an `.env` file following the `.env_sample` file.

`make docker/run/api`

# how to lint

```
make lint
```

# how to test

```bash
make test
```
