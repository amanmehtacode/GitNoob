# GitNoob

This repository contains a collection of Git-related tools implemented in Go.

## Tools

- **autobranch**: Automatically creates a new Git branch based on the current ticket or feature.
- **lazypush**: Stages, commits, and pushes changes to the current branch, optionally pulling the latest changes first.
- **gitdiff**: Runs diff --git a/.DS_Store b/.DS_Store
index 7d78bb8..855ad1c 100644
Binary files a/.DS_Store and b/.DS_Store differ
diff --git a/cmd/main.go b/cmd/main.go
deleted file mode 100644
index e69de29..0000000
diff --git a/go.mod b/go.mod
deleted file mode 100644
index 42513f3..0000000
--- a/go.mod
+++ /dev/null
@@ -1,5 +0,0 @@
-module github.com/amanmehtacode/GitNoob
-
-go 1.20
-
-// Any additional dependencies will be listed here
diff --git a/go.sum b/go.sum
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/backup.go b/pkg/gitutils/backup.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/bisect.go b/pkg/gitutils/bisect.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/branch.go b/pkg/gitutils/branch.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/commit.go b/pkg/gitutils/commit.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/config.go b/pkg/gitutils/config.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/diff.go b/pkg/gitutils/diff.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/flow.go b/pkg/gitutils/flow.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/hooks.go b/pkg/gitutils/hooks.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/prune.go b/pkg/gitutils/prune.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/rollback.go b/pkg/gitutils/rollback.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/stash.go b/pkg/gitutils/stash.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/sync.go b/pkg/gitutils/sync.go
deleted file mode 100644
index e69de29..0000000
diff --git a/pkg/gitutils/utils.go b/pkg/gitutils/utils.go
deleted file mode 100644
index e69de29..0000000
diff --git a/readme.md b/readme.md
deleted file mode 100644
index e69de29..0000000 with custom formatting.
- **autocommit**: Stages all changes and commits with a user-provided or default message.
- **gitsync**: Pulls the latest changes for each Git repository in a directory and optionally pushes local changes.
- **newrepo**: Initializes a new Git repository and sets up a remote.
- **gitcleanup**: Lists and deletes merged local and remote branches.
- **commitdiary**: Generates a Markdown log of commit history for the day.
- **autorebase**: Fetches the latest remote changes and attempts to rebase the current branch.
- **precommitlint**: Automatically runs a linter on all staged files before committing.
- **stashmanager**: Manages Git stashes interactively.
- **autoupdate**: Automatically pulls and rebases updates across all branches.
- **gitflowhelper**: Automates Git Flow operations for starting features, releases, and hotfixes.
- **gitbisecthelper**: Simplifies the Git bisect process.
- **gitbackups**: Creates backups of Git repositories.
- **gitpruner**: Prunes old or unused Git branches.
- **githooksmanager**: Manages Git hooks for the repository.
- **rollbackhelper**: Rolls back to a previous commit or branch state.
- **autogitconfig**: Automatically sets up Git configurations.

