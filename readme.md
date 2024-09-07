# GitNoob

GitNoob is a collection of Git-related tools implemented in Go, designed to simplify and enhance your Git workflow. This repository provides a suite of utilities for automating common Git tasks, improving productivity, and managing repositories effectively.

## Table of Contents

- [Installation](#installation)
- [Tools](#tools)
  - [autobranch](#autobranch)
  - [lazypush](#lazypush)
  - [autocommit](#autocommit)
  - [gitsync](#gitsync)
  - [newrepo](#newrepo)
  - [gitcleanup](#gitcleanup)
  - [commitdiary](#commitdiary)
  - [autorebase](#autorebase)
  - [precommitlint](#precommitlint)
  - [stashmanager](#stashmanager)
  - [autoupdate](#autoupdate)
  - [gitflowhelper](#gitflowhelper)
  - [gitbisecthelper](#gitbisecthelper)
  - [gitbackups](#gitbackups)
  - [gitpruner](#gitpruner)
  - [githooksmanager](#githooksmanager)
  - [rollbackhelper](#rollbackhelper)
  - [autogitconfig](#autogitconfig)
- [Usage Examples](#usage-examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install GitNoob, clone the repository and build the tools:


git clone https://github.com/amanmehtacode/GitNoob.git
git clone https://github.com/amanmehtacode/GitNoob.git
git clone https://github.com/amanmehtacode/GitNoob.git
git clone https://github.com/amanmehtacode/GitNoob.git

bash
git clone https://github.com/amanmehtacode/GitNoob.git
cd GitNoob
go build -o autobranch ./cmd/autobranch
go build -o lazypush ./cmd/lazypush
Repeat for other tools as needed