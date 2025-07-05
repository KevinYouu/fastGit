package i18n

// enTranslations contains all English translations
var enTranslations = map[string]string{
	// Root command
	"root.short":       "fastGit is a tool that helps you quickly submit code with a command line interface.",
	"root.description": "A fast and efficient Git workflow tool",

	// Version command
	"version.short":   "Show fastGit version",
	"version.version": "Version:",
	"version.github":  "Github:",
	"version.about":   "To know more about me, you can visit:",

	// Status command
	"status.short": "Show git status",

	// Push commands
	"push.all.short":      "Push all changes to remote repository",
	"push.selected.short": "Select and push specific changes",

	// Remote commands
	"remotes.short": "Manage git remotes",

	// Reset command
	"reset.short": "Reset git repository to a specific state",

	// Tag commands
	"tag.short":        "Manage git tags",
	"tag.delete.short": "Delete git tags",

	// Merge command
	"merge.short": "Merge branches",

	// Update command
	"update.short": "Update fastGit to the latest version",

	// Init command
	"init.short": "Initialize a new git repository",

	// Common messages
	"error.general":    "An error occurred:",
	"success.general":  "Operation completed successfully",
	"confirm.continue": "Do you want to continue?",
	"select.option":    "Please select an option:",
	"input.required":   "This field is required",

	// Git specific
	"git.branch":           "Branch:",
	"git.commit":           "Commit:",
	"git.status.clean":     "Working tree clean",
	"git.status.modified":  "Modified files:",
	"git.status.untracked": "Untracked files:",
	"git.push.success":     "Successfully pushed to remote",
	"git.push.failed":      "Failed to push to remote",

	// File operations
	"file.select":   "Select files:",
	"file.selected": "Selected files:",
	"file.none":     "No files selected",
	"file.all":      "All files",

	// Progress messages
	"progress.pushing":  "Pushing changes...",
	"progress.fetching": "Fetching updates...",
	"progress.merging":  "Merging branches...",
	"progress.loading":  "Loading...",
	"progress.complete": "Complete!",

	// Form components
	"form.input.placeholder": "Please enter...",
	"form.input.empty.error": "Input cannot be empty",
	"form.confirm.title":     "Confirmation",
	"form.select.title":      "Please select an option",
	"form.multiselect.title": "Please select options",

	// Git commands and operations
	"git.remotes.title":     "Remotes:",
	"git.remotes.failed":    "Failed to get remotes:",
	"git.status.no_changes": "No files changed.",
	"git.status.title":      "File statuses:",

	// Reset command
	"reset.title":           "🔄 Git Reset",
	"reset.select.commit":   "Choose a commit to reset to",
	"reset.select.mode":     "Select reset mode",
	"reset.mode.soft":       "Soft (keep changes staged)",
	"reset.mode.mixed":      "Mixed (keep changes unstaged)",
	"reset.mode.hard":       "Hard (discard all changes)",
	"reset.confirm.title":   "Reset Confirmation",
	"reset.confirm.message": "Reset to: %s\nMode: %s",
	"reset.confirm.warning": "⚠️  WARNING: Hard reset will permanently delete all uncommitted changes!",
	"reset.cancelled":       "🚫 Reset operation cancelled.",
	"reset.executing":       "Executing git reset...",
	"reset.success":         "✅ Git reset completed successfully!",

	// Reset - additional keys for implementation
	"reset.error.select.commit": "Error selecting commit:",
	"reset.error.select.mode":   "Error selecting reset mode:",
	"reset.mode.soft.label":     "Soft - Keep working directory and staging area",
	"reset.mode.mixed.label":    "Mixed - Keep working directory, clear staging area",
	"reset.mode.hard.label":     "Hard - Discard all uncommitted changes",
	"reset.mode.soft.desc":      " (keep all)",
	"reset.mode.mixed.desc":     " (default)",
	"reset.mode.hard.desc":      " (dangerous)",
	"reset.confirm.to":          "Confirm reset to",
	"reset.confirm.mode":        "mode",
	"reset.hard.warning":        "⚠️ Will lose all uncommitted changes!",
	"reset.executing.mode":      "Resetting (%s)...",
	"reset.completed.to":        "Reset to %s (%s)",
	"reset.success.prefix":      "Reset completed (HEAD → %s)",
	"reset.hint.soft":           "💡 Changes preserved in staging area",
	"reset.hint.mixed":          "💡 Changes preserved in working directory",
	"reset.hint.hard":           "💡 All uncommitted changes discarded",
	"reset.cancelled.msg":       "Cancelled",
	"reset.error.git.reset":     "Error executing git reset command:",

	// Tag operations
	"tag.create.title":     "🏷️  Create and Push Tag",
	"tag.input.name":       "Enter tag name:",
	"tag.input.message":    "Enter tag message (optional):",
	"tag.confirm.create":   "Create and push tag '%s'?",
	"tag.creating":         "Creating tag...",
	"tag.pushing":          "Pushing tag to remote...",
	"tag.success":          "✅ Tag created and pushed successfully!",
	"tag.delete.title":     "🗑️  Delete Tag",
	"tag.delete.select":    "Choose a tag to delete",
	"tag.delete.confirm":   "Are you sure you want to delete tag '%s'?\nThis will remove the tag both locally and from the remote repository.",
	"tag.delete.cancelled": "🚫 Tag deletion cancelled.",
	"tag.delete.local":     "Deleting local tag",
	"tag.delete.remote":    "Deleting remote tag",
	"tag.delete.success":   "Tag deleted successfully",
	"tag.get.error":        "get tags error:",

	// Push operations
	"push.all.title":      "🚀 Push All Changes",
	"push.selected.title": "📋 Push Selected Changes",
	"push.select.files":   "Select files to push:",
	"push.no.changes":     "No changes to push.",
	"push.preparing":      "Preparing to push...",
	"push.success":        "✅ Push completed successfully!",

	// Merge operations
	"merge.title":         "🔀 Merge Branch",
	"merge.select.branch": "Select branch to merge:",
	"merge.confirm":       "Merge '%s' into current branch?",
	"merge.executing":     "Merging branch...",
	"merge.success":       "✅ Merge completed successfully!",

	// Update operations
	"update.checking":         "Checking for updates...",
	"update.downloading":      "Downloading update...",
	"update.installing":       "Installing update...",
	"update.success":          "✅ Update completed successfully!",
	"update.restart.required": "Please restart fastGit manually.",
	"update.windows.script":   "🔄 Running Windows update script...",
	"update.script.success":   "Update script executed successfully",
	"update.failed.script":    "failed to run install script:",
	"update.unsupported":      "unsupported platform:",

	// Update operations - detailed
	"update.checking.version":   "🔍 Checking for latest version...",
	"update.latest.version":     "📦 Latest version: %s",
	"update.downloading.asset":  "Downloading %s...",
	"update.download.success":   "Download completed successfully",
	"update.download.failed":    "failed to download with both curl and wget:",
	"update.curl.failed":        "⚠️  curl failed, trying wget...",
	"update.wget.downloading":   "Downloading %s with wget...",
	"update.extracting":         "📂 Extracting downloaded file...",
	"update.extract.failed":     "failed to extract zip:",
	"update.installing.to":      "📥 Installing to %s...",
	"update.sudo.required":      "⚠️  Root permissions required for installation",
	"update.password.prompt":    "💡 You may be prompted for your password...",
	"update.direct.install":     "✓ Direct installation (sufficient permissions)",
	"update.no.sudo.required":   "✓ Direct installation (no sudo required)",
	"update.install.failed":     "failed to install binary:",
	"update.restart.terminal":   "💡 Please restart your terminal or run 'source ~/.bashrc' (or equivalent) to use the updated version.",
	"update.sudo.installing":    "🔐 Installing with sudo...",
	"update.copy.failed":        "failed to copy binary:",
	"update.permissions.failed": "failed to set permissions:",

	// Update - missing keys for new implementation
	"update.running_windows_script":                "🔄 Running Windows update script...",
	"update.downloading_running_script":            "Downloading and running update script...",
	"update.script_executed_success":               "Update script executed successfully",
	"update.failed_run_script":                     "failed to run install script",
	"update.complete_restart_manual":               "✅ Update complete. Please restart fastGit manually.",
	"update.unsupported_platform":                  "unsupported platform",
	"update.checking_latest_version":               "🔍 Checking for latest version...",
	"update.failed_get_latest_version":             "failed to get latest version",
	"update.latest_version":                        "📦 Latest version",
	"update.failed_create_temp_dir":                "failed to create temp directory",
	"update.downloading_latest_release":            "Downloading latest release",
	"update.download_completed":                    "Download completed successfully",
	"update.curl_failed_try_wget":                  "⚠️  curl failed, trying wget...",
	"update.downloading_with_wget":                 "Downloading %s with wget...",
	"update.failed_download_both":                  "failed to download with both curl and wget",
	"update.extracting_file":                       "📂 Extracting downloaded file...",
	"update.failed_extract_zip":                    "failed to extract zip",
	"update.installing_to":                         "📥 Installing to",
	"update.root_permissions_required":             "⚠️  Root permissions required for installation",
	"update.password_prompt_hint":                  "💡 You may be prompted for your password...",
	"update.direct_install_sufficient_permissions": "✓ Direct installation (sufficient permissions)",
	"update.direct_install_no_sudo":                "✓ Direct installation (no sudo required)",
	"update.failed_install_binary":                 "failed to install binary",
	"update.completed_successfully":                "🎉 Update completed successfully!",
	"update.restart_terminal_hint":                 "💡 Please restart your terminal or run 'source ~/.bashrc' (or equivalent) to use the updated version.",
	"update.installing_binary_system":              "Installing binary to system directory",
	"update.installing_binary":                     "Installing fastGit binary...",
	"update.binary_installed_success":              "Binary installed successfully",
	"update.setting_executable_permissions":        "Setting executable permissions",
	"update.setting_permissions":                   "Setting permissions...",
	"update.permissions_set_success":               "Permissions set successfully",
	"update.installing_with_sudo":                  "🔐 Installing with sudo...",
	"update.installing_binary_sudo":                "Installing binary with sudo...",
	"update.failed_copy_binary":                    "failed to copy binary",
	"update.failed_set_permissions":                "failed to set permissions",

	// Command execution
	"cmd.failed.step": "Failed at step %d: %s",
	"cmd.command":     "Command:",
	"cmd.executing":   "Executing:",

	// Time and date
	"time.created":  "Created:",
	"time.modified": "Modified:",
	"time.format":   "01-02 15:04",

	// File status
	"status.modified":  "Modified",
	"status.added":     "Added",
	"status.deleted":   "Deleted",
	"status.untracked": "Untracked",
	"status.unknown":   "Unknown",

	// Push operations - detailed
	"push.select.commit.type":   "Choose a commit type",
	"push.input.commit.message": "Enter your commit message:",
	"push.input.tag.name":       "Enter tag name:",
	"push.input.tag.message":    "Enter tag message:",
	"push.no.files.selected":    "No files selected for push.",
	"push.files.adding":         "Adding files...",
	"push.committing":           "Committing changes...",
	"push.to.remote":            "Pushing to remote...",

	// Form validation
	"validation.required": "This field is required",
	"validation.invalid":  "Invalid input",

	// Common actions
	"action.continue": "Continue",
	"action.cancel":   "Cancel",
	"action.confirm":  "Confirm",
	"action.select":   "Select",

	// UI Components
	"ui.executing.commands": "Executing commands...",
	"ui.progress":           "Progress:",
	"ui.status":             "Status:",
	"ui.error.details":      "Error details:",
	"ui.operation.success":  "Operation completed successfully!",
	"ui.operation.failed":   "Operation failed",
	"ui.exiting.error":      "💡 Exiting to show error details...",
	"ui.exiting.success":    "💡 Exiting...",
	"ui.step":               "Step %d: %s",

	// Spinner messages
	"spinner.fastgit.operation":  "🚀 FastGit Operation in Progress...",
	"spinner.operation.complete": "Operation completed successfully!",
	"spinner.operation.failed":   "Operation failed",
	"spinner.error.details":      "Error details: %v",
	"spinner.elapsed.time":       "⏱️ Elapsed: %v",
	"spinner.step.progress":      "Step progress:",
	"spinner.loading":            "loading",
	"spinner.pending":            "pending",
	"spinner.success":            "success",

	// Table operations
	"table.user.aborted": "user aborted",
	"table.no.selection": "no selection made",

	// Git command operations - pushAll
	"git.add.all.description": "Adding all files to staging area",
	"git.add.all.loading":     "Adding files...",
	"git.add.all.success":     "Files added successfully",
	"git.commit.description":  "Creating commit with message",
	"git.commit.loading":      "Creating commit...",
	"git.commit.success":      "Commit created successfully",
	"git.pull.description":    "Pulling latest changes from remote",
	"git.pull.loading":        "Pulling changes...",
	"git.pull.success":        "Pull completed successfully",
	"git.push.description":    "Pushing changes to remote repository",
	"git.push.loading":        "Pushing to remote...",

	// Git command operations - pushSelected
	"push.selected.no.files":       "No files to push.",
	"push.selected.no.selection":   "No files selected.",
	"git.add.selected.description": "Adding selected files to staging area",
	"git.add.selected.loading":     "Adding selected files...",
	"git.add.selected.success":     "Selected files added successfully",

	// Tag operations - detailed
	"tag.input.version":         "Enter your version:",
	"tag.input.commit.message":  "Enter your commit message:",
	"tag.create.description":    "Creating annotated tag",
	"tag.create.loading":        "Creating tag...",
	"tag.create.success":        "Tag %s created successfully",
	"tag.push.description":      "Pushing tag to remote repository",
	"tag.push.loading":          "Pushing tag to remote...",
	"tag.push.success":          "Tag %s pushed successfully",
	"tag.no.tags":               "no tags found in repository",
	"tag.delete.local.loading":  "Deleting local tag %s...",
	"tag.delete.local.success":  "Local tag %s deleted successfully",
	"tag.delete.remote.loading": "Deleting remote tag %s...",
	"tag.delete.remote.success": "Remote tag %s deleted successfully",

	// Merge operations - detailed
	"merge.no.branches":                     "No branches to merge.",
	"merge.select.target":                   "Select branch to merge into current branch:",
	"merge.select.strategy":                 "Select merge strategy:",
	"merge.success.message":                 "Merge completed successfully.",
	"merge.failed":                          "Failed to merge",
	"merge.starting":                        "Starting merge of '%s' using %s strategy...",
	"merge.warning.dirty.working.directory": "⚠️  Warning: You have uncommitted changes in your working directory.",
	"merge.confirm.continue.with.changes":   "Continue with merge anyway?",
	"merge.conflict.detected":               "🔀 Merge conflict detected!",
	"merge.conflict.instructions":           "💡 Resolve conflicts manually, then run 'git add <file>' and 'git commit'",
	"merge.fast.forward.failed":             "❌ Fast-forward merge not possible",
	"merge.fast.forward.suggestion":         "💡 Try using 'No fast-forward' strategy or resolve any conflicts",
	"merge.uncommitted.changes":             "❌ You have uncommitted changes that would be overwritten",

	// Merge strategies
	"merge.strategy.default.name":        "Default",
	"merge.strategy.default.description": "Default merge behavior",
	"merge.strategy.ff.only.name":        "Fast-forward only",
	"merge.strategy.ff.only.description": "Only merge if fast-forward is possible",
	"merge.strategy.no.ff.name":          "No fast-forward",
	"merge.strategy.no.ff.description":   "Always create a merge commit",
	"merge.strategy.squash.name":         "Squash",
	"merge.strategy.squash.description":  "Squash all commits into a single commit",

	// Error messages - detailed
	"error.get.options":        "Failed to get options:",
	"error.get.file.status":    "Failed to get file statuses",
	"error.multiselect.form":   "Failed to get file statuses:",
	"error.select.form.detail": "error selecting branch:",
	"error.current.branch":     "Failed to get current branch:",
}
