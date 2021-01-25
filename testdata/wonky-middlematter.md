+++
labels = ["doc"]
+++

# Cards

A card is just markdown contents with metadata read from frontmatter.

```
---
labels: [doc]
title: Foo
---
# Foo
...
```

becomes

```json
{
   "name": "path/to/document.md"
   "properties": {
      "labels": ["doc"]
      "title": "Foo"
   "contents": "# Foo\n..."

}
```

## Name

Name is a slug, including path:

/path/to/document-name.md

## Contents

Contents is file contents with frontmatter block removed.
