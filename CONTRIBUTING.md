CONTRIBUTING
============

The Cayley-Cookbook is made of two parts:
- The book (this repo)
- The [source code repository](https://github.com/tombenke/cayley-cookbook-src) for executable examples.

## Maintenance of the book

This book is using [Hugo static site generator](https://gohugo.io/) to build the content.

The content is placed into the [`website/content/`](website/content/) folder.
The source code examples are mirrored from the [source code repository](https://github.com/tombenke/cayley-cookbook-src) into the [`website/static/src`](website/static/src/) folder, as a git submodule.

To update the source code and the themes, execute:

```bash
    git submodule update --remote
```

To build the site, run in the `website folder:

```bash
    hugo -D
```

The site is published as a github page, from the [`docs/`](docs/) folder on the `master` branch.

