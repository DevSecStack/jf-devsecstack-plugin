# JFrog CLI DevSecStack Plugin

The JFrog DevSecStack Plugin is a custom plugin for JFrog CLI that adds DevSec related capabilities to your pipeline. 

## Features

- Add Cargo (Rust) dependencies to build-info from Cargo.lock file.

## Installation

To install the plugin, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/devsecstack/jf-devsecstack-plugin.git
    cd jf-devsecstack-plugin
    ```

2. Build the plugin:
    ```sh
    go build -o jf-devsecstack-plugin
    ```

3. Publish the plugin:
    ```sh
    export JFROG_CLI_PLUGINS_SERVER=<SERVER_ID>
    jf plugin publish devsecstack v1.0.0
    ```
4. Install the plugin:
    ```sh
    export JFROG_CLI_PLUGINS_SERVER=<SERVER_ID>
    jf plugin install devsecstack
    ```


### Cargo Add Dependencies

To use the plugin, run the following command:

```sh
jf devsecstack cargo-add-dependencies [flags]
```
#### Example

```sh
jf devsecstack cad --build-name my-build --build-number 1
```
#### Flags

- `--build-name` (required): Build name.
- `--build-number` (required): Build number.
- `--project` (optional): JFrog project key.
- `--module` (optional): Optional module name in the build-info for adding the dependency.
- `--server` (required): Artifactory server ID.
- `--dry-run` (optional): Set to true to disable communication with Artifactory. Default is false.

#### Environment Variables

- `CARGO_SKIP` (optional): Set to true to skip crago commands (```cargo generate-lockfile```). Default is false.
- `CARGO_LOCKFILE` (optional): Path to the Cargo.lock file. Default is `Cargo.lock`.

