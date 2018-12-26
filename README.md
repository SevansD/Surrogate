# Surrogate ![](https://travis-ci.org/SevansD/Surrogate.svg?branch=master) [![codecov](https://codecov.io/gh/SevansD/Surrogate/branch/master/graph/badge.svg)](https://codecov.io/gh/SevansD/Surrogate)



Package for replace character in string by replacement map.

```surrogate.Sub("somestring", map[string]string{"s": "c"}) 
// returns []string{"somestring", "somectring", "comestring", "comectring"}
surrogate.SubMultiMap("a", map[string][]string{
  "a" : {"b", "c"},
})
// returns []string{"a", "b", "c"}
```
