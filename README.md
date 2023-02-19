<h1 align="center">natter 🐍</h1>

<p align="center">
  A small tool that aims to help debugging NAT on a firewall.
  <br><br>
  This tool listens on the specified interfaces and prints all packets, that match the given filter and target.
</p>


## 📦 Requirements

- Golang (>=1.19)
- libpcap
- Make
- go vet (✅ tests/linting)
- staticcheck (✅ tests/linting)
- gosec (✅ tests/linting)

*or*

- nix


## 🔧 Usage

First build the binary according to the build instructions.
Then you should be able to run the binary and configure it according to the help instructions (`./bin/natter --help`).

Here's an example, that displays all icmp packets on interface en0 and lo0.

```bash
sudo bin/natter -i en0,lo0 -f "icmp"
```


## 🏗 Build

In order to build the binary, just run `make build`. The binary will be placed in the `./bin` directory.

*Hint* Since libpcap is used, the tool should be compiled at least on a machine the target OS.


## 🤝 Contributing

If you are missing features or find some annoying bugs please feel free to submit an issue or a bugfix within a pull request :)


## 📝 License

© 2023 Daniel Jankowski


This project is licensed under the MIT license.


Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:


The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.


THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
