# template

This repo is a template for all other repos.

Create a new repo using this repo as the template.

Additionally update the avid-global-terraform repo by copying the
entry for 'template':

```
  "template" = {
    branch_protection_enabled                      = true
    repo_name                                      = "template"
    default_branch                                 = "main"
    default_branch_required_status_checks_contexts = [
      "check_pr / Check Commit Message",
    ]
    azuredevops_environments = {}
    azuredevops_pipelines = {}
  },
```

 to an entry with the same name as the new repo:

```
  "new-repo" = {
    branch_protection_enabled                      = true
    repo_name                                      = "new-repo"
    default_branch                                 = "main"
    default_branch_required_status_checks_contexts = [
      "check_pr / Check Commit Message",
    ]
    azuredevops_environments = {}
    azuredevops_pipelines = {}
  },
```

 and run:

    task terraform-apply-repos

