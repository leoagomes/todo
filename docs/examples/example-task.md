# An example task

This will be the task summary.

Tasks can include anything that you'd include in markdown files. You can use
your favorite markdown flavor (as long as your renderer supports it).

Somewhere in the task (we recommend the bottom of it), you need a YAML block
marked `yaml .todo` (notice the space). This is so regular YAML blocks are not
confused with todo metadata blocks.

```yaml .todo
# The following is unnecessary, but left here for display.
title: An example task
summary: This will be the task summary.

# The following are required
status: open

# And these are optional
stage: backlog
tags:
  - examples
  - demo
references:
  - other-example-task

# And you can add your own data here
my_custom_flag: true
```
