# Commit guidelines

## Commit and PR Messages

* **Reference issues, wiki pages, and pull requests liberally!**
* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move button left..." not "Moves button left...")
* Limit the first line to 72 characters or less

## Make relevant commits

Don't make your personal workflow part of the commit history. Make concise commits that describe the changes appropriately

For example, a merge to master might look like this:

* Add Register functionality
* Fix lint issues
* Add unit tests
* Address review comments

This is really only your own personal work history, and it has no value for others.

Once you merge to master, a single commit will be doing the job just fine:

* Add Register functionality

This makes it easier to read the commit history, because there is less noise. This also makes it easier to revert a specific change, because you don't have to guess which commits are part of that change.

## Don't squeeze everything into the same commit

Create separate PRs for each individual change that you do, or create separate commits within the same PR.

For example, a merge to master might look like this:

* Added Register functionality & fixed a bug in the list screen

This will make the PR harder to read for reviewers, because it is difficult to see which changes contribute to each feature. It might also be confusing if several unrelated changes are part of the same file.

It will also be really difficult to revert one of the changes, if it is baked together with a bunch of other changes. No commit should break the build, and every commit should be revertible.

If you are fixing several related bugs that are hard to split into individual commits, it could be ok to merge them as one commit.
