### Initialization

`todo` keeps indexes of tasks to make tracking more efficient, so it needs to
keep some data that should not be tracked by your versioning system. This data
lives (by default) in a `.todo` directory in your project, internally called the
`meta` directory.

The initialization process creates the `.todo`, initializes the indices, and,
if git is detected (or you ask for it via a flag), adds `.todo/` to your
`.gitignore`.

To initialize, run `todo init`.
