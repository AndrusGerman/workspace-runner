# Workspace Runner

## Description

This is a simple CLI tool that allows you to run multiple projects in parallel.

## Usage

Run workspace
```bash
workspace-runner run <workspace-name>
```

Run workspace with projects
```bash
workspace-runner run <workspace-name> <project-name>
```

List workspaces
```bash
workspace-runner list
```

Run server
```bash
workspace-runner server
```

## Reason to Exist

I often faced the challenge of working on projects that required running multiple environments simultaneously -
running npm commands, starting databases, Redis instances, etc. Managing all these services can become quite complex,
especially when dealing with numerous services.

That's why I created this tool - you just create a workspace and it automatically starts everything you need.