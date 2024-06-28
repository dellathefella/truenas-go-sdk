# TrueNAS Go SDK Development

This SDK supports both `angelfish`,`dragonfish` and `bluefin` release trains. Substitude `<release>` with your release train below.

Repo layout is as follows:

* `<release>/` is the generated Go SDK
* `cfg/<release>_original.yaml` is the raw openapi spec as returned by a TrueNAS instance
* `cfg/<release>.yaml` is a manually curated openapi spec compatible with `openapi-generator`
* `cfg/<release>_config.yaml` contains config options for `openapi-generator`

Unfortunately the raw openapi spec (`cfg/<release>_original.yaml`) is both incompatible with openapi generator, but also missing a lot of response/input definitions.

To fetch the remote spec:

```bash
./cfg/fetch-remote.sh core_13.0-u6_apiv2 dragonfish
openapi-generator-cli validate -i cfg/dragonfish.yaml
```

To generate go SDK from current (curated) openapi spec use this command:
```bash
npm install @openapitools/openapi-generator-cli -g
npm install -g swagger2openapi
npm install -g @stoplight/spectral-cli
```

```bash
$ openapi-generator-cli generate -i cfg/bluefin.yaml -c cfg/bluefin_config.yaml -o ./truenas -g go --git-user-id dellathefella --git-repo-id truenas-go-sdk
```

Workaround for TrueNAS content negotiation issue:

```bash
patch client.go < client.patch
```

The full release build can be built using the `generate.sh` script:

```bash
./generate.sh <release>
```
## Code quality

* Install git pre-commit: [https://pre-commit.com/](https://pre-commit.com/) and run `pre-commit install` in project root
* Install commitizen and make commits with `git cz`:

```bash
npm install -g commitizen cz-conventional-changelog
echo '{ "path": "cz-conventional-changelog" }' > ~/.czrc
```