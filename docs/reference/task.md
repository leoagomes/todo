# `todo` tasks

> You can check an example in [`docs/examples/example-task.md`](../examples/example-task.md)

A task in `todo` is simply a markdown file in the `tasks/` directory. Somewhere
in the task there must be a YAML block where its metadata is kept.

````
# My task

Small task summary.

```yaml .todo
status: open
```
````

## Task identifiers

A task's "main identifier" (or "main ID") is derived from its file name. A task
named `rewrite-in-rust.md` will have `rewrite-in-rust` as its main ID. Tasks can
also have alternative identifiers (or "alt IDs"), defined in their metadata block.

> Because many file systems are case-insensitive, main IDs are also case-insensitive.

If `rewrite-in-rust.md`'s content was the following, its alt IDs would be
`rust-rewrite` and `123`.

````
# Rewrite everything in Rust

```yaml .todo
status: open
alt-ids:
  - rust-rewrite
  - 123
```
````

In your code, you can refer to tasks by their main or alt IDs.

```C
int my_func(void) {
    // TODO(rewrite-in-rust)
    // or
    // TODO(@rust-rewrite)
    // TODO(@123)
    // note the @ before alt IDs
}
```

## Task metadata

The task metadata block has a few required and optional keys. You can also add
your custom data to it, which can be queried with `todo`.

The following keys have meaning or are otherwise reserved for `todo` handling:
- `status`
  - Required.
  - One of `open`, `closed`, or `completed`.
- `stage`
  - Optional.
  - Must be one of your defined "stages".
- `tags`
  - Optional.
  - An array of string "tags" for this task. Can be queried, and will be
    displayed in the UI.
- `alt-ids`
  - Optional.
  - An array of string "alternative identifiers".
  - Can be used to refer to this task in code.
- `title`
  - Optional.
  - This will be your task's title.
  - If not provided, it'll be derived from the task's first heading.
- `summary`
  - Optional.
  - This will be your task's summary/description.
  - If not provided, it'll be derived from the task's first paragraph.

Any other keys will be queriable from `todo`, but are otherwise ignored.
