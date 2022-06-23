[![stars](https://img.shields.io/github/stars/barbarbar338/perman?color=yellow&logo=github&style=for-the-badge)](https://github.com/barbarbar338/perman)
[![license](https://img.shields.io/github/license/barbarbar338/perman?logo=github&style=for-the-badge)](https://github.com/barbarbar338/perman)
[![supportServer](https://img.shields.io/discord/711995199945179187?color=7289DA&label=Support&logo=discord&style=for-the-badge)](https://discord.gg/BjEJFwh)
[![forks](https://img.shields.io/github/forks/barbarbar338/perman?color=green&logo=github&style=for-the-badge)](https://github.com/barbarbar338/perman)
[![issues](https://img.shields.io/github/issues/barbarbar338/perman?color=red&logo=github&style=for-the-badge)](https://github.com/barbarbar338/perman)

<p align="center">
  <img src="https://raw.githubusercontent.com/barbarbar338/readme-template/main/icon.png" alt="Logo" width="160" height="160" />
  <h3 align="center">Perman.GO</h3>

  <p align="center">
    🔑 Permission management made easy
    <br />
    <a href="https://discord.gg/BjEJFwh"><strong>Get support »</strong></a>
    <br />
    <br />
    <a href="https://github.com/barbarbar338/perman/issues">Report Bug</a>
    ·
    <a href="https://github.com/barbarbar338/perman/issues">Request Feature</a>
    ·
    <a href="https://338.rocks">Webpage</a>
  </p>
</p>

# 🔑 Perman

Permission management made easy

# 📦 Installation

-   Run `go get github.com/barbarbar338/perman`

# 🤓 Usage

See [API](#-api) for all methods

```go
package main

import (
	"fmt"

	"github.com/barbarbar338/perman"
)

func main() {
    p := perman.Factory([]string{"user", "verified", "admin"})

    user := p.Serialize([]string{"user"})
    verified := p.Serialize([]string{"user", "verified"})
    admin := p.Serialize([]string{"user", "admin"})

    p.Has(user, "user") // true
    p.Has(user, "admin") // false
    p.Has(verified, "verified") // true
    p.Has(verified, "admin") // false
    p.Has(admin, "admin") // true

    // add permissions
    p.Has(user, "verified") // false
    user = p.Add(user, "verified")
    p.Has(user, "verified") // true

    // remove permissions
    p.Has(verified, "verified") // true
    verified = p.Remove(verified, "verified")
    p.Has(verified, "verified") // false
}
```

# 📄 License

Copyright © 2022 [Barış DEMİRCİ](https://github.com/barbarbar338).

Distributed under the [GPL-3.0](https://www.gnu.org/licenses/gpl-3.0.html) License. See `LICENSE` for more information.

# 🧦 Contributing

Fell free to use GitHub's features.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/my-feature`)
3. Commit your Changes (`git commit -m 'my awesome feature my-feature'`)
4. Push to the Branch (`git push origin feature/my-feature`)
5. Open a Pull Request

# 🔥 Show your support

Give a ⭐️ if this project helped you!

# 📞 Contact

-   Mail: demirci.baris38@gmail.com
-   Discord: https://discord.gg/BjEJFwh
-   Instagram: https://www.instagram.com/ben_baris.d/
-   Webpage: https://338.rocks

# 📜 API

| Method                        | Description                                                                         | Usage                                              | Output       |
| ----------------------------- | ----------------------------------------------------------------------------------- | -------------------------------------------------- | ------------ |
| `Factory`                     | Creates a new Perman instance                                                       | `perman.Factory(flags []string{})`                 | `Perman`     |
| `Keys`                        | Returns all flag names                                                              | `p.Keys()`                                         | `[]string{}` |
| `Values`                      | Returns all flag values                                                             | `p.Values()`                                       | `[]int64`    |
| `Get`                         | Returns the numeric value of flag                                                   | `p.Get(flag string)`                               | `int64`      |
| `Serialize`                   | Serializes the flags                                                                | `p.Serialize(flags []string{})`                    | `int64`      |
| `Deserialize`                 | Deserializes the permission                                                         | `p.Deserialize(permissions int64)`                 | `[]string{}` |
| `Match`                       | Matches permissions with flags, if permissions has all flags, returns true          | `p.Match(permissions int64, flags []string{})`     | `bool`       |
| `MatchAll` (alias of `Match`) | Matches permissions with flags, if permissions has all flags, returns true          | `p.MatchAll(permissions: int64, flags []string{})` | `bool`       |
| `HasAll` (alias of `match`)   | Matches permissions with flags, if permissions has all flags, returns true          | `p.HasAll(permissions: int64, flags []string{})`   | `bool`       |
| `Some`                        | Matches permissions with flags, if permissions has at least one flag, returns true  | `p.Some(permissions: int64, flags []string{})`     | `bool`       |
| `HasSome` (alias of `some`)   | Matches permissions with flags, if permissions has at least one flag, returns true  | `p.HasSome(permissions: int64, flags []string{})`  | `bool`       |
| `HasNone`                     | Matches permissions with flags, if permissions has at least one flag, returns false | `p.HasNone(permissions: int64, flags []string{})`  | `bool`       |
| `None` (alias of `hasNone`)   | Matches permissions with flags, if permissions has at least one flag, returns false | `p.None(permissions: int64, flags []string{})`     | `bool`       |
| `Has`                         | Checks if the given permission is granted                                           | `p.Has(permission int64, flag string)`             | `bool`       |
| `Test` (alias of `has`)       | Checks if the given permission is granted                                           | `p.Test(permission int64, flag string)`            | `bool`       |
| `Add`                         | Adds a new flag to given permission                                                 | `p.Add(permission int64, flag string)`             | `int64`      |
| `Remove`                      | Removes a flag from given permission                                                | `p.Remove(permission int64, flag string)`          | `int64`      |
| `Full`                        | Creates a permission with all flags                                                 | `p.Full()`                                         | `int64`      |
