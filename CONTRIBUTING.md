# Contributing

First off, thank you for considering contributing to Vaddy.

If you are interested in contributing but are stuck getting started post a [discussion] for assistance.

## Find a bug?

Are you seeing a panic from a built-in validation?
Maybe a value is passing when it should fail?

Head on over to the [issues page](https://github.com/miniscruff/vaddy/issues) and let us know what you did.

## Need an additional feature?

Post a [discussion] to present your idea and use cases.
If the new feature falls in line with the goals of the project it may be approved and implemented.

## Making changes

Before making changes please create an issue to discuss if it is a good fit.
Besides just adding features or fixing bugs, documentation changes are also very helpful.
Some ideas are example validations or benchmarks.

If you do decide to contribute code or documentation you will need the following:

* [Go](https://golang.org/doc/install): check [go.mod](go.mod) for which version is used.
    * Newer versions should also work
* [Golang CI Lint](https://golangci-lint.run/): for linting

### Fork and create a branch

If there is an issue you want to propose changes for start by [creating a fork](https://help.github.com/articles/fork-a-repo).
On your new forked repository create a branch.
Pick any branch name you like that describes your change, but really anything works.

Now you are ready to start working on your change.

Development commands are listed in the [Makefile](/Makefile) and can be run using make.
Make without arguments, or the help recipe can be used to list possible commands.

### Get the tests running

```bash
make test
```

Running tests should be possible immediately after cloning with just `go`.

### Creating a change file

Vaddy uses [changie](https://changie.dev) to generate the [CHANGELOG](CHANGELOG.md).
If you are making a change that would effect end users you will need to create a change file.
Change files are generated by Changie using the "new" command.

If you are not sure your work needs a change file you can ask.

`changie new`

Gitlab created a guide for [good changelog entires](https://docs.gitlab.com/ee/development/changelog.html#writing-good-changelog-entries)
that can be referenced when writing your body message.

Be sure to commit the generated file to your branch as part of your pull request.

### Prepare your pull request

```bash
make lint
```

## Creating your pull request

When you are ready to get your work merged into the main repository you can
[open a pull request](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests).

Some tips to streamline this process:
* Keep the change small.
* Start your pull request as a draft for feedback.
* Run all checks before opening your pull request.
* Add comments to your pull request if you have questions about implementation details.
* Link your pull request to the issue it resolves [using GitHub keywords](https://docs.github.com/en/github/managing-your-work-on-github/linking-a-pull-request-to-an-issue#linking-a-pull-request-to-an-issue-using-a-keyword).

A maintainer should review your pull request in a timely manner.

If your work is approved it will be "squash and merged" so feel free to use as many commits as you need.

[discussion]: https://github.com/miniscruff/vaddy/discussions