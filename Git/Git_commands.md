Git Configuration
Set username globally:

git config --global user.name "Your Name"
Set email globally:

git config --global user.email "your.email@example.com"
List all configurations:

git config --list
Show help for Git configuration:

git config help
Show help for configuring Git:

git help config
Simple Git Commands
Repository Initialization and Status
Initialize a new Git repository:

git init
Display the status of the repository:

git status
Adding and Committing Changes
Add all changes to the staging area:

git add .
Add specific file(s) to the staging area:

git add <filename>
Add all changes including untracked files to the staging area:

git add -A
Commit changes with a specified message:

git commit -m "Your commit message here"
Amend the last commit message:

git commit --amend -m "New commit message"
Amend the last commit without changing the message:

git commit --amend
Branch Operations
Create a new branch:

git branch branch_name
Switch to the specified branch:

git checkout branch_name
Create and switch to a new branch:

git checkout -b branch_name
List all branches, including remote-tracking branches:

git branch -a
Delete the specified branch:

git branch -d branch_name
Remote Operations
Add a remote repository:

git remote add origin remote_repository_url
List all remote repositories:

git remote -v
Clone a repository from a remote URL:

git clone <remote-repo-url>
Push changes to a remote repository and set the upstream branch:

git push -u origin branch_name
Fetch and merge changes from a remote repository:

git pull origin branch_name
Intermediate Git Commands
Viewing and Comparing Commits
Display commit history:

git log
Display compact commit history:

git log --oneline
Display commit history with file changes:

git log --stat
Show the difference between two commits:

git diff commit-hash1 commit-hash2
Show the difference in unstaged files:

git diff
Merging and Rebasing
Merge changes from another branch into the current branch:

git merge branch_name
Rebase the current branch onto another:

git rebase branch_name
Stashing and Cleaning
Stash changes in a dirty working directory:

git stash
Apply stashed changes and remove them from the stash list:

git stash pop
Clean up untracked files and directories:

git clean -df
Resetting and Reverting
Unstage a file from the staging area:

git reset <filename>
Unstage all files from the staging area:

git reset
Soft reset to a specific commit, keeping changes in the staging area:

git reset --soft commit-hash
Mixed reset (default) to a specific commit, keeping changes in the working directory:

git reset commit-hash
Hard reset to a specific commit, discarding all changes:

git reset --hard commit-hash
Revert a specific commit by creating a new commit:

git revert commit-hash
Cherry-Picking
Apply the changes from a specific commit to the current branch:
git cherry-pick commit-hash
Advanced Git Commands
Managing Remote Repositories and Tags
Fetch changes from all remotes:

git fetch --all
Fetch and merge changes from the remote "main" branch:

git pull origin main
Push changes to the remote "main" branch:

git push origin main
Fetch and prune deleted remote branches:

git fetch --prune
Push tags to the remote repository:

git push --tags
Delete a branch from the remote repository:

git push origin --delete branch_name
Advanced Branch Operations
Start an interactive rebase:

git rebase -i HEAD~n
Squash commits during a merge:

git merge --squash branch_name
List branches that have been merged into the current branch:

git branch --merged
Show branches that contain a specific commit:

git branch --contains commit-hash
Miscellaneous Advanced Commands
View reflog to see the history of all references to HEAD:

git reflog
Use bisect to find the commit that introduced a bug:

git bisect start
Manage multiple working trees:

git worktree add <path> <branch>
Show the details of a specific commit:

git show commit-hash
Pack unpacked objects in the repository:

git gc
Verify the connectivity and validity of objects in the database:

git fsck
Create an archive of files from a named tree:

git archive -o <file>.zip HEAD
This list covers essential Git commands across different levels of complexity, providing a comprehensive overview for learners at various stages. Use these commands to manage your repositories efficiently and effectively.
