### Tasks

By default, tasks are markdown files that live in the `tasks/` directory. The
file name (excluding its extension) will be used as its main identifier (its
"main ID").

Tasks carry some metadata. Currently, `tags`, `status`, `stage`, and `alt-ids`
have their own semantics. Other keys are not used by `todo`, but can be queried.

A task's first H1 (`#` heading) will be used as its title.
