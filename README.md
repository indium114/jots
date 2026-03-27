# jots

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/indium114/jots)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/indium114/jots/build.yml)

[![asciicast](https://asciinema.org/a/831290.svg)](https://asciinema.org/a/831290)

## Usage

### Adding an entry

To add an entry, run `jots add`.

You will be prompted for the entry text. Write your entry, and press enter.

You will then be asked if you want to attach a file. If not, press No (n, or enter with 'No' selected)
If you do want to, press Yes (y, or enter with 'Yes' selected) and type in the canonical (absolute) path to the file.

### Viewing the entries for a day.

#### For today

If you want to see the entries for the present day, just run:

```shell
jots list
```

#### For a specific day

To see the entries for a specific day, run:

```shell
jots list YYYY-MM-DD
```

### Viewing a specific entry

To view a specific entry, run:

```shell
jots view <UUID>
```

The UUID can be seen with the `list` command from earlier, it will be *eight characters long*.

### Open an attachment

To open an attachment, run:

```shell
jots open <UUID>
```

The attachment UUID can be seen with the `view` command from earlier, it will also be *eight caracters long*.

**jots** is a minimal CLI journaling tool

## Installation

### with Nix

Simply add the repo to your flake inputs...

```nix
inputs = {
  spyglass.url = "github:indium114/jots";
};
```

...and pass it into your `environment.systemPackages`...

```nix
environment.systemPackages = [
  inputs.jots.packages.${pkgs.stdenv.hostPlatform.system}.jots
];
```

### with Go

To install, simply run:

```shell
go install github.com/indium114/jots@latest
```

## Usage

## Inspirations

**jots** was mainly inspired by [jrnl](https://jrnl.sh)
