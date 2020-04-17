# Card cabinet

This repository is a Go library for use in other projects.

Card cabinet is a simple system of markdown files with properties in frontmatter, together with boards that collect cards in different views.

The cards (files) are fairly stand alone, leaving the boards to link or structure cards based on labels or other 
properties.

The intention is use this in different projects, for example a simple kanban-style issue tracker.

## Non goals

Card cabinet is **not** a system where notes can reference each other in a infinite hierarchy like [Org mode](orgmode.org) or [Roam](roamresearch.com). The idea here is to actually simplify and remove relationships from cards themselves. 
