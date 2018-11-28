package versioning

import (
  "strings"
)

// Check if one semantic version is greater than another
func IsGreater (semver1 string, semver2 string) bool {
  semver1 = strings.TrimLeft(semver1, "v")
  semver2 = strings.TrimLeft(semver2, "v")
  split1 := strings.Split(semver1, ".")
  split2 := strings.Split(semver2, ".")
  return split1[0] > split2[0] || split1[1] > split2[1] || split1[2] > split2[2]
}

