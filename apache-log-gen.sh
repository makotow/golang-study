#!/bin/sh

## 前提として gem install apache-loggenをしておくこと


apache-loggen --progress --limit=10000000 access.log
