# tmpl
Generate projects from powerful templates

## Installation

```sh
go get github.com/stevenmatthewt/tmpl
```

Alternatively, you can download one of the binaries available in the releases section.

## Use

1. Download the template. GitHub is a greate place to store templates, but anywhere will work.
2. Run `tmpl --template <template folder> --destination <desired destination folder>`

## Creating a template

A template is simply a collection of files that will be copied from the template folder to the destination folder. Of course, there are a couple of additional features. To create a template, simply create a folder that has all of your template files, and these additional files:

1. a `tmpl.yml` file in the root of the template folder.
    - This file defines the variables that will be used to populate the template. See below.
2. any number of `*.tmpl` files.
    - adding the `.tmpl` extension to a file in the template folder will cause it to be processed as a template. This allows you to fill in variables, or entire sections of the file conditionally.

## tmpl.yml

The `tmpl.yml` file is how you can define the options that will be used when generating a project from your template. For any project that you want to use a template, simply include a `tmpl.yml` file in the root of that project. 

```yaml
# `prompts` is a list of all of the parameters that
# will be required when building form this template
prompts:
  # create one object per parameter. In this case, you
  # will be have access to the `name` and `monitoring_needed`
  # fields for your template files.
  - name: name
    description: Please enter your name
  - name: monitoring_needed
    description: Do you need New Relic monitoring?
    # you can specify the type of the parameter.
    # currently, only string and bool are supported.
    # `string` is the default.
    type: bool
```

## Conditional File/Folder generation

`tmpl` supports conditionally generating files and folders! File and folder names will be run through template processing automatically. If you have folders/files with the following names, here's what will happen:

- `root/{{.folderName}}`
  - This folder will be created with the `folderName` parameter that was passed in.
  - I.e. if you specify `example` as `folderName`, then the folder `root/example` will be created.
- `root/{{if .need_extra_folder}}extra_folder{{end}}`
  - This folder will **only** be created if the parameter `need_extra_folder` is true.

Any valid GoLang templating is supported for file/folder names. Be creative!