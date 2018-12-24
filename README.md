# Surrogate ![](https://travis-ci.org/SevansD/Surrogate.svg?branch=master)

Package for replace character in string by replacement map.

```surrogate.Sub("somestring", map[string]string{"s": "c"}) // returns []string{"somestring", "somectring", "comestring", "comectring"}```